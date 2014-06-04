package main

import (
	"gobase"
	"os"
	"testing"
	"time"
)

var sleepTime, _ = time.ParseDuration("1000ms")

var path string = os.TempDir() + "/touch_test.dat"

func removeIfPresent(s string) error {
	if gobase.FileExists(s) {
		return os.Remove(s)
	}
	return nil
}

func TestTouchCreatesNewFile(t *testing.T) {
	removeIfPresent(path)
	Touch(path)
	if !gobase.FileExists(path) {
		t.Errorf("file not created %s", path)
	}
	removeIfPresent(path)
}

func TestTouchUpdatesTimestamp(t *testing.T) {
	if !gobase.FileExists(path) {
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
