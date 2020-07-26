package encrypt

import (
	"testing"
)

func Test_GetHash(t *testing.T) {
	testdata := "aaaaaaaccccvvvvv"
	hashdata := DataHash{}.GetHash(testdata)
	t.Log(hashdata)
	if hashdata == testdata {
		t.Fatal("failed GetHash ")
	}
	t.Log("success GetHash")
}
