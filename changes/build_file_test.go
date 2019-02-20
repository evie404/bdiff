package changes

import (
	"reflect"
	"testing"
)

func TestIsBuildFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"a directory",
			args{
				"foo/hi",
			},
			false,
		},
		{
			"a file",
			args{
				"foo.go",
			},
			false,
		},
		{
			"a build file",
			args{
				"BUILD",
			},
			true,
		},
		{
			"a bazel build file",
			args{
				"BUILD.bazel",
			},
			true,
		},
		{
			"a bzl file",
			args{
				"repos.bzl",
			},
			true,
		},
		{
			"a build file inside a directory",
			args{
				"foo/BUILD",
			},
			true,
		},
		{
			"a bazel build file inside a directory",
			args{
				"foo/BUILD.bazel",
			},
			true,
		},
		{
			"a bzl file inside a directory",
			args{
				"foo/repos.bzl",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBuildFile(tt.args.file); got != tt.want {
				t.Errorf("IsBuildFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildFileChanges(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"a directory",
			args{
				"foo/hi",
			},
			nil,
		},
		{
			"a file",
			args{
				"foo.go",
			},
			nil,
		},
		{
			"a build file",
			args{
				"BUILD",
			},
			[]string{"//..."},
		},
		{
			"a bazel build file",
			args{
				"BUILD.bazel",
			},
			[]string{"//..."},
		},
		{
			"a bzl file",
			args{
				"repos.bzl",
			},
			[]string{"//..."},
		},
		{
			"a build file inside a directory",
			args{
				"foo/BUILD",
			},
			[]string{"//foo/..."},
		},
		{
			"a bazel build file inside a directory",
			args{
				"foo/BUILD.bazel",
			},
			[]string{"//foo/..."},
		},
		{
			"a bzl file inside a directory",
			args{
				"foo/repos.bzl",
			},
			[]string{"//foo/..."},
		},
		{
			"a build file inside a subdirectory",
			args{
				"foo/test/zzz/BUILD",
			},
			[]string{"//foo/test/zzz/..."},
		},
		{
			"a bazel build file inside a subdirectory",
			args{
				"foo/test/zzz/BUILD.bazel",
			},
			[]string{"//foo/test/zzz/..."},
		},
		{
			"a bzl file inside a subdirectory",
			args{
				"foo/test/zzz/repos.bzl",
			},
			[]string{"//foo/test/zzz/..."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildFileChanges(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildFileChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
