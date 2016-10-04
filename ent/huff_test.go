package ent

import (
	"testing"
)

// thanks to @spheenik and @invokr for the expected huffman codes. these are
// ripped from the huffman trees that are known to be working in clarity and
// manta.
var expected_codes = map[string]string{
	"PlusOne":               "0",
	"FieldPathEncodeFinish": "10",
	"PlusTwo":               "1110",
	"PushOneLeftDeltaNRightNonZeroPack6Bits": "1111",
	"PushOneLeftDeltaOneRightNonZero":        "11000",
	"PlusN":                                  "11010",
	"PlusThree":                              "110010",
	"PopAllButOnePlusOne":                    "110011",
	"PushOneLeftDeltaNRightNonZero":          "11011001",
	"PushOneLeftDeltaOneRightZero":           "11011010",
	"PushOneLeftDeltaNRightZero":             "11011100",
	"PopAllButOnePlusNPack6Bits":             "11011110",
	"PlusFour":                               "11011111",
	"PopAllButOnePlusN":                      "110110000",
	"PushOneLeftDeltaNRightNonZeroPack8Bits": "110110110",
	"NonTopoPenultimatePlusOne":              "110110111",
	"PopAllButOnePlusNPack3Bits":             "110111010",
	"PushNAndNonTopological":                 "110111011",
	"NonTopoComplexPack4Bits":                "1101100010",
	"NonTopoComplex":                         "11011000111",
	"PushOneLeftDeltaZeroRightZero":          "110110001101",
	"PopOnePlusOne":                          "110110001100001",
	"PushOneLeftDeltaZeroRightNonZero":       "110110001100101",
	"PopNAndNonTopographical":                "1101100011000000",
	"PopNPlusN":                              "1101100011000001",
	"PushN":                                  "1101100011000100",
	"PushThreePack5LeftDeltaN":               "1101100011000101",
	"PopNPlusOne":                            "1101100011000110",
	"PopOnePlusN":                            "1101100011000111",
	"PushTwoLeftDeltaZero":                   "1101100011001000",
	"PushThreeLeftDeltaZero":                 "11011000110010010",
	"PushTwoPack5LeftDeltaZero":              "11011000110010011",
	"PushTwoLeftDeltaN":                      "11011000110011000",
	"PushThreePack5LeftDeltaOne":             "11011000110011001",
	"PushThreeLeftDeltaN":                    "11011000110011010",
	"PushTwoPack5LeftDeltaN":                 "11011000110011011",
	"PushTwoLeftDeltaOne":                    "11011000110011100",
	"PushThreePack5LeftDeltaZero":            "11011000110011101",
	"PushThreeLeftDeltaOne":                  "11011000110011110",
	"PushTwoPack5LeftDeltaOne":               "11011000110011111",
}

func TestTree(t *testing.T) {
	var testWalk func(node, string)
	testWalk = func(n node, code string) {
		switch v := n.(type) {
		case lNode:
			if expected_codes[v.name] != code {
				t.Errorf("op %s has code %s, expected %s", v.name, code, expected_codes[v.name])
			} else {
				t.Logf("op %s has expected code %s", v.name, code)
			}
		case iNode:
			testWalk(v.left, code+"0")
			testWalk(v.right, code+"1")
		}
	}

	testWalk(huffRoot, "")
}
