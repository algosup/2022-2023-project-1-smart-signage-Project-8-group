package main

import (
	"testing"
)

func TestADCSenso(t *testing.T) {
	t.Run("test ChangeLight()", func(t *testing.T) {
		e := float32(1.1852446)
		result := ADCSensor(uint16(15535))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ChangeLight()", func(t *testing.T) {
		e := float32(3.05180438)
		result := ADCSensor(uint16(40000))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ChangeLight()", func(t *testing.T) {
		e := float32(5)
		result := ADCSensor(uint16(65535))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}
