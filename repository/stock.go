package repository

import "github.com/XEFF09/calculator/domain"

type StockRepository interface {
	GetByName(name domain.Item) (*domain.ItemInfo, error)
	UpdateStockByName(name domain.Item, reqQty int) error
}
