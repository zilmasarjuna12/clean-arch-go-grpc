package main

import (
	product_grpc "clean-arch-go-grpc/internal/delivery/grpc/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	product := NewProduct()
	defer product.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := product.GetStream(ctx)

	if err != nil {
		return
	}

	for r := range res {
		log.Printf("ID: %s", r)
	}
}

type product struct {
	conn    *grpc.ClientConn
	handler product_grpc.ProductHandlerClient
}

func NewProduct() *product {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &product{
		conn:    conn,
		handler: product_grpc.NewProductHandlerClient(conn),
	}
}

func (p *product) Close() {
	p.conn.Close()
}

func (p *product) Create(ctx context.Context, name, description string, price float32) (*product_grpc.Product, error) {
	req := &product_grpc.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	return p.handler.Create(ctx, req)
}

func (p *product) GetList(ctx context.Context) (*product_grpc.Products, error) {
	return p.handler.GetList(ctx, &product_grpc.Empty{})
}

func (p *product) Get(ctx context.Context, id string) (*product_grpc.Product, error) {
	return p.handler.Get(ctx, &product_grpc.GetRequest{ID: id})
}

func (p *product) GetStream(ctx context.Context) (chan *product_grpc.Product, error) {
	stream, err := p.handler.GetStream(ctx, &product_grpc.Empty{})
	if err != nil {
		return nil, err
	}

	ch := make(chan *product_grpc.Product)

	go func() {
		defer close(ch)

		for {
			res, err := stream.Recv()
			if err != nil {
				return
			}

			ch <- res
		}
	}()

	return ch, nil
}

func (p *product) BatchCreate(ctx context.Context) error {
	reqs := []*product_grpc.Product{
		{
			Name:        "Product 1",
			Description: "Product 1",
			Price:       1000,
		},
		{
			Name:        "Product 1",
			Description: "Product 1",
			Price:       1000,
		},
	}

	stream, err := p.handler.BatchCreate(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", req, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)

		return err
	}

	log.Printf("Route summary: %v", reply.TotalSuccess)

	return nil
}
