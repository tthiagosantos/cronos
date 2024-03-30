package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

var jwtKey = []byte("SECRET")

func JWTInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor", info.FullMethod)

	if info.FullMethod == "/pb.UserService/Authentication" {
		// Se for o serviço de autenticação, permitir o acesso público sem validação de token JWT
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(grpc.Code(errors.New("403")), "metadata não está presente na solicitação")
	}

	authHeaders, ok := md["authorization"]
	if !ok {
		return nil, grpc.Errorf(grpc.Code(errors.New("403")), "token de autorização não está presente na solicitação")
	}

	authHeader := authHeaders[0]
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, grpc.Errorf(grpc.Code(errors.New("403")), "falha ao analisar o token: %v", err)
	}

	if !token.Valid {
		return nil, grpc.Errorf(grpc.Code(errors.New("403")), "token inválido")
	}

	return handler(ctx, req)
}
