package usecase

import (
	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/repository"
)

type OrderService interface {
	SubTotal() (float64, error)
}

type order struct {
	Items     map[domain.Item]int
	IsMember  bool
	stockRepo repository.StockRepository
}

func NewOrder(items map[domain.Item]int, isMember bool, stockRepo repository.StockRepository) OrderService {
	return &order{
		Items:     items,
		IsMember:  isMember,
		stockRepo: stockRepo,
	}
}

func (o *order) SubTotal() (float64, error) {
	total := 0.0
	for item, qty := range o.Items {
		stockItem, err := o.stockRepo.GetByName(item)
		if err != nil {
			return 0, err
		}

		total += stockItem.Price * float64(qty)
	}
	return total, nil
}
