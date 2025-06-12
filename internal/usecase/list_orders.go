package usecase

import (
	"github.com/black-dev-x/clean-architecture-go/internal/entity"
)

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (*[]*OrderDTO, error) {
	println("ListOrdersUseCase.Execute", "start")
	println("ListOrdersUseCase.Execute", "OrderRepository", c.OrderRepository)
	orders, err := c.OrderRepository.FindAll()
	println("ListOrdersUseCase.Execute", "orders", len(*orders), "err", err)
	dtos := make([]*OrderDTO, 0, len(*orders))
	for _, order := range *orders {
		dtos = append(dtos, &OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return &dtos, err
}
