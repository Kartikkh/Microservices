package custom_error

type CustomerError struct {
	error
	Message string
}

func New(err error) *CustomerError {
	return &CustomerError{error: err}
}

func (c *CustomerError) SetMessage(message string) *CustomerError {
	c.Message = message
	return c
}

func (c *CustomerError) GetMessage() string {
	return c.Message
}
