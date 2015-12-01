package randstr

import "testing"

func TestRandString(t *testing.T) {
	l := 96
	str := Get(l)
	t.Logf("Generated string: %v", str)

	if len(str) != l {
		t.Fatalf("Expected a random string of length %v but got %v", l, len(str))
	}

	// this many collisions can't be
	if str[1] == str[2] && str[3] == str[4] && str[5] == str[6] && str[7] == str[8] && str[9] == str[10] {
		t.Fatal("Expected a random string, got repeated characters")
	}
}
