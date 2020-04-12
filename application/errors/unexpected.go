package errors

type UnexpectedError struct {
	Message string
	Cause   error
}

func (s *UnexpectedError) Error() string {
	return s.Message
}
