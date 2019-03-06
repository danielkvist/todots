package copier

import (
	"os"
	"testing"
)

func TestCopyFiles(t *testing.T) {
	src := "test.txt"
	content := "Hello, Gophers!\n"

	createSrcFile(t, src, content)

	dst := "dup/"
	if err := Copy(src, dst); err != nil {
		t.Error(err)
	}

	defer rmTestFiles(t, src, dst)

	di, err := os.Stat(dst + "/" + src)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("destination file not found: %v", err)
		}

		t.Errorf("%v", err)
	}

	si, err := os.Stat(src)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("source file not found: %v", err)
		}

		t.Errorf("%v", err)
	}

	if si.Size() != di.Size() {
		t.Fatalf("source file (%q) and destination file (%q) have different sizes. source=%v and destination=%v", si.Name(), di.Name(), si.Size(), di.Size())
	}
}

func createSrcFile(t *testing.T, src, content string) {
	sf, err := os.Create(src)
	if err != nil {
		t.Fatalf("while creating %q: %v", src, err)
	}
	defer sf.Close()

	if _, err := sf.WriteString(content); err != nil {
		t.Fatalf("while writing content to %q: %v", src, err)
	}
}

func rmTestFiles(t *testing.T, src, dst string) {
	if err := os.Remove(src); err != nil {
		t.Errorf("while removing %q: %v", src, err)
	}

	if err := os.Remove(dst + "/" + src); err != nil {
		t.Errorf("while removing %q: %v", dst+"/"+src, err)
	}

	if err := os.Remove(dst); err != nil {
		t.Errorf("while removing %q: %v", dst, err)
	}
}
