package list

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	l := NewLinkList()
	l.Insert(11, 1)
	l.Insert(12, 2)
	l.Insert(13, 3)
	l.Insert(14, 4)
	l.Insert(15, 5)
	l.Insert(33, 3)
	l.Display()

	fmt.Println("删除元素：", l.Delete(5))
	l.Display()
	fmt.Println("删除元素：", l.Delete(3))
	l.Display()

	l.Update(3, 44)
	l.Display()

	fmt.Println("查询元素 44, 结果：", l.Search(44).data)

	l.Clear()
	l.Display()
}