package apache_jena

import (
	"backend/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUpdateParameterQuery(t *testing.T) {
	service := NewService("http://example.com/", "http://example.com/sparql", "test", "test")

	tests := []struct {
		name       string
		parameter  models.ParameterView
		wantUpdate string
	}{
		{
			name: "No classes or contradictions",
			parameter: models.ParameterView{
				ID:                      "param1",
				Title:                   "Parameter 1",
				AllowedClasses:          []uint{},
				ContradictionParameters: []string{},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:param_param1 a :Parameter .
}`,
		},
		{
			name: "With classes",
			parameter: models.ParameterView{
				ID:                      "param2",
				Title:                   "Parameter 2",
				AllowedClasses:          []uint{1, 33},
				ContradictionParameters: []string{},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:param_param2 a :Parameter ;
	:allowedClass :class_1 ;
	:allowedClass :class_33 .
}`,
		},
		{
			name: "With contradictions",
			parameter: models.ParameterView{
				ID:                      "param3",
				Title:                   "Parameter 3",
				AllowedClasses:          []uint{},
				ContradictionParameters: []string{"paramA", "paramB"},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:param_param3 a :Parameter ;
	:hasContradictionParameter :param_paramA ;
	:hasContradictionParameter :param_paramB .
}`,
		},
		{
			name: "With classes and contradictions",
			parameter: models.ParameterView{
				ID:                      "param4",
				Title:                   "Parameter 4",
				AllowedClasses:          []uint{1},
				ContradictionParameters: []string{"paramA"},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:param_param4 a :Parameter ;
	:allowedClass :class_1 ;
	:hasContradictionParameter :param_paramA .
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUpdate := service.buildUpdateParameterQuery(tt.parameter, false)
			assert.Equal(t, tt.wantUpdate, gotUpdate)
		})
	}
}

func TestBuildUpdateClassQuery(t *testing.T) {
	service := NewService("http://example.com/", "http://example.com/sparql", "test", "test")

	tests := []struct {
		name       string
		class      models.ClassView
		wantUpdate string
	}{
		{
			name: "No allowed parameters",
			class: models.ClassView{
				ID:                1,
				AllowedParameters: []string{},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:class_1 a :Class .
}`,
		},
		{
			name: "With allowed parameters",
			class: models.ClassView{
				ID:                2,
				AllowedParameters: []string{"param1", "param2"},
			},
			wantUpdate: `PREFIX : <http://example.com/>
INSERT DATA {
	:class_2 a :Class ;
	:hasAllowedParameter :param_param1 ;
	:hasAllowedParameter :param_param2 .
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUpdate := service.buildUpdateClassQuery(tt.class, false)
			assert.Equal(t, tt.wantUpdate, gotUpdate)
		})
	}
}
