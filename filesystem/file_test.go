package filesystem

import (
	"log"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"file exists",
			args{
				"testdata/testfile",
			},
			true,
		},
		{
			"file exists as a dir",
			args{
				"testdata",
			},
			true,
		},
		{
			"file does not exists",
			args{
				"nothing",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(dir)

			if got := FileExists(tt.args.file); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
