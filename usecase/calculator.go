package usecase

import (
	"github.com/XEFF09/calculator/domain"
)

type Calculator struct {
	orderService OrderService
	order        domain.Order
	promotions   []Promotion
}

func NewCalculator(
	order domain.Order,
	orderService OrderService,
	promos []Promotion,
) *Calculator {
	return &Calculator{
		order:        order,
		orderService: orderService,
		promotions:   promos,
	}
}

func (c *Calculator) Calculate() (float64, error) {
	subTotal, err := c.orderService.SubTotal()
	if err != nil {
		return 0, err
	}

	totalDiscount := 0.0

	for _, p := range c.promotions {
		promoDiscount, err := p.Apply(c.order, subTotal)
		if err != nil {
			return 0, err
		}

		totalDiscount += promoDiscount
	}

	return subTotal - totalDiscount, nil
}
