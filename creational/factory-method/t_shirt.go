package factory_method

type TShirt struct {
	Clothes
}

func (t *TShirt) SetColor(color string) {
	t.Color = color
}

func (t *TShirt) SetSize(size ClothesSize) {
	t.Size = size
}

func (t *TShirt) SetPrice(price int64) {
	t.Price = price
}

func (t *TShirt) GetData() TShirt {
	return TShirt{
		Clothes: Clothes{
			Color: t.Color,
			Price: t.Price,
			Size:  t.Size,
		},
	}
}

func NewTShirt() ClothesItf[TShirt] {
	return &TShirt{
		Clothes{},
	}
}
