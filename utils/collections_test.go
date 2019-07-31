package utils_test

import (
	"errors"
	"github.com/nicklarsennz/terrain/utils"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	slice := &[]string{"one", "two", "three", "go"}
	expect_to_find := "go"
	expect_not_to_find := "something not in the slice"

	if !utils.StringInSlice(expect_to_find, slice) {
		t.Fatalf("expected to find the string: '%s' in %v", expect_to_find, *slice)
	}

	if utils.StringInSlice(expect_not_to_find, slice) {
		t.Fatalf("expected not to find the string: '%s' in %v", expect_not_to_find, *slice)
	}
}

func TestAppendError(t *testing.T) {
	errs := &[]error{
		errors.New("first error"),
		errors.New("second error"),
	}

	expected_errors := len(*errs) + 1

	utils.AppendError(errs, errors.New("third error"))

	actual_errors := len(*errs)

	if len(*errs) != expected_errors {
		t.Fatalf("expected there to be %d errors, but only see %d", expected_errors, actual_errors)
	}
}
