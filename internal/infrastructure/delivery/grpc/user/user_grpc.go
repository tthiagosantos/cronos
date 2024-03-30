package user

import (
	"context"
	"cronos/internal/domain/user/usecase"
	usecase2 "cronos/internal/domain/wfm/usecase"
	"cronos/internal/infrastructure/delivery/grpc/pb"
	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc"
	"net/http"

	"cronos/internal/infrastructure/middleware"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	UserUsecase               usecase.UserUsecase
	UserAuthenticationUseCase usecase.UserAuthenticateUsecase
	LoginWFM                  usecase2.NewLoginUseCase
	ListWFM                   usecase2.ListTicketWFMUseCase
	SaveWFM                   usecase2.TicketSaveWFMUseCase
	FlowUseCase               usecase2.FlowUseCase
	FlowSaveUseCase           usecase2.FlowSaveUseCase
	Jwt                       *jwtauth.JWTAuth
	JwtExperiesIn             int
}

func NewUserServiceServer(userUsecase usecase.UserUsecase) *UserServiceServer {
	return &UserServiceServer{
		UserUsecase: userUsecase,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, err := s.UserUsecase.Execute(req.Email, req.Password)
	if err != nil {
		return &pb.CreateUserResponse{
			Success: false,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}, nil
	}
	return &pb.CreateUserResponse{
		Success: true,
		Message: "Registrado com sucesso",
		Status:  http.StatusCreated,
	}, nil
}

func (s *UserServiceServer) Authentication(ctx context.Context, in *pb.CreateUserRequest) (*pb.TokenResponse, error) {
	token, err := s.UserAuthenticationUseCase.Execute(in.Email, in.Password)
	if err != nil {
		return &pb.TokenResponse{
			Token:   "",
			Message: err.Error(),
		}, err
	}
	tokenWFM := s.LoginWFM.Execute()
	lisTikectWFM, flowIds, err := s.ListWFM.Execute(tokenWFM)
	if err != nil {
		return &pb.TokenResponse{
			Token:   "",
			Message: err.Error(),
		}, err
	}

	err = s.SaveWFM.Execute(*lisTikectWFM)
	if err != nil {
		return nil, err
	}

	flowResponse, err := s.FlowUseCase.Execute(*flowIds, tokenWFM)
	if err != nil {
		return &pb.TokenResponse{
			Token:   "",
			Message: err.Error(),
		}, err
	}
	s.FlowSaveUseCase.Execute(flowResponse)
	return &pb.TokenResponse{Token: token, Message: tokenWFM}, nil
}

func (s *UserServiceServer) Register(server *grpc.Server) {
	pb.RegisterUserServiceServer(server, s)
}

func (s *UserServiceServer) Middleware() grpc.UnaryServerInterceptor {
	return middleware.JWTInterceptor
}
