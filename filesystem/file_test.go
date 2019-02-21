package filesystem

import (
	"log"
	"os"
	"reflect"
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

func TestFilesExists(t *testing.T) {
	type args struct {
		files []string
	}
	tests := []struct {
		name         string
		args         args
		wantExist    []string
		wantNotExist []string
	}{
		{
			"splits files into exists and notExists slices",
			args{
				[]string{
					"testdata/testfile",
					"testdata",
					"nothing",
					"nothing/nothing",
				},
			},
			[]string{
				"testdata/testfile",
				"testdata",
			},
			[]string{
				"nothing",
				"nothing/nothing",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, gotNotExist := FilesExists(tt.args.files)
			if !reflect.DeepEqual(gotExist, tt.wantExist) {
				t.Errorf("FilesExists() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
			if !reflect.DeepEqual(gotNotExist, tt.wantNotExist) {
				t.Errorf("FilesExists() gotNotExist = %v, want %v", gotNotExist, tt.wantNotExist)
			}
		})
	}
}
