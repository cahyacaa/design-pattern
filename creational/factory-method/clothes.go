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

func (t *Clothes) SetColor(color string) {
	t.Color = color
}

func (t *Clothes) SetSize(size ClothesSize) {
	t.Size = size
}

func (t *Clothes) SetPrice(price int64) {
	t.Price = price
}

func (t *Clothes) GetData() Clothes {
	return Clothes{
		Color: t.Color,
		Price: t.Price,
		Size:  t.Size,
	}
}
