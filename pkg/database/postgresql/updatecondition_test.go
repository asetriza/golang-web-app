package postgresql

import "testing"

type TestStruct struct {
	Name     *string `json:"name" db:"name"`
	Username *string `json:"username" db:"username"`
	Password *string `json:"password" db:"password"`
}

func TestUpdateConditionFromStruct(t *testing.T) {
	name := "name"
	username := "username"
	password := "password"
	testStruct := &TestStruct{
		Name:     &name,
		Username: &username,
		Password: &password,
	}

	got := UpdateConditionFromStruct(testStruct)
	want := "name = :name, username = :username, password = :password"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
