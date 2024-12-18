package apache_jena

import (
	"backend/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	prefix   string
	baseURL  string
	login    string
	password string
	client   *http.Client
}

func NewService(prefix string, baseURL string, login, password string) *Service {
	return &Service{
		prefix:   prefix,
		baseURL:  baseURL,
		login:    login,
		password: password,
		client: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (s *Service) AddParameter(ctx context.Context, parameter models.ParameterView) error {
	update := s.buildUpdateParameterQuery(parameter, false)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) buildUpdateParameterQuery(parameter models.ParameterView, forUpdate bool) string {
	allowedClasses := make([]string, 0, len(parameter.AllowedClasses))
	for _, class := range parameter.AllowedClasses {
		allowedClasses = append(allowedClasses, fmt.Sprintf("\t:allowedClass :class_%d", class))
	}
	classes := strings.Join(allowedClasses, " ;\n")
	if len(classes) > 0 {
		if len(parameter.ContradictionParameters) > 0 {
			classes += " ;\n"
		} else {
			classes += " .\n"
		}
	}

	contradictions := make([]string, 0, len(parameter.ContradictionParameters))
	for _, contradiction := range parameter.ContradictionParameters {
		contradictions = append(contradictions, fmt.Sprintf("\t:hasContradictionParameter :param_%s", contradiction))
	}
	contradictionsStr := strings.Join(contradictions, " ;\n")
	if len(contradictionsStr) > 0 {
		contradictionsStr += " .\n"
	}

	var update strings.Builder
	update.WriteString("PREFIX : <")
	update.WriteString(s.prefix)
	update.WriteString(">\n")
	if !forUpdate {
		update.WriteString("INSERT DATA {\n")
	} else {
		update.WriteString("INSERT {\n")
	}
	update.WriteString("\t:param_")
	update.WriteString(parameter.ID)
	update.WriteString(" a :Parameter")
	if len(parameter.AllowedClasses) == 0 && len(parameter.ContradictionParameters) == 0 {
		update.WriteString(" .\n")
	} else {
		update.WriteString(" ;\n")
	}
	if len(classes) > 0 {
		update.WriteString(classes)
	}
	if len(contradictionsStr) > 0 {
		update.WriteString(contradictionsStr)
	}
	update.WriteString("}")

	return update.String()
}

func (s *Service) UpdateParameter(ctx context.Context, parameter models.ParameterView) error {
	update := s.buildUpdateParameterQuery(parameter, true)

	update += fmt.Sprintf(`
		WHERE {
		  :param_%s ?p ?o .
		}
	`, parameter.ID)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) DeleteParameter(ctx context.Context, id string) error {
	update := fmt.Sprintf(`
		PREFIX : <%s>
		DELETE {
		  :param_%s ?p ?o .
		}
		WHERE {
		  :param_%s ?p ?o .
		}
	`, s.prefix, id, id)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) GetParameterConstraints(ctx context.Context, parameterID string) ([]uint, []string, error) {
	query := fmt.Sprintf(`
		PREFIX : <%s>
		SELECT ?contradictionParam
		WHERE {
		    :param_%s :hasContradictionParameter ?contradictionParam .
		}
	`, s.prefix, parameterID)

	contrParams, err := s.query(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	var contradictions []string
	for _, binding := range contrParams.Results.Bindings {
		if contradiction, ok := binding["contradictionParam"]; ok {
			contradictionValue := strings.TrimPrefix(contradiction["value"], s.prefix+"param_")
			contradictions = append(contradictions, contradictionValue)
		}
	}

	query = fmt.Sprintf(`
		PREFIX : <%s>
		SELECT ?class
		WHERE {
			{ 
				?class :hasAllowedParameter :param_%s .
			}
			UNION
			{
				:param_%s :allowedClass ?class .
			}  
		}
	`, s.prefix, parameterID, parameterID)

	allowedClasses, err := s.query(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	var classes []uint
	for _, binding := range allowedClasses.Results.Bindings {
		if class, ok := binding["class"]; ok {
			classIDStr := strings.TrimPrefix(class["value"], s.prefix+"class_")
			classID, err := strconv.ParseUint(classIDStr, 10, 64)
			if err != nil {
				return nil, nil, err
			}
			classes = append(classes, uint(classID))
		}
	}

	return classes, contradictions, nil
}

func (s *Service) AddClass(ctx context.Context, class models.ClassView) error {
	update := s.buildUpdateClassQuery(class, false)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) buildUpdateClassQuery(class models.ClassView, forUpdate bool) string {
	allowedParameters := make([]string, 0, len(class.AllowedParameters))
	for _, parameter := range class.AllowedParameters {
		allowedParameters = append(allowedParameters, fmt.Sprintf("\t:hasAllowedParameter :param_%s", parameter))
	}
	parameters := strings.Join(allowedParameters, " ;\n")
	if len(parameters) > 0 {
		parameters += " .\n"
	}

	var update strings.Builder
	update.WriteString("PREFIX : <")
	update.WriteString(s.prefix)
	update.WriteString(">\n")
	if !forUpdate {
		update.WriteString("INSERT DATA {\n")
	} else {
		update.WriteString("INSERT {\n")
	}
	update.WriteString("\t:class_")
	update.WriteString(strconv.FormatUint(uint64(class.ID), 10))
	update.WriteString(" a :Class")
	if len(class.AllowedParameters) > 0 {
		update.WriteString(" ;\n")
	} else {
		update.WriteString(" .\n")
	}
	if len(parameters) > 0 {
		update.WriteString(parameters)
	}
	update.WriteString("}")

	return update.String()
}

func (s *Service) UpdateClass(ctx context.Context, class models.ClassView) error {
	update := s.buildUpdateClassQuery(class, true)

	update += fmt.Sprintf(`
		WHERE {
		  :class_%d ?p ?o .
		}
	`, class.ID)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) DeleteClass(ctx context.Context, id uint) error {
	update := fmt.Sprintf(`
		PREFIX : <%s>
		DELETE {
		  :class_%d ?p ?o .
		}
		WHERE {
		  :class_%d ?p ?o .
		}
	`, s.prefix, id, id)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) GetClassConstraints(ctx context.Context, classID uint) ([]string, error) {
	query := fmt.Sprintf(`
		PREFIX : <%s>
		SELECT ?allowedParam
		WHERE {
			{
				:class_%d :hasAllowedParameter ?allowedParam .
			}
			UNION
			{
				?allowedParam :allowedClass :class_%d .
			}
		}
	`, s.prefix, classID, classID)

	result, err := s.query(ctx, query)
	if err != nil {
		return nil, err
	}

	var allowedParams []string
	for _, binding := range result.Results.Bindings {
		if allowedParam, ok := binding["allowedParam"]; ok {
			allowedParamValue := strings.TrimPrefix(allowedParam["value"], s.prefix+"param_")
			allowedParams = append(allowedParams, allowedParamValue)
		}
	}

	return allowedParams, nil
}

func (s *Service) AddService(ctx context.Context, service *models.Service) error {
	update := s.buildUpdateServiceQuery(service, false)

	return s.runSparqlUpdate(ctx, update)
}

func (s *Service) buildUpdateServiceQuery(service *models.Service, forUpdate bool) string {
	parameters := make([]string, 0, len(service.Parameters))
	for _, parameter := range service.Parameters {
		parameters = append(parameters, fmt.Sprintf(":param_%s", parameter.ID))
	}
	parametersStr := strings.Join(parameters, " , ")

	var update strings.Builder
	update.WriteString("PREFIX : <")
	update.WriteString(s.prefix)
	update.WriteString(">\n")
	if !forUpdate {
		update.WriteString("INSERT DATA {\n")
	} else {
		update.WriteString("INSERT {\n")
	}
	update.WriteString("\t:service_")
	update.WriteString(strconv.FormatUint(uint64(service.ID), 10))
	update.WriteString(" a :Service")
	if service.ClassID != nil {
		update.WriteString(" ;\n")
		update.WriteString("\t\t:hasClass :class_")
		update.WriteString(strconv.FormatUint(uint64(*service.ClassID), 10))
	}
	if len(parameters) > 0 {
		update.WriteString(" ;\n")
		update.WriteString("\t\t:hasParameter ")
		update.WriteString(parametersStr)
	}
	update.WriteString(" .\n")
	update.WriteString("}")

	return update.String()
}

type ProposedClass struct {
	ClassID               uint
	MatchingParameterNums int
	SimilarServices       []uint
}

func (s *Service) ProposedClasses(ctx context.Context, service *models.Service) ([]ProposedClass, error) {
	var serviceParams []string
	for _, param := range service.Parameters {
		serviceParams = append(serviceParams, fmt.Sprintf(":param_%s", param.ID))
	}
	serviceParam := strings.Join(serviceParams, " ")
	allowedParam := strings.Join(serviceParams, ", ")

	query := fmt.Sprintf(`
		PREFIX : <%s>
		SELECT ?class (COUNT(DISTINCT ?commonParam) AS ?matching_parameter_numbers) (GROUP_CONCAT(DISTINCT ?similarService; SEPARATOR=",") AS ?similar_services)
		WHERE {
		  VALUES ?serviceParam { %s } # service parameters
		  ?class a :Class ;
		            :hasAllowedParameter ?allowedParam .
		  FILTER(?allowedParam IN (%s))
		  BIND(?allowedParam AS ?commonParam)
		  OPTIONAL {
		    ?similarService a :Service ;
		                    :hasParameter ?allowedParam ;
		                    :hasClass ?class .
		    FILTER(?similarService != :service_%d)
		  }
		}
		GROUP BY ?class
		ORDER BY DESC(?matching_parameter_numbers)
	`, s.prefix, serviceParam, allowedParam, service.ID)

	result, err := s.query(ctx, query)
	if err != nil {
		return nil, err
	}

	classes := make([]ProposedClass, 0, len(result.Results.Bindings))
	for _, binding := range result.Results.Bindings {
		var class ProposedClass
		classIDStr := strings.TrimPrefix(binding["class"]["value"], s.prefix+"class_")
		classID, err := strconv.ParseUint(classIDStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse class ID: %w", err)
		}
		class.ClassID = uint(classID)

		class.MatchingParameterNums, err = strconv.Atoi(binding["matching_parameter_numbers"]["value"])
		if err != nil {
			return nil, fmt.Errorf("failed to parse matching parameter numbers: %w", err)
		}

		similarServices := strings.Split(binding["similar_services"]["value"], ",")
		for _, similarService := range similarServices {
			similarServiceIDStr := strings.TrimPrefix(similarService, s.prefix+"service_")
			if similarServiceIDStr == "" {
				continue
			}

			similarServiceID, err := strconv.ParseUint(similarServiceIDStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse similar service ID: %w", err)
			}
			class.SimilarServices = append(class.SimilarServices, uint(similarServiceID))
		}

		classes = append(classes, class)
	}

	return classes, nil
}

func (s *Service) ValidateClass(ctx context.Context, service *models.Service, chosenClass uint) (bool, error) {
	serviceParams := make([]string, 0, len(service.Parameters))
	for _, param := range service.Parameters {
		serviceParams = append(serviceParams, fmt.Sprintf(":param_%s", param.ID))
	}
	serviceParam := strings.Join(serviceParams, " ")

	query := fmt.Sprintf(`
		PREFIX : <%s>
		SELECT (IF((COUNT(?allowedParam) = %d), true, false) AS ?allAllowed)
		WHERE {
		  VALUES ?requiredParam { %s } # service parameters
		  :class_%d :hasAllowedParameter ?allowedParam .
		  FILTER(?allowedParam = ?requiredParam)
		}
	`, s.prefix, len(service.Parameters), serviceParam, chosenClass)

	result, err := s.query(ctx, query)
	if err != nil {
		return false, err
	}

	return result.Results.Bindings[0]["allAllowed"]["value"] == "true", nil
}

func (s *Service) ValidateService(ctx context.Context, service *models.Service) ([]string, error) {
	serviceParams := make([]string, 0, len(service.Parameters))
	for _, param := range service.Parameters {
		serviceParams = append(serviceParams, fmt.Sprintf(":param_%s", param.ID))
	}
	serviceParam := strings.Join(serviceParams, " ")
	query := fmt.Sprintf(`
		PREFIX : <%s>
		SELECT ?p1 ?p2
		WHERE {
		  VALUES ?p1 { %s }
		  VALUES ?p2 { %s }
		  FILTER(?p1 != ?p2)
		  ?p1 :hasContradictionParameter ?p2 .
		}
	`, s.prefix, serviceParam, serviceParam)

	result, err := s.query(ctx, query)
	if err != nil {
		return nil, err
	}

	contradictions := make([]string, 0, len(result.Results.Bindings))
	for _, binding := range result.Results.Bindings {
		value := strings.TrimPrefix(binding["p1"]["value"], s.prefix+"param_")
		contradictions = append(contradictions, value)
	}

	return contradictions, nil
}

func (s *Service) runSparqlUpdate(ctx context.Context, update string) error {

	req, err := http.NewRequestWithContext(ctx, "POST", s.baseURL+"/update", bytes.NewBufferString(update))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/sparql-update")
	req.Header.Set("Accept", "application/sparql-results+json")
	req.SetBasicAuth(s.login, s.password)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("SPARQL update failed with code %d", resp.StatusCode)
		}
		return fmt.Errorf("SPARQL update failed with code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (s *Service) query(ctx context.Context, sparql string) (*SparqlResult, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", s.baseURL+"/query", bytes.NewBufferString(sparql))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/sparql-query")
	req.Header.Set("Accept", "application/sparql-results+json")
	req.SetBasicAuth(s.login, s.password)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result SparqlResult
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling response: %v\nResponse: %s", err, string(bodyBytes))
	}

	return &result, nil
}

type SparqlResult struct {
	Head struct {
		Vars []string `json:"vars"`
	} `json:"head"`
	Results struct {
		Bindings []map[string]map[string]string `json:"bindings"`
	} `json:"results"`
}
