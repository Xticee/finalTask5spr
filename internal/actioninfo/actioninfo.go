package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		return
	}

	for _, val := range dataset {
		err := dp.Parse(val)

		if err != nil {
			log.Println(err)
		}

		str, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
		}

		fmt.Println(str)
	}
}
