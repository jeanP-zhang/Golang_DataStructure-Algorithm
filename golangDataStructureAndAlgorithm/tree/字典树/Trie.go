package main

import (
	"fmt"
	"sort"
	"sync"
)

type KeyWordTreeNode struct {
	//1，2，3
	KeyWordIDs map[int64]bool
	//百度一下
	Char string
	//父亲节点
	ParentKeyWordTreeNode *KeyWordTreeNode
	//子节点集合
	SubKeyWord map[string]*KeyWordTreeNode
}

func NewKeyWordTreeNode() *KeyWordTreeNode {
	return &KeyWordTreeNode{make(map[int64]bool, 0),
		"",
		nil,
		make(map[string]*KeyWordTreeNode, 0)}
}

func NewKeyWordTreeNodeWithParams(ch string, parent *KeyWordTreeNode) *KeyWordTreeNode {

	return &KeyWordTreeNode{make(map[int64]bool, 0),
		ch,
		parent,
		make(map[string]*KeyWordTreeNode, 0)}
}

type KeyWordTree struct {
	root        *KeyWordTreeNode //根节点
	kv          KeyWordKV        //映射关系
	charBeginKV CharBeginKV      //开始映射
	rw          *sync.RWMutex    //保护线程安全
}

func NewKeyWordTree() *KeyWordTree {
	return &KeyWordTree{NewKeyWordTreeNode(),
		KeyWordKV{},
		CharBeginKV{},
		new(sync.RWMutex)}
}
func (sTree *KeyWordTree) DeBugOut() {
	fmt.Println("s.kv", sTree.kv) //输出节点
	tempRoot := sTree.root
	dfs(tempRoot)
}

//遍历字典树
func dfs(root *KeyWordTreeNode) {
	if root == nil {
		return
	} else {
		fmt.Println("s.root=", root.Char)
		fmt.Println("s.KeyWordIds=", root.KeyWordIDs)
		for _, v := range root.SubKeyWord {
			dfs(v)
		}
	}
}
func (sTree *KeyWordTree) Put(id int64, keyWord string) {
	sTree.rw.Lock()
	defer sTree.rw.Unlock()
	sTree.kv[id] = keyWord
	tmpRoot := sTree.root //备份root节点
	//keyword反转
	for _, v := range keyWord { //循环每一个字符
		ch := string(v) //处理每个字符转化为字符串
		if tmpRoot.SubKeyWord[ch] == nil {
			//{百：{读}}
			node := NewKeyWordTreeNodeWithParams(ch, tmpRoot)           //开辟节点插入
			tmpRoot.SubKeyWord[ch] = node                               //赋值
			sTree.charBeginKV[ch] = append(sTree.charBeginKV[ch], node) //加入每一个节点
		} else {
			keyWordTreeNode := tmpRoot.SubKeyWord[ch]
			keyWordTreeNode.KeyWordIDs[id] = true //生效
			tmpRoot = tmpRoot.SubKeyWord[ch]      //节点向前推进
		}
	}
}

//百度搜索自动提示，返回字符串，limit限制深度
func (sTree *KeyWordTree) Search(keyWord string, limit int) []string {
	sTree.rw.Lock()
	defer sTree.rw.Unlock()
	ids := make(map[int64]int64, 0)
	for pos, v := range keyWord {
		ch := string(v)
		begins := sTree.charBeginKV[ch] //取得映射字符的所有节点
		for _, begin := range begins {
			key_word_tmp_pt := begin
			next_pos := pos + 1 //标记下一个位置
			if len(key_word_tmp_pt.SubKeyWord) > 0 && next_pos < len(keyWord) {
				//最大匹配
				nextCh := string(keyWord[next_pos]) //下一个字符
				if key_word_tmp_pt.SubKeyWord[nextCh] == nil {
					break
				}
				key_word_tmp_pt = key_word_tmp_pt.SubKeyWord[nextCh] //递推前进
				next_pos++
			}
			//保存结果
			for id, _ := range key_word_tmp_pt.KeyWordIDs {
				ids[id] = ids[id] + 1 //保存结果
			}
		}
	}
	list := PairList{} //列表
	for id, count := range ids {
		list = append(list, Pair{id, count}) //加载数据
	}
	if !sort.IsSorted(list) {
		sort.Sort(list) //排序
	}
	//limit限制出现的数量
	if len(list) > limit {
		list = list[:limit] //数据进行截取
	}
	ret := make([]string, 0)
	for _, item := range list {
		ret = append(ret, sTree.kv[item.K]) //返回数组叠加
	}

	return ret
}

//搜索
func (sTree *KeyWordTree) Sugg(keyword string, limit int) []string {
	sTree.rw.Lock()
	defer sTree.rw.Unlock()
	key_word_tmp_pt := sTree.root //根节点
	isEnd := true                 //判断是否结束
	//a abc acd axs
	for _, v := range keyword {
		ch := string(v)
		if key_word_tmp_pt.SubKeyWord[ch] == nil {
			isEnd = false
			break //提前结束判断
		}
		//循环条件
		key_word_tmp_pt = key_word_tmp_pt.SubKeyWord[ch] //子集合
		//abc ab 就无意义
		if isEnd {
			ret := make([]string, 0)
			ids := key_word_tmp_pt.KeyWordIDs //编号
			for id, _ := range ids {
				ret = append(ret, sTree.kv[id]) //结果的追加
				limit--
				if limit == 0 {
					break
				}
			}
			return ret
		}

	}
	return make([]string, 0)
}
