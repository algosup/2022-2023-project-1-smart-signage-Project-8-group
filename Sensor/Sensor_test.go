package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCheckPowerStatusH(t *testing.T) {
	t.Run("test checkPowerStatusH()", func(t *testing.T) {
		e := true
		result := checkPowerStatusH()
		if result == e {
			t.Errorf("expected false, got %v", result)
		}
		time.Sleep(time.Millisecond * 10)
		powerLED()
		result = checkPowerStatusH()
		if result != e {
			t.Errorf("expected true, got %v", result)
		}
	})
}

func TestCheckPowerStatusL(t *testing.T) {
	t.Run("test checkPowerStatusL()", func(t *testing.T) {
		e := true
		result := checkPowerStatusL()
		if result == e {
			t.Errorf("expected false, got %v", result)
		}
		time.Sleep(time.Millisecond * 10)
		powerLED()
		result = checkPowerStatusL()
		if result != e {
			t.Errorf("expected true, got %v", result)
		}
	})
}

func TestCheckHeat(t *testing.T) {
	t.Run("test checkHeat()", func(t *testing.T) {
		e := 0.0
		result := checkHeat()
		if reflect.TypeOf(result) == reflect.TypeOf(e) {
			t.Errorf("expected a float, got %v", reflect.TypeOf(result))
		}
	})
}

func TestCheckLight(t *testing.T) {
	t.Run("test checkLight()", func(t *testing.T) {
		now, _ = strconv.Atoi(time.Now().String()[11:13])
		e := false
		if (now > 8) || (now < 20) {
			e = true
		}

		result := checkLight()
		if result != e {
			t.Errorf("expected %v, got %v", e, result)
		}
	})
}
