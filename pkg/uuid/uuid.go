package uuid

import "github.com/google/uuid"

type UUID interface {
	Create() string
	Parse(uuid string) error
}

func Create() string {
	return uuid.NewString()
}

func Parse(uuidString string) error {
	_, err := uuid.Parse(uuidString)
	return err
}
