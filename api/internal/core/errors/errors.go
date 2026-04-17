package errors

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound      = errors.New("record not found")
	ErrAlreadyExists = errors.New("record already exists")
)

func MapRepoError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrAlreadyExists
	}
	return err
}
