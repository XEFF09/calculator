package usecase

import (
	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/domain/exception"
	"github.com/XEFF09/calculator/repository"
)

type OrderService interface {
	SubTotal() (float64, error)
}

type storeOrder struct {
	order     domain.Order
	stockRepo repository.StockRepository
}

func NewOrder(order domain.Order, stockRepo repository.StockRepository) OrderService {
	return &storeOrder{
		order:     order,
		stockRepo: stockRepo,
	}
}

func (so *storeOrder) SubTotal() (float64, error) {
	total := 0.0
	for item, qty := range so.order.Items {

		if qty <= 0 {
			return 0, exception.ErrInvalidQuantity
		}

		stockItem, err := so.stockRepo.GetByName(item)
		if err != nil {
			return 0, err
		}

		if qty > stockItem.Quantity {
			return 0, exception.ErrNotEnoughStock
		}

		total += stockItem.Price * float64(qty)
	}
	return total, nil
}
