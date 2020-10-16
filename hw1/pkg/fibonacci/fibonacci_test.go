package fibonacci

import "testing"

func TestFib(t *testing.T) {
	data := make(map[int]int)
	data[1] = 0
	data[5] = 3
	data[10] = 34
	for k, v := range data {
		got, _ := Fib(k)
		want := v
		if got != want {
			t.Errorf("Got: %d; want: %d", got, want)
		}
	}
}
