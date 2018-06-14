package main

import (
	"bytes"
	"testing"
)

const testFullResult = `test
├───dir1
│	├───dir1_1
│	│	└───text1_1_1.txt (665b)
│	└───text1_1.txt (835b)
├───dir2
│	└───text2_1.txt (empty)
└───text1.txt (551b)
`

func TestFullTree(t *testing.T) {
	out := new(bytes.Buffer)
	err := tree(out, "test", "", true)
	if err != nil {
		t.Errorf("Full tree test failed, error:\n%v", err)
	}

	result := out.String()
	if result != testFullResult {
		t.Errorf("Full tree test failed, results don't match:\nGot:\n%v\nWant:\n%v", result, testFullResult)
	}
}

const testDirResult = `test
├───dir1
│	└───dir1_1
└───dir2
`

func TestDirTree(t *testing.T) {
	out := new(bytes.Buffer)
	err := tree(out, "test", "", false)
	if err != nil {
		t.Errorf("Dir tree test failed, error:\n%v", err)
	}

	result := out.String()
	if result != testDirResult {
		t.Errorf("Dir tree test failed, results don't match:\nGot:\n%v\nWant:\n%v", result, testDirResult)
	}
}
