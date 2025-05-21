package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUser_FullName - тест с использованием стандартной библиотеки.
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

// TestUser_FullName - тест с использованием сторонней библиотеки testify.
func TestUser_FullName2(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple test",
			fields: fields{
				FirstName: "Misha",
				LastName:  "Popov",
			},
			want: "Misha Popov",
		},
		{
			name: "long name",
			fields: fields{
				FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan" +
					" Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
				LastName: "Picasso",
			},
			want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno" +
				" Krispin Krispiano de la Santisima Trinidad Ruiz Picasso",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			v := u.FullName()
			// как и в предыдущем тесте сроки сравниваются с помощью функции Equal
			assert.Equal(t, tt.want, v)
		})
	}
}
