package main

import (
	"cronos/configs"
	"cronos/internal/domain/user/repository"
	"cronos/internal/domain/user/usecase"
	repository2 "cronos/internal/domain/wfm/repository"
	usecase2 "cronos/internal/domain/wfm/usecase"
	"cronos/internal/infrastructure/database"
	"cronos/internal/infrastructure/delivery/grpc"
	"cronos/internal/infrastructure/delivery/grpc/flowstates"
	"cronos/internal/infrastructure/delivery/grpc/user"
	"fmt"
	"log"
)

func init() {
	var err error
	configs.ConfigEnv, err = configs.LoadConfig("./cmd/application")
	if err != nil {
		panic(err)
	}

}

func main() {
	mongoConnectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		configs.ConfigEnv.DBUser, configs.ConfigEnv.DBPassword, configs.ConfigEnv.DBHost, configs.ConfigEnv.DBPort)
	if err := database.ConnectMongoDB(mongoConnectionString, configs.ConfigEnv.DBName); err != nil {
		log.Println("Error connect mondoDB")
		panic(err)
	}
	userRepo := repository.NewUserRepository()
	flowRepo := repository2.NewFlowStateRepository()
	userUsecase := &usecase.UserUsecase{
		UserRepository: userRepo,
	}
	flowUseCase := &usecase2.FlowStateUseCaseList{
		FlowStateListRepository: flowRepo,
	}
	grpc.StartServer(
		configs.ConfigEnv.GrpcServerPort,
		user.NewUserServiceServer(*userUsecase),
		flowstates.NewFlowStateServer(*flowUseCase),
	)
}
