package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	got := NewUser("Nick", 25)

	if got.Name != "Nick" {
		t.Errorf("name = %v, want Nick", got.Name)
	}

	if got.Age != 25 {
		t.Errorf("age = %d, want 25", got.Age)
	}

}

func TestUserIsAdult(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{name: "minor", user: NewUser("Nick", 17), want: false},
		{name: "secondary", user: NewUser("James", 18), want: true},
		{name: "third", user: NewUser("Valentin", 25), want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.user.IsAdult()

			if got != test.want {
				t.Errorf("got %t, want %t", got, test.want)
			}
		})
	}
}
