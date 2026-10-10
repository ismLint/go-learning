package user

import "testing"

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr string
	}{
		{name: "negative", age: -1, wantErr: "age cannot be negative"},
		{name: "havn't error 1", age: 0, wantErr: ""},
		{name: "havn't error 2", age: 18, wantErr: ""},
		{name: "havn't error 3", age: 150, wantErr: ""},
		{name: "age is too large", age: 151, wantErr: "age is too large"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ValidateAge(test.age)
			gotErr := ""
			if got != nil {
				gotErr = got.Error()
			}

			if gotErr != test.wantErr {
				t.Errorf("got %v, want %v", gotErr, test.wantErr)
			}
		})
	}
}
