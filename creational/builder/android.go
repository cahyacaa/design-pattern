package builder

type Android struct {
	phone Smartphone
}

func (i *Android) SetBattery() BuildElementProcess {
	i.phone.Battery = "5000 mAh"
	return i
}

func (i *Android) SetOS() BuildElementProcess {
	i.phone.OS = "Tiramisu"
	return i
}

func (i *Android) SetCamera() BuildElementProcess {
	i.phone.Camera = 4
	return i
}
func (i *Android) SetDisplay() BuildElementProcess {
	i.phone.OS = "AMOLED"
	return i
}

func (i *Android) GetDetails() Smartphone {
	i.phone.Name = "Google Pixes"
	return i.phone
}
