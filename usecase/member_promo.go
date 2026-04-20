package usecase

import "github.com/XEFF09/calculator/domain"

type memberPromotion struct{}

func NewMemberPromotion() Promotion {
	return &memberPromotion{}
}

func (m *memberPromotion) Apply(order domain.Order, subTotal float64) (float64, error) {
	if !order.IsMember {
		return 0, nil
	}
	return subTotal * 0.10, nil
}
