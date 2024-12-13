package main

import (
	"backend/config"
	"backend/docs"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/router"
	"backend/internal/services"
	"fmt"
	"log"
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
	err = db.AutoMigrate(&models.Group{}, &models.Parameter{}, &models.Service{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	serviceRepo := repositories.NewServiceRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	paramRepo := repositories.NewParameterRepository(db)
	handler := handlers.NewHandler(serviceRepo, groupRepo, paramRepo)

	migrate(groupRepo, paramRepo)

	docs.SwaggerInfo.Host = os.Getenv("PUBLIC_HOST") + ":8080"
	docs.SwaggerInfo.Description = "This is a backend server."

	r := router.SetupRouter(handler)

	log.Fatal(r.Run(":8080"))
}

func migrate(groups repositories.GroupRepository, parameters repositories.ParameterRepository) {
	parameterService := services.NewParameterService(parameters)
	groupService := services.NewGroupService(groups)

	parameterService.CreateParameter(services.ParameterView{Title: "Mobile Internet", ID: "mob_inet"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed Internet", ID: "fix_inet"})
	parameterService.CreateParameter(services.ParameterView{Title: "Home television via CTV technology (digital TV)", ID: "fix_ctv"})
	parameterService.CreateParameter(services.ParameterView{Title: "Home television via ICTV technology (interactive digital TV)", ID: "fix_ictv"})
	parameterService.CreateParameter(services.ParameterView{Title: "Mobile voice communication services", ID: "voice_mob"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed voice communication services", ID: "voice_fix"})
	parameterService.CreateParameter(services.ParameterView{Title: "Short Message Service (SMS)", ID: "sms"})
	parameterService.CreateParameter(services.ParameterView{Title: "Circuit Switched Data (CSD)", ID: "csd"})
	parameterService.CreateParameter(services.ParameterView{Title: "Internet of Things", ID: "iot"})
	parameterService.CreateParameter(services.ParameterView{Title: "Multimedia Messaging Service (MMS)", ID: "mms"})
	parameterService.CreateParameter(services.ParameterView{Title: "Provision of mobile communication services in other operators’ networks", ID: "roaming"})
	parameterService.CreateParameter(services.ParameterView{Title: "International communication services", ID: "mn"})
	parameterService.CreateParameter(services.ParameterView{Title: "Call/SMS routing within the network", ID: "mts"})
	parameterService.CreateParameter(services.ParameterView{Title: "Call/SMS routing to other operators’ networks (competitor networks)", ID: "conc"})
	parameterService.CreateParameter(services.ParameterView{Title: "Call/SMS routing to fixed operators’ networks", ID: "fix_op"})
	parameterService.CreateParameter(services.ParameterView{Title: "Intra-network roaming (domestic travel within Russia)", ID: "vsr_roam"})
	parameterService.CreateParameter(services.ParameterView{Title: "Roaming in other operators’ networks (domestic travel within Russia)", ID: "national_roam"})
	parameterService.CreateParameter(services.ParameterView{Title: "Roaming in other operators’ networks (international travel)", ID: "mn_roam"})
	parameterService.CreateParameter(services.ParameterView{Title: "Long-distance communication services", ID: "mg"})
	parameterService.CreateParameter(services.ParameterView{Title: "Mobile voice communication services, subscription fee", ID: "voice_ap"})
	parameterService.CreateParameter(services.ParameterView{Title: "Mobile voice communication services, connection fee", ID: "voice_fee"})
	parameterService.CreateParameter(services.ParameterView{Title: "Recurring services", ID: "period_service"})
	parameterService.CreateParameter(services.ParameterView{Title: "One-time services (services related to adding/removing recurring services in billing, and unique services such as owner change, number change, etc.)", ID: "one_time_service"})
	parameterService.CreateParameter(services.ParameterView{Title: "Additional services", ID: "dop_service"})
	parameterService.CreateParameter(services.ParameterView{Title: "Content-based informational and entertainment services", ID: "content"})
	parameterService.CreateParameter(services.ParameterView{Title: "Additional subscriber support services", ID: "services_service"})
	parameterService.CreateParameter(services.ParameterView{Title: "Sale of express payment cards", ID: "keo_sale"})
	parameterService.CreateParameter(services.ParameterView{Title: "Discounts", ID: "discount"})
	parameterService.CreateParameter(services.ParameterView{Title: "Incoming calls only", ID: "only_inbound"})
	parameterService.CreateParameter(services.ParameterView{Title: "A2P (Application-to-Person) SMS Messaging", ID: "sms_a2p"})
	parameterService.CreateParameter(services.ParameterView{Title: "Same as A2P (Application-to-Person) SMS Messaging but with a different revenue allocation model with a partner", ID: "sms_gross"})
	parameterService.CreateParameter(services.ParameterView{Title: "Other additional services", ID: "other_service"})
	parameterService.CreateParameter(services.ParameterView{Title: "Voicemail services", ID: "voice_mail"})
	parameterService.CreateParameter(services.ParameterView{Title: "Geoanalytics services (subscriber clustering, etc.)", ID: "geo"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed telephony services, monthly fee per number", ID: "ep_for_number"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed telephony services, one-time fee per line", ID: "ep_for_line"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed telephony services, one-time fee per number", ID: "onetime_fee_for_number"})
	parameterService.CreateParameter(services.ParameterView{Title: "Fixed communication services, equipment rental", ID: "equipment rent"})

	groupService.CreateGroup(services.GroupView{ID: 1, Title: "Subscription fee", AllowedParameters: []string{"mg", "voice_ap", "period_service", "one_time_service", "group_id"}})
	groupService.CreateGroup(services.GroupView{ID: 4, Title: "Long-distance communication", AllowedParameters: []string{"voice_mob", "mg", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 5, Title: "International communication", AllowedParameters: []string{"voice_mob", "mn", "national_roam", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 6, Title: "Intra-network roaming", AllowedParameters: []string{"voice_mob", "roaming", "vsr_roam", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 7, Title: "National roaming", AllowedParameters: []string{"voice_mob", "roaming", "national_roam", "period_service"}})
	groupService.CreateGroup(services.GroupView{ID: 8, Title: "International roaming", AllowedParameters: []string{"voice_mob", "roaming", "mn_roam", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 9, Title: "SMS", AllowedParameters: []string{"sms", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 10, Title: "Other additional services", AllowedParameters: []string{"mob_inet", "period_service", "one_time_service", "other_service"}})
	groupService.CreateGroup(services.GroupView{ID: 11, Title: "SMS content services", AllowedParameters: []string{"sms", "roaming", "vsr_roam", "one_time_service", "dop_service", "other_service"}})
	groupService.CreateGroup(services.GroupView{ID: 12, Title: "Data transfer and telematics services", AllowedParameters: []string{"voice_mob", "csd", "mg", "services_service"}})
	groupService.CreateGroup(services.GroupView{ID: 15, Title: "Other subscriber services", AllowedParameters: []string{"services_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 16, Title: "Goods and accessories", AllowedParameters: []string{"one_time_service", "dop_service", "services_service"}})
	groupService.CreateGroup(services.GroupView{ID: 22, Title: "Discounts", AllowedParameters: []string{"voice_fix", "period_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 29, Title: "GPRS revenues", AllowedParameters: []string{"mob_inet", "period_service"}})
	groupService.CreateGroup(services.GroupView{ID: 33, Title: "Incoming traffic modifiers in roaming", AllowedParameters: []string{"voice_mob", "roaming", "mn_roam", "only_inbound"}})
	groupService.CreateGroup(services.GroupView{ID: 90, Title: "SMS A2P", AllowedParameters: []string{"services_service", "sms_a2p"}})
	groupService.CreateGroup(services.GroupView{ID: 91, Title: "Revenues from SMS mailings (gross)", AllowedParameters: []string{"one_time_service", "services_service", "sms_gross"}})
	groupService.CreateGroup(services.GroupView{ID: 99, Title: "Undefined.", AllowedParameters: []string{}})
	groupService.CreateGroup(services.GroupView{ID: 100, Title: "Per-minute fee for incoming calls", AllowedParameters: []string{"voice_mob", "only_inbound"}})
	groupService.CreateGroup(services.GroupView{ID: 101, Title: "Per-minute fee for intra-network zonal traffic", AllowedParameters: []string{"voice_mob", "mts", "voice_fee", "discount", "other_service"}})
	groupService.CreateGroup(services.GroupView{ID: 102, Title: "Per-minute fee for zonal traffic to other operators", AllowedParameters: []string{"voice_mob", "conc", "services_service"}})
	groupService.CreateGroup(services.GroupView{ID: 103, Title: "Per-minute fee for long-distance intra-network traffic", AllowedParameters: []string{"voice_mob", "roaming", "mg", "mts", "vsr_roam", "services_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 104, Title: "Per-minute fee for long-distance traffic to other operators", AllowedParameters: []string{"voice_mob", "csd", "mg", "conc"}})
	groupService.CreateGroup(services.GroupView{ID: 105, Title: "Per-minute fee for international traffic", AllowedParameters: []string{"voice_mob", "csd", "roaming", "mn", "vsr_roam", "services_service"}})
	groupService.CreateGroup(services.GroupView{ID: 1001, Title: "Subscription fee for voice services", AllowedParameters: []string{"voice_mob", "voice_ap", "voice_fee", "period_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1004, Title: "SMS in international roaming", AllowedParameters: []string{"voice_mob", "sms", "roaming", "conc", "fix_op", "mn_roam", "period_service"}})
	groupService.CreateGroup(services.GroupView{ID: 1006, Title: "GPRS in national roaming", AllowedParameters: []string{"mob_inet", "roaming", "national_roam", "period_service", "one_time_service", "services_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1007, Title: "GPRS in international roaming", AllowedParameters: []string{"mob_inet", "roaming", "mn_roam", "period_service", "services_service"}})
	groupService.CreateGroup(services.GroupView{ID: 1016, Title: "MMS", AllowedParameters: []string{"mms", "mts", "period_service", "dop_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1027, Title: "Revenues from activation of voice services valid for over 6 months", AllowedParameters: []string{"voice_mob", "voice_fee", "period_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1030, Title: "Revenues from voicemail", AllowedParameters: []string{"voice_mob", "period_service", "discount", "voice_mail"}})
	groupService.CreateGroup(services.GroupView{ID: 1033, Title: "Revenues from Internet access services", AllowedParameters: []string{"fix_inet", "one_time_service", "dop_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1034, Title: "Revenues from telephony services", AllowedParameters: []string{"voice_fix", "period_service", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1100, Title: "Revenues from combined Voice/SMS/GPRS services", AllowedParameters: []string{"mob_inet", "voice_mob", "sms", "period_service", "services_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1121, Title: "Revenues from IoT geoanalytics", AllowedParameters: []string{"iot", "period_service", "dop_service", "discount", "geo"}})
	groupService.CreateGroup(services.GroupView{ID: 1124, Title: "Revenues from IoT SMS", AllowedParameters: []string{"sms", "iot", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 1125, Title: "Revenues from IoT GPRS", AllowedParameters: []string{"mob_inet", "iot", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 1126, Title: "Revenues from voice + CSD IoT", AllowedParameters: []string{"voice_mob", "csd", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 3001, Title: "_FB. Telephony - subscription fee per number", AllowedParameters: []string{"voice_fix", "voice_ap", "dop_service", "discount", "ep_for_number"}})
	groupService.CreateGroup(services.GroupView{ID: 3002, Title: "_FB. Telephony - subscription fee per line", AllowedParameters: []string{"voice_fix", "one_time_service", "dop_service", "ep_for_line"}})
	groupService.CreateGroup(services.GroupView{ID: 3003, Title: "_FB. Telephony - setup fee for subscriber number", AllowedParameters: []string{"voice_fix", "period_service", "one_time_service", "discount", "one-time_fee_for_number"}})
	groupService.CreateGroup(services.GroupView{ID: 3009, Title: "_FB. Telephony - zonal outgoing calls to fixed operators", AllowedParameters: []string{"voice_fix", "fix_op", "voice_ap", "period_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3010, Title: "_FB. Telephony - zonal outgoing calls to SPC", AllowedParameters: []string{"voice_fix", "conc"}})
	groupService.CreateGroup(services.GroupView{ID: 3019, Title: "_FB. Telephony - long-distance services from own subscribers", AllowedParameters: []string{"voice_fix", "mg", "period_service", "one_time_service", "dop_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3020, Title: "_FB. Telephony - international services from own subscribers", AllowedParameters: []string{"voice_fix", "mn", "one_time_service", "dop_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3031, Title: "_FB. Telephony - additional services", AllowedParameters: []string{"voice_fix", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3309, Title: "FB_Television (CTV) - subscription fee per line", AllowedParameters: []string{"fix_ctv", "period_service", "dop_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3311, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (CTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_ctv", "voice_fix", "period_service", "one_time_service"}})
	groupService.CreateGroup(services.GroupView{ID: 3312, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (CTV) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ctv", "period_service", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3313, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + CTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ctv", "voice_fix", "period_service", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3314, Title: "FB_Sale of goods, works, and services - Individuals - Equipment rental for CTV service", AllowedParameters: []string{"fix_ctv", "period_service", "equipment rent"}})
	groupService.CreateGroup(services.GroupView{ID: 3320, Title: "FB_Sale of goods, works, and services - Individuals - Television (ICTV) - Subscription fee", AllowedParameters: []string{"fix_ictv", "period_service", "one_time_service", "dop_service"}})
	groupService.CreateGroup(services.GroupView{ID: 3321, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (ICTV) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ictv", "period_service", "one_time_service", "dop_service"}})
	groupService.CreateGroup(services.GroupView{ID: 3322, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (ICTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_ictv", "voice_fix", "period_service", "one_time_service", "discount"}})
	groupService.CreateGroup(services.GroupView{ID: 3323, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + ICTV + Telephony) - Subscription fee", AllowedParameters: []string{"fix_inet", "fix_ictv", "voice_fix", "period_service", "one_time_service"}})
}
