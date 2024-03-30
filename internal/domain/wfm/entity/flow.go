package entity

type Flow struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	FlowState   []FlowState `json:"flowState"`
}

type FlowState struct {
	FlowId                    int    `json:"flowId"`
	FlowStateId               int    `json:"flowStateId"`
	FlowStateName             string `json:"flowStateName"`
	FlowStateDescription      string `json:"flowStateDescription"`
	FlowOrder                 int    `json:"flowOrder"`
	MacroFlowStateId          int    `json:"macroFlowStateId"`
	MacroFlowStateDescription string `json:"macroFlowStateDescription"`
}

func NewFlow(flowState FlowState) *FlowState {
	return &FlowState{
		FlowId:                    flowState.FlowId,
		FlowStateId:               flowState.FlowStateId,
		FlowStateName:             flowState.FlowStateName,
		FlowStateDescription:      flowState.FlowStateDescription,
		FlowOrder:                 flowState.FlowOrder,
		MacroFlowStateId:          flowState.MacroFlowStateId,
		MacroFlowStateDescription: flowState.MacroFlowStateDescription,
	}
}
