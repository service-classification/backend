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

	groupService.CreateGroup(services.GroupView{ID: 1, Title: "Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 4, Title: "Long-distance communication"})
	groupService.CreateGroup(services.GroupView{ID: 5, Title: "International communication"})
	groupService.CreateGroup(services.GroupView{ID: 6, Title: "Intra-network roaming"})
	groupService.CreateGroup(services.GroupView{ID: 7, Title: "National roaming"})
	groupService.CreateGroup(services.GroupView{ID: 8, Title: "International roaming"})
	groupService.CreateGroup(services.GroupView{ID: 9, Title: "SMS"})
	groupService.CreateGroup(services.GroupView{ID: 10, Title: "Other additional services"})
	groupService.CreateGroup(services.GroupView{ID: 11, Title: "SMS content services"})
	groupService.CreateGroup(services.GroupView{ID: 12, Title: "Data transfer and telematics services"})
	groupService.CreateGroup(services.GroupView{ID: 15, Title: "Other subscriber services"})
	groupService.CreateGroup(services.GroupView{ID: 16, Title: "Goods and accessories"})
	groupService.CreateGroup(services.GroupView{ID: 22, Title: "Discounts"})
	groupService.CreateGroup(services.GroupView{ID: 29, Title: "GPRS revenues"})
	groupService.CreateGroup(services.GroupView{ID: 33, Title: "Incoming traffic modifiers in roaming"})
	groupService.CreateGroup(services.GroupView{ID: 90, Title: "SMS A2P"})
	groupService.CreateGroup(services.GroupView{ID: 91, Title: "Revenues from SMS mailings (gross)"})
	groupService.CreateGroup(services.GroupView{ID: 99, Title: "Undefined."})
	groupService.CreateGroup(services.GroupView{ID: 100, Title: "Per-minute fee for incoming calls"})
	groupService.CreateGroup(services.GroupView{ID: 101, Title: "Per-minute fee for intra-network zonal traffic"})
	groupService.CreateGroup(services.GroupView{ID: 102, Title: "Per-minute fee for zonal traffic to other operators"})
	groupService.CreateGroup(services.GroupView{ID: 103, Title: "Per-minute fee for long-distance intra-network traffic"})
	groupService.CreateGroup(services.GroupView{ID: 104, Title: "Per-minute fee for long-distance traffic to other operators"})
	groupService.CreateGroup(services.GroupView{ID: 105, Title: "Per-minute fee for international traffic"})
	groupService.CreateGroup(services.GroupView{ID: 1001, Title: "Subscription fee for voice services"})
	groupService.CreateGroup(services.GroupView{ID: 1004, Title: "SMS in international roaming"})
	groupService.CreateGroup(services.GroupView{ID: 1006, Title: "GPRS in national roaming"})
	groupService.CreateGroup(services.GroupView{ID: 1007, Title: "GPRS in international roaming"})
	groupService.CreateGroup(services.GroupView{ID: 1016, Title: "MMS"})
	groupService.CreateGroup(services.GroupView{ID: 1027, Title: "Revenues from activation of voice services valid for over 6 months"})
	groupService.CreateGroup(services.GroupView{ID: 1030, Title: "Revenues from voicemail"})
	groupService.CreateGroup(services.GroupView{ID: 1033, Title: "Revenues from Internet access services"})
	groupService.CreateGroup(services.GroupView{ID: 1034, Title: "Revenues from telephony services"})
	groupService.CreateGroup(services.GroupView{ID: 1100, Title: "Revenues from combined Voice/SMS/GPRS services"})
	groupService.CreateGroup(services.GroupView{ID: 1121, Title: "Revenues from IoT geoanalytics"})
	groupService.CreateGroup(services.GroupView{ID: 1124, Title: "Revenues from IoT SMS"})
	groupService.CreateGroup(services.GroupView{ID: 1125, Title: "Revenues from IoT GPRS"})
	groupService.CreateGroup(services.GroupView{ID: 1126, Title: "Revenues from voice + CSD IoT"})
	groupService.CreateGroup(services.GroupView{ID: 3001, Title: "_FB. Telephony - subscription fee per number"})
	groupService.CreateGroup(services.GroupView{ID: 3002, Title: "_FB. Telephony - subscription fee per line"})
	groupService.CreateGroup(services.GroupView{ID: 3003, Title: "_FB. Telephony - setup fee for subscriber number"})
	groupService.CreateGroup(services.GroupView{ID: 3009, Title: "_FB. Telephony - zonal outgoing calls to fixed operators"})
	groupService.CreateGroup(services.GroupView{ID: 3010, Title: "_FB. Telephony - zonal outgoing calls to SPC"})
	groupService.CreateGroup(services.GroupView{ID: 3019, Title: "_FB. Telephony - long-distance services from own subscribers"})
	groupService.CreateGroup(services.GroupView{ID: 3020, Title: "_FB. Telephony - international services from own subscribers"})
	groupService.CreateGroup(services.GroupView{ID: 3031, Title: "_FB. Telephony - additional services"})
	groupService.CreateGroup(services.GroupView{ID: 3309, Title: "FB_Television (CTV) - subscription fee per line"})
	groupService.CreateGroup(services.GroupView{ID: 3311, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (CTV + Telephony) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3312, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (CTV) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3313, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + CTV + Telephony) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3314, Title: "FB_Sale of goods, works, and services - Individuals - Equipment rental for CTV service"})
	groupService.CreateGroup(services.GroupView{ID: 3320, Title: "FB_Sale of goods, works, and services - Individuals - Television (ICTV) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3321, Title: "FB_Sale of goods, works, and services - Individuals - Internet Access - Double Play (ICTV) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3322, Title: "FB_Sale of goods, works, and services - Individuals - Double Play (ICTV + Telephony) - Subscription fee"})
	groupService.CreateGroup(services.GroupView{ID: 3323, Title: "FB_Sale of goods, works, and services - Individuals - Triple Play (Internet + ICTV + Telephony) - Subscription fee"})
}
