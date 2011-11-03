package main

import (
	"testing"
	"os"
	"time"
)

const sleepTime int64 = 500 * int64(1e6)

var path string = os.TempDir() + "/touch_test.dat"

func exist(s string) bool {
	fi, _ := os.Stat(s)
	return fi != nil
}

func removeIfPresent(s string) os.Error {
	if exist(s) {
		return os.Remove(s)
	}
	return nil
}

func TestTouchCreatesNewFile(t *testing.T) {
	removeIfPresent(path)
	Touch(path)
	if !exist(path) {
		t.Errorf("file not created %s", path)
	}
	removeIfPresent(path)
}

func TestTouchUpdatesTimestamp(t *testing.T) {
	if !exist(path) {
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
	if before.Mtime_ns >= after.Mtime_ns {
		t.Errorf("modified time not updated")
	}
	removeIfPresent(path)
}
