package ppot

import (
	"errors"
	// "log"
	"pgu_status/types"
)

// Msg содержит атрибуты резульата загрузки в ППОТ
type Msg struct {
	externalSystemID string
	externalCaseID   string
	ummsID           string
	isError          bool
	errorText        string
}

// ExternalSystemID interface method
func (msg Msg) ExternalSystemID() string {
	return msg.externalSystemID
}

// ExternalCaseID interface method
func (msg Msg) ExternalCaseID() string {
	return msg.externalCaseID
}

// UmmsID interface method
func (msg Msg) UmmsID() string {
	return msg.ummsID
}

// IsError interface method
func (msg Msg) IsError() bool {
	return msg.isError
}

// ErrorText interface method
func (msg Msg) ErrorText() string {
	return msg.errorText
}

// ResultParser object for parse msg
type ResultParser struct{}

// NewResultParser create new IPpotResult
func NewResultParser() types.IResultParser {
	return &ResultParser{}
}

// Parse input xml into Response struct or error
func (parser ResultParser) Parse(data []byte) (types.IPpotResultMsg, error) {
	if msg, err := ParseSuccess(data); err == nil {
		return msg, nil
	}

	if msg, err := ParseIntegrityError(data); err == nil {
		return msg, nil
	}

	return nil, errors.New("Не удалось разобрать ответ из ППОТ")
}
