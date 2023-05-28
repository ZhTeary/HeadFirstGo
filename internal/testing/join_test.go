package testing

/*
	_test.go 为结尾
*/
import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{[]string{""}, ""},
		{[]string{"apple"}, "apple"},
		{[]string{"apple", "orange"}, "apple and orange"},
		{[]string{"apple", "orange", "banana"}, "apple, orange, and banana"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) =\"%s\", want \"%s\"", list, got, want)
}
