package exception

import "errors"

var (
	ErrMenuItemNotFound = errors.New("menu item not found")
	ErrInvalidQuantity  = errors.New("invalid quantity")
	ErrNotEnoughStock   = errors.New("not enough stock")
)
