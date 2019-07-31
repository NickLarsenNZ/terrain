package utils_test

import (
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
