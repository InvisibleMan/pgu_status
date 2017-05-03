package sx

import (
	"encoding/xml"
	"io/ioutil"
	// "log"
	"errors"
	"net/http"
	"pgu_status/types"
	"strings"
)

// Service wsdl-сервис для изменения дела на ПГУ в СК
type Service struct {
	Endpoint string
}

// NewSXService create new instance SX
// sample url http://1.99.30.38:8080/
func NewSXService(endpoint string) types.ISxService {
	return &Service{endpoint}
}

// Success описывает успешный результат импорта
type Success struct {
	ExternalSystemID string `xml:"externalSystemId"`
	ExternalCaseID   string `xml:"externalCaseId"`
	UmmsID           string `xml:"ummsId"`
}

// Envelop описывает Полную XML успешного ответа
type Envelop struct {
	XMLName xml.Name `xml:"Envelope"`

	Body struct {
		Response struct {
			ErrorCode  string `xml:"errorCode"`
			Parameters []struct {
				Name  string `xml:"name"`
				Value string `xml:"value"`
			} `xml:"taskResult>parameters"`
		} `xml:"processOutgoingTaskResponse"`
	} `xml:"Body"`
}

// ChangePguCaseStatus update PGU Case
func (service Service) ChangePguCaseStatus(msg types.IPguStatusMsg) error {
	req := CreateRequestBody(msg)
	resp, err := http.Post(service.Endpoint, " text/xml;charset=UTF-8", strings.NewReader(req))

	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	res := string(b)
	return Parse(res)
}

// Parse result
func Parse(data string) error {
	v := Envelop{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		return err
	}

	if v.Body.Response.ErrorCode != "0" {
		return errors.New("Что-то пошло не так при POST-запросе в СК")
	}

	params := make(map[string]string)
	for _, p := range v.Body.Response.Parameters {
		params[p.Name] = p.Value
	}

	if params["RESULT"] == "0" {
		return nil
	} else if desc, ok := params["DESCRIPTION"]; ok {
		return errors.New(desc)
	}

	return errors.New("Что-то пошло не так при POST-запросе в СК")
}

// CreateRequestBody create simple request
func CreateRequestBody(msg types.IPguStatusMsg) string {
	elements := map[string]string{
		"{ORDER_ID}":        msg.OrderID(),
		"{COMMENT}":         msg.Comment(),
		"{TECH_STATE_CODE}": msg.TechStatus(),
		"{REQUEST_ID_REF}":  msg.RequestID(),
		"{SERVICE_CODE}":    msg.ReasonServiceCode(),
	}

	body := SoapTemplate
	for k, v := range elements {
		body = strings.Replace(body, k, v, -1)
	}

	return body
}

// SoapTemplate Шаблон SOAP-сообщения для изменения статуса дела на ПГУ
var SoapTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<soapenv:Envelope xmlns:bas="http://baseTypes.border.webservices.kernel.sx.fms.ru" xmlns:inc="http://www.w3.org/2004/08/xop/include" xmlns:out="http://outgoingRequests.webservices.kernel.sx.fms.ru" xmlns:rev="http://smev.gosuslugi.ru/rev111111" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xd="http://www.w3.org/2000/09/xmldsig#">
  <soapenv:Header/>
  <soapenv:Body>
    <out:processOutgoingTask>
      <bas:user>
        <bas:organization>FMS001001</bas:organization>
        <bas:person>
          <bas:id></bas:id>
          <bas:firstName/>
          <bas:secondName/>
          <bas:lastName>ППО Территория</bas:lastName>
        </bas:person>
      </bas:user>
      <bas:serviceCode>ПГУ05</bas:serviceCode>
      <bas:versionCode>001</bas:versionCode>
      <bas:parameters>
        <bas:name>ORDER_ID</bas:name>
        <bas:value>{ORDER_ID}</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>COMMENT</bas:name>
        <bas:value>{COMMENT}</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>AUTHOR</bas:name>
        <bas:value>ППО Т</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>TECH_STATE_CODE</bas:name>
        <bas:value>{TECH_STATE_CODE}</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>SEND_MESSAGE_ALLOWED</bas:name>
        <bas:value>false</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>CANCEL_ALLOWED</bas:name>
        <bas:value>false</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>ORIGIN_REQUEST_ID_REF</bas:name>
        <bas:value>{REQUEST_ID_REF}</bas:value>
      </bas:parameters>
      <bas:parameters>
        <bas:name>REQUEST_ID_REF</bas:name>
        <bas:value>{REQUEST_ID_REF}</bas:value>
      </bas:parameters>
      <bas:reasonServiceCode>{SERVICE_CODE}</bas:reasonServiceCode>
      <bas:reasonCaseNumber>{ORDER_ID}</bas:reasonCaseNumber>
    </out:processOutgoingTask>
  </soapenv:Body>
</soapenv:Envelope>
`
