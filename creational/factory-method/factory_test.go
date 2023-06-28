package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunTShirt(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test Factory Method Pattern", func(t *testing.T) {
		polo := getClothes(PoloType)
		tshirt := getClothes(TShirtType)

		polo.SetPrice(1000.000)
		polo.SetColor("blue")
		polo.SetSize(XXL)

		assert.Equal(Clothes{
			Color: "blue",
			Price: 1000.000,
			Size:  XXL,
		}, polo.GetData())

		tshirt.SetPrice(100.000)
		tshirt.SetColor("red")
		tshirt.SetSize(XL)

		assert.Equal(Clothes{
			Color: "red",
			Price: 100.000,
			Size:  XL,
		}, tshirt.GetData())
	})

	t.Run("Test TShirt", func(t *testing.T) {
		shirt := NewTShirt()
		shirt.SetPrice(100.000)
		shirt.SetColor("red")
		shirt.SetSize(XL)

		assert.Equal(Clothes{
			Color: "red",
			Price: 100.000,
			Size:  XL,
		}, shirt.GetData())
	})

	t.Run("Test Polo Shirt", func(t *testing.T) {
		shirt := NewPoloShirt()
		shirt.SetPrice(1000.000)
		shirt.SetColor("blue")
		shirt.SetSize(XXL)

		assert.Equal(Clothes{
			Color: "blue",
			Price: 1000.000,
			Size:  XXL,
		}, shirt.GetData())
	})
}
