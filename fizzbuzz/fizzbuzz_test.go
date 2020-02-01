package fizzbuzz

import "testing"

func TestGivenOneSayOne(t *testing.T) {
	var want string = "1"
	var given int = 1

	var get string = Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenTwoSayTwo(t *testing.T) {
	var want string = "2"
	var given int = 2

	var get string = Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenFourayFour(t *testing.T) {
	var want string = "4"
	var given int = 4

	var get string = Say(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}
