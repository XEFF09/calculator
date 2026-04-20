package usecase

import "github.com/XEFF09/calculator/domain"

type Promotion interface {
	Apply(order domain.Order) float64
}
