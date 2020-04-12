package errors

type NoImplementation struct {
	Message string
	Cause   error
}

func (s *NoImplementation) Error() string {
	return s.Message
}
