package builder

import "testing"

func TestBuilder(t *testing.T) {
	t.Run("Testing Builder Design Pattern", func(t *testing.T) {
		directorManufacture := ManufacturingDirector{}

		android := Android{}
		directorManufacture.SetBuilder(&android)
		directorManufacture.Construct()

		if android.phone.OS != "Tiramisu" {
			t.Errorf("expected must be 'Tiramisu', actual :,  %s\n", android.phone.OS)
		}

		iPhone := Iphone{}
		directorManufacture.SetBuilder(&iPhone)
		directorManufacture.Construct()

		if iPhone.phone.OS != "iOS" {
			t.Errorf("expected must be 'Tiramisu', actual :,  %s\n", iPhone.phone.OS)
		}
	})

}
