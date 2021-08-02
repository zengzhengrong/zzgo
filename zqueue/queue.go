package zqueue

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

//获取队列长度
func (q *Queue) Len() int {
	return q.length
}

//返回true队列不为空
func (q *Queue) Any() bool {
	return q.length > 0
}

//返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

//入队操作
func (q *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if q.length == 0 {
		// 当 queue 的长度时0的时候，push的第一个原素即是头也是尾
		q.top = n
		q.rear = q.top
	} else {
		// queue 的长度为0，则入队的时，队列的最后一个，将成为n的前一个元素，
		n.pre = q.rear
		q.rear.next = n // 连接队列最后一个元素和入队元素之间的关系
		q.rear = n      // 此时队列的最后一个就是入队的元素
	}
	q.length++
}

//出队操作
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.top
	if q.top.next == nil {
		q.top = nil
	} else {
		q.top = q.top.next
		// 切断出队的节点与下一个节点的关系
		q.top.pre.next = nil
		q.top.pre = nil
	}
	q.length--
	return n.value
}
