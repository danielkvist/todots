package copier

import (
	"os"
	"testing"
)

func TestCopyFiles(t *testing.T) {
	src := "test.txt"
	content := "Hello, Gophers!\n"

	createFile(t, src, content)
	defer rmFile(t, src)

	dst := "dup/"
	if err := Copy(src, dst); err != nil {
		t.Error(err)
	}
	defer rmFile(t, dst)

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

func TestCopyDirs(t *testing.T) {
	srcD := "test/"
	srcF := "test.txt"
	content := "Hello, Gophers!\n"

	createDir(t, srcD)
	createFile(t, srcD+srcF, content)
	defer rmFile(t, srcD)

	dst := "dup/"
	if err := Copy(srcD, dst); err != nil {
		t.Error(err)
	}

	defer rmFile(t, dst)
	di, err := os.Stat(dst + "/" + srcF)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("destination file not found: %v", err)
		}

		t.Errorf("%v", err)
	}

	si, err := os.Stat(srcD + srcF)
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

func createDir(t *testing.T, d string) {
	if err := os.Mkdir(d, os.ModePerm); err != nil {
		t.Errorf("while creating dir %q: %v", d, err)
	}
}

func createFile(t *testing.T, f, content string) {
	fi, err := os.Create(f)
	if err != nil {
		t.Fatalf("while creating %q: %v", f, err)
	}
	defer fi.Close()

	if _, err := fi.WriteString(content); err != nil {
		t.Fatalf("while writing content (%q) to %q: %v", content, f, err)
	}
}

func rmFile(t *testing.T, f string) {
	fi, err := os.Stat(f)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !fi.IsDir() {
		if err := os.Remove(f); err != nil {
			t.Errorf("while removing %q: %v", f, err)
			return
		}
	}

	if err := os.RemoveAll(f); err != nil {
		t.Errorf("while removing directory %q: %v", f, err)
	}
}
