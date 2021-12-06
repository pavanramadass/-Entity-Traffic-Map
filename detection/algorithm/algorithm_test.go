package algorithm

import "testing"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// TestOpenFile tests whether the json file is openeable
func TestOpenFile(t *testing.T) {
	got, file := openFile("test.json")
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

// TestWriteFile tests whether the one can write to the json file
func TestWriteFile(t *testing.T) {
	got, file := writeFile("test.json")
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

// TestOpenVideoCapture tests whether the deviceID can capture the video footage
func TestOpenVideoCapture(t *testing.T) {
	deviceID = nil
	got := openVideoCapture(deviceID)
	wanted = true

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}
