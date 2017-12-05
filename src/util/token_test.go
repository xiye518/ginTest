package util

import "testing"

func TestSetToken(t *testing.T) {
	token := SetToken("xiye")
	t.Log(token)
}
