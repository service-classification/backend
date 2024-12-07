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

	parameterService.CreateParameter(services.NewParameter{Title: "Мобильный интернет", ID: "mob_inet"})
	parameterService.CreateParameter(services.NewParameter{Title: "Фиксированный интернет", ID: "fix_inet"})
	parameterService.CreateParameter(services.NewParameter{Title: "Телевидение домашнее по технологии ЦТВ (цифровое ТВ)", ID: "fix_ctv"})
	parameterService.CreateParameter(services.NewParameter{Title: "Телевидение домашнее по технологии ИЦТВ (интерактивное цифровое ТВ)", ID: "fix_ictv"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги голосовой мобильной связи", ID: "voice_mob"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги голосовой фиксированной связи", ID: "voice_fix"})
	parameterService.CreateParameter(services.NewParameter{Title: "Короткие текстовые сообщения (Short Message Service)", ID: "sms"})
	parameterService.CreateParameter(services.NewParameter{Title: "Circuit Switched Data (CSD)", ID: "csd"})
	parameterService.CreateParameter(services.NewParameter{Title: "Интернет вещей (IoT)", ID: "iot"})
	parameterService.CreateParameter(services.NewParameter{Title: "Multimedia Messaging Service (MMS)", ID: "mms"})
	parameterService.CreateParameter(services.NewParameter{Title: "Оказание услуг мобильной связи в сетях других операторов", ID: "roaming"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги междугородней связи", ID: "mg"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги международной связи", ID: "mn"})
	parameterService.CreateParameter(services.NewParameter{Title: "Направление звонка/sms внутри сети", ID: "mts"})
	parameterService.CreateParameter(services.NewParameter{Title: "Направление звонка/sms на сети других операторов (сети конкурентов)", ID: "conc"})
	parameterService.CreateParameter(services.NewParameter{Title: "Направление звонка/sms на сети фиксированных операторов", ID: "fix_op"})
	parameterService.CreateParameter(services.NewParameter{Title: "Роуминг внутри сети (поездки по РФ)", ID: "vsr_roam"})
	parameterService.CreateParameter(services.NewParameter{Title: "Роуминг в сетях других операторов (поездки по РФ)", ID: "national_roam"})
	parameterService.CreateParameter(services.NewParameter{Title: "Роуминг в сетях других операторов (поездки за границу)", ID: "mn_roam"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги голосовой мобильной связи, абон. плата", ID: "voice_ap"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги голосовой мобильной связи, плата за подключение", ID: "voice_fee"})
	parameterService.CreateParameter(services.NewParameter{Title: "Периодические услуги", ID: "period_service"})
	parameterService.CreateParameter(services.NewParameter{Title: "Разовые услуги (услуги сопровождащие добавление/удаление периодических услуг в биллинге и искоючительно разовые услуги, такие как смена владельца, смена номера и т.д.)", ID: "one_time_service"})
	parameterService.CreateParameter(services.NewParameter{Title: "Дополнительные услуги", ID: "dop_service"})
	parameterService.CreateParameter(services.NewParameter{Title: "Контентные информационно-развлекательные услуги", ID: "content"})
	parameterService.CreateParameter(services.NewParameter{Title: "Доп. услуги по абон. обслуживанию", ID: "services_service"})
	parameterService.CreateParameter(services.NewParameter{Title: "Продажа карт экспресс-оплаты", ID: "keo_sale"})
	parameterService.CreateParameter(services.NewParameter{Title: "Скидки", ID: "discount"})
	parameterService.CreateParameter(services.NewParameter{Title: "Только входящие звонки", ID: "only_inbound"})
	parameterService.CreateParameter(services.NewParameter{Title: "A2P (Application-to-Person) SMS Messaging", ID: "sms_a2p"})
	parameterService.CreateParameter(services.NewParameter{Title: "Тоже самое что и A2P (Application-to-Person) SMS Messaging, только другой вариант аллокации доходов с партнером", ID: "sms_gross"})
	parameterService.CreateParameter(services.NewParameter{Title: "Скоринг", ID: "skoring"})
	parameterService.CreateParameter(services.NewParameter{Title: "Прочие дополнительные услуги", ID: "other_service"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги голосовой почты", ID: "voice_mail"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги геоаналитики (точки скопления абонентов и т.д.)", ID: "geo"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги фикс.телефонии, ежемесячная плата за номер", ID: "ep_for_number"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги фикс.телефонии, разовая плата за линию", ID: "ep_for_line"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги фикс.телефонии, разовая плата за номер", ID: "one_time_fee_for_number"})
	parameterService.CreateParameter(services.NewParameter{Title: "Услуги фикс. связи, аренда оборудования", ID: "equipment_rent"})
	parameterService.CreateParameter(services.NewParameter{Title: "Дополнительные пакеты", ID: "add_package"})

	groupService.CreateGroup(services.NewGroup{ID: 1, Title: "Абонентская плата", AllowedParameters: []string{"voice_ap", "voice_fee", "ep_for_number", "ep_for_line", "one_time_fee_for_number"}})
	groupService.CreateGroup(services.NewGroup{ID: 4, Title: "Междугородная связь", AllowedParameters: []string{"mg"}})
	groupService.CreateGroup(services.NewGroup{ID: 5, Title: "Международная связь", AllowedParameters: []string{"mn"}})
	groupService.CreateGroup(services.NewGroup{ID: 6, Title: "Роуминг внутрисетевой", AllowedParameters: []string{"vsr_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 7, Title: "Роуминг национальный", AllowedParameters: []string{"national_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 8, Title: "Роуминг международный", AllowedParameters: []string{"mn_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 9, Title: "SMS", AllowedParameters: []string{"sms", "sms_a2p", "sms_gross"}})
	groupService.CreateGroup(services.NewGroup{ID: 10, Title: "Прочие дополнительные услуги", AllowedParameters: []string{"other_service"}})
	groupService.CreateGroup(services.NewGroup{ID: 11, Title: "Контентные услуги SMS", AllowedParameters: []string{"content"}})
	groupService.CreateGroup(services.NewGroup{ID: 12, Title: "Передача данных и телематические услуги", AllowedParameters: []string{"mob_inet", "fix_inet", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 15, Title: "Прочие услуги по абон.обслуживанию", AllowedParameters: []string{"services_service"}})
	groupService.CreateGroup(services.NewGroup{ID: 16, Title: "Товары и аксессуары", AllowedParameters: []string{"equipment_rent"}})
	groupService.CreateGroup(services.NewGroup{ID: 22, Title: "Скидки", AllowedParameters: []string{"discount"}})
	groupService.CreateGroup(services.NewGroup{ID: 29, Title: "Доходы от GPRS", AllowedParameters: []string{"mob_inet", "fix_inet", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 33, Title: "Модификаторы входящего трафика в роуминге", AllowedParameters: []string{"roaming"}})
	groupService.CreateGroup(services.NewGroup{ID: 90, Title: "SMS A2P", AllowedParameters: []string{"sms_a2p"}})
	groupService.CreateGroup(services.NewGroup{ID: 91, Title: "Доходы от SMS рассылок (gross)", AllowedParameters: []string{"sms_gross"}})
	groupService.CreateGroup(services.NewGroup{ID: 99, Title: "Не определено.", AllowedParameters: []string{"skoring"}})
	groupService.CreateGroup(services.NewGroup{ID: 100, Title: "Повременная плата за входящий вызов", AllowedParameters: []string{"only_inbound"}})
	groupService.CreateGroup(services.NewGroup{ID: 101, Title: "Повременная плата за зоновый внутрисетевой трафик", AllowedParameters: []string{"mts"}})
	groupService.CreateGroup(services.NewGroup{ID: 102, Title: "Повременная плата за зоновый трафик на др. операторов", AllowedParameters: []string{"conc"}})
	groupService.CreateGroup(services.NewGroup{ID: 103, Title: "Повременная плата за междугородний внутрисетевой трафик", AllowedParameters: []string{"mg"}})
	groupService.CreateGroup(services.NewGroup{ID: 104, Title: "Повременная плата за междугородний трафик на др. операторов", AllowedParameters: []string{"conc"}})
	groupService.CreateGroup(services.NewGroup{ID: 105, Title: "Повременная плата за международный трафик", AllowedParameters: []string{"mn"}})
	groupService.CreateGroup(services.NewGroup{ID: 1001, Title: "Абонентская плата за голосовые услуги", AllowedParameters: []string{"voice_ap", "voice_fee"}})
	groupService.CreateGroup(services.NewGroup{ID: 1004, Title: "SMS в роуминге международный", AllowedParameters: []string{"mn_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 1006, Title: "GPRS в роуминге национальный", AllowedParameters: []string{"national_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 1007, Title: "GPRS в роуминге международный", AllowedParameters: []string{"mn_roam"}})
	groupService.CreateGroup(services.NewGroup{ID: 1016, Title: "MMS", AllowedParameters: []string{"mms"}})
	groupService.CreateGroup(services.NewGroup{ID: 1027, Title: "Доходы от активации голосовых услуг со сроком действия свыше 6 мес", AllowedParameters: []string{"voice_ap"}})
	groupService.CreateGroup(services.NewGroup{ID: 1030, Title: "Доходы от голосовой почты", AllowedParameters: []string{"voice_mail"}})
	groupService.CreateGroup(services.NewGroup{ID: 1033, Title: "Доходы от предоставления услуг доступа в Internet", AllowedParameters: []string{"mob_inet", "fix_inet", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 1034, Title: "Доходы от предоставления услуг телефонии", AllowedParameters: []string{"voice_mob", "voice_fix"}})
	groupService.CreateGroup(services.NewGroup{ID: 1100, Title: "Доход от совместных услуг Голос/SMS/GPRS", AllowedParameters: []string{"voice_mob", "voice_fix", "sms", "mob_inet", "fix_inet", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 1121, Title: "Доходы от геоаналитики IoT", AllowedParameters: []string{"geo"}})
	groupService.CreateGroup(services.NewGroup{ID: 1124, Title: "Доходы от SMS IoT", AllowedParameters: []string{"sms", "sms_a2p", "sms_gross"}})
	groupService.CreateGroup(services.NewGroup{ID: 1125, Title: "Доходы от GPRS IoT", AllowedParameters: []string{"mob_inet", "fix_inet", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 1126, Title: "Доходы от голос + CSD IoT", AllowedParameters: []string{"voice_mob", "voice_fix", "csd", "iot"}})
	groupService.CreateGroup(services.NewGroup{ID: 3001, Title: "_ФБ. Телефония - абонентская плата за номер", AllowedParameters: []string{"ep_for_number"}})
	groupService.CreateGroup(services.NewGroup{ID: 3002, Title: "_ФБ. Телефония - абонентская плата за линию", AllowedParameters: []string{"ep_for_line"}})
	groupService.CreateGroup(services.NewGroup{ID: 3003, Title: "_ФБ. Телефония - установочная плата за абонентский номер", AllowedParameters: []string{"one_time_fee_for_number"}})
	groupService.CreateGroup(services.NewGroup{ID: 3009, Title: "_ФБ. Телефония - зоновые исх. вызовы на фикс. операторов", AllowedParameters: []string{"fix_op"}})
	groupService.CreateGroup(services.NewGroup{ID: 3010, Title: "_ФБ. Телефония - зоновые исх. вызовы на СПС", AllowedParameters: []string{"conc"}})
	groupService.CreateGroup(services.NewGroup{ID: 3019, Title: "_ФБ. Телефония - услуги МГ связи от абон. собств. сети", AllowedParameters: []string{"mg"}})
	groupService.CreateGroup(services.NewGroup{ID: 3020, Title: "_ФБ. Телефония - услуги МН связи от абон. собств. сети", AllowedParameters: []string{"mn"}})
	groupService.CreateGroup(services.NewGroup{ID: 3031, Title: "_ФБ. Телефония - дополнительные услуги", AllowedParameters: []string{"dop_service"}})
	groupService.CreateGroup(services.NewGroup{ID: 3309, Title: "ФБ_Телевидение (ЦТВ) - Абонентская плата за линию", AllowedParameters: []string{"ep_for_line"}})
	groupService.CreateGroup(services.NewGroup{ID: 3311, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Double Play (ЦТВ + Телефония) - Абонентская плата", AllowedParameters: []string{"fix_ctv", "voice_fix"}})
	groupService.CreateGroup(services.NewGroup{ID: 3312, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Доступ в Интернет - DoublePlay (ЦТВ) - Абонентская плата", AllowedParameters: []string{"fix_inet"}})
	groupService.CreateGroup(services.NewGroup{ID: 3313, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Triple Play (Интернет + ЦТВ + Телефония) - Абонентская плата", AllowedParameters: []string{"mob_inet", "fix_ctv", "voice_fix"}})
	groupService.CreateGroup(services.NewGroup{ID: 3314, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Аренда оборудования по услуге ЦТВ", AllowedParameters: []string{"equipment_rent"}})
	groupService.CreateGroup(services.NewGroup{ID: 3320, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Телеведение (ИЦТВ) - Абонентская плата", AllowedParameters: []string{"fix_ictv"}})
	groupService.CreateGroup(services.NewGroup{ID: 3321, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Доступ в Интернет - Double Play (ИЦТВ) - Абонентская плата", AllowedParameters: []string{"fix_inet"}})
	groupService.CreateGroup(services.NewGroup{ID: 3322, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Double Play (ИЦТВ + Телефония) - Абонентская плата", AllowedParameters: []string{"fix_ictv", "voice_fix"}})
	groupService.CreateGroup(services.NewGroup{ID: 3323, Title: "ФБ_Реализация товаров, работ и услуг - Физ. лица - Triple Play (Интернет + ИЦТВ + Телефония) - Абонентская плата", AllowedParameters: []string{"mob_inet", "fix_ictv", "voice_fix"}})
}
