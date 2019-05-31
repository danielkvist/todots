package copier

import (
	"bytes"
	"testing"
)

func TestCheck(t *testing.T) {
	if err := Check("./copier.go"); err != nil {
		t.Fatalf("while checking path \"./copier.go\": %v", err)
	}
}

func TestCopyFrom(t *testing.T) {
	tt := []struct {
		data string
	}{
		{"aaa"},
		{"bbb"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
		{"42"},
		{"2019-05-31"},
	}

	for _, tc := range tt {
		t.Run(tc.data, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(tc.data))
			d := NewDotfile("test")
			if err := d.CopyFrom(buf); err != nil {
				t.Fatalf("while copying data into a new Dotfile: %v", err)
			}

			if string(d.data) != tc.data {
				t.Fatalf("expected copied data to be %q. got=%q", tc.data, string(d.data))
			}
		})
	}
}

func TestWriteTo(t *testing.T) {
	tt := []struct {
		name string
		data string
	}{
		{"one letter", "a"},
		{"two letters", "aa"},
		{"five letters", "aaaaa"},
		{"lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
		{"response to live", "42"},
		{"abc", "abcdefghijklmnopqrstuvwxyz"},
	}

	for _, tc := range tt {
		var buf bytes.Buffer
		d := NewDotfile(tc.name)
		d.data = []byte(tc.data)

		if err := d.WriteTo(&buf); err != nil {
			t.Fatalf("while writing data: %v", err)
		}

		if buf.String() != tc.data {
			t.Fatalf("expected writed data to be %q. got=%q", tc.data, buf.String())
		}
	}
}
