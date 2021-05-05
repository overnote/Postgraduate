const Linklist = require('./LinkList')

const l = new Linklist()
l.display()

l.insert(11, 1)
l.insert(12, 2)
l.insert(13, 3)
l.insert(14, 4)
l.insert(15, 5)
l.insert(33, 3)
l.display()

console.log('删除元素：', l.delete(5))
l.display()
console.log('删除元素：', l.delete(3))
l.display()

l.update(3, 44)
l.display()

console.log('查询元素 44，结果：', l.search(44).data)

l.clear()
l.display()
