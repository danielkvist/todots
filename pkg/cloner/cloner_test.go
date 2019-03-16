package cloner

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestClone(t *testing.T) {
	src := "test.txt"
	dst := "test/"

	sf, err := os.Create(src)
	if err != nil {
		t.Errorf("while creating testing file %q: %v", src, err)
	}

	if _, err := sf.WriteString("testing"); err != nil {
		t.Errorf("while writing into the testing file %q: %v", src, err)
	}

	defer func() {
		sf.Close()
		if err := os.Remove(src); err != nil {
			t.Errorf("while removing the testing file %q: %v", src, err)
		}
	}()

	if err := Clone(src, dst); err != nil {
		t.Errorf("while cloning %q into %q: %v", src, dst, err)
	}

	defer func(dst string) {
		if err := os.RemoveAll(dst); err != nil {
			t.Errorf("while removing the testing file %q: %v", dst+src, err)
		}
	}(dst)

	dst += src
	if !testEquality(src, dst, t) {
		t.Errorf("source file (%q) and destination file (%q) expected to have the same content", src, dst)
	}

}

func testEquality(f1, f2 string, t *testing.T) bool {
	var eq bool

	f1d, err := ioutil.ReadFile(f1)
	if err != nil {
		t.Errorf("while reading %q: %v", f1, err)
	}

	f2d, err := ioutil.ReadFile(f2)
	if err != nil {
		t.Errorf("while reading %q: %v", f2, err)
	}

	if bytes.Equal(f1d, f2d) {
		eq = true
	}

	return eq
}
