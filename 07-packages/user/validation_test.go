package user

import "testing"

func TestIsAdult(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{name: "midor", age: 17, want: false},
		{name: "exactly 18", age: 18, want: true},
		{name: "adult", age: 25, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsAdult(test.age)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
