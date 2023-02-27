package uuid

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func Test_GetNonce(t *testing.T) {
	uid := GetUUID()
	t.Log("uid:", uid)
	AssertEqual(t, 32, len(uid))
}
