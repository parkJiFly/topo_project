package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gookit/goutil/strutil"
	"log"
)

func main() {
	arraylist.New()
	list := createTestList()
	//printTree(list)
	flag, deviceList := getLengthAndDeviceList(list, "101")
	if flag {
		value, found1 := deviceList.lengthMap.Get("length")
		if found1 {
			fmt.Println(value.(int))
		}
		fmt.Println("--------------------------------------------------------------")
		listValue, found2 := deviceList.lengthMap.Get("deviceList")
		if found2 {
			ll := listValue.(*arraylist.List)
			it := ll.Iterator()
			for it.Next() {
				node := it.Value().(*Node) // 类型断言
				fmt.Println(node.psrId)    // 只打印 psrId
			}
		}
	} else {
		log.Println("错啦错啦")
	}

}

func getLengthAndDeviceList(list arraylist.List, parentId string) (bool, *Node) {
	relist := arraylist.New()
	deviceList := arraylist.New()
	length := 0
	nodeFlag, node := findNode(list, parentId)
	if nodeFlag == false {
		return false, &Node{}
	} else {
		rightConn := node.rightConn
		if strutil.IsNotBlank(rightConn) {
			childConn := strutil.Split(rightConn, ",")
			for i := range childConn {
				childFlag, childNode := getLengthAndDeviceList(list, childConn[i])
				if childFlag == true {
					relist.Add(childNode)
				}
			}

			if relist.Size() != 0 {
				log.Println(relist.Size())
				it := relist.Iterator()
				for it.Next() {
					lengthNode := it.Value().(*Node)
					listValue, listFound := lengthNode.lengthMap.Get("deviceList")
					lengthValue, lengthFound := lengthNode.lengthMap.Get("length")
					if lengthFound {
						// 进行类型断言，将 value 转换为 arraylist.List
						comLength, ok := lengthValue.(int)
						if ok {
							if comLength >= length {
								if listFound {
									comList, ok := listValue.(*arraylist.List)
									if ok {
										deviceList = comList
										length = comLength
									}
								}
							}
						}
					}
				}
			}
			node.child = *relist
		}
		if node.deviceType == 1 {
			deviceList.Add(node)
			length = length + node.length
		}
		node.lengthMap.Put("deviceList", deviceList)
		node.lengthMap.Put("length", length)
		return true, node
	}

}
