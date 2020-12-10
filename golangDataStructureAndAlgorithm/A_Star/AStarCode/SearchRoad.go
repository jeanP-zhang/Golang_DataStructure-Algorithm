package main

import (
	"container/heap"
	"fmt"
)

type SearchRoad struct {
	theMap  *Map
	start   AstarPoint
	end     AstarPoint
	closeLi map[string]*AstarPoint //关闭不通的路
	openLi  OpenList               //通路
	openSet map[string]*AstarPoint //去掉重复
	TheRoad []*AstarPoint          //通道
}

//打印地图，按照寻找好的通路
func (this *Map) printMap(path *SearchRoad) {
	fmt.Println("地图的边界", this.maxX, this.maxY)
	for x := 0; x < this.maxX; x++ {
		for y := 0; y < this.maxY; y++ {
			if path != nil {
				if x == path.start.x && y == path.start.y {
					fmt.Printf("%3s", "S")
					goto NEXT
				} //S代表开始
			}
			if x == path.end.x && y == path.end.y {
				fmt.Printf("%3s", "E") //E代表结束
				goto NEXT
			}
			for i := 0; i < len(path.TheRoad); i++ { //循环找路
				if path.TheRoad[i].x == x && path.TheRoad[i].y == y {
					fmt.Printf("%3s", "*")
					goto NEXT
				}
			}
			fmt.Printf("%3s", this.points[x][y].view)
		NEXT:
		}
	}
	fmt.Println()
}
func NewSearchRoad(startX, startY, endX, endY int, m *Map) *SearchRoad {
	sr := &SearchRoad{}
	sr.theMap = m
	sr.start = *NewAstarPoint(&Point{startX, startY, "S"}, nil, nil) //结束的节点
	sr.end = *NewAstarPoint(&Point{endX, endY, "E"}, nil, nil)       //结束的节点
	sr.TheRoad = make([]*AstarPoint, 0)                              //路,开辟内存
	sr.openSet = make(map[string]*AstarPoint, m.maxY+m.maxX)         //开放集合
	sr.closeLi = make(map[string]*AstarPoint, m.maxY+m.maxX)
	heap.Init(&sr.openLi)                                      //初始化栈
	heap.Push(&sr.openLi, &sr.start)                           //压入开始节点
	sr.openSet[PointAsKey(sr.start.x, sr.start.y)] = &sr.start //开放集合，压入开始节点
	for k, v := range m.blocks {
		sr.closeLi[k] = NewAstarPoint(v, nil, nil) //所有障碍加入blocks
	}
	return sr
}

//A*算法的核心
func (sr *SearchRoad) FindOutShortestPath() bool {
	//如果开放节点大于零，永远循环下去

	for len(sr.openLi) > 0 {
		//从开放节点中取出，放入关闭节点
		x := heap.Pop(&sr.openLi)
		curPoint := x.(*AstarPoint)                            //取得当前节点
		delete(sr.openSet, PointAsKey(curPoint.x, curPoint.y)) //删除开放列表
		sr.closeLi[PointAsKey(curPoint.x, curPoint.y)] = curPoint
		adjacs := sr.theMap.GetAdjPoint(&curPoint.Point) //取出所有邻居的节点
		for _, p := range adjacs {
			theAp := NewAstarPoint(p, curPoint, &sr.end) //创建A*节点
			//			找到了结束的节点
			if PointAsKey(theAp.x, theAp.y) == PointAsKey(sr.end.x, sr.end.y) {
				for theAp.father != nil {
					sr.TheRoad = append(sr.TheRoad, theAp) //加入节点
					theAp.view = "*"
					theAp = theAp.father //返回上一个节点
				}
				return true //结束
			}
			_, ok := sr.closeLi[PointAsKey(p.x, p.y)] //节点关闭就提前结束本次循环
			if ok {
				continue
			}
			if existAp, ok := sr.openSet[PointAsKey(p.x, p.y)]; !ok { //取出开放的节点,如果节点存在则处理，不存在则不处理
				heap.Push(&sr.openLi, theAp)                     //节点不存在就压入
				sr.openSet[PointAsKey(theAp.x, theAp.y)] = theAp //放入开放列表
			} else {
				//如果节点存在，经过对比要取得最短路径
				oldgVar, oldFather := existAp.gVal, existAp.father
				existAp.father = curPoint
				existAp.CalcGval() //计算最短的值
				//如果新节点距离比老节点还要短
				if existAp.gVal > oldgVar {
					existAp.father = oldFather
					existAp.gVal = oldgVar
				}
			}
		}

	}

	return false
}
