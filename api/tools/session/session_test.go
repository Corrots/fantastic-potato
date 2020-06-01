package session

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	b, err := newSessionID(32)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", b)

}
