package resterrors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("This is the message.")

	if err == nil {
		t.Error("Error shouldn't be nil")
	}

	if err.Status != http.StatusBadRequest {
		t.Error("Error should have status as Bad Request.")
	}

	if err.Message == "This is the message." {
		t.Error("Error should have the message as: This is the message.")
	}

	if err.Error == "bad_request" {
		t.Error("Error should have Error as: bad_request.")
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("This is the message.")

	if err == nil {
		t.Error("Error shouldn't be nil")
	}

	if err.Status != http.StatusNotFound {
		t.Error("Error should have status as Not Found.")
	}

	if err.Message == "This is the message." {
		t.Error("Error should have the message as: This is the message.")
	}

	if err.Error == "not_found" {
		t.Error("Error should have Error as: not_found.")
	}
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is the message.", errors.New("database error"))

	if err == nil {
		t.Error("Error shouldn't be nil")
	}

	if err.Status != http.StatusInternalServerError {
		t.Error("Error should have status as Internal Server Error.")
	}

	if err.Message == "This is the message." {
		t.Error("Error should have the message as: This is the message.")
	}

	if err.Error == "internal_server_error" {
		t.Error("Error should have Error as: internal_server_error.")
	}

	if err.Causes != nil {
		t.Error("Error shoudn't have a Causes as nil.")
	}

	if len(err.Causes) != 1 {
		t.Error("Error Causes should have a length of 1.")
	}

	if err.Causes[0] != "database error" {
		t.Error("First Error Cause should be: database error.")
	}
}
