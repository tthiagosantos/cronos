package dto

type FlowState struct {
	FlowStateId      int         `json:"flowStateId"`
	FlowStateName    string      `json:"flowStateName"`
	FlowStateStartAt string      `json:"flowStateStartAt,omitempty"`
	FlowStateEndAt   string      `json:"flowStateEndAt,omitempty"`
	QuizQuantity     interface{} `json:"quizQuantity"`
	Quizzes          interface{} `json:"quizzes"`
}

type Ticket struct {
	TicketId   int         `json:"ticketId"`
	Code       string      `json:"code"`
	FlowStates []FlowState `json:"flowStates"`
}

type FlowStateOutput struct {
	Errors  []string `json:"errors"`
	Error   string   `json:"error"`
	Data    []Ticket `json:"data"`
	IsValid bool     `json:"isValid"`
}
