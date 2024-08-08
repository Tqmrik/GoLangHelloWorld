package main

import (
	"errors"
)

func validateStatus(status string) error {
	switch {
	case status == "":
		return errors.New("status cannot be empty")
	case len(status) > 140:
		return errors.New("status exceeds 140 characters")
	default:
		return nil
	}

}
