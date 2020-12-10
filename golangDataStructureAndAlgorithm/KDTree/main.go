package main

import (
	"code.qiangqiang.com/studygo/golang-数据结构与算法/KDTree/points"
	"fmt"
)

func main() {
	test := []struct {
		Name   string
		Input  []KDPoint
		Output []KDPoint
	}{
		{Name: "nil",
			Input:  nil,
			Output: []KDPoint{},
		},
		{Name: "empty",
			Input:  []KDPoint{},
			Output: []KDPoint{},
		},
		{Name: "1",
			Input:  []KDPoint{&points.Point2D{X: 1.0, Y: 2.0}},
			Output: []KDPoint{&points.Point2D{X: 1.0, Y: 2.0}},
		},
		{Name: "2 equal",
			Input:  []KDPoint{&points.Point2D{X: 1.0, Y: 2.0}, &points.Point2D{X: 1.0, Y: 2.0}},
			Output: []KDPoint{&points.Point2D{X: 1.0, Y: 2.0}, &points.Point2D{X: 1.0, Y: 2.0}},
		},
		{Name: "sort1 dim equal",
			Input:  []KDPoint{&points.Point2D{X: 1.1, Y: 1.2}, &points.Point2D{X: 1.3, Y: 1.0}, &points.Point2D{X: 0.9, Y: 1.30}},
			Output: []KDPoint{&points.Point2D{X: 0.9, Y: 1.3}, &points.Point2D{X: 1.1, Y: 1.2}, &points.Point2D{X: 1.3, Y: 1.0}},
		},
	}
	mytree := NewKDTree(test[4].Input)
	fmt.Println(mytree.Strings())
}
