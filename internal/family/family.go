package family

import "errors"

type Relationship string

const (
	Father      Relationship = "father"
	Mother      Relationship = "mother"
	Child       Relationship = "child"
	GrandMother Relationship = "grandMother"
	GrandFather Relationship = "grandFather"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Family struct {
	Members map[Relationship]Person
}

var (
	// ErrRelationshipAlreadyExists возвращает ошибку если роль уже занята
	ErrRelationshipAlreadyExists = errors.New("relationship already exists")
)

// AddNew добавляет нового члена семьи. Если в семье еще нет людей, создается пустая мапа.
// Если роль уже занята, метод выдает ошибку.
func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; !ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[r] = p
	return nil
}
