package types

// IPpotResultMsg описание сообщения от ППОТ
type IPpotResultMsg interface {
	ExternalSystemID() string
	ExternalCaseID() string
	UmmsID() string
	IsError() bool
	ErrorText() string
}

// ISxMsg описывает сообщения в СК
type ISxMsg interface {
	ReasonCaseNumber() string
	ExtNumber() string
	ReasonServiceCode() string
}

// IPguStatusMsg описывает сообщение о статусе дела на ПГУ
type IPguStatusMsg interface {
	OrderID() string
	ReasonServiceCode() string
	RequestID() string
	TechStatus() string
	Comment() string
}

// IResultParser interface
type IResultParser interface {
	Parse(msg []byte) (IPpotResultMsg, error)
}

// ITaskFinder interface
type ITaskFinder interface {
	Find(string) (ISxMsg, error)
	Close()
}

////////////////// Implementing /////////

// PguStatusMsg содержит данные для отправки статуса на ПГУ через СК
type PguStatusMsg struct {
	orderID           string
	reasonServiceCode string
	requestID         string
	techStatus        string
	comment           string
}

// OrderID return order ID
func (msg PguStatusMsg) OrderID() string {
	return msg.orderID
}

// Comment return order ID
func (msg PguStatusMsg) Comment() string {
	return msg.comment
}

// ReasonServiceCode return order ID
func (msg PguStatusMsg) ReasonServiceCode() string {
	return msg.reasonServiceCode
}

// TechStatus return order ID
func (msg PguStatusMsg) TechStatus() string {
	return msg.requestID
}

// RequestID return order ID
func (msg PguStatusMsg) RequestID() string {
	return msg.requestID
}

// // IConfig предоставляет доступ к конфигу приложения
// type IConfig interface {
// 	GetString(string) string
// }

// IListener предостваляет доступ к слушателю очереди
type IListener interface {
	Start(IResultParser, ITaskFinder, ISxService)
}

// ISxService предостваляет доступ к сервису обновления дела на ПГУ
type ISxService interface {
	ChangePguCaseStatus(IPguStatusMsg) error
}

// MakePguStatusMsg create new PguStatusMsg
func MakePguStatusMsg(ppotMsg IPpotResultMsg, taskMsg ISxMsg) IPguStatusMsg {
	status := "3"
	comment := "Исполнено"
	if ppotMsg.IsError() {
		status = "15"
		comment = "Заявка требует дополнительной корректировки:\n" + ppotMsg.ErrorText()
	}

	return &PguStatusMsg{
		orderID:           taskMsg.ReasonCaseNumber(),
		reasonServiceCode: taskMsg.ReasonServiceCode(),
		requestID:         taskMsg.ExtNumber(),
		techStatus:        status,
		comment:           comment,
	}
}

// MakePguStatusMsgStub create msg for test
func MakePguStatusMsgStub() IPguStatusMsg {
	return &PguStatusMsg{
		orderID:           "175851555",
		reasonServiceCode: "10000022975",
		requestID:         "5c3908e1-7125-4b4c-8af6-64d79e425a16",
		techStatus:        "15",
		comment:           "Заявка требует дополнительной корректировки",
	}
}
