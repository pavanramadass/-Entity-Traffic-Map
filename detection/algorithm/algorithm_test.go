package main

import "testing"
import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	"io/ioutil"
)

// TestOpenFile tests the open file function from main.go. 
func TestOpenFile(t *testing.T) {
	got, file := openFile("test.json") 
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted) 
	}
}

// TestWriteFile tests the write file function from main.go. 
func TestWriteFile(t *testing.T) {
	got, file := writeFile("test.json")
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

// TestOpenVideoCapture tests the open video capture function from main.go.
func TestOpenVideoCapture(t *testing.T) {
	deviceID = nil
	got := openVideoCapture(deviceID) 
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted) 
	}
}