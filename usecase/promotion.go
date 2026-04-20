package usecase

import "github.com/XEFF09/calculator/domain"

type Promotion interface {
	Apply(order domain.Order, subTotal float64) (float64, error)
}
