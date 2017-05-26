package ppot

import (
	"encoding/xml"
	"errors"
	"pgu_status/types"
	// "log"
)

// Success описывает часть XML успешного ответа
type Success struct {
	XMLName xml.Name `xml:"success"`

	UmmsID           string `xml:"ummsId"`
	ExternalSystemID string `xml:"externalSystemId"`
	ExternalCaseID   string `xml:"externalCaseId"`
}

// Error описывает часть XML успешного ответа
type Error struct {
	XMLName xml.Name `xml:"error"`

	ErrorMsg         string `xml:"errorMsg"`
	ExternalSystemID string `xml:"externalSystemId"`
	ExternalCaseID   string `xml:"externalCaseId"`
}

// ResponseSuccess описывает Полную XML успешного ответа
type ResponseSuccess struct {
	XMLName    xml.Name `xml:"response"`
	EntityType string   `xml:"entityType"`

	Success Success `xml:"success"`
	// UmmsID           string `xml:"success>ummsId"`
	// ExternalSystemID string `xml:"success>externalSystemId"`
	// ExternalCaseID   string `xml:"success>externalCaseId"`
}

// ResponseError описывает Полную XML успешного ответа
type ResponseError struct {
	XMLName    xml.Name `xml:"response"`
	EntityType string   `xml:"entityType"`

	Error Error `xml:"error"`
}

// ParseSuccess input xml into Response struct or error
func ParseSuccess(data []byte) (types.IPpotResultMsg, error) {
	r1 := ResponseSuccess{}
	err := xml.Unmarshal(data, &r1)

	if err == nil && r1.Success.XMLName.Local != "" && r1.Success.ExternalCaseID != "" {
		// log.Printf("[INFO] Parse OBJ: '%v'", r1.Success.ExternalCaseID)
		return Msg{
			externalSystemID: r1.Success.ExternalSystemID,
			externalCaseID:   r1.Success.ExternalCaseID,
			ummsID:           "",
			isError:          false,
			errorText:        "",
		}, nil
	}

	r2 := ResponseError{}
	err = xml.Unmarshal(data, &r2)

	if err == nil && r2.Error.XMLName.Local != "" && r2.Error.ExternalCaseID != "" {
		return Msg{
			externalSystemID: r2.Error.ExternalSystemID,
			externalCaseID:   r2.Error.ExternalCaseID,
			ummsID:           "",
			isError:          true,
			errorText:        r2.Error.ErrorMsg,
		}, nil
	}

	if err == nil {
		err = errors.New("Не удалось разобрать стандартный ответ из ППО Т")
	}
	// log.Printf("[ERROR] Parse error: '%v'", err)
	return nil, err
}
