// Package actioninfo выводит общую информацию об активностях через интерфейс DataParser.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser разбирает строку с данными об активности и формирует по ней сводку.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит сводку по каждой строке dataset; ошибки логируются
// и не прерывают обработку.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("не удалось разобрать строку %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось сформировать информацию по строке %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
