package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const dataDir = "../data/compact-boxes"

func Test_run(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		debug   bool
	}{
		{
			"1",
			args{strings.NewReader(`5
11 12
+----------+
|A+---+....|
|.|B52|....|
|.+---+....|
|+-------+.|
||r9.+--+|.|
||+-+|Ip||.|
|||7||..||.|
||+-++--+|.|
|+-------+.|
+----------+
3 3
...
...
...
3 5
+---+
|I63|
+---+
8 9
+------+.
|256...|.
|......|.
|......|.
+------+.
....+---+
....|R..|
....+---+
3 9
+-++-++-+
|2||5||6|
+-++-++-+`)},
			[]byte(`[
  {
    "A": {
      "B52": 3,
      "r9": {
        "7": 1,
        "Ip": 4
      }
    }
  },
  {},
  {
    "I63": 3
  },
  {
    "256": 18,
    "R": 3
  },
  {
    "2": 1,
    "5": 1,
    "6": 1
  }
]
`),
			true,
		},
		// 		{
		// 			"2",
		// 			args{strings.NewReader(``)},
		// 			[]byte(``),
		// 			true,
		// 		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug

			gotOut := &bytes.Buffer{}
			run(tt.args.in, gotOut)

			wantM := []any{}
			if err := json.Unmarshal(tt.wantOut, &wantM); err != nil {
				panic(err)
			}

			gotM := []any{}
			if err := json.Unmarshal(gotOut.Bytes(), &gotM); err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(gotM, wantM) {
				t.Errorf("run() = %v, want %v", gotM, wantM)
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
		wantOut []byte
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
			wantOut,
		}

		t.Run(tt.name, func(t *testing.T) {
			gotOut := &bytes.Buffer{}
			run(tt.args.in, gotOut)

			wantM := []any{}
			if err := json.Unmarshal(tt.wantOut, &wantM); err != nil {
				panic(err)
			}

			gotM := []any{}
			if err := json.Unmarshal(gotOut.Bytes(), &gotM); err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(gotM, wantM) {
				t.Errorf("run() = %v, want %v", gotM, wantM)
			}
		})
	}
}

func Benchmark_run(b *testing.B) {

	for i := 0; i < b.N; i++ {
		func() {
			f, err := os.Open(filepath.Join(dataDir, "35"))
			if err != nil {
				panic(err)
			}
			defer f.Close()
			run(f, io.Discard)
		}()
	}
}
