package rstrings

type ArgumentError struct {
}

// No argument was given to the rstrings command.
func (e *ArgumentError) Error() string {
	return "no argument was given to the rstrings command"
}
