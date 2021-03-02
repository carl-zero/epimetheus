package log

import "testing"

func TestAppend(t *testing.T)  {
	name := "testwritefile.txt"
	content := "Hello, xxbandy.github.io!\n"
	logSegmentAppend(name,content)
}