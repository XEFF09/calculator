package domain

type Order struct {
	Items    map[Item]int
	IsMember bool
}
