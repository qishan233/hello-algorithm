package kmp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNormalSearch(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple found",
			args: args{
				s: "abc",
				t: "a",
			},
			want: 0,
		},
		{
			name: "simple found 2",
			args: args{
				s: "abc",
				t: "b",
			},
			want: 1,
		},
		{
			name: "simple found 3",
			args: args{
				s: "abc",
				t: "c",
			},
			want: 2,
		},

		{
			name: "found 4",
			args: args{
				s: "ababababc",
				t: "abc",
			},
			want: 6,
		},

		{
			name: "found 5",
			args: args{
				s: "abababcab",
				t: "abc",
			},
			want: 4,
		},
		{
			name: "found 6",
			args: args{
				s: "abcababab",
				t: "abc",
			},
			want: 0,
		},
		{
			name: "not found 1",
			args: args{
				s: "abcababab",
				t: "abcd",
			},
			want: -1,
		},
		{
			name: "not found 2",
			args: args{
				s: "abcabdabeab",
				t: "abab",
			},
			want: -1,
		},
		{
			name: "not found 3",
			args: args{
				s: "abcababab",
				t: "abcdabc",
			},
			want: -1,
		},
		{
			name: "some error",
			args: args{
				s: "aabaaabaaac",
				t: "aabaaac",
			},
			want: 4,
		},
	}
	Convey("kmp: self", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				got := NormalSearch(tt.args.s, tt.args.t)
				So(got, ShouldEqual, tt.want)
			})
		}
	})
}
