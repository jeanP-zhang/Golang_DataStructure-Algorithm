package main

import "fmt"

func _compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -2
	}
	if newB, ok = b.(int); !ok {
		return -2
	}

	if newA > newB {
		return 1
	} else if newA < newB {
		return -1
	} else {
		return 0
	}

}

func main() {
	myVal, err := NewAvlTree(3, _compare)
	if err != nil {
		fmt.Println("创建函数失败")
	}
	myVal = myVal.Insert(2)
	myVal = myVal.Insert(1)
	myVal = myVal.Insert(4)
	myVal = myVal.Insert(5)
	myVal = myVal.Insert(6)
	myVal = myVal.Insert(7)
	myVal = myVal.Insert(15)
	myVal = myVal.Insert(26)
	myVal = myVal.Insert(17)
	myVal = myVal.Delete(7)
	a := myVal.GetAll()
	fmt.Println(a)
}
