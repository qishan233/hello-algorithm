package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_sort1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "升序数组",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "降序数组",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "重复数据",
			args: args{
				nums: []int{5, 5, 4, 3, 3, 2, 1},
			},
			want: []int{1, 2, 3, 3, 4, 5, 5},
		},
		{
			name: "空数组",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "随机数组",
			args: args{
				nums: []int{5, 55, 6, 3, 5, 7, 5, 9, 10, 5, 4, 3, 3, 2, 1},
			},
			want: []int{1, 2, 3, 3, 3, 4, 5, 5, 5, 5, 6, 7, 9, 10, 55},
		},
	}

	Convey("test for fast sort algorithm", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				sort1(tt.args.nums)
				So(tt.args.nums, ShouldEqual, tt.want)
			})
		}

	})

}

func Test_sort2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "升序数组",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "降序数组",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "重复数据",
			args: args{
				nums: []int{5, 5, 4, 3, 3, 2, 1},
			},
			want: []int{1, 2, 3, 3, 4, 5, 5},
		},
		{
			name: "空数组",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "随机数组",
			args: args{
				nums: []int{5, 55, 6, 3, 5, 7, 5, 9, 10, 5, 4, 3, 3, 2, 1},
			},
			want: []int{1, 2, 3, 3, 3, 4, 5, 5, 5, 5, 6, 7, 9, 10, 55},
		},
	}
	Convey("test for fast sort algorithm", t, func() {
		for _, tt := range tests {
			Convey(tt.name, func() {
				sort2(tt.args.nums)
				So(tt.args.nums, ShouldEqual, tt.want)
			})
		}
	})

}
