package actioninfo

import "log"

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info парсит и логирует информацию об активностях из dataset через DataParser.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("ошибка парсинга %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации для %q: %v", data, err)
			continue
		}

		log.Println(info)
	}
}
