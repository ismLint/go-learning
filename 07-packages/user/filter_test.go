package user

import "testing"

func TestFilterAdults(t *testing.T) {
	tests := []struct {
		name  string
		users []User
		want  []User
	}{
		{
			name:  "empty slice",
			users: []User{},
			want:  []User{},
		},
		{
			name: "no adults",
			users: []User{
				{Name: "James", Age: 17},
				{Name: "Mike", Age: 12},
			},
			want: []User{},
		},
		{
			name: "normal case",
			users: []User{
				{Name: "Nick", Age: 25},
				{Name: "James", Age: 17},
				{Name: "Anna", Age: 18},
			},
			want: []User{
				{Name: "Nick", Age: 25},
				{Name: "Anna", Age: 18},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original := make([]User, len(test.users))
			copy(original, test.users)

			got := FilterAdults(test.users)

			if len(got) != len(test.want) {
				t.Fatalf("got %d users, want %d", len(got), len(test.want))
			}

			for index := range test.want {
				if got[index].Name != test.want[index].Name {
					t.Errorf("at index %d: got name %q, want %q", index, got[index].Name, test.want[index].Name)
				}

				if got[index].Age != test.want[index].Age {
					t.Errorf("at index %d: got age %d, want %d", index, got[index].Age, test.want[index].Age)
				}
			}

			for index := range original {
				if test.users[index] != original[index] {
					t.Errorf("original slice was modified at index %d: got %v, want %v",
						index, test.users[index], original[index])
				}
			}
		})
	}
}
