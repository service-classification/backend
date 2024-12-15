package main

import (
	"backend/config"
	"backend/docs"
	"backend/internal/apache_jena"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/router"
	"backend/internal/services"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "backend/docs" // This line is necessary for go-swagger to find docs

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			MyApp API
//	@version		1.0
//	@description	This is a backend server.

//	@host		194.135.25.202:8080
//	@BasePath	/

//	@schemes	http

func main() {
	cfg := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Class{}, &models.Parameter{}, &models.Service{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	baseURL := fmt.Sprintf("http://%s:%s/%s", os.Getenv("KB_HOST"), os.Getenv("KB_PORT"), os.Getenv("KB_DATASET"))
	publicURL := fmt.Sprintf("http://%s:%s/%s#", os.Getenv("PUBLIC_HOST"), os.Getenv("KB_PORT"), os.Getenv("KB_DATASET"))
	jenaService := apache_jena.NewService(publicURL, baseURL, os.Getenv("KB_LOGIN"), os.Getenv("KB_PASSWORD"))

	serviceRepo := repositories.NewServiceRepository(db)
	classRepo := repositories.NewClassRepository(db)
	paramRepo := repositories.NewParameterRepository(db)
	parameterService := services.NewParameterService(paramRepo, jenaService)
	classService := services.NewClassService(classRepo, jenaService)
	handler := handlers.NewHandler(serviceRepo, classRepo, paramRepo, jenaService, parameterService, classService)

	migrate(classService, parameterService)

	host := os.Getenv("PUBLIC_HOST")
	docs.SwaggerInfo.Host = host + ":8080"
	docs.SwaggerInfo.Description = "This is a backend server."

	r := router.SetupRouter(handler)

	log.Fatal(r.Run(":8080"))
}

func migrate(classService *services.ClassService, parameterService *services.ParameterService) {

	_, err := parameterService.CreateParameter(models.ParameterView{Title: "Mobile Internet", ID: "mob_inet"}, false)
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed Internet", ID: "fix_inet"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Home television via CTV technology (digital TV)", ID: "fix_ctv"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Home television via ICTV technology (interactive digital TV)", ID: "fix_ictv"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Mobile voice communication services", ID: "voice_mob"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed voice communication services", ID: "voice_fix"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Short Message Service (SMS)", ID: "sms"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Circuit Switched Data (CSD)", ID: "csd"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Internet of Things", ID: "iot"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Multimedia Messaging Service (MMS)", ID: "mms"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Provision of mobile communication services in other operators’ networks", ID: "roaming"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "International communication services", ID: "mn"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Call/SMS routing within the network", ID: "mts"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Call/SMS routing to other operators’ networks (competitor networks)", ID: "conc"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Call/SMS routing to fixed operators’ networks", ID: "fix_op"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Intra-network roaming (domestic travel within Russia)", ID: "vsr_roam"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Roaming in other operators’ networks (domestic travel within Russia)", ID: "national_roam"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Roaming in other operators’ networks (international travel)", ID: "mn_roam"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Long-distance communication services", ID: "mg"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Mobile voice communication services, subscription fee", ID: "voice_ap"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Mobile voice communication services, connection fee", ID: "voice_fee"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Recurring services", ID: "period_service"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "One-time services (services related to adding/removing recurring services in billing, and unique services such as owner change, number change, etc.)", ID: "one_time_service"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Additional services", ID: "dop_service"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Content-based informational and entertainment services", ID: "content"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Additional subscriber support services", ID: "services_service"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Sale of express payment cards", ID: "keo_sale"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Discounts", ID: "discount"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Incoming calls only", ID: "only_inbound"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "A2P (Application-to-Person) SMS Messaging", ID: "sms_a2p"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Same as A2P (Application-to-Person) SMS Messaging but with a different revenue allocation model with a partner", ID: "sms_gross"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Other additional services", ID: "other_service"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Voicemail services", ID: "voice_mail"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Geoanalytics services (subscriber clustering, etc.)", ID: "geo"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed telephony services, monthly fee per number", ID: "ep_for_number"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed telephony services, one-time fee per line", ID: "ep_for_line"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed telephony services, one-time fee per number", ID: "one_time_fee_for_number"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Fixed communication services, equipment rental", ID: "equipment_rent"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Scoring", ID: "skoring"}, false)))
	err = errors.Join(err, second(parameterService.CreateParameter(models.ParameterView{Title: "Additional package", ID: "add_package"}, false)))

	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1, Title: "Subscription fee", AllowedParameters: []string{"mg", "voice_ap", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 4, Title: "Long-distance communication", AllowedParameters: []string{"voice_mob", "mg", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 5, Title: "International communication", AllowedParameters: []string{"voice_mob", "mn", "national_roam", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 6, Title: "Intra-network roaming", AllowedParameters: []string{"voice_mob", "roaming", "vsr_roam", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 7, Title: "National roaming", AllowedParameters: []string{"voice_mob", "roaming", "national_roam", "period_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 8, Title: "International roaming", AllowedParameters: []string{"voice_mob", "roaming", "mn_roam", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 9, Title: "SMS", AllowedParameters: []string{"sms", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 10, Title: "Other additional services", AllowedParameters: []string{"mob_inet", "period_service", "one_time_service", "other_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 11, Title: "SMS content services", AllowedParameters: []string{"sms", "roaming", "vsr_roam", "one_time_service", "dop_service", "other_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 12, Title: "Data transfer and telematics services", AllowedParameters: []string{"voice_mob", "csd", "mg", "services_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 15, Title: "Other subscriber services", AllowedParameters: []string{"services_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 16, Title: "Goods and accessories", AllowedParameters: []string{"one_time_service", "dop_service", "services_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 22, Title: "Discounts", AllowedParameters: []string{"voice_fix", "period_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 29, Title: "GPRS revenues", AllowedParameters: []string{"mob_inet", "period_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 33, Title: "Incoming traffic modifiers in roaming", AllowedParameters: []string{"voice_mob", "roaming", "mn_roam", "only_inbound"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 90, Title: "SMS A2P", AllowedParameters: []string{"services_service", "sms_a2p"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 91, Title: "Revenues from SMS mailings (gross)", AllowedParameters: []string{"one_time_service", "services_service", "sms_gross"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 99, Title: "Undefined.", AllowedParameters: []string{}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 100, Title: "Per-minute fee for incoming calls", AllowedParameters: []string{"voice_mob", "only_inbound"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 101, Title: "Per-minute fee for intra-network zonal traffic", AllowedParameters: []string{"voice_mob", "mts", "voice_fee", "discount", "other_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 102, Title: "Per-minute fee for zonal traffic to other operators", AllowedParameters: []string{"voice_mob", "conc", "services_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 103, Title: "Per-minute fee for long-distance intra-network traffic", AllowedParameters: []string{"voice_mob", "roaming", "mg", "mts", "vsr_roam", "services_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 104, Title: "Per-minute fee for long-distance traffic to other operators", AllowedParameters: []string{"voice_mob", "csd", "mg", "conc"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 105, Title: "Per-minute fee for international traffic", AllowedParameters: []string{"voice_mob", "csd", "roaming", "mn", "vsr_roam", "services_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1001, Title: "Subscription fee for voice services", AllowedParameters: []string{"voice_mob", "voice_ap", "voice_fee", "period_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1004, Title: "SMS in international roaming", AllowedParameters: []string{"voice_mob", "sms", "roaming", "conc", "fix_op", "mn_roam", "period_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1006, Title: "GPRS in national roaming", AllowedParameters: []string{"mob_inet", "roaming", "national_roam", "period_service", "one_time_service", "services_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1007, Title: "GPRS in international roaming", AllowedParameters: []string{"mob_inet", "roaming", "mn_roam", "period_service", "services_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1016, Title: "MMS", AllowedParameters: []string{"mms", "mts", "period_service", "dop_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1027, Title: "Revenues from activation of voice services valid for over 6 months", AllowedParameters: []string{"voice_mob", "voice_fee", "period_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1030, Title: "Revenues from voicemail", AllowedParameters: []string{"voice_mob", "period_service", "discount", "voice_mail"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1033, Title: "Revenues from Internet access services", AllowedParameters: []string{"fix_inet", "one_time_service", "dop_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1034, Title: "Revenues from telephony services", AllowedParameters: []string{"voice_fix", "period_service", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1100, Title: "Revenues from combined Voice/SMS/GPRS services", AllowedParameters: []string{"mob_inet", "voice_mob", "sms", "period_service", "services_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1121, Title: "Revenues from IoT geoanalytics", AllowedParameters: []string{"iot", "period_service", "dop_service", "discount", "geo"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1124, Title: "Revenues from IoT SMS", AllowedParameters: []string{"sms", "iot", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1125, Title: "Revenues from IoT GPRS", AllowedParameters: []string{"mob_inet", "iot", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 1126, Title: "Revenues from voice + CSD IoT", AllowedParameters: []string{"voice_mob", "csd", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3001, Title: "_FB. Telephony - subscription fee per number", AllowedParameters: []string{"voice_fix", "voice_ap", "dop_service", "discount", "ep_for_number"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3002, Title: "_FB. Telephony - subscription fee per line", AllowedParameters: []string{"voice_fix", "one_time_service", "dop_service", "ep_for_line"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3003, Title: "_FB. Telephony - setup fee for subscriber number", AllowedParameters: []string{"voice_fix", "period_service", "one_time_service", "discount", "one-time_fee_for_number"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3009, Title: "_FB. Telephony - zonal outgoing calls to fixed operators", AllowedParameters: []string{"voice_fix", "fix_op", "voice_ap", "period_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3010, Title: "_FB. Telephony - zonal outgoing calls to SPC", AllowedParameters: []string{"voice_fix", "conc"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3019, Title: "_FB. Telephony - long-distance services from own subscribers", AllowedParameters: []string{"voice_fix", "mg", "period_service", "one_time_service", "dop_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3020, Title: "_FB. Telephony - international services from own subscribers", AllowedParameters: []string{"voice_fix", "mn", "one_time_service", "dop_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3031, Title: "_FB. Telephony - additional services", AllowedParameters: []string{"voice_fix", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3309, Title: "FB_Television (CTV) - subscription fee per line", AllowedParameters: []string{"fix_ctv", "period_service", "dop_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3311, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (CTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_ctv", "voice_fix", "period_service", "one_time_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3312, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (CTV) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ctv", "period_service", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3313, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + CTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ctv", "voice_fix", "period_service", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3314, Title: "FB_Sale of goods, works, and services - Individuals - Equipment rental for CTV service", AllowedParameters: []string{"fix_ctv", "period_service", "equipment_rent"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3320, Title: "FB_Sale of goods, works, and services - Individuals - Television (ICTV) - Subscription fee", AllowedParameters: []string{"fix_ictv", "period_service", "one_time_service", "dop_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3321, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (ICTV) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ictv", "period_service", "one_time_service", "dop_service"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3322, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (ICTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_ictv", "voice_fix", "period_service", "one_time_service", "discount"}}, false)))
	err = errors.Join(err, second(classService.CreateClass(models.ClassView{ID: 3323, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + ICTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ictv", "voice_fix", "period_service", "one_time_service"}}, false)))

	if err != nil {
		slog.Error("Error while migration", slog.Any("error", err))
	}
}

func second[T any, T2 any](a T, b T2) T2 {
	return b
}
