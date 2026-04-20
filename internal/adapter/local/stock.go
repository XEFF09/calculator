package local

import (
	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/domain/exception"
	"github.com/XEFF09/calculator/repository"
)

type stock struct {
	stockItems map[domain.Item]domain.ItemInfo
}

func NewStock() repository.StockRepository {
	stockItems := map[domain.Item]domain.ItemInfo{
		domain.RedSet: {
			Price:    50,
			Quantity: 100,
		},
		domain.GreenSet: {
			Price:    40,
			Quantity: 100,
		},
		domain.BlueSet: {
			Price:    30,
			Quantity: 100,
		},
		domain.YellowSet: {
			Price:    50,
			Quantity: 100,
		},
		domain.PinkSet: {
			Price:    80,
			Quantity: 100,
		},
		domain.PurpleSet: {
			Price:    90,
			Quantity: 100,
		},
		domain.OrangeSet: {
			Price:    120,
			Quantity: 100,
		},
	}

	return &stock{
		stockItems: stockItems,
	}
}

func (s *stock) GetByName(name domain.Item) (*domain.ItemInfo, error) {
	data, found := s.stockItems[domain.Item(name)]
	if !found {
		return nil, exception.ErrMenuItemNotFound
	}

	return &data, nil
}

func (s *stock) UpdateStockByName(name domain.Item, reqQty int) error {
	data, found := s.stockItems[domain.Item(name)]
	if !found {
		return exception.ErrMenuItemNotFound
	}

	data.Quantity -= reqQty
	s.stockItems[domain.Item(name)] = data

	return nil
}
