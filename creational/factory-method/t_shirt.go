package factory_method

type TShirt struct {
	Clothes
}

func NewTShirt() ClothesItf[Clothes] {
	return &TShirt{
		Clothes{},
	}
}
