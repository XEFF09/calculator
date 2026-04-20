package main

import (
	"fmt"

	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/internal/adapter/local"
	"github.com/XEFF09/calculator/usecase"
)

func main() {
	// mock stock repo
	stockRepo := local.NewStock()

	order := domain.Order{
		IsMember: true,
		Items: map[domain.Item]int{
			domain.OrangeSet: 7,
			domain.RedSet:    1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	calculator := usecase.NewCalculator(
		order,
		orderService,
		[]usecase.Promotion{
			usecase.NewBundlePromotion(stockRepo),
			usecase.NewMemberPromotion(),
		},
	)

	total, err := calculator.Calculate()
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Price:", total)
}
