package main

import (
	"testing"
)

func TestChangeLight(t *testing.T) {
	t.Run("test ChangeLight()", func(t *testing.T) {
		e := float32(1.18524452583)
		result := ADCSensor(uint16(15535))
		if result == e {
			t.Errorf("expected 1.18524452583, got %v", result)
		}
	})
}
