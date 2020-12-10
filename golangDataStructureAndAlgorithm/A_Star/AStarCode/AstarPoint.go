package main

import "math"

//A星算法，地图上的点结构
type AstarPoint struct {
	Point
	father *AstarPoint
	gVal   int //g(n)表示初始节点到任意节点的代价
	hVal   int //h(n)表示从节点n到目标点的启发式评估代价
	fVal   int //综合评估，f(n)最小的节点n，其中f(n)=g(n)+h(n)
}

func NewAstarPoint(p *Point, father *AstarPoint, end *AstarPoint) (ap *AstarPoint) {
	ap = &AstarPoint{*p, father, 0, 0, 0} //初始化
	if end != nil {
		ap.CalcFval(end) //创建的时候就计算节点的评估
	}
	return ap
}
func (asp *AstarPoint) CalcGval() int {
	if asp.father != nil {
		deltaX := math.Abs(float64(asp.father.x - asp.x))
		deltaY := math.Abs(float64(asp.father.y - asp.y)) //从父亲节点走过来需要的代价
		if deltaX == 1 && deltaY == 0 {                   //移动一步
			asp.gVal = asp.father.gVal + 10
		} else if deltaY == 1 && deltaX == 0 {
			asp.gVal = asp.father.gVal + 10
		} else if deltaY == 1 && deltaX == 1 {
			asp.gVal = asp.father.gVal + 14
		} else {
			panic("run error")
		}
	}
	return asp.gVal
}

//计算当前节点与目标节点的差距
func (asp *AstarPoint) CalcHval(end *AstarPoint) int {
	asp.hVal = int(math.Abs(float64(end.x-asp.x)) + math.Abs(float64(end.y-asp.y)))
	return asp.hVal
}

func (asp *AstarPoint) CalcFval(end *AstarPoint) int {
	asp.fVal = asp.CalcGval() + asp.CalcHval(end)
	return asp.fVal
}
