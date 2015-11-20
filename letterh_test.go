package letterhistogram

import  (
	"testing"
)

func TestShouldPass(t *testing.T) {
}

func TestShouldGetBackSingleLetter(t *testing.T) {
	a := Hist("a");
	if a != "a" {
		t.Fail()
	}
}

func TestShouldGetBackTwoSameLetters(t *testing.T) {
	a := Hist("aa");
	if a != "aa" {
		t.Fail()
	}
}

func TestShouldGetBackTwoDiffLetters(t *testing.T) {
	a := Hist("ba");
	if a != "ab" {
		t.Fail()
	}
}

func TestAab(t *testing.T) {
	if Hist("aba") != "aab" {
		t.Fail()
	}
}

func TestAbc(t *testing.T) {
	if Hist("abc") != "abc" {
		t.Fail()
	}
}


