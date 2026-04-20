package usecase

import (
	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/repository"
)

type bundlePromo struct {
	stockRepo repository.StockRepository
}

func NewBundlePromotion(repo repository.StockRepository) Promotion {
	return &bundlePromo{stockRepo: repo}
}

func (b *bundlePromo) Apply(order domain.Order, subtotal float64) (float64, error) {
	discount := 0.0

	target := map[domain.Item]bool{
		domain.OrangeSet: true,
		domain.PinkSet:   true,
		domain.GreenSet:  true,
	}

	for item, qty := range order.Items {
		if !target[item] {
			continue
		}

		pairs := qty / 2
		if pairs == 0 {
			continue
		}

		stockItem, err := b.stockRepo.GetByName(item)
		if err != nil {
			return 0, err
		}

		discount += float64(pairs*2) * stockItem.Price * 0.05
	}

	return discount, nil
}
