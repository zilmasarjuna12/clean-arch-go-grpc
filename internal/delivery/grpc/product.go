package delivery_grpc

import (
	product_grpc "clean-arch-go-grpc/internal/delivery/grpc/proto"
	"clean-arch-go-grpc/internal/entity"
	"clean-arch-go-grpc/internal/usecase"
	"context"
	"io"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewProductServerGrpc(
	gserver *grpc.Server,
	log *logrus.Logger,
	productUsecase usecase.IProductUsecase,
) {
	productServer := &server{productUsecase: productUsecase, log: log}

	product_grpc.RegisterProductHandlerServer(gserver, productServer)
	reflection.Register(gserver)
}

type server struct {
	product_grpc.UnimplementedProductHandlerServer
	productUsecase usecase.IProductUsecase
	log            *logrus.Logger
}

func (s *server) Create(ctx context.Context, reqProduct *product_grpc.Product) (*product_grpc.Product, error) {
	req := &entity.Product{
		Name:        reqProduct.Name,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
	}

	product, err := s.productUsecase.Create(context.Background(), req)
	if err != nil {
		return nil, err
	}

	res := &product_grpc.Product{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	return res, nil
}

func (s *server) GetList(context.Context, *product_grpc.Empty) (*product_grpc.Products, error) {
	products, err := s.productUsecase.Gets(context.Background())
	if err != nil {
		return nil, err
	}

	res := &product_grpc.Products{
		Products: []*product_grpc.Product{},
	}

	for _, product := range products {
		res.Products = append(res.Products, &product_grpc.Product{
			ID: product.ID.String(),
		})
	}

	return res, nil
}

func (s *server) Get(ctx context.Context, in *product_grpc.GetRequest) (*product_grpc.Product, error) {
	product, err := s.productUsecase.GetByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	res := &product_grpc.Product{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	return res, nil
}

func (s *server) GetStream(in *product_grpc.Empty, stream product_grpc.ProductHandler_GetStreamServer) error {
	products, err := s.productUsecase.Gets(context.Background())
	if err != nil {
		return err
	}

	for _, product := range products {
		res := &product_grpc.Product{
			ID:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) BatchCreate(stream product_grpc.ProductHandler_BatchCreateServer) error {
	errs := make([]*product_grpc.ErrorMessage, 0)

	totalSuccess := int64(0)

	for {
		product, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&product_grpc.BatchCreateResponse{
				TotalSuccess: totalSuccess,
				Errors:       errs,
			})
		}

		if err != nil {
			return err
		}

		req := &entity.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}

		res, err := s.productUsecase.Create(context.Background(), req)
		if err != nil {
			e := &product_grpc.ErrorMessage{
				Message: err.Error(),
			}

			errs = append(errs, e)
		}

		if res != nil {
			totalSuccess++
		}
	}
}
