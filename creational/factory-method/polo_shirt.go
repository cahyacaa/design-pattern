package factory_method

type PoloShirt struct {
	Clothes
	CollarColor string
}

func (t *PoloShirt) SetColor(color string) {
	t.Color = color
	t.CollarColor = "white"
}

func (t *PoloShirt) SetSize(size ClothesSize) {
	t.Size = size
}

func (t *PoloShirt) SetPrice(price int64) {
	t.Price = price
}

func (t *PoloShirt) GetData() PoloShirt {
	return PoloShirt{
		Clothes: Clothes{
			Color: t.Color,
			Price: t.Price,
			Size:  t.Size,
		},
		CollarColor: t.CollarColor,
	}
}

func NewPoloShirt() ClothesItf[PoloShirt] {
	return &PoloShirt{}
}
