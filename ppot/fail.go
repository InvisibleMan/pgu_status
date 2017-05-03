package ppot

import (
	"encoding/xml"
	"html"
	"pgu_status/types"
)

// IntegrityError representation for Message
type IntegrityError struct {
	Xmlns         string `xml:"xmlns,attr"`
	Description   string `xml:"description"`
	Body          string `xml:"body"`
	SchemaVersion string `xml:"schemaVersion,attr"`
}

// Form5 representation
type Form5 struct {
	UmmsID           string `xml:"uid"`
	ExternalSystemID string `xml:"supplierInfo"`
	ExternalCaseID   string `xml:"number"`
}

// ParseIntegrityError into Response struct
func ParseIntegrityError(data string) (types.IPpotResultMsg, error) {
	v := IntegrityError{}
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		return nil, err
	}

	f5 := Form5{}
	err = xml.Unmarshal([]byte(html.UnescapeString(v.Body)), &f5)
	if err != nil {
		return nil, err
	}

	return Msg{
		externalSystemID: f5.ExternalSystemID,
		externalCaseID:   f5.ExternalCaseID,
		ummsID:           f5.UmmsID,
		isError:          true,
		errorText:        v.Description,
	}, nil
}
