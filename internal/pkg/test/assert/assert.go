// simple drop-in assert package
// @author jacek.wysocki@gmail.com
package assert

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func NoError(t *testing.T, err error) {
	if err := CheckNoError(err); err != nil {
		t.Fatal(err)
	}
}

func ErrorIs(t *testing.T, err error, expectedError error) {
	if err != expectedError {
		t.Fatalf("Expected error '%v' but got '%v'", expectedError, err)
	}
}

func CheckNoError(err error) error {
	if err != nil {
		return fmt.Errorf("Expected no error but got: %v", err)
	}
	return nil
}

func Error(t *testing.T, err error) {
	if err := CheckError(err); err != nil {
		t.Fatal(err)
	}
}

func CheckError(err error) error {
	if err == nil {
		return fmt.Errorf("Expected error but got: %+v", err)
	}
	return nil
}

func Assert(t *testing.T, message string, assumption bool) {
	if !assumption {
		t.Fatalf(message)
	}
}

func Check(message string, assumption bool) error {
	if !assumption {
		return fmt.Errorf(message)
	}
	return nil
}

func Equals(t *testing.T, expected, result interface{}) {
	if err := CheckEquals(expected, result); err != nil {
		t.Fatal(err)
	}
}

func CheckEquals(expected, result interface{}) error {
	if diff := cmp.Diff(expected, result); diff != "" {
		return fmt.Errorf("AssertEquals() mismatch (-want +got):\n%s", diff)
	}
	return nil
}

func Range(t *testing.T, expectedFrom, expectedTo, result int) {
	if err := CheckRange(expectedFrom, expectedTo, result); err != nil {
		t.Fatal(err)
	}
}

func CheckRange(expectedFrom, expectedTo, result int) error {
	if result < expectedFrom || result > expectedTo {
		return fmt.Errorf("Value should be between %d and %d but got %d", expectedFrom, expectedTo, result)
	}
	return nil
}

func NotEquals(t *testing.T, expected, result interface{}) {
	if err := CheckNotEquals(expected, result); err != nil {
		t.Fatal(err)
	}
}

func CheckNotEquals(expected, result interface{}) error {
	if diff := cmp.Diff(expected, result); diff == "" {
		return fmt.Errorf("AssertEquals() mismatch (-want +got):\n%s", diff)
	}
	return nil
}

func NotNil(t *testing.T, result interface{}) {
	if result == nil {
		t.Fatalf("Value should be nil but got: %s", result)
	}
}

func Nil(t *testing.T, result interface{}) {
	if result != nil {
		t.Fatalf("Value should be nil but got: %s", result)
	}
}

func GtInt(t *testing.T, expected int, result int) {
	if expected > result {
		return
	}
	t.Fatalf("AssertGtInt '%d > %d' is not true", expected, result)
}

func StringMatch(t *testing.T, result, regexpString string) {
	match, _ := regexp.MatchString(regexpString, result)
	message := fmt.Sprintf("Given string (%s) doesn't match regexp (%s)", result, regexpString)
	Assert(t, message, match)
}

func Contains(t *testing.T, result, search string) {
	if !strings.Contains(result, search) {
		Assert(t, fmt.Sprintf("Given string doesn't contain '%s' (string: %s)", search, result), false)
	}
}
