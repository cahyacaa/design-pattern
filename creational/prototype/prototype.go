package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(m ShirtColor) (ItemInfoGetter, error)
}

type ShirtColor string

const (
	White ShirtColor = "white"
	Black            = "black"
	Blue             = "blue"
)

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m ShirtColor) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
	GetPrice() float32
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
