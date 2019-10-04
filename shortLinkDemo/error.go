package main

type Error interface {
	error
	Status() int
}

type StausError struct {
	Code int
	Err  error
}

func (s StausError) Error() string {
	return s.Err.Error()
}

func (s StausError) Status() int {
	return s.Code
}
