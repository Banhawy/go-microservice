package dberrors

type ConflictError struct{}

// Create new instance of Error so we don't have to leak DB logic into web layer
func (e ConflictError) Error() string {
	return "attempted to create a record that already exists"
}
