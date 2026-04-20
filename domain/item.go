package domain

type Item string

const (
	RedSet    Item = "Red Set"
	GreenSet  Item = "Green Set"
	BlueSet   Item = "Blue Set"
	YellowSet Item = "Yellow Set"
	PinkSet   Item = "Pink Set"
	PurpleSet Item = "Purple Set"
	OrangeSet Item = "Orange Set"
)

type ItemInfo struct {
	Price    float64
	Quantity int
}
