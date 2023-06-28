package factory_method

type PoloShirtItf interface {
	AddCollar(color string)
}

type PoloShirt struct {
	Clothes
	CollarColor string
}

func (s *PoloShirt) AddCollar(color string) {
	s.CollarColor = color
}
func NewPoloShirt() ClothesItf[Clothes] {
	return &PoloShirt{}
}
func NewPoloShirt2() PoloShirtItf {
	return &PoloShirt{}
}
