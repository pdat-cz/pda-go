package math

import "testing"

func TestContains(t *testing.T) {
	check := []struct {
		input   []interface{}
		element interface{}
		want    bool
	}{
		{[]interface{}{1, 2, 3}, 1, true},
		{[]interface{}{1, 2, 3}, 0, false},
		{[]interface{}{"a", "b", "c"}, "a", true},
		{[]interface{}{"a", "b", "c"}, "d", false},
	}

	for _, c := range check {
		received := Contains(c.input, c.element)
		if received != c.want {
			t.Errorf("Contains(%v,%v) != %v, received: %v", c.input, c.element, c.want, received)
		}
	}
}
