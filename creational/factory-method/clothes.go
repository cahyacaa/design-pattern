package factory_method

type ClothesSize string
type ClothesType string

const (
	PoloType   ClothesType = "polo"
	TShirtType             = "t-shirt"
)

const (
	S   ClothesSize = "S"
	M               = "M"
	L               = "L"
	XL              = "XL"
	XXL             = "XXL"
)

type Clothes struct {
	Color string
	Price int64
	Size  ClothesSize
}
