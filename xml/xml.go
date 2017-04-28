package xml

import (
	"encoding/xml"
	"log"
)

// "encoding/xml"
// "fmt"

// TODO: Добавить разбор xml в котором пришла ошибка
//      + реализовать диспатчинг (когда без обишки xml и с ошибкой)

// Parse input raw-byte xml into Response struct or error
func Parse(raw []byte) (*Response, error) {
	succsess := &Success{}
	err := xml.Unmarshal(raw, succsess)
	if err != nil {
		log.Println("[ERROR]:", err)
		return nil, err
	}

	failure := &IntegrityError{}
	err = xml.Unmarshal(raw, failure)
	if err != nil {
		log.Println("[ERROR]:", err)
		return nil, err
	}
	// if len(failure.Body) == 0 && failure.Description == "" {
	// 	return nil, nil
	// }

	return NewResponse(succsess, failure), nil
}
