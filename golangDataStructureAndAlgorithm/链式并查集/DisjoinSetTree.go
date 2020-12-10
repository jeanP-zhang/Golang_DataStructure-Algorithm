package main

const offLineMinimumExtract = -1

//集合归并，并查集去重
func OffLineMinimum(seq []int) []int {
	//第一个不能等于被挖空的运算
	if len(seq) == 0 || seq[0] == offLineMinimumExtract {
		panic("seq 输入错误")
	}

	insertSet := make([]*DisjoinSetTree, 0, 0)            //构造并查集数组
	values := make([]*DisjoinSetTree, len(seq), cap(seq)) //开辟内存
	n, m := 1, 0
	insertSet = append(insertSet, NewDisjoinSetTree(0)) //新添加一个节点值为0的并查集节点
	for i := range seq {
		if seq[i] == offLineMinimumExtract {
			m++ //统计系数
			insertSet = append(insertSet, NewDisjoinSetTree(m))
		} else {
			values[seq[i]] = NewDisjoinSetTree(seq[i])    //设置数据，重新导入
			insertSet[m] = Union(insertSet[m], values[i]) //数据的批量插入
			n++
		}
	}
	extractSeq := make([]int, m, m) //获取最短的
	for i := 1; i < n; i++ {
		j := FindSet(values[i]).Value.(int) //实例化
		if j != m {
			extractSeq[j] = i
			for l := j + 1; l <= m; l++ {
				if insertSet[l] != nil {
					insertSet[l] = Union(insertSet[l], insertSet[j])
					insertSet[l].Value = l
					insertSet[j] = nil
					break
				}
			}
		}
	}
	return extractSeq
}

//并查集--基本结构
type DisjoinSetTree struct {
	parent *DisjoinSetTree //父节点指针
	rank   int             //指数
	Value  interface{}     //数据
}

//环状并查集
func NewDisjoinSetTree(value interface{}) *DisjoinSetTree {
	t := new(DisjoinSetTree)
	t.Value = value
	t.parent = t //指向自己
	t.rank = 0
	return t
}

//查找集合
func FindSet(dst *DisjoinSetTree) *DisjoinSetTree {
	if dst.parent != dst {
		dst.parent = FindSet(dst.parent)
	}
	return dst
}

//合并
func Union(dst1, dst2 *DisjoinSetTree) *DisjoinSetTree {
	return Link(FindSet(dst1), FindSet(dst2))
}

func Link(dst1, dst2 *DisjoinSetTree) *DisjoinSetTree {
	if dst1 != dst2 {
		if dst1.rank < dst2.rank {
			dst1.parent = dst2
			return dst2
		}
		dst2.parent = dst1 //链接
		if dst1.rank > dst2.rank {
			dst1.rank++
		}
		return dst1
	}
	return dst1
}
