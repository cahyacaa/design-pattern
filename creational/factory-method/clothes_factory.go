package factory_method

func getClothes(cType ClothesType) ClothesItf[Clothes] {
	switch cType {
	case PoloType:
		return NewPoloShirt()
	case TShirtType:
		return NewTShirt()
	default:
		return nil
	}
}
