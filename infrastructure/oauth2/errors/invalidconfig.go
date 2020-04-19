package errors

type InvalidConfigError struct {
	Message string
	Cause   error
}

func (s *InvalidConfigError) Error() string {
	return s.Message
}
