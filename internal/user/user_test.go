package user

import "testing"

func TestUser_FullName(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "John Doe", fields: fields{FirstName: "John", LastName: "Doe"}, want: "John Doe"},
		{name: "Ekaterina Petrova", fields: fields{FirstName: "Ekaterina", LastName: "Petrova"}, want: "Ekaterina Petrova"},
		{name: "Евгений Амплеев", fields: fields{FirstName: "Евгений", LastName: "Амплеев"}, want: "Евгений Амплеев"},
		{name: "Long name", fields: fields{
			FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
			LastName:  "Picasso",
		}, want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz Picasso"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := User{FirstName: test.fields.FirstName, LastName: test.fields.LastName}
			if got := u.FullName(); got != test.want {
				t.Errorf("%s fail.. User.FullName() = %v, want %v", test.name, got, test.want)
			}
		})
	}
}
