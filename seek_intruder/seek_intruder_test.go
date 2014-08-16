package main

import "fmt"
import "testing"

func Test_findIPs(t *testing.T) {
	tests := []struct {
		in       string
		expected []string
	}{
		{"9Pl.1011001111001011000001011100001,Y,HNAiSzL;?BU_UQlCvyzRU^\"R]{kVJ\"[+3%PK`]\"V?;Y'8CjJ<&QGmESP6W7&P,@$tFtL",
			[]string{"1011001111001011000001011100001"}},
		{"VA,8Z%z-AYzp6o{qeX3Q|\\`Zw7{78:Y80qP-,b0BDVvZh60x59.0xe5.0x82.0xe1uptW8eF8C]nKJ9c(AtXa9>Dy}nF'Jr",
			[]string{"0x59.0xe5.0x82.0xe1"}},
		{"`z6DR}/>gLfLX[1&]Vr8\"EG-_+wy?sw4beHIp^oTtZzvWBwY{[89.229.130.225R,?B;\"?[ix4^9D$fVaJ_V\\)N`B",
			[]string{"89.229.130.225"}},
		{"HydTS#@@864.599.341.917Q&.DVb\ails}C101100100110010011011000101110100&@KyFK\"7}u3\\63?u][~zz>-r$_OqUbE*uv,3puL.A'].BA>reos{U0x360.0x257.0x155.0x395y^UGXh^'`|I.CV}R>a}RAhO%Vw",
			[]string{"10110010011001001101100010111010"}},
	}
	for _, test := range tests {
		actual := findIPs(test.in)
		if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", test.expected) {
			t.Errorf("findIPs(%s): expected %s, got %s", test.in, test.expected, actual)
		}
	}
}
