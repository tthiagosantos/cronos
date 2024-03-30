package flowstates

import (
	"context"
	"cronos/internal/domain/wfm/usecase"
	"cronos/internal/infrastructure/delivery/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

type FlowStateServer struct {
	pb.UnimplementedFlowDataServiceServer
	ListUseCase usecase.FlowStateUseCaseList
}

func (s *FlowStateServer) Register(server *grpc.Server) {
	pb.RegisterFlowDataServiceServer(server, s)
}

func NewFlowStateServer(flowStateUsecase usecase.FlowStateUseCaseList) *FlowStateServer {
	return &FlowStateServer{
		ListUseCase: flowStateUsecase,
	}
}

func (fs *FlowStateServer) FindAll(ctx context.Context, req *pb.FlowDataRequest) ([]*pb.FlowDataResponse, error) {
	responses, err := fs.ListUseCase.Execute()
	if err != nil {
		return nil, err
	}
	log.Println("PASSANDO AQUI")
	var flowDataResponses []*pb.FlowDataResponse
	for _, flowData := range responses {
		flowDataResponse := &pb.FlowDataResponse{
			XId:                       "",
			Flowid:                    strconv.Itoa(flowData.FlowId),
			Floworder:                 int32(flowData.FlowOrder),
			Flowstatedescription:      flowData.FlowStateDescription,
			Flowstateid:               int32(flowData.FlowStateId),
			Flowstatename:             flowData.FlowStateName,
			Macroflowstatedescription: flowData.MacroFlowStateDescription,
			Macroflowstateid:          int32(flowData.MacroFlowStateId),
		}
		flowDataResponses = append(flowDataResponses, flowDataResponse)
	}
	return flowDataResponses, nil
}
