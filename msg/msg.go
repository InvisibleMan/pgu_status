package msg

import (
	"encoding/xml"
	"log"
)

// Msg осуществляет поиск атрибутов
// Дела с ЕПГУ
type Msg struct {
	ExternalSystemID string
	ExternalCaseID   string
	UmmsID           string
	IsError          bool
	Error            string
}

type Success struct {
	ExternalSystemID string `xml:"externalSystemId"`
	ExternalCaseID   string `xml:"externalCaseId"`
	UmmsID           string `xml:"ummsId"`
}

// Response from EPGU (??)
type Response struct {
	XMLName    xml.Name `xml:"response"`
	EntityType string   `xml:"entityType"`

	Success Success `xml:"success"`
}

// Parse xml-msg (form5) from RabbitMQ to Msg
func Parse(data string) (Msg, error) {
	v := Response{}
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Println("[ERROR]:", err)
		return Msg{}, err
	}

	return Msg{IsError: true}, nil
}
