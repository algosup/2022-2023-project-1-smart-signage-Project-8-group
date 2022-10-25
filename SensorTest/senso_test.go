package main

import (
	"testing"
)

func TestADCSensor(t *testing.T) {
	t.Run("test ADCSensor()", func(t *testing.T) {
		e := float32(1.1852446)
		result := ADCSensor(uint16(15535))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ADCSensor()", func(t *testing.T) {
		e := float32(3.05180438)
		result := ADCSensor(uint16(40000))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ADCSensor()", func(t *testing.T) {
		e := float32(5)
		result := ADCSensor(uint16(65535))
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}

func TestChangeLight(t *testing.T) {
	t.Run("test ChangeLight()", func(t *testing.T) {
		e := uint16(0)
		result := changeLight(float32(5), 100, 0)
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ChangeLight()", func(t *testing.T) {
		e := uint16(32767)
		result := changeLight(float32(2.5), 100, 0)
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test ChangeLight()", func(t *testing.T) {
		e := uint16(65403)
		result := changeLight(float32(0.01), 100, 0)
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}

func TestMainProg(t *testing.T) {
	t.Run("test MainProg()", func(t *testing.T) {
		e1 := uint16(0)
		e2 := uint16(1)
		e3 := uint16(1)
		e4 := "00360"

		r1, r2, r3, r4 := mainProg(uint16(65535), uint16(21000), uint16(20000), uint8(122))
		if r1 != e1 {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if r2 != e2 {
			t.Errorf("expected %v, got %v", e2, r2)
		}
		if r3 != e3 {
			t.Errorf("expected %v, got %v", e3, r3)
		}
		if r4 != e4 {
			t.Errorf("expected %v, got %v", e4, r4)
		}
	})

	t.Run("test MainProg()", func(t *testing.T) {
		e1 := uint16(52913)
		e2 := uint16(1)
		e3 := uint16(4)
		e4 := "500c3600"

		r1, r2, r3, r4 := mainProg(uint16(12622), uint16(25000), uint16(65000), uint8(12))
		if r1 != e1 {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if r2 != e2 {
			t.Errorf("expected %v, got %v", e2, r2)
		}
		if r3 != e3 {
			t.Errorf("expected %v, got %v", e3, r3)
		}
		if r4 != e4 {
			t.Errorf("expected %v, got %v", e4, r4)
		}
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("test SendMessage()", func(t *testing.T) {
		e := `AT+MSG= "00360"` + "\r\n"
		result := SendMessage("00360")
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test SendMessage()", func(t *testing.T) {
		e := `AT+MSG= "500c3600"` + "\r\n"
		result := SendMessage("500c3600")
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}

func TestSendMessage(t *testing.T) {
	t.Run("test SendMessage()", func(t *testing.T) {
		e := `AT+MSG= "00360"` + "\r\n"
		result := SendMessage("00360")
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})

	t.Run("test SendMessage()", func(t *testing.T) {
		e := `AT+MSG= "500c3600"` + "\r\n"
		result := SendMessage("500c3600")
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}

/*
func TestMsgTreating(t *testing.T) {
	t.Run("test msgTreating()", func(t *testing.T) {
		e1 := false
		e2 := false
		e3 := false
		e4 := false
		r1, r2, r3, r4 := SendMessage("01")
		if r1 != e1 {
			t.Errorf("expected %v, got %v", e1, r1)
		}
	})
}*/
