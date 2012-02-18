package gobase

import (
	"os"
	"testing"
	"time"
)

const sleepTime int64 = 500 * int64(1e6)

var path string = os.TempDir() + "/touch_test.dat"

func removeIfPresent(s string) error {
	if FileExists(s) {
		return os.Remove(s)
	}
	return nil
}

func TestTouchCreatesNewFile(t *testing.T) {
	removeIfPresent(path)
	Touch(path)
	if !FileExists(path) {
		t.Errorf("file not created %s", path)
	}
	removeIfPresent(path)
}

func TestTouchUpdatesTimestamp(t *testing.T) {
	if !FileExists(path) {
		Touch(path)
	}
	before, _ := os.Stat(path)
	time.Sleep(sleepTime)
	er := Touch(path)
	if er != nil {
		t.Errorf("error returned by Touch(path) %s", er)
	}
	after, _ := os.Stat(path)
	if before.Atime_ns >= after.Atime_ns {
		t.Errorf("access time not updated")
	}
	if before.ModTime() >= after.ModTime() {
		t.Errorf("modified time not updated")
	}
	removeIfPresent(path)
}
