package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const dataDir = "./data/validate-result"

func Test_run(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
		{
			"1",
			args{strings.NewReader(`9
3
a 1
b 2
c 3
a:1,c:3,b:2
3
a 1
b 2
c 2
c:2,a:1
3
a 1
b 2
c 2
b:2,c:2
3
a 1
b 2
c 2
a:1,a:1,a:1,a:1
3
a 1
b 2
c 2
b:1
3
a 1
b 2
c 2
d:4,a:1,c:2
3
a 1
b 2
c 2
abcdef
3
a 1
b 2
c 2
a:12345678901234567890,c:2
1
abc 123
abc:0123
`)},
			`YES
YES
NO
NO
NO
NO
NO
NO
NO
`,
			true,
		},
		// 		{
		// 			"2",
		// 			args{strings.NewReader(``)},
		// 			``,
		// 			true,
		// 		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug

			out := &bytes.Buffer{}
			run(tt.args.in, out)
			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(tt.wantOut, "\n") {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		in io.Reader
	}
	type test struct {
		name    string
		args    args
		wantOut string
	}

	files, err := os.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if ok, err := filepath.Match("*.a", fileName); err != nil {
			panic(err)
		} else if !ok {
			continue
		}

		wantOut, err := os.ReadFile(filepath.Join(dataDir, fileName))
		if err != nil {
			panic(err)
		}

		in, err := os.Open(filepath.Join(dataDir, fileName[:len(fileName)-2]))
		if err != nil {
			panic(err)
		}

		tt := test{
			fileName,
			args{in},
			string(wantOut),
		}

		if !t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.in, out)
			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(tt.wantOut, "\n") {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		}) {
			return
		}
	}
}
