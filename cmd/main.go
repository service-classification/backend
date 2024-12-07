package main

import (
	"backend/config"
	"backend/docs"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/router"
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
	migrate(db)

	serviceRepo := repositories.NewServiceRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	paramRepo := repositories.NewParameterRepository(db)
	handler := handlers.NewHandler(serviceRepo, groupRepo, paramRepo)

	docs.SwaggerInfo.Host = os.Getenv("PUBLIC_HOST") + ":8080"
	docs.SwaggerInfo.Description = "This is a backend server."

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

	db.Create(&models.Group{ID: 1, Title: "Абонентская плата"})
	db.Create(&models.Group{ID: 4, Title: "Междугородная связь"})
	db.Create(&models.Group{ID: 5, Title: "Международная связь"})
	db.Create(&models.Group{ID: 6, Title: "Роуминг внутрисетевой"})
	db.Create(&models.Group{ID: 7, Title: "Роуминг национальный"})
	db.Create(&models.Group{ID: 8, Title: "Роуминг международный"})
	db.Create(&models.Group{ID: 9, Title: "SMS"})
	db.Create(&models.Group{ID: 10, Title: "Прочие дополнительные услуги"})
	db.Create(&models.Group{ID: 11, Title: "Контентные услуги SMS"})
	db.Create(&models.Group{ID: 12, Title: "Передача данных и телематические услуги"})
	db.Create(&models.Group{ID: 15, Title: "Прочие услуги по абон.обслуживанию"})
	db.Create(&models.Group{ID: 16, Title: "Товары и аксессуары"})
	db.Create(&models.Group{ID: 22, Title: "Скидки"})
	db.Create(&models.Group{ID: 29, Title: "Доходы от GPRS"})
	db.Create(&models.Group{ID: 33, Title: "Модификаторы входящего трафика в роуминге"})
	db.Create(&models.Group{ID: 90, Title: "SMS A2P"})
	db.Create(&models.Group{ID: 91, Title: "Доходы от SMS рассылок (gross)"})
	db.Create(&models.Group{ID: 99, Title: "Не определено."})
	db.Create(&models.Group{ID: 100, Title: "Повременная плата за входящий вызов"})
	db.Create(&models.Group{ID: 101, Title: "Повременная плата за зоновый внутрисетевой трафик"})
	db.Create(&models.Group{ID: 102, Title: "Повременная плата за зоновый трафик на др. операторов"})
	db.Create(&models.Group{ID: 103, Title: "Повременная плата за междугородний внутрисетевой трафик"})
	db.Create(&models.Group{ID: 104, Title: "Повременная плата за междугородний трафик на др. операторов"})
	db.Create(&models.Group{ID: 105, Title: "Повременная плата за международный трафик"})
	db.Create(&models.Group{ID: 1001, Title: "Абонентская плата за голосовые услуги"})
	db.Create(&models.Group{ID: 1004, Title: "SMS в роуминге международный"})
	db.Create(&models.Group{ID: 1006, Title: "GPRS в роуминге национальный"})
	db.Create(&models.Group{ID: 1007, Title: "GPRS в роуминге международный"})
	db.Create(&models.Group{ID: 1016, Title: "MMS"})
	db.Create(&models.Group{ID: 1027, Title: "Доходы от активации голосовых услуг со сроком действия свыше 6 мес"})
	db.Create(&models.Group{ID: 1030, Title: "Доходы от голосовой почты"})
	db.Create(&models.Group{ID: 1033, Title: "Доходы от предоставления услуг доступа в Internet"})
	db.Create(&models.Group{ID: 1034, Title: "Доходы от предоставления услуг телефонии"})
	db.Create(&models.Group{ID: 1100, Title: "Доход от совместных услуг Голос/SMS/GPRS"})
	db.Create(&models.Group{ID: 1121, Title: "Доходы от геоаналитики IoT"})
	db.Create(&models.Group{ID: 1124, Title: "Доходы от SMS IoT"})
	db.Create(&models.Group{ID: 1125, Title: "Доходы от GPRS IoT"})
	db.Create(&models.Group{ID: 1126, Title: "Доходы от голос + CSD IoT"})
	db.Create(&models.Group{ID: 3001, Title: "_ФБ. Телефония - абонентская плата за номер"})
	db.Create(&models.Group{ID: 3002, Title: "_ФБ. Телефония - абонентская плата за линию"})
	db.Create(&models.Group{ID: 3003, Title: "_ФБ. Телефония - установочная плата за абонентский номер"})
	db.Create(&models.Group{ID: 3009, Title: "_ФБ. Телефония - зоновые исх. вызовы на фикс. операторов"})
	db.Create(&models.Group{ID: 3010, Title: "_ФБ. Телефония - зоновые исх. вызовы на СПС"})
	db.Create(&models.Group{ID: 3019, Title: "_ФБ. Телефония - услуги МГ связи от абон. собств. сети"})
	db.Create(&models.Group{ID: 3020, Title: "_ФБ. Телефония - услуги МН связи от абон. собств. сети"})
	db.Create(&models.Group{ID: 3031, Title: "_ФБ. Телефония - дополнительные услуги"})
	db.Create(&models.Group{ID: 3309, Title: "ФБ_Телевидение (ЦТВ) - Абонентская плата за линию"})
	db.Create(&models.Group{ID: 3311, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Double Play (ЦТВ + Телефония) - Абонентская плата"})
	db.Create(&models.Group{ID: 3312, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Доступ в Интернет - DoublePlay (ЦТВ) - Абонентская плата"})
	db.Create(&models.Group{ID: 3313, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Triple Play (Интернет + ЦТВ + Телефония) - Абонентская плата"})
	db.Create(&models.Group{ID: 3314, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Аренда оборудования по услуге ЦТВ"})
	db.Create(&models.Group{ID: 3320, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Телеведение (ИЦТВ) - Абонентская плата"})
	db.Create(&models.Group{ID: 3321, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Доступ в Интернет - Double Play (ИЦТВ) - Абонентская плата"})
	db.Create(&models.Group{ID: 3322, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Double Play (ИЦТВ + Телефония) - Абонентская плата"})
	db.Create(&models.Group{ID: 3323, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Triple Play (Интернет + ИЦТВ + Телефония) - Абонентская плата"})
}
