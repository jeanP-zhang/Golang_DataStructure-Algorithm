package main

type Pair struct {
	K int64
	V int64
}
type KeyWordKV map[int64]string                //set相当于不要value的map
type CharBeginKV map[string][]*KeyWordTreeNode //字典树结构
type PairList []Pair

func (p PairList) Len() int {
	return len(p) //求长度
}

func (p PairList) Less(i, j int) bool {
	return p[i].V > p[j].V //比大小
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i] //实现数据交换
}
