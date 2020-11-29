/*Package webcolors Utility functions for working with the color names and color value
formats defined by the HTML and CSS specifications for use in
documents on the Web.

See documentation in godoc for
details of the supported formats, conventions and conversions.*/
package webcolors

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	// HTML4 html4 spec
	HTML4 = "html4"
	// CSS2 css2 spec
	CSS2 = "css2"
	// CSS21 css21 spec
	CSS21 = "css21"
	// CSS3 css3 spec
	CSS3 = "css3"
)

// SupportedSpecifications supported specifications
var SupportedSpecifications = []string{HTML4, CSS2, CSS21, CSS3}

// HexColorRegex a regexp for hex colors
var HexColorRegex = regexp.MustCompile(`^#([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$`)

// reverseMap Internal helper for generating reverse mappings; given a
// dictionary, returns a new dictionary with keys and values swapped.
func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

// HTML4NamesToHex mapping of html4 color names to hex colors
//
// The HTML 4 named colors.
//
// The canonical source for these color definitions is the HTML 4
// specification:
//
// http://www.w3.org/TR/html401/types.html#h-6.5
var HTML4NamesToHex = map[string]string{
	"aqua":    "#00ffff",
	"black":   "#000000",
	"blue":    "#0000ff",
	"fuchsia": "#ff00ff",
	"green":   "#008000",
	"gray":    "#808080",
	"grey":    "#808080",
	"lime":    "#00ff00",
	"maroon":  "#800000",
	"navy":    "#000080",
	"olive":   "#808000",
	"purple":  "#800080",
	"red":     "#ff0000",
	"silver":  "#c0c0c0",
	"teal":    "#008080",
	"white":   "#ffffff",
	"yellow":  "#ffff00",
}

// CSS2NamesToHex mapping of css2 color names to hex colors
//
// CSS 2 used the same list as HTML 4.
var CSS2NamesToHex = HTML4NamesToHex

// CSS21NamesToHex mapping of css21 color names to hex colors
//
// CSS 2.1 added orange.
var CSS21NamesToHex = make(map[string]string) // initialized in init()

// CSS3NamesToHex mapping of css3 color names to hex colors
//
// The CSS 3/SVG named colors.
//
// The canonical source for these color definitions is the SVG
// specification's color list (which was adopted as CSS 3's color
// definition):
//
// http://www.w3.org/TR/SVG11/types.html#ColorKeywords
//
// CSS 3 also provides definitions of these colors:
//
// http://www.w3.org/TR/css3-color/#svg-color
//
// SVG provides the definitions as RGB triplets. CSS 3 provides them
// both as RGB triplets and as hexadecimal. Since hex values are more
// common in real-world HTML and CSS, the mapping below is to hex
// values instead.
var CSS3NamesToHex = map[string]string{
	"aliceblue":            "#f0f8ff",
	"antiquewhite":         "#faebd7",
	"aqua":                 "#00ffff",
	"aquamarine":           "#7fffd4",
	"azure":                "#f0ffff",
	"beige":                "#f5f5dc",
	"bisque":               "#ffe4c4",
	"black":                "#000000",
	"blanchedalmond":       "#ffebcd",
	"blue":                 "#0000ff",
	"blueviolet":           "#8a2be2",
	"brown":                "#a52a2a",
	"burlywood":            "#deb887",
	"cadetblue":            "#5f9ea0",
	"chartreuse":           "#7fff00",
	"chocolate":            "#d2691e",
	"coral":                "#ff7f50",
	"cornflowerblue":       "#6495ed",
	"cornsilk":             "#fff8dc",
	"crimson":              "#dc143c",
	"cyan":                 "#00ffff",
	"darkblue":             "#00008b",
	"darkcyan":             "#008b8b",
	"darkgoldenrod":        "#b8860b",
	"darkgray":             "#a9a9a9",
	"darkgrey":             "#a9a9a9",
	"darkgreen":            "#006400",
	"darkkhaki":            "#bdb76b",
	"darkmagenta":          "#8b008b",
	"darkolivegreen":       "#556b2f",
	"darkorange":           "#ff8c00",
	"darkorchid":           "#9932cc",
	"darkred":              "#8b0000",
	"darksalmon":           "#e9967a",
	"darkseagreen":         "#8fbc8f",
	"darkslateblue":        "#483d8b",
	"darkslategray":        "#2f4f4f",
	"darkslategrey":        "#2f4f4f",
	"darkturquoise":        "#00ced1",
	"darkviolet":           "#9400d3",
	"deeppink":             "#ff1493",
	"deepskyblue":          "#00bfff",
	"dimgray":              "#696969",
	"dimgrey":              "#696969",
	"dodgerblue":           "#1e90ff",
	"firebrick":            "#b22222",
	"floralwhite":          "#fffaf0",
	"forestgreen":          "#228b22",
	"fuchsia":              "#ff00ff",
	"gainsboro":            "#dcdcdc",
	"ghostwhite":           "#f8f8ff",
	"gold":                 "#ffd700",
	"goldenrod":            "#daa520",
	"gray":                 "#808080",
	"grey":                 "#808080",
	"green":                "#008000",
	"greenyellow":          "#adff2f",
	"honeydew":             "#f0fff0",
	"hotpink":              "#ff69b4",
	"indianred":            "#cd5c5c",
	"indigo":               "#4b0082",
	"ivory":                "#fffff0",
	"khaki":                "#f0e68c",
	"lavender":             "#e6e6fa",
	"lavenderblush":        "#fff0f5",
	"lawngreen":            "#7cfc00",
	"lemonchiffon":         "#fffacd",
	"lightblue":            "#add8e6",
	"lightcoral":           "#f08080",
	"lightcyan":            "#e0ffff",
	"lightgoldenrodyellow": "#fafad2",
	"lightgray":            "#d3d3d3",
	"lightgrey":            "#d3d3d3",
	"lightgreen":           "#90ee90",
	"lightpink":            "#ffb6c1",
	"lightsalmon":          "#ffa07a",
	"lightseagreen":        "#20b2aa",
	"lightskyblue":         "#87cefa",
	"lightslategray":       "#778899",
	"lightslategrey":       "#778899",
	"lightsteelblue":       "#b0c4de",
	"lightyellow":          "#ffffe0",
	"lime":                 "#00ff00",
	"limegreen":            "#32cd32",
	"linen":                "#faf0e6",
	"magenta":              "#ff00ff",
	"maroon":               "#800000",
	"mediumaquamarine":     "#66cdaa",
	"mediumblue":           "#0000cd",
	"mediumorchid":         "#ba55d3",
	"mediumpurple":         "#9370db",
	"mediumseagreen":       "#3cb371",
	"mediumslateblue":      "#7b68ee",
	"mediumspringgreen":    "#00fa9a",
	"mediumturquoise":      "#48d1cc",
	"mediumvioletred":      "#c71585",
	"midnightblue":         "#191970",
	"mintcream":            "#f5fffa",
	"mistyrose":            "#ffe4e1",
	"moccasin":             "#ffe4b5",
	"navajowhite":          "#ffdead",
	"navy":                 "#000080",
	"oldlace":              "#fdf5e6",
	"olive":                "#808000",
	"olivedrab":            "#6b8e23",
	"orange":               "#ffa500",
	"orangered":            "#ff4500",
	"orchid":               "#da70d6",
	"palegoldenrod":        "#eee8aa",
	"palegreen":            "#98fb98",
	"paleturquoise":        "#afeeee",
	"palevioletred":        "#db7093",
	"papayawhip":           "#ffefd5",
	"peachpuff":            "#ffdab9",
	"peru":                 "#cd853f",
	"pink":                 "#ffc0cb",
	"plum":                 "#dda0dd",
	"powderblue":           "#b0e0e6",
	"purple":               "#800080",
	"red":                  "#ff0000",
	"rosybrown":            "#bc8f8f",
	"royalblue":            "#4169e1",
	"saddlebrown":          "#8b4513",
	"salmon":               "#fa8072",
	"sandybrown":           "#f4a460",
	"seagreen":             "#2e8b57",
	"seashell":             "#fff5ee",
	"sienna":               "#a0522d",
	"silver":               "#c0c0c0",
	"skyblue":              "#87ceeb",
	"slateblue":            "#6a5acd",
	"slategray":            "#708090",
	"slategrey":            "#708090",
	"snow":                 "#fffafa",
	"springgreen":          "#00ff7f",
	"steelblue":            "#4682b4",
	"tan":                  "#d2b48c",
	"teal":                 "#008080",
	"thistle":              "#d8bfd8",
	"tomato":               "#ff6347",
	"turquoise":            "#40e0d0",
	"violet":               "#ee82ee",
	"wheat":                "#f5deb3",
	"white":                "#ffffff",
	"whitesmoke":           "#f5f5f5",
	"yellow":               "#ffff00",
	"yellowgreen":          "#9acd32",
}

// # Mappings of Normalized hexadecimal color Values to color Names.
// #################################################################

// HTML4HexToNames html4 color map of hex color values to color names
var HTML4HexToNames = reverseMap(HTML4NamesToHex)

// CSS2HexToNames css2 color map of hex color values to color names
var CSS2HexToNames = HTML4HexToNames

// CSS21HexToNames css21 color map of hex color values to color names
var CSS21HexToNames = reverseMap(CSS21NamesToHex)

// CSS3HexToNames css3 color map of hex color values to color names
var CSS3HexToNames = reverseMap(CSS3NamesToHex)

func init() {
	// copy map
	for k, v := range HTML4NamesToHex {
		CSS21NamesToHex[k] = v
	}
	// add orange
	CSS21NamesToHex["orange"] = "#ffa500"

	// CSS3 defines both 'gray' and 'grey', as well as defining either
	// variant for other related colors like 'darkgray'/'darkgrey'. For a
	// 'forward' lookup from name to hex, this is straightforward, but a
	// 'reverse' lookup from hex to name requires picking one spelling.
	//
	// The way in which reverseMap() generates the reverse mappings will
	// pick a spelling based on the ordering of map keys,
	// here we manually pick a single spelling that will
	// consistently be returned. Since 'gray' was the only spelling
	// supported in HTML 4, CSS1, and CSS2, 'gray' and its varients are
	// chosen.
	CSS3HexToNames["#a9a9a9"] = "darkgray"
	CSS3HexToNames["#2f4f4f"] = "darkslategray"
	CSS3HexToNames["#696969"] = "dimgray"
	CSS3HexToNames["#808080"] = "gray"
	CSS3HexToNames["#d3d3d3"] = "lightgray"
	CSS3HexToNames["#778899"] = "lightslategray"
	CSS3HexToNames["#708090"] = "slategray"
}

// Normalization routines.
// #################################################################

// NormalizeHex Normalize a hexadecimal color value to 6 digits, lowercase.
func NormalizeHex(HexValue string) string {
	hexDigits := HexColorRegex.FindStringSubmatch(HexValue)[1]
	if len(hexDigits) == 3 {
		finalhex := []string{}
		for i := range hexDigits {
			finalhex = append(finalhex, strings.Repeat(string(hexDigits[i]), 2))
		}
		return "#" + strings.ToLower(strings.Join(finalhex, ""))
	}
	return "#" + strings.ToLower(hexDigits)
}

// NormalizeIntegerTriplet Normalize an integer rgb triplet so that all values are within the range 0-255 inclusive.
func NormalizeIntegerTriplet(RGBTriplet []int) []int {
	integerTriplet := []int{normalizeIntegerRGB(RGBTriplet[0]), normalizeIntegerRGB(RGBTriplet[1]), normalizeIntegerRGB(RGBTriplet[2])}
	return integerTriplet
}

// normalizeIntegerRGB Normalize value for use in an integer rgb triplet
func normalizeIntegerRGB(value int) int {
	if value >= 0 && value <= 255 {
		return value
	} else if value < 0 {
		return 0
	} else if value > 255 {
		return 255
	}
	return 0
}

// normalizePercentRGB Normalize value for use in a percentage rgb triplet
func normalizePercentRGB(value string) (string, error) {
	var percent = strings.Split(value, "%")[0]
	var err error

	if strings.Contains(percent, ".") {
		var percentf float64
		percentf, err = strconv.ParseFloat(percent, 64)
		if err == nil {
			if percentf >= 0 && percentf <= 100 {
				return strconv.FormatFloat(percentf, 'g', 4, 64) + "%", nil
			} else if percentf < 0 {
				return "0%", nil
			} else if percentf > 100 {
				return "100%", nil
			}
		}
	} else {
		var percenti int
		percenti, err = strconv.Atoi(percent)
		if err == nil {
			if percenti >= 0 && percenti <= 100 {
				return strconv.Itoa(percenti) + "%", nil
			} else if percenti < 0 {
				return "0%", nil
			} else if percenti > 100 {
				return "100%", nil
			}
		}
	}
	return "", err
}

// NormalizePercentTriplet Normalize a percentage rgb triplet to that all values are within the range 0%-100% inclusive.
func NormalizePercentTriplet(rgbTriplet []string) ([]string, error) {
	finalTriplet := []string{}
	for i := range rgbTriplet {
		np, err := normalizePercentRGB(rgbTriplet[i])
		if err != nil {
			return nil, err
		}
		finalTriplet = append(finalTriplet, np)
	}
	return finalTriplet, nil
}

//# Conversions from color Names to various formats.
// #################################################################

// contains checks if a array of strings contains a string
func contains(s []string, e string) bool {
	for _, output := range s {
		if output == e {
			return true
		}
	}
	return false
}

// NameToHex Convert a color name to a normalized hexadecimal color value
func NameToHex(name string, spec string) (string, error) {
	if contains(SupportedSpecifications, spec) == true {
		normalized := strings.ToLower(name)
		var hexStr string
		var ok bool
		switch spec {
		case HTML4:
			hexStr, ok = HTML4NamesToHex[normalized]
		case CSS2:
			hexStr, ok = CSS2NamesToHex[normalized]
		case CSS21:
			hexStr, ok = CSS21NamesToHex[normalized]
		case CSS3:
			hexStr, ok = CSS3NamesToHex[normalized]
		default:
			return "", errors.New(spec + "is not output supported Specification for color name lookups")
		}
		if !ok {
			return "", errors.New(name + "has no defined color name in " + spec)
		}
		return hexStr, nil
	}
	return "", errors.New(spec + "is not output supported Specification for color name lookups")
}

// NameToRGB Convert a color name to a 3-tuple of integers suitable for use in an rgb triplet specifying that color
func NameToRGB(name string, spec string) ([]int, error) {
	hx, err := NameToHex(name, spec)
	if err != nil {
		return []int{}, err
	}
	return HexToRGB(hx)
}

// NameToRGBPercent Convert a color name to a 3-tuple of percentages suitable for use in an rgb triplet specifying that color
func NameToRGBPercent(name string, spec string) ([]string, error) {
	rgb, err := NameToRGB(name, spec)
	if err != nil {
		return []string{}, err
	}
	return RGBToRGBPercent(rgb)
}

// # Conversions from hexadecimal color Values to various formats.
// #################################################################

// HexToName Convert a hexadecimal color value to its corresponding normalized color name, if any such name exists
func HexToName(hexValue string, spec string) (string, error) {
	if contains(SupportedSpecifications, spec) == true {
		normalized := NormalizeHex(hexValue)
		var name string
		var ok bool
		switch spec {
		case HTML4:
			name, ok = HTML4HexToNames[normalized]
		case CSS2:
			name, ok = CSS2HexToNames[normalized]
		case CSS21:
			name, ok = CSS21HexToNames[normalized]
		case CSS3:
			name, ok = CSS3HexToNames[normalized]
		default:
			return "", errors.New(spec + "is not output supported Specification for color name lookups")
		}
		if !ok {
			return "", errors.New(hexValue + "has no defined color name in " + spec)
		}
		return name, nil
	}
	return "", errors.New(spec + "is not output supported Specification for color Name lookups")
}

// ByteToInt converts a hex bytearray to hex integer
func ByteToInt(input []byte) int {
	var output uint32
	l := len(input)
	for i, b := range input {
		shift := uint32((l - i - 1) * 8)
		output |= uint32(b) << shift
	}
	return int(output)
}

// HexToRGB Convert a hexadecimal color value to a 3-tuple of integers suitable for use in an rgb triplet specifying that color
func HexToRGB(hexValue string) ([]int, error) {
	hexDigits := NormalizeHex(hexValue)
	rgbTuple := []int{}
	partialHex1, err := hex.DecodeString(hexDigits[1:3])
	if err != nil {
		return rgbTuple, err
	}
	rgbTuple = append(rgbTuple, ByteToInt(partialHex1))
	partialHex2, err := hex.DecodeString(hexDigits[3:5])
	if err != nil {
		return rgbTuple, err
	}
	rgbTuple = append(rgbTuple, ByteToInt(partialHex2))
	partialHex3, err := hex.DecodeString(hexDigits[5:7])
	if err != nil {
		return rgbTuple, err
	}
	rgbTuple = append(rgbTuple, ByteToInt(partialHex3))
	return rgbTuple, err
}

// HexToRGBPercent Convert a hexadecimal color value to a 3-tuple of percentages suitable for use in an rgb triplet representing that color
func HexToRGBPercent(hexValue string) ([]string, error) {
	hx, err := HexToRGB(hexValue)
	if err != nil {
		return []string{}, err
	}
	return RGBToRGBPercent(hx)
}

// # Conversions from  integer rgb() triplets to various formats.
// #################################################################

// RGBToName Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to its corresponding normalized color name, if any such name exists
func RGBToName(rgbTriplet []int, spec string) (string, error) {
	return HexToName(RGBToHex(NormalizeIntegerTriplet(rgbTriplet)), spec)
}

// RGBToHex Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to a normalized hexadecimal value for that color
func RGBToHex(rgbTriplet []int) string {
	integerTriplet := NormalizeIntegerTriplet(rgbTriplet)
	hexString := "#"
	for i := range integerTriplet {
		byteCoded := make([]byte, 2)
		binary.BigEndian.PutUint16(byteCoded, uint16(integerTriplet[i]))
		hexString = hexString + hex.EncodeToString(byteCoded[1:2])
	}
	return hexString
}

// RGBToRGBPercent Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to a 3-tuple of percentages suitable for use in representing that color
func RGBToRGBPercent(rgbTriplet []int) ([]string, error) {
	specials := map[int]string{
		255: "100%",
		128: "50%",
		64:  "25%",
		32:  "12.50%",
		16:  "6.25%",
		0:   "0%",
	}
	rgbPercentTriplet := []string{}
	normalizedTriplet := NormalizeIntegerTriplet(rgbTriplet)

	for i := range normalizedTriplet {
		if name, ok := specials[normalizedTriplet[i]]; ok {
			rgbPercentTriplet = append(rgbPercentTriplet, name)
		} else {
			percentVal := (float64(normalizedTriplet[i]) / 255.0) * 100
			rgbPercentTriplet = append(rgbPercentTriplet, strconv.FormatFloat(percentVal, 'g', 4, 64)+"%")
		}
	}
	return rgbPercentTriplet, nil
}

// # Conversions from Percentage rgb() triplets to various formats.
// #################################################################

// RGBPercentToName Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to its corresponding normalized color name, if any such name exists
func RGBPercentToName(rgbPercentTriplet []string, spec string) (string, error) {
	npt, err := NormalizePercentTriplet(rgbPercentTriplet)
	if err != nil {
		return "", err
	}
	rgb, err := RGBPercentToRGB(npt)
	if err != nil {
		return "", err
	}
	return RGBToName(rgb, spec)
}

// RGBPercentToHex Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to a normalized hexadecimal color value for that color
func RGBPercentToHex(rgbPercentTriplet []string) (string, error) {
	npt, err := NormalizePercentTriplet(rgbPercentTriplet)
	if err != nil {
		return "", err
	}
	rgb, err := RGBPercentToRGB(npt)
	if err != nil {
		return "", err
	}
	return RGBToHex(rgb), nil
}

// percentToInteger Internal helper for converting a percentage value to an integer between 0 and 255 inclusive
func percentToInteger(percent string) (int, error) {
	num, err := strconv.ParseFloat(strings.Split(percent, "%")[0], 64)
	if err == nil {
		num = 255 * (num / 100.0)
		e := num - math.Floor(num)
		if e < 0.5 {
			return int(math.Floor(num)), nil
		}
		return int(math.Ceil(num)), nil
	}
	return 0, err
}

// RGBPercentToRGB Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to a 3-tuple of integers suitable for use in representing that color
func RGBPercentToRGB(rgbPercentTriplet []string) ([]int, error) {
	rgbTriplet := []int{}
	normalizedTriplet, err := NormalizePercentTriplet(rgbPercentTriplet)
	if err != nil {
		return rgbTriplet, err
	}
	for i := range normalizedTriplet {
		perI, err := percentToInteger(normalizedTriplet[i])
		if err != nil {
			return rgbTriplet, err
		}
		rgbTriplet = append(rgbTriplet, perI)
	}
	return rgbTriplet, nil
}
