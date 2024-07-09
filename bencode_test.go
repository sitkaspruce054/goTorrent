package main

import (
	"bytes"
	"testing"
)

func TestBencodeInt(t *testing.T) {
	bf := new(bytes.Buffer)

	tst_num := 42
	bEncode(bf, tst_num)
	print(bf)
	t.Log(bf)
	print("hello")
}

func TestBencodeString(t *testing.T) {
	bf := new(bytes.Buffer)

	tst_str := "testing this string"

	bEncode(bf, tst_str)
	t.Log(bf)

}

func TestBencodeDict(t *testing.T) {
	bf := new(bytes.Buffer)

	dict := make(map[string]interface{})

	dict["3"] = 45
	dict["announce"] = "bruh"

	bEncode(bf, dict)
	t.Log(bf)
}
