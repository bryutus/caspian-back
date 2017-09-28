package main

import "testing"

func TestParseDatetime(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		want string
	}{
		{"空文字列を渡して2006-01-02 15:04:05が返却されること", "", "2006-01-02 15:04:05"},
		{"文字列の2017-09-26T04:21:40Zを渡して2017-09-26 04:21:40が返却されること", "2017-09-26T04:21:40Z", "2017-09-26 04:21:40"},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := parseDatetime(test.in); got != test.want {
				t.Errorf("parseDatetime(\"%v\"): got [%v] want [%v]", test.in, got, test.want)
			}
		})
	}
}
