package cloner

import (
	"fmt"
	"io"
	"os"
)

func verify(file string) (os.FileInfo, error) {
	fi, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) || os.IsExist(err) {
			return fi, nil
		}

		return nil, fmt.Errorf("while reading %q: %v", fi.Name(), err)
	}

	if !fi.Mode().IsRegular() && !fi.Mode().IsDir() {
		return nil, fmt.Errorf("there is a problem with the file permissions for %q (%v)", fi.Name(), fi.Mode())
	}

	return fi, nil
}

func createDir(dst string) (os.FileInfo, error) {
	_, err := verify(dst)
	if err != nil {
		return nil, err
	}

	di, err := os.Stat(dst)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("while creating directory %q: %v", di.Name(), err)
		}
	}

	if os.IsExist(err) {
		return di, nil
	}

	return di, nil
}

func Clone(src, dst string) error {
	si, err := verify(src)
	if err != nil {
		return err
	}

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("while opening %q: %v", si.Name(), err)
	}
	defer sf.Close()

	if si.IsDir() {
		entries, err := sf.Readdir(0)
		if err != nil {
			return fmt.Errorf("while scanning %q directory: %v", si.Name(), err)
		}

		for _, e := range entries {
			Clone(src+e.Name(), dst)
		}

		return nil
	}

	if _, err := createDir(dst); err != nil {
		return err
	}

	df, err := os.Create(dst + si.Name())
	if err != nil {
		return fmt.Errorf("while creating %q: %v", dst+si.Name(), err)
	}

	if _, err := io.Copy(df, sf); err != nil {
		return fmt.Errorf("while cloning %q on %q: %v", si.Name(), dst+si.Name(), err)
	}

	return nil
}
