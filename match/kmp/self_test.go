package kmp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getMoveInfo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "场景1",
			args: args{
				s: "abcd",
			},
			want: map[int]int{
				0: 0,
				1: 0,
				2: 0,
				3: 0,
			},
		},
		{
			name: "场景2",
			args: args{
				s: "abab",
			},
			want: map[int]int{
				0: 0,
				1: 0,
				2: 1,
				3: 2,
			},
		},
		{
			name: "场景3",
			args: args{
				s: "abcab",
			},
			want: map[int]int{
				0: 0,
				1: 0,
				2: 0,
				3: 1,
				4: 2,
			},
		},
		{
			name: "场景4",
			args: args{
				s: "aaaaa", // 真子串的含义需要理解
			},
			want: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3,
				4: 4,
			},
		},
		{
			name: "场景5",
			args: args{
				s: "aabaaac",
			},
			want: map[int]int{
				0: 0, // a
				1: 1, // a
				2: 0, // b
				3: 1, // a
				4: 2, // a
				5: 2, // a
				6: 0, // c

			},
		},
	}
	Convey("getMoveInfo", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				got := getMoveInfo(tt.args.s)
				So(got, ShouldEqual, tt.want)
			})
		}
	})
}

func Test_findSubString(t *testing.T) {
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
				got := findSubString(tt.args.s, tt.args.t)
				So(got, ShouldEqual, tt.want)
			})
		}
	})
}
