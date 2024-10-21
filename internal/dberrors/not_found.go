package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     string
}

// Create new instance of Error so we don't have to leak DB logic into web layer
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("could not find %s with ID %s", e.Entity, e.ID)
}
