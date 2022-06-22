package exception

type NotFoundError struct {
	Message string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message, 
	}
}

func (self NotFoundError) Error() string {
	return self.Message
}
