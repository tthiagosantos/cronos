package entity

import (
	"cronos/internal/dto"
	"cronos/pkg/entity"
)

type Contact struct {
	RequesterContactID    int    `json:"requesterContactId"`
	RequesterContactName  string `json:"requesterContactName"`
	RequesterContactEmail string `json:"requesterContactEmail"`
	RequesterContactPhone string `json:"requesterContactPhone"`
}

type Ticket struct {
	ID                         int             `json:"id"`
	ShortDescription           string          `json:"shortDescription"`
	Description                interface{}     `json:"description"`
	Code                       string          `json:"code"`
	ExternalCode               string          `json:"externalCode"`
	ParentTicketCode           interface{}     `json:"parentTicketCode"`
	RequesterName              interface{}     `json:"requesterName"`
	RequesterEmail             string          `json:"requesterEmail"`
	ActivityDescription        string          `json:"activityDescription"`
	EventReport                string          `json:"eventReport"`
	AdditionalInfo             interface{}     `json:"additionalInfo"`
	TicketType                 string          `json:"ticketType"`
	TicketTypeID               int             `json:"ticketTypeId"`
	SituationType              interface{}     `json:"situationType"`
	SituationTypeID            interface{}     `json:"situationTypeId"`
	Repeat                     string          `json:"repeat"`
	RepeatTypeID               int             `json:"repeatTypeId"`
	TypeName                   string          `json:"typeName"`
	SituationName              interface{}     `json:"situationName"`
	CriticalityID              int             `json:"criticalityId"`
	CriticalityName            string          `json:"criticalityName"`
	StationID                  interface{}     `json:"stationId"`
	StationName                interface{}     `json:"stationName"`
	StationCity                interface{}     `json:"stationCity"`
	StationUF                  interface{}     `json:"stationUf"`
	StationLatitude            interface{}     `json:"stationLatitude"`
	StationLongitude           interface{}     `json:"stationLongitude"`
	MaintenanceCenterID        int             `json:"maintenanceCenterId"`
	MaintenanceCenterName      string          `json:"maintenanceCenterName"`
	MaintenanceCenterCity      string          `json:"maintenanceCenterCity"`
	MaintenanceCenterUF        interface{}     `json:"maintenanceCenterUf"`
	MaintenanceCenterLatitude  int             `json:"maintenanceCenterLatitude"`
	MaintenanceCenterLongitude int             `json:"maintenanceCenterLongitude"`
	ActivityID                 int             `json:"activityId"`
	ActivityName               string          `json:"activityName"`
	CategoryID                 int             `json:"categoryId"`
	CategoryName               string          `json:"categoryName"`
	SubcategoryID              int             `json:"subcategoryId"`
	SubcategoryName            string          `json:"subcategoryName"`
	ClientID                   int             `json:"clientId"`
	ClientName                 string          `json:"clientName"`
	ContractID                 int             `json:"contractId"`
	ContractName               string          `json:"contractName"`
	TeamID                     interface{}     `json:"teamId"`
	TeamName                   interface{}     `json:"teamName"`
	QueueID                    int             `json:"queueId"`
	QueueName                  string          `json:"queueName"`
	TechnicianID               int             `json:"technicianId"`
	TechnicianName             string          `json:"technicianName"`
	TechnicianPhotoURL         string          `json:"technicianPhotoUrl"`
	FlowID                     interface{}     `json:"flowId"`
	FlowName                   interface{}     `json:"flowName"`
	FlowStateID                interface{}     `json:"flowStateId"`
	FlowStateName              interface{}     `json:"flowStateName"`
	FlowStartAt                interface{}     `json:"flowStartAt"`
	ScheduleStartAt            string          `json:"scheduleStartAt"`
	ScheduleEndAt              string          `json:"scheduleEndAt"`
	AssignEndAt                interface{}     `json:"assignEndAt"`
	CreateAt                   string          `json:"createAt"`
	PushScheduledAt            string          `json:"pushScheduledAt"`
	PushSentAt                 interface{}     `json:"pushSentAt"`
	Contacts                   []Contact       `json:"contacts"`
	FlowStates                 []dto.FlowState `json:"flowStates"`
}

type TicketInput struct {
	ID                         entity.ID       `json:"id"`
	IDTikect                   int             `json:"id_ticket"`
	ShortDescription           string          `json:"shortDescription"`
	Description                interface{}     `json:"description"`
	Code                       string          `json:"code"`
	ExternalCode               string          `json:"externalCode"`
	ParentTicketCode           interface{}     `json:"parentTicketCode"`
	RequesterName              interface{}     `json:"requesterName"`
	RequesterEmail             string          `json:"requesterEmail"`
	ActivityDescription        string          `json:"activityDescription"`
	EventReport                string          `json:"eventReport"`
	AdditionalInfo             interface{}     `json:"additionalInfo"`
	TicketType                 string          `json:"ticketType"`
	TicketTypeID               int             `json:"ticketTypeId"`
	SituationType              interface{}     `json:"situationType"`
	SituationTypeID            interface{}     `json:"situationTypeId"`
	Repeat                     string          `json:"repeat"`
	RepeatTypeID               int             `json:"repeatTypeId"`
	TypeName                   string          `json:"typeName"`
	SituationName              interface{}     `json:"situationName"`
	CriticalityID              int             `json:"criticalityId"`
	CriticalityName            string          `json:"criticalityName"`
	StationID                  interface{}     `json:"stationId"`
	StationName                interface{}     `json:"stationName"`
	StationCity                interface{}     `json:"stationCity"`
	StationUF                  interface{}     `json:"stationUf"`
	StationLatitude            interface{}     `json:"stationLatitude"`
	StationLongitude           interface{}     `json:"stationLongitude"`
	MaintenanceCenterID        int             `json:"maintenanceCenterId"`
	MaintenanceCenterName      string          `json:"maintenanceCenterName"`
	MaintenanceCenterCity      string          `json:"maintenanceCenterCity"`
	MaintenanceCenterUF        interface{}     `json:"maintenanceCenterUf"`
	MaintenanceCenterLatitude  int             `json:"maintenanceCenterLatitude"`
	MaintenanceCenterLongitude int             `json:"maintenanceCenterLongitude"`
	ActivityID                 int             `json:"activityId"`
	ActivityName               string          `json:"activityName"`
	CategoryID                 int             `json:"categoryId"`
	CategoryName               string          `json:"categoryName"`
	SubcategoryID              int             `json:"subcategoryId"`
	SubcategoryName            string          `json:"subcategoryName"`
	ClientID                   int             `json:"clientId"`
	ClientName                 string          `json:"clientName"`
	ContractID                 int             `json:"contractId"`
	ContractName               string          `json:"contractName"`
	TeamID                     interface{}     `json:"teamId"`
	TeamName                   interface{}     `json:"teamName"`
	QueueID                    int             `json:"queueId"`
	QueueName                  string          `json:"queueName"`
	TechnicianID               int             `json:"technicianId"`
	TechnicianName             string          `json:"technicianName"`
	TechnicianPhotoURL         string          `json:"technicianPhotoUrl"`
	FlowID                     interface{}     `json:"flowId"`
	FlowName                   interface{}     `json:"flowName"`
	FlowStateID                interface{}     `json:"flowStateId"`
	FlowStateName              interface{}     `json:"flowStateName"`
	FlowStartAt                interface{}     `json:"flowStartAt"`
	ScheduleStartAt            string          `json:"scheduleStartAt"`
	ScheduleEndAt              string          `json:"scheduleEndAt"`
	AssignEndAt                interface{}     `json:"assignEndAt"`
	CreateAt                   string          `json:"createAt"`
	PushScheduledAt            string          `json:"pushScheduledAt"`
	PushSentAt                 interface{}     `json:"pushSentAt"`
	Contacts                   []Contact       `json:"contacts"`
	FlowStates                 []dto.FlowState `json:"flowStates"`
}

type TicketResponse struct {
	Errors []interface{} `json:"errors"`
	Error  interface{}   `json:"error"`
	Data   struct {
		TotalQuantity int      `json:"totalQuantity"`
		Tickets       []Ticket `json:"tickets"`
	} `json:"data"`
	IsValid bool `json:"isValid"`
}

func NewTicket(ticket Ticket) TicketInput {
	return TicketInput{
		ID:                         entity.NewID(),
		IDTikect:                   ticket.ID,
		ShortDescription:           ticket.ShortDescription,
		Description:                ticket.Description,
		Code:                       ticket.Code,
		ExternalCode:               ticket.ExternalCode,
		ParentTicketCode:           ticket.ParentTicketCode,
		RequesterName:              ticket.RequesterName,
		RequesterEmail:             ticket.RequesterEmail,
		ActivityDescription:        ticket.ActivityDescription,
		EventReport:                ticket.EventReport,
		AdditionalInfo:             ticket.AdditionalInfo,
		TicketType:                 ticket.TicketType,
		TicketTypeID:               ticket.TicketTypeID,
		SituationType:              ticket.SituationType,
		SituationTypeID:            ticket.SituationTypeID,
		Repeat:                     ticket.Repeat,
		RepeatTypeID:               ticket.RepeatTypeID,
		TypeName:                   ticket.TypeName,
		SituationName:              ticket.SituationName,
		CriticalityID:              ticket.CriticalityID,
		CriticalityName:            ticket.CriticalityName,
		StationID:                  ticket.StationID,
		StationName:                ticket.StationName,
		StationCity:                ticket.StationCity,
		StationUF:                  ticket.StationUF,
		StationLatitude:            ticket.StationLatitude,
		StationLongitude:           ticket.StationLongitude,
		MaintenanceCenterID:        ticket.MaintenanceCenterID,
		MaintenanceCenterName:      ticket.MaintenanceCenterName,
		MaintenanceCenterCity:      ticket.MaintenanceCenterCity,
		MaintenanceCenterUF:        ticket.MaintenanceCenterUF,
		MaintenanceCenterLatitude:  ticket.MaintenanceCenterLatitude,
		MaintenanceCenterLongitude: ticket.MaintenanceCenterLongitude,
		ActivityID:                 ticket.ActivityID,
		ActivityName:               ticket.ActivityName,
		CategoryID:                 ticket.CategoryID,
		CategoryName:               ticket.CategoryName,
		SubcategoryID:              ticket.SubcategoryID,
		SubcategoryName:            ticket.SubcategoryName,
		ClientID:                   ticket.ClientID,
		ClientName:                 ticket.ClientName,
		ContractID:                 ticket.ContractID,
		ContractName:               ticket.ContractName,
		TeamID:                     ticket.TeamID,
		TeamName:                   ticket.TeamName,
		QueueID:                    ticket.QueueID,
		QueueName:                  ticket.QueueName,
		TechnicianID:               ticket.TechnicianID,
		TechnicianName:             ticket.TechnicianName,
		TechnicianPhotoURL:         ticket.TechnicianPhotoURL,
		FlowID:                     ticket.FlowID,
		FlowName:                   ticket.FlowName,
		FlowStateID:                ticket.FlowStateID,
		FlowStateName:              ticket.FlowStateName,
		FlowStartAt:                ticket.FlowStartAt,
		ScheduleStartAt:            ticket.ScheduleStartAt,
		ScheduleEndAt:              ticket.ScheduleEndAt,
		AssignEndAt:                ticket.AssignEndAt,
		CreateAt:                   ticket.CreateAt,
		PushScheduledAt:            ticket.PushScheduledAt,
		PushSentAt:                 ticket.PushSentAt,
		Contacts:                   ticket.Contacts,
		FlowStates:                 ticket.FlowStates,
	}
}
