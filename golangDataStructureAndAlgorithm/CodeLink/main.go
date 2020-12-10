package main

func AS0(array []int)  {

}
func main() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node1.data = 1
	node1.pNext = node2
	node2.data = 2
	node2.pNext = node3
	node3.data = 3
	node3.pNext = node4
	node4.data = 4
	node4.pNext = nil
}
