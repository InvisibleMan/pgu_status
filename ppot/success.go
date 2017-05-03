package ppot

import (
	"encoding/xml"
	"pgu_status/types"
)

// Response описывает Полную XML успешного ответа
type Response struct {
	XMLName    xml.Name `xml:"response"`
	EntityType string   `xml:"entityType"`

	UmmsID           string `xml:"success>ummsId"`
	ExternalSystemID string `xml:"success>externalSystemId"`
	ExternalCaseID   string `xml:"success>externalCaseId"`

	// SchemaVersion    string `xml:"schemaVersion,attr"`
	// Xmlns            string `xml:"xmlns,attr"`
	// Ns2              string `xml:"ns2,attr"`
}

// ParseSuccess input xml into Response struct or error
func ParseSuccess(data []byte) (types.IPpotResultMsg, error) {
	v := Response{}
	err := xml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return Msg{
		externalSystemID: v.ExternalSystemID,
		externalCaseID:   v.ExternalCaseID,
		ummsID:           v.UmmsID,
		isError:          false,
		errorText:        "",
	}, nil
}
