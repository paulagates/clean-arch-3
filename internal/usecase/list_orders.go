package usecase

import (
	"github.com/paulagates/clean-arch-3/internal/entity"
	"github.com/paulagates/clean-arch-3/pkg/events"
)

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var output ListOrdersOutputDTO
	for _, order := range orders {
		output.Orders = append(output.Orders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	l.OrdersListed.SetPayload(orders)
	l.EventDispatcher.Dispatch(l.OrdersListed)

	return output, nil
}
