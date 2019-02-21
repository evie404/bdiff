package filesystem

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"file exists but not dir",
			args{
				"testdata/testfile",
			},
			false,
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
			if got := IsDir(tt.args.file); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirOfFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"file exists but not dir",
			args{
				"testdata/testfile",
			},
			"testdata",
		},
		{
			"file exists as a dir",
			args{
				"testdata",
			},
			"testdata",
		},
		{
			"file does not exists",
			args{
				"foo",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirOfFile(tt.args.file); got != tt.want {
				t.Errorf("DirOfFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
