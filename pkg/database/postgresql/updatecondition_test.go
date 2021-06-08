package postgresql

import "testing"

type TestStruct struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func TestUpdateConditionFromStruct(t *testing.T) {
	id := 0
	name := "name"
	username := "username"
	password := "password"
	testStruct := &TestStruct{
		ID:       id,
		Name:     name,
		Username: username,
		Password: password,
	}

	got := UpdateConditionFromStruct(testStruct)
	want := "name = :name, username = :username, password = :password"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
