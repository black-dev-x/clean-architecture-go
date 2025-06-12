package service

import (
	"context"

	"github.com/black-dev-x/clean-architecture-go/internal/infra/grpc/pb"
	"github.com/black-dev-x/clean-architecture-go/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Empty) (*pb.ListOrders, error) {
	dto, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	orders := make([]*pb.Order, 0, len(*dto))
	for _, order := range *dto {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrders{Orders: orders}, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
