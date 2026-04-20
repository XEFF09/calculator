package test

import (
	"testing"

	"github.com/XEFF09/calculator/domain"
	"github.com/XEFF09/calculator/domain/exception"
	"github.com/XEFF09/calculator/internal/adapter/local"
	"github.com/XEFF09/calculator/usecase"
)

func Test_SubTotal(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			domain.RedSet:   1,
			domain.GreenSet: 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	result, err := orderService.SubTotal()
	if err != nil {
		t.Fatal(err)
	}

	expected := 90.0

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Test_BundlePromotion(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			domain.OrangeSet: 5,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	calculator := usecase.NewCalculator(
		order,
		orderService,
		[]usecase.Promotion{
			usecase.NewBundlePromotion(stockRepo),
		},
	)

	result, err := calculator.Calculate()
	if err != nil {
		t.Fatal(err)
	}

	expected := 576.0

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Test_MemberPromotion(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		IsMember: true,
		Items: map[domain.Item]int{
			domain.RedSet:   1,
			domain.GreenSet: 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	calculator := usecase.NewCalculator(
		order,
		orderService,
		[]usecase.Promotion{
			usecase.NewMemberPromotion(),
		},
	)

	result, err := calculator.Calculate()
	if err != nil {
		t.Fatal(err)
	}

	expected := 90 * 0.9

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Test_CombinedPromotions(t *testing.T) {
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

	result, err := calculator.Calculate()
	if err != nil {
		t.Fatal(err)
	}

	expected := 765.0

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Test_BundleOddQuantity(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			domain.GreenSet: 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	calculator := usecase.NewCalculator(
		order,
		orderService,
		[]usecase.Promotion{
			usecase.NewBundlePromotion(stockRepo),
		},
	)

	result, err := calculator.Calculate()
	if err != nil {
		t.Fatal(err)
	}

	expected := 40.0

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func Test_InvalidQuantityRequest(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			domain.RedSet:   0,
			domain.GreenSet: 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	_, err := orderService.SubTotal()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != exception.ErrInvalidQuantity.Error() {
		t.Errorf("expected %v, got %v", exception.ErrInvalidQuantity, err)
	}
}

func Test_NotEnoughStock(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			domain.RedSet:   101,
			domain.GreenSet: 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	_, err := orderService.SubTotal()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != exception.ErrNotEnoughStock.Error() {
		t.Errorf("expected %v, got %v", exception.ErrNotEnoughStock, err)
	}
}

func Test_ItemNotFound(t *testing.T) {
	stockRepo := local.NewStock()

	order := domain.Order{
		Items: map[domain.Item]int{
			"WhiteSet": 1,
		},
	}

	orderService := usecase.NewOrder(order, stockRepo)

	_, err := orderService.SubTotal()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != exception.ErrMenuItemNotFound.Error() {
		t.Errorf("expected %v, got %v", exception.ErrMenuItemNotFound, err)
	}
}
