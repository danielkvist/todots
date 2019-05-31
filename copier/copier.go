// Package copier provides utilities to copy simple configuration files.
package copier

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Dotfile represents a simple configuration file.
type Dotfile struct {
	name string
	data []byte
}

// Check verifies that there are no errors in the provided path
// as restricted file permissions. It also verifies that the provided path
// ends in a file and not in a directory.
func Check(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("while checking path %q: %v", path, err)
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("while checking path %q found that file permissions are restricted: %v", path, fi.Mode())
	}

	if fi.Mode().IsDir() {
		return fmt.Errorf("while checking path %q found a directory instead of a file", path)
	}

	return nil
}

// NewDotfile return a *Dotfile with the name provided.
func NewDotfile(name string) *Dotfile {
	return &Dotfile{
		name: name,
		data: []byte{},
	}
}

// CopyFrom copies all the data from the provided io.Reader
// to the *Dotfile. It reports an error if there are any problems
// while copying the data.
func (d *Dotfile) CopyFrom(r io.Reader) error {
	data := &bytes.Buffer{}
	if _, err := io.Copy(data, r); err != nil {
		return fmt.Errorf("while copying data to file %q: %v", d.name, err)
	}

	d.data = data.Bytes()
	return nil
}

// WriteTo copies all the data from the *Dotfile to the
// provided io.Writer. It reports an error if there are any problems
// while copying the data.
//
// If there are no errors it returns the number of bytes copied.
func (d *Dotfile) WriteTo(w io.Writer) (int64, error) {
	buf := bytes.NewBuffer(d.data)
	bc, err := io.Copy(w, buf)
	if err != nil {
		return 0, fmt.Errorf("while writing data from %q: %v", d.name, err)
	}

	return bc, nil
}
