package singleton

import "testing"

func TestSingleton(t *testing.T) {
	t.Run("Singleton With Result OK", func(t *testing.T) {
		instanceData := DBObjectExample{TotalCalled: 1}

		if err := InitDBExample(instanceData); err != nil {
			t.Error(err.Error())
		}

		instanceData.TotalCalled = 2
		if err := InitDBExample(instanceData); err != nil {
			t.Error(err.Error())
		}

		instanceInMemory := GetInstance()
		if instanceInMemory.TotalCalled != 1 {
			t.Error("Total Called is more than once")
		}
	})
}
