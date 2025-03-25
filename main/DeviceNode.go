package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gookit/goutil/strutil"
	"strings"
)

type Node struct {
	//如果deviceType是1的话就是线段设备  不是1就是其他设备
	deviceType int
	psrId      string
	length     int
	leftConn   string
	rightConn  string
	child      arraylist.List
	lengthMap  hashmap.Map
}

func createTestList() arraylist.List {
	// 创建 List
	nodeList := arraylist.New()

	// 第1层
	node1 := &Node{
		deviceType: 1,
		psrId:      "A",
		length:     0,
		leftConn:   "101",
		rightConn:  "B,C,D",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	// 第2层
	node2 := &Node{
		deviceType: 1,
		psrId:      "B",
		length:     5,
		leftConn:   "A",
		rightConn:  "E,F",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node3 := &Node{
		deviceType: 1,
		psrId:      "C",
		length:     10,
		leftConn:   "A",
		rightConn:  "G,H",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node4 := &Node{
		deviceType: 1,
		psrId:      "D",
		length:     15,
		leftConn:   "A",
		rightConn:  "I,J",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	// 第3层
	node5 := &Node{
		deviceType: 1,
		psrId:      "E",
		length:     20,
		leftConn:   "B",
		rightConn:  "K,L",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node6 := &Node{
		deviceType: 1,
		psrId:      "F",
		length:     25,
		leftConn:   "B",
		rightConn:  "M,N",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node7 := &Node{
		deviceType: 1,
		psrId:      "G",
		length:     30,
		leftConn:   "C",
		rightConn:  "O,P",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node8 := &Node{
		deviceType: 1,
		psrId:      "H",
		length:     35,
		leftConn:   "C",
		rightConn:  "Q,R",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node9 := &Node{
		deviceType: 1,
		psrId:      "I",
		length:     40,
		leftConn:   "D",
		rightConn:  "S,T",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	node10 := &Node{
		deviceType: 1,
		psrId:      "J",
		length:     45,
		leftConn:   "D",
		rightConn:  "U,V",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(),
	}

	// 第4层
	node11 := &Node{
		deviceType: 1,
		psrId:      "K",
		length:     50,
		leftConn:   "E",
		rightConn:  "",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(), // K 是终端节点
	}

	node12 := &Node{
		deviceType: 1,
		psrId:      "L",
		length:     55,
		leftConn:   "E",
		rightConn:  "",
		child:      *arraylist.New(),
		lengthMap:  *hashmap.New(), // L 是终端节点
	}

	// 添加设备节点到 List
	nodeList.Add(node1)
	nodeList.Add(node2)
	nodeList.Add(node3)
	nodeList.Add(node4)
	nodeList.Add(node5)
	nodeList.Add(node6)
	nodeList.Add(node7)
	nodeList.Add(node8)
	nodeList.Add(node9)
	nodeList.Add(node10)
	nodeList.Add(node11)
	nodeList.Add(node12)

	return *nodeList
}

// 打印树形结构
func printTree(list arraylist.List) {
	// 从根节点开始打印
	_, rootNode := findNode(list, "101")
	if rootNode != nil {
		printNode(list, rootNode, 0) // 传递list
	}
}

// 打印节点，递归显示树的每一层
func printNode(list arraylist.List, node *Node, depth int) {
	// 打印当前节点
	printIndent(depth)
	fmt.Printf("%s\n", node.psrId)

	// 打印与当前节点连接的子节点（递归）
	if node.deviceType == 1 && node.rightConn != "" {
		childConn := strutil.Split(node.rightConn, ",")
		for _, conn := range childConn {
			childFlag, childNode := findNode(list, conn)
			if childFlag {
				printIndent(depth + 1) // 子节点应该有缩进
				fmt.Print("/")
				printNode(list, childNode, depth+1) // 递归调用打印子节点，传递 list
			}
		}
	}
}

// 找到指定节点
func findNode(list arraylist.List, parentId string) (bool, *Node) {
	_, value := list.Find(func(index int, value interface{}) bool {
		p, ok := value.(*Node)
		if ok {
			return p.leftConn == parentId
		}
		return false
	})
	if value != nil {
		return true, value.(*Node)
	}
	return false, nil
}

// 打印缩进（根据层数打印不同数量的空格）
func printIndent(depth int) {
	indent := strings.Repeat("  ", depth) // 每一层增加两个空格缩进
	fmt.Print(indent)
}
