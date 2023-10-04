package main

import (
	"testing"
)

func HandleTestError(t *testing.T, err error, logType string) {
	if err != nil {
		switch logType {
		case "log":
			t.Log(err)
		case "error":
			t.Error(err)
		case "fatal":
			t.Fatal(err)
		default:
			t.Error(err)
		}
	}
}
