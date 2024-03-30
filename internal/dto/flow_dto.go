package dto

type FlowStates struct {
	FlowId                    int    `json:"flowId"`
	FlowStateId               int    `json:"flowStateId"`
	FlowStateName             string `json:"flowStateName"`
	FlowStateDescription      string `json:"flowStateDescription"`
	FlowOrder                 int    `json:"flowOrder"`
	MacroFlowStateId          int    `json:"macroFlowStateId"`
	MacroFlowStateDescription string `json:"macroFlowStateDescription"`
}

type Flow struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	FlowState   []FlowStates `json:"flowState"`
}

type FlowResponse struct {
	Errors  []string `json:"errors"`
	Error   string   `json:"error"`
	Data    []Flow   `json:"data"`
	IsValid bool     `json:"isValid"`
}
