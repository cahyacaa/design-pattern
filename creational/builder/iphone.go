package builder

type Iphone struct {
	phone Smartphone
}

func (i *Iphone) SetBattery() BuildElementProcess {
	i.phone.Battery = "3600 mAh"
	return i
}

func (i *Iphone) SetOS() BuildElementProcess {
	i.phone.OS = "iOS"
	return i
}

func (i *Iphone) SetCamera() BuildElementProcess {
	i.phone.Camera = 3
	return i
}
func (i *Iphone) SetDisplay() BuildElementProcess {
	i.phone.OS = "OLED"
	return i
}

func (i *Iphone) GetDetails() Smartphone {
	i.phone.Name = "iPhone"
	return i.phone
}
