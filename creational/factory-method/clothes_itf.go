package factory_method

type ClothesItf[T any] interface {
	SetSize(size ClothesSize)
	SetColor(color string)
	SetPrice(price int64)
	GetData() T
}
