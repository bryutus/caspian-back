package handler

import "testing"

func TestIsNumeric(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		want error
	}{
		{"文字列の0を渡してnilが返却されること", "0", nil},
		{"文字列の1を渡してnilが返却されること", "1", nil},
		{"文字列の10を渡してnilが返却されること", "10", nil},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := isNumeric(test.in); got != test.want {
				t.Errorf("isNumeric(\"%v\"): got [%v] want [%v]", test.in, got, test.want)
			}
		})
	}

	errTestCases := []struct {
		desc string
		in   string
		want string
	}{
		{"空文字列を渡してエラーメッセージが返却されること", "", "limit: Invalid value `` is specified"},
		{"文字列のGを渡してエラーメッセージが返却されること", "G", "limit: Invalid value `G` is specified"},
		{"文字列の_を渡してエラーメッセージが返却されること", "_", "limit: Invalid value `_` is specified"},
	}
	for _, test := range errTestCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := isNumeric(test.in); got.Error() != test.want {
				t.Errorf("isNumeric(\"%v\"): got [%v] want [%v]", test.in, got, test.want)
			}
		})
	}

}
