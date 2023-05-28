package testing

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Error(errorStringLarger(want, got))
	}
}
func TestSecondLarger(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if got != want {
		t.Error(errorStringLarger(want, got))
	}
}
func errorStringLarger(want int, got int) string {
	return fmt.Sprintf("Testing func Larger, want: %d , got: %d", want, got)
}
