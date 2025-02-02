package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const dataDir = "../data/even-strings"

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
			args{strings.NewReader(`7
3
ababa
ababa
ababa
3
asd
das
sda
2
abca
abc
4
aaaa
aaaa
aaaa
aaa
2
aa
aa
2
a
a
2
a
b
`)},
			`3
0
1
6
1
1
0
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

func Benchmark_run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(filepath.Join(dataDir, "10"))
		if err != nil {
			panic(err)
		}
		run(f, io.Discard)
		f.Close()
	}
}
