package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

//需要改变加*，无*不需要加*
type VEB struct {
	u, min, max int
	summary     *VEB
	cluster     []*VEB
}

func (V VEB) Max() int {
	return V.max
}

func (V VEB) Min() int {
	return V.min
}

//log2.log2
func LowerSqrt(u int) int {
	return int(math.Pow(2.0, math.Floor(math.Log2(float64(u))/2)))
}
func HigherSqrt(u int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(u))/2)))
}

//给我一个x计算存储的深度的簇号
func (V VEB) High(x int) int {
	return int(math.Floor(float64(x) / float64(LowerSqrt(V.u))))
}

//给我一个x计算存储的位置
func (V VEB) Low(x int) int {
	return x % LowerSqrt(V.u)
}

//计算并返回x,y的索引
func (V VEB) Index(x, y int) int {
	return x*LowerSqrt(V.u) + y
}

//VEB树所有接口
//创造树
func CreatTree(size int) *VEB {
	if size <= 0 {
		return nil
	}
	x := math.Ceil(math.Log2(float64(size))) //假定size==100，2^6=64
	u := int(math.Pow(2, x))
	V := new(VEB) //新建一个节点
	V.min, V.max = -1, -1
	V.u = u
	if u == 2 {
		return V //构造完成一个子树
	}
	clusterCount := HigherSqrt(u) //计算出来高度了,计算cluster数量
	clusterSize := LowerSqrt(u)   //计算大小
	for i := 0; i < clusterCount; i++ {
		V.cluster = append(V.cluster, CreatTree(clusterSize)) //递归插入
	}
	summarySize := HigherSqrt(u)
	V.summary = CreatTree(summarySize)
	return V
}

//判断节点是否存在
func (V VEB) IsMember(x int) bool {
	if x == V.min || x == V.max {
		return true
	} else if V.u == 2 {
		return false
	} else {
		return V.cluster[V.High(x)].IsMember(V.Low(x))
	}
}

//插入节点
func (V *VEB) Insert(x int) { //需要加指针，否则不会生效
	if V.min == -1 {
		V.min, V.max = x, x //存储了一个元素
	} else {
		if x < V.min {
			V.min, x = x, V.min //交换数据
		}
		if V.u > 2 {
			if V.cluster[V.High(x)].Min() == -1 {
				V.summary.Insert(V.High(x))
				V.cluster[V.High(x)].min, V.cluster[V.High(x)].max = V.Low(x), V.Low(x)
			} else {
				V.cluster[V.High(x)].Insert(V.Low(x)) //降维处理
			}
		}
		if x > V.max {
			V.max = x
		}
	}
}

//删除节点
func (V *VEB) Delete(x int) {
	if V.summary == nil && V.summary.Min() == -1 { //无非空簇
		if x == V.min && x == V.max {
			V.min, V.max = -1, -1

		} else if x == V.min { //两个元素，x最小
			V.min = V.max //重合删除
		} else {
			V.max = V.min //两个元素，x最大
		}

	} else { //有非空簇
		if x == V.min {
			//取得最小的在cluster
			y := V.Index(V.summary.min, V.cluster[V.summary.min].min) //取得Y所在的索引
			V.min = y                                                 //取得最接近的，赋值V.min
			V.cluster[V.High(x)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 { //仅有的数据
				V.summary.Delete(V.High(y))
			}
		} else if x == V.max {
			//
			y := V.Index(V.summary.max, V.cluster[V.summary.min].min) //取得Y所在的索引
			V.cluster[V.High(y)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 {
				V.summary.Delete(V.High(y))
			}
			if V.summary == nil || V.summary.min == -1 {
				if V.min == y { //直接删除
					V.min, V.max = -1, -1
				}
			} else {
				V.max = V.min
				V.max = V.Index(V.summary.max, V.cluster[V.summary.max].max)
			}
		} else { //
			V.cluster[V.High(x)].Delete(V.Low(x)) //删除了节点
			if V.cluster[V.High(x)].min == -1 {
				V.summary.Delete(V.High(x))
			}
		}

	}
}

//后继,找到x位置
func (V VEB) Successor(x int) int {
	if V.u == 2 {
		if x == 0 && V.max == 1 {
			return 1
		} else {
			return -1 //找不到
		}
	} else if V.max != -1 && x > V.min {
		return V.max //找不到
	} else {
		minLow := V.cluster[V.High(x)].Min() //最大
		if minLow != -1 && V.Low(x) > minLow {
			offset := V.cluster[V.High(x)].Predecessor(V.Low(x)) //解决偏移量
			return V.Index(V.High(x), offset)                    //求索引
		} else { //进入子节点
			preCluster := V.summary.Predecessor(V.High(x))
			if preCluster == -1 {
				if V.min != -1 && x > V.min {
					return V.min
				}
				return -1
			} else {
				offset := V.cluster[preCluster].Max()
				return V.Index(preCluster, offset)
			}
		}
	}

}

//找到X的前驱位置
func (V VEB) Predecessor(x int) int {
	if V.u == 2 {
		if x == 1 && V.min == 0 {
			return 1
		} else {
			return -1 //找不到
		}
	} else if V.min != -1 && x < V.min {
		return V.min //找不到
	} else {
		maxLow := V.cluster[V.High(x)].Max() //最大
		if maxLow != -1 && V.Low(x) < maxLow {
			offset := V.cluster[V.High(x)].Successor(V.Low(x)) //解决偏移量
			return V.Index(V.High(x), offset)                  //求索引
		} else { //进入子节点
			succCluster := V.summary.Successor(V.High(x))
			if succCluster == -1 {
				return -1
			} else {
				offset := V.cluster[succCluster].Min()
				return V.Index(succCluster, offset)
			}
		}
	}

}

//统计树的节点
func (V VEB) Count() int {
	if V.u == 2 {
		return 1
	}
	sum := 1
	for i := 0; i < len(V.cluster); i++ {
		sum += V.cluster[i].Count() //统计次数
	}
	sum += V.summary.Count()
	return sum
}

func (V VEB) Print() {
	V.PrintFunc(0, 0, false)
}

//递归实现

func (V VEB) PrintFunc(level int, clusterNo int, summary bool) {
	space := "  "
	for i := 0; i < level; i++ {
		space += "\t"
	}
	if level == 0 {
		fmt.Printf("%vR:{U:%v,min:%v;max:%v,cluster:%v}\n", space, V.u, V.min, V.max, len(V.cluster))
	} else {
		if summary {
			fmt.Printf("%vS:{U:%v,min:%v;max:%v,cluster:%v}\n", space, V.u, V.min, V.max, len(V.cluster))

		} else {
			fmt.Printf("%vC[%v]:{U:%v,min:%v;max:%v,cluster:%v}\n", space, clusterNo, V.u, V.min, V.max, len(V.cluster))

		}
	}
	if len(V.cluster) > 0 {
		V.summary.PrintFunc(level+1, 0, true)
		for i := 0; i < len(V.cluster); i++ {
			V.cluster[i].PrintFunc(level, i, false)
		}
	}
}

//清空
func (V *VEB) Clear() {
	V.min, V.max = -1, -1
	if V.u == 2 {
		return
	}
	for i := 0; i < len(V.cluster); i++ {
		V.cluster[i].Clear()
	}
	V.summary.Clear()
}

//填充
func (V *VEB) Fill() {
	for i := 0; i < V.u; i++ {
		V.Insert(i)
	}
}

//返回元素
func (V VEB) Members() []int {
	members := []int{}
	for i := 0; i < V.u; i++ {
		if V.IsMember(i) {
			members = append(members, i)
		}
	}
	return members
}
func arrayContains(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return true
		}
	}
	return false
}

//创建不重复的随机数数组
func makeRandom(max int) []int {
	myrand := rand.New(rand.NewSource(int64(time.Now().Nanosecond()))) //时间随机数
	keys := make([]int, 0)
	keyNo := myrand.Intn(max) //创建随机数
	for i := 0; i < keyNo; i++ {
		myKey := myrand.Intn(max - 1) //取得随机数
		if !arrayContains(keys, myKey) {
			keys = append(keys, myKey)

		}
	}
	sort.Ints(keys) //排序号

	return keys
}
func main1() {
	maxUpPower := 10
	for i := 1; i < maxUpPower; i++ {
		u := int(math.Pow(2.0, float64(i)))
		V := CreatTree(u)
		keys := makeRandom(u) //随机数组
		for j := 0; j < u; j++ {
			fmt.Println(V.IsMember(j))
		}
		fmt.Println("keys:", keys)
		for j := 0; j < len(keys); j++ {
			V.Insert(keys[j]) //插入数据
		}

		V.Print()
		fmt.Println("---------------------------------------------")
	}
}
