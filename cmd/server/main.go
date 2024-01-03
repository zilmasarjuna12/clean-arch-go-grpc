package main

import (
	delivery_grpc "clean-arch-go-grpc/internal/delivery/grpc"
	"clean-arch-go-grpc/internal/repository"
	"clean-arch-go-grpc/internal/usecase"
	"clean-arch-go-grpc/pkg/gorm"
	"clean-arch-go-grpc/pkg/logrus"
	"clean-arch-go-grpc/pkg/viper"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	viper := viper.NewViper()
	logrus := logrus.NewLogger()
	gorm := gorm.NewDatabase(viper, logrus)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pr := repository.NewProductRepository(gorm, logrus)

	pu := usecase.NewProductUsecase(logrus, pr)

	delivery_grpc.NewProductServerGrpc(server, logrus, pu)

	err = server.Serve(lis)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}
