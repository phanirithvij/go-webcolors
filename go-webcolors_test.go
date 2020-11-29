package webcolors

import "testing"

func TestNormalizeHex(t *testing.T) {
	value := NormalizeHex("#0099CC")
	if value != "#0099cc" {
		t.Error("expected #0099cc, got", value)
	}
}

func TestNormalizeIntegerTriplet(t *testing.T) {
	value := NormalizeIntegerTriplet([]int{270, -20, 128})
	expected := []int{255, 0, 128}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestNormalizeIntegerRGB(t *testing.T) {
	value := normalizeIntegerRGB(270)
	if value != 255 {
		t.Error("expected 255, got", value)
	}
}

func TestNormalizePercentTriplet(t *testing.T) {
	value, _ := NormalizePercentTriplet([]string{"-10%", "250%", "500%"})
	expected := []string{"0%", "100%", "100%"}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestNormalizePercentRGB(t *testing.T) {
	value, _ := normalizePercentRGB("-5%")
	if value != "0%" {
		t.Error("expected 0%, got", value)
	}
}

func TestNamesToHex(t *testing.T) {
	value, _ := NameToHex("white", "css3")
	if value != "#ffffff" {
		t.Error("expected white, got", value)
	}
}

func TestNameToRGB(t *testing.T) {
	value, _ := NameToRGB("navy", "css3")
	expected := []int{0, 0, 128}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestNameToRGBPercent(t *testing.T) {
	value, _ := NameToRGBPercent("navy", "css3")
	expected := []string{"0%", "0%", "50%"}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestHexToName(t *testing.T) {
	value, _ := HexToName("#daa520", "css3")
	if value != "goldenrod" {
		t.Error("expected goldenrod, got", value)
	}
}

func TestHexToRGB(t *testing.T) {
	value, _ := HexToRGB("#000080")
	expected := []int{0, 0, 128}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestHexToRGBPercent(t *testing.T) {
	value, _ := HexToRGBPercent("#000080")
	expected := []string{"0%", "0%", "50%"}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestRGBToName(t *testing.T) {
	value, _ := RGBToName([]int{0, 0, 128}, "css3")
	if value != "navy" {
		t.Error("expected navy, got", value)
	}
}

func TestRGBToHex(t *testing.T) {
	value := RGBToHex([]int{0, 0, 128})
	if value != "#000080" {
		t.Error("expected #000080, got", value)
	}
}

func TestRGBToRGBPercent(t *testing.T) {
	value, _ := RGBToRGBPercent([]int{218, 165, 32})
	expected := []string{"85.49%", "64.71%", "12.50%"}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}

func TestRGBPercentToName(t *testing.T) {
	value, _ := RGBPercentToName([]string{"85.49%", "64.71%", "12.5%"}, "css3")
	if value != "goldenrod" {
		t.Error("expected goldenrod, got", value)
	}
}

func TestRGBPercentToHex(t *testing.T) {
	value, _ := RGBPercentToHex([]string{"100%", "100%", "0%"})
	if value != "#ffff00" {
		t.Error("expected #ffff00, got", value)
	}
}

func TestRGBPercentToRGB(t *testing.T) {
	value, _ := RGBPercentToRGB([]string{"0%", "0%", "50%"})
	expected := []int{0, 0, 128}
	for i := range value {
		if value[i] != expected[i] {
			t.Error("expected", expected[i], " got", value[i])
		}
	}
}
