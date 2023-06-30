package factory_method

func getClothes(cType ClothesType) interface{} {
	switch cType {
	case PoloType:
		return NewPoloShirt()
	case TShirtType:
		return NewTShirt()
	default:
		return nil
	}
}
