package main

import (
	"testing"
)

func TestChangeLight(t *testing.T) {
	t.Run("test ChangeLight()", func(t *testing.T) {
		e := uint16(12000)
		result := changeLight(2.84)
		if result == e {
			t.Errorf("expected 12000, got %v", result)
		}
	})

	t.Run("test ChangeLight()", func(t *testing.T) {
		e := uint16(42000)
		result := changeLight(4.76)
		if result == e {
			t.Errorf("expected 42000, got %v", result)
		}
	})
}

func TestADCSensorConv(t *testing.T) {
	t.Run("test ADCSensorConv()", func(t *testing.T) {
		e := float32(2.82711528191)
		result := ADCSensorConv(37055)
		if result == e {
			t.Errorf("expected 2.82711528191, got %v", result)
		}
	})

	t.Run("test ADCSensorConv()", func(t *testing.T) {
		e := float32(3.13992523079)
		result := ADCSensorConv(41155)
		if result == e {
			t.Errorf("expected 2.82711528191, got %v", result)
		}
	})
}
