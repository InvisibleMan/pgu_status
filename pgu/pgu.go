package pgu

// import (
// "encoding/xml"
// "fmt"
// )

// CaseFinder осуществляет поиск атрибутов
// Дела с ЕПГУ
type CaseFinder struct {
}

// Case содержит информацию о Заявке с ПГУ
type Case struct {
	Number     string
	MessageID  string
	StatusCode string
	Comment    string
}

// Find deal in DB
func (finder CaseFinder) Find() (deal Case, err error) {
	return Case{}, nil
}
