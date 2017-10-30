package models

//https://github.com/hflabs/suggestions-csharp/blob/master/src/SuggestModel.cs

type SuggestionsRespond struct {
	Suggestions []SuggestionsAddress `json:"suggestions"`
}
type SuggestionsAddress struct {
	Value              string  `json:"value"`
	Unrestricted_value string  `json:"unrestricted_value"`
	Data               Address `json:"data"`
}

type (
	Address struct {
		Source               string  `json:"source"`                  // Исходный адрес одной строкой
		Result               string  `json:"result"`                  // Стандартизованный адрес одной строкой
		PostalCode           string  `json:"postal_code"`             // Индекс
		Country              string  `json:"country"`                 // Страна
		RegionFiasId         string  `json:"region_fias_id"`          // Код ФИАС региона
		RegionKladrId        string  `json:"region_kladr_id"`         // Код КЛАДР региона
		RegionWithType       string  `json:"region_with_type"`        // Регион с типом
		RegionType           string  `json:"region_type"`             // Тип региона (сокращенный)
		RegionTypeFull       string  `json:"region_type_full"`        // Тип региона
		Region               string  `json:"region"`                  // Регион
		AreaFiasId           string  `json:"area_fias_id"`            // Код ФИАС района в регионе
		AreaKladrId          string  `json:"area_kladr_id"`           // Код КЛАДР района в регионе
		AreaWithType         string  `json:"area_with_type"`          // Район в регионе с типом
		AreaType             string  `json:"area_type"`               // Тип района в регионе (сокращенный)
		AreaTypeFull         string  `json:"area_type_full"`          // Тип района в регионе
		Area                 string  `json:"area"`                    // Район в регионе
		CityFiasId           string  `json:"city_fias_id"`            // Код ФИАС города
		CityKladrId          string  `json:"city_kladr_id"`           // Код КЛАДР города
		CityWithType         string  `json:"city_with_type"`          // Город с типом
		CityType             string  `json:"city_type"`               // Тип города (сокращенный)
		CityTypeFull         string  `json:"city_type_full"`          // Тип города
		City                 string  `json:"city"`                    // Город
		CityArea             string  `json:"city_area"`               // Административный округ (только для Москвы)
		CityDistrictFiasId   string  `json:"city_district_fias_id"`   // Код ФИАС района города (заполняется, только если район есть в ФИАС)
		CityDistrictKladrId  string  `json:"city_district_kladr_id"`  // Код КЛАДР района города (не заполняется)
		CityDistrictWithType string  `json:"city_district_with_type"` // Район города с типом
		CityDistrictType     string  `json:"city_district_type"`      // Тип района города (сокращенный)
		CityDistrictTypeFull string  `json:"city_district_type_full"` // Тип района города
		CityDistrict         string  `json:"city_district"`           // Район города
		SettlementFiasId     string  `json:"settlement_fias_id"`      // Код ФИАС нас. пункта
		SettlementKladrId    string  `json:"settlement_kladr_id"`     // Код КЛАДР нас. пункта
		SettlementWithType   string  `json:"settlement_with_type"`    // Населенный пункт с типом
		SettlementType       string  `json:"settlement_type"`         // Тип населенного пункта (сокращенный)
		SettlementTypeFull   string  `json:"settlement_type_full"`    // Тип населенного пункта
		Settlement           string  `json:"settlement"`              // Населенный пункт
		StreetFiasId         string  `json:"street_fias_id"`          // Код ФИАС улицы
		StreetKladrId        string  `json:"street_kladr_id"`         // Код КЛАДР улицы
		StreetWithType       string  `json:"street_with_type"`        // Улица с типом
		StreetType           string  `json:"street_type"`             // Тип улицы (сокращенный)
		StreetTypeFull       string  `json:"street_type_full"`        // Тип улицы
		Street               string  `json:"street"`                  // Улица
		HouseFiasId          string  `json:"house_fias_id"`           // Код ФИАС дома
		HouseKladrId         string  `json:"house_kladr_id"`          // Код КЛАДР дома
		HouseType            string  `json:"house_type"`              // Тип дома (сокращенный)
		HouseTypeFull        string  `json:"house_type_full"`         // Тип дома
		House                string  `json:"house"`                   // Дом
		BlockType            string  `json:"block_type"`              // Тип корпуса/строения (сокращенный)
		BlockTypeFull        string  `json:"block_type_full"`         // Тип корпуса/строения
		Block                string  `json:"block"`                   // Корпус/строение
		FlatType             string  `json:"flat_type"`               // Тип квартиры (сокращенный)
		FlatTypeFull         string  `json:"flat_type_full"`          // Тип квартиры
		Flat                 string  `json:"flat"`                    // Квартира
		FlatArea             string  `json:"flat_area"`               // Площадь квартиры
		SquareMeterPrice     string  `json:"square_meter_price"`      // Рыночная стоимость м²
		FlatPrice            string  `json:"flat_price"`              // Рыночная стоимость квартиры
		PostalBox            string  `json:"postal_box"`              // Абонентский ящик
		FiasId               string  `json:"fias_id"`                 // Код ФИАС
		FiasLevel            string  `json:"fias_level"`              // Уровень детализации, до которого адрес найден в ФИАС
		KladrId              string  `json:"kladr_id"`                // Код КЛАДР
		CapitalMarker        string  `json:"capital_marker"`          // Статус центра
		Okato                string  `json:"okato"`                   // Код ОКАТО
		Oktmo                string  `json:"oktmo"`                   // Код ОКТМО
		TaxOffice            string  `json:"tax_office"`              // Код ИФНС для физических лиц
		Timezone             string  `json:"timezone"`                // Часовой пояс
		GeoLat               string  `json:"geo_lat"`                 // Координаты: широта
		GeoLon               string  `json:"geo_lon"`                 // Координаты: долгота
		BeltwayHit           string  `json:"beltway_hit"`             // Внутри кольцевой?
		BeltwayDistance      string  `json:"beltway_distance"`        // Расстояние от кольцевой в км.
		QualityCodeGeo       string  `json:"qc_geo"`                  // Код точности координат: 0 —	точные координаты	1 — ближайший дом	2 — улица	3 — населенный пункт	4 — город	5 — координаты не определены
		QualityCodeComplete  int     `json:"qc_complete"`             // Код полноты
		QualityCodeHouse     int     `json:"qc_house"`                // Код проверки дома
		QualityCode          int     `json:"qc"`                      // Код качества
		UnparsedParts        string  `json:"unparsed_parts"`          // Нераспознанная часть адреса. Для адреса
		qc                   string  `json:"qc"`
		Metro                []Metro `json:"metro"`
	}
	
	Metro struct {
		Name     string `json:"name"`
		Line     string `json:"line"`
		Distance string `json:"distance"`
	}
)
