package main

import (
	"strconv"
	"strings"
)

type Map struct {
	points [][]Point         //地图，保存一个矩阵
	blocks map[string]*Point //字符串对应每一个节点
	maxX   int               //最大的X坐标
	maxY   int               //最大的Y坐标
}

func NewMap(charMap []string) (m Map) {
	m.points = make([][]Point, len(charMap))           //开辟内存存储二维数组
	m.blocks = make(map[string]*Point, len(charMap)*2) //两倍边长
	for x, row := range charMap {
		cols := strings.Split(row, " ")
		m.points[x] = make([]Point, len(cols))
		for y, view := range cols {
			m.points[x][y] = Point{x, y, view}
			if view == "X" {
				//标记障碍
				//m.blocks
				m.blocks[PointAsKey(x, y)] = &m.points[x][y]
			}
		}
	}
	m.maxX = len(m.points)
	m.maxY = len(m.points[0])
	return m
}
func PointAsKey(x, y int) (key string) {
	key = strconv.Itoa(x) + "," + strconv.Itoa(y) //坐标转化为字符串
	return key
}

//抓取相邻的节点
//x-1    y+1   x,   y+1    x+1,  y+1
//x-1,   y       x,   y        x+1,   y
//x-1,   y-1    x,   y-1      x+1,  y-1
func (this *Map) GetAdjPoint(curPoint *Point) (Adjaoent []*Point) {
	if x, y := curPoint.x, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x+1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y+1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}
	if x, y := curPoint.x-1, curPoint.y-1; x >= 0 && x < this.maxX && y >= 0 && y < this.maxY {
		Adjaoent = append(Adjaoent, &this.points[x][y])
	}

	return Adjaoent
}
