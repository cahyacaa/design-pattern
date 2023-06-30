package builder

type BuildElementProcess interface {
	SetCamera() BuildElementProcess
	SetBattery() BuildElementProcess
	SetDisplay() BuildElementProcess
	SetOS() BuildElementProcess
	GetDetails() Smartphone
}
