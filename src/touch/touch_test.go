package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var sleepTime, _ = time.ParseDuration("1000ms")

var path string = filepath.Join(os.TempDir(), "touch_test.dat")

// FileExists returns true if path points to an existing file, and false
// otherwise.
func fileExists(path string) bool {
	fi, _ := os.Stat(path)
	return fi != nil
}

func removeIfPresent(s string) error {
	if fileExists(s) {
		return os.Remove(s)
	}
	return nil
}

func TestTouchCreatesNewFile(t *testing.T) {
	removeIfPresent(path)
	Touch(path)
	if !fileExists(path) {
		t.Errorf("file not created %s", path)
	}
	removeIfPresent(path)
}

func TestTouchUpdatesTimestamp(t *testing.T) {
	if !fileExists(path) {
		Touch(path)
	}
	before, _ := os.Stat(path)
	btime := before.ModTime().UnixNano()
	time.Sleep(sleepTime)
	er := Touch(path)
	if er != nil {
		t.Errorf("error returned by Touch(path) %s", er)
	}
	after, _ := os.Stat(path)
	if btime >= after.ModTime().UnixNano() {
		t.Errorf("access time not updated")
	}
	removeIfPresent(path)
}
