package basic_programming

import (
	"reflect"
	"testing"
)

func TestGetPrimeNumbers(t *testing.T) {
	if !reflect.DeepEqual(GetPrimeNumbers(20), []int{2,3,5,7}) {
		t.Fail()
	}
}
