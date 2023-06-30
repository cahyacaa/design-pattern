package builder

type ManufacturingDirector struct {
	builder BuildElementProcess
}

func (d *ManufacturingDirector) Construct() {
	d.builder.SetDisplay().SetOS().SetBattery().SetCamera()
}

func (d *ManufacturingDirector) SetBuilder(b BuildElementProcess) {
	d.builder = b
}
