package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/router"
	"fmt"
	"log"

	_ "backend/docs" // This line is necessary for go-swagger to find docs

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			MyApp API
//	@version		1.0
//	@description	This is a sample server.

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
	err = db.AutoMigrate(&models.Service{}, &models.Group{}, &models.ClassifiedService{}, &models.Parameter{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	migrate(db)

	serviceRepo := repositories.NewServiceRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	csRepo := repositories.NewClassifiedServiceRepository(db)
	paramRepo := repositories.NewParameterRepository(db)
	handler := handlers.NewHandler(serviceRepo, groupRepo, csRepo, paramRepo)

	r := router.SetupRouter(handler)

	log.Fatal(r.Run(":8080"))
}

func migrate(db *gorm.DB) {
	db.Create(&models.Parameter{Title: "Мобильный интернет", Code: "mob_inet"})
	db.Create(&models.Parameter{Title: "Фиксированный интернет", Code: "fix_inet"})
	db.Create(&models.Parameter{Title: "Телевидение домашнее по технологии ЦТВ (цифровое ТВ)", Code: "fix_ctv"})
	db.Create(&models.Parameter{Title: "Телевидение домашнее по технологии ИЦТВ (интерактивное цифровое ТВ)", Code: "fix_ictv"})
	db.Create(&models.Parameter{Title: "Услуги голосовой мобильной связи", Code: "voice_mob"})
	db.Create(&models.Parameter{Title: "Услуги голосовой фиксированной связи", Code: "voice_fix"})
	db.Create(&models.Parameter{Title: "Короткие текстовые сообщения (Short Message Service)", Code: "sms"})
	db.Create(&models.Parameter{Title: "Circuit Switched Data (CSD)", Code: "csd"})
	db.Create(&models.Parameter{Title: "Интернет вещей (IoT)", Code: "iot"})
	db.Create(&models.Parameter{Title: "Multimedia Messaging Service (MMS)", Code: "mms"})
	db.Create(&models.Parameter{Title: "Оказание услуг мобильной связи в сетях других операторов", Code: "roaming"})
	db.Create(&models.Parameter{Title: "Услуги междугородней связи", Code: "mg"})
	db.Create(&models.Parameter{Title: "Услуги международной связи", Code: "mn"})
	db.Create(&models.Parameter{Title: "Направление звонка/sms внутри сети", Code: "mts"})
	db.Create(&models.Parameter{Title: "Направление звонка/sms на сети других операторов (сети конкурентов)", Code: "conc"})
	db.Create(&models.Parameter{Title: "Направление звонка/sms на сети фиксированных операторов", Code: "fix_op"})
	db.Create(&models.Parameter{Title: "Роуминг внутри сети (поездки по РФ)", Code: "vsr_roam"})
	db.Create(&models.Parameter{Title: "Роуминг в сетях других операторов (поездки по РФ)", Code: "national_roam"})
	db.Create(&models.Parameter{Title: "Роуминг в сетях других операторов (поездки за границу)", Code: "mn_roam"})
	db.Create(&models.Parameter{Title: "Услуги голосовой мобильной связи, абон. плата", Code: "voice_ap"})
	db.Create(&models.Parameter{Title: "Услуги голосовой мобильной связи, плата за подключение", Code: "voice_fee"})
	db.Create(&models.Parameter{Title: "Периодические услуги", Code: "period_service"})
	db.Create(&models.Parameter{Title: "Разовые услуги (услуги сопровождащие добавление/удаление периодических услуг в биллинге и искоючительно разовые услуги, такие как смена владельца, смена номера и т.д.)", Code: "one_time_service"})
	db.Create(&models.Parameter{Title: "Дополнительные услуги", Code: "dop_service"})
	db.Create(&models.Parameter{Title: "Контентные информационно-развлекательные услуги", Code: "content"})
	db.Create(&models.Parameter{Title: "Доп. услуги по абон. обслуживанию", Code: "services_service"})
	db.Create(&models.Parameter{Title: "Продажа карт экспресс-оплаты", Code: "keo_sale"})
	db.Create(&models.Parameter{Title: "Скидки", Code: "discount"})
	db.Create(&models.Parameter{Title: "Только входящие звонки", Code: "only_inbound"})
	db.Create(&models.Parameter{Title: "A2P (Application-to-Person) SMS Messaging", Code: "sms_a2p"})
	db.Create(&models.Parameter{Title: "Тоже самое что и A2P (Application-to-Person) SMS Messaging, только другой вариант аллокации доходов с партнером", Code: "sms_gross"})
	db.Create(&models.Parameter{Title: "Скоринг", Code: "skoring"})
	db.Create(&models.Parameter{Title: "Прочие дополнительные услуги", Code: "other_service"})
	db.Create(&models.Parameter{Title: "Услуги голосовой почты", Code: "voice_mail"})
	db.Create(&models.Parameter{Title: "Услуги геоаналитики (точки скопления абонентов и т.д.)", Code: "geo"})
	db.Create(&models.Parameter{Title: "Услуги фикс.телефонии, ежемесячная плата за номер", Code: "ep_for_number"})
	db.Create(&models.Parameter{Title: "Услуги фикс.телефонии, разовая плата за линию", Code: "ep_for_line"})
	db.Create(&models.Parameter{Title: "Услуги фикс.телефонии, разовая плата за номер", Code: "one_time_fee_for_number"})
	db.Create(&models.Parameter{Title: "Услуги фикс. связи, аренда оборудования", Code: "equipment_rent"})
	db.Create(&models.Parameter{Title: "Дополнительные пакеты", Code: "add_package"})
}
