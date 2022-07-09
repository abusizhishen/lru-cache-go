package lrc

type LruCache struct {
	cap        int
	count      int
	hash       map[string]*Node
	head, tail *Node
}

type Node struct {
	Key       string
	Value     interface{}
	Pre, Next *Node
}

func New(cap int) *LruCache {
	var head, tail = &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head

	return &LruCache{
		cap:  cap,
		hash: map[string]*Node{},
		head: head,
		tail: tail,
	}
}

func (l *LruCache) Put(key string, value interface{}) {
	node, ok := l.hash[key]
	if ok {
		node.Value = value
		l.moveToHead(node)
		return
	}

	node = &Node{Key: key, Value: value}
	if l.count == l.cap {
		l.removeTail()
	}

	l.count++
	l.hash[key] = node
	l.addToHead(node)
}

func (l *LruCache) Get(key string) (bool, *Node) {
	node, ok := l.hash[key]
	if !ok {
		return false, nil
	}

	if node != l.head.Next {
		l.moveToHead(node)
	}
	return true, node
}

func (l *LruCache) moveToHead(node *Node) {
	pre, next := node.Pre, node.Next
	pre.Next = l.tail
	next.Pre = pre
	l.addToHead(node)
}

func (l *LruCache) removeTail() *Node {
	if l.tail == l.head.Next {
		return nil
	}

	tail := l.tail.Pre
	pre := tail.Pre
	pre.Next = l.tail
	l.tail.Pre = pre

	delete(l.hash, tail.Key)
	l.count--
	return tail
}
func (l *LruCache) addToHead(node *Node) {
	next := l.head.Next
	node.Next = next
	next.Pre = node
	l.head.Next = node
}

func (l *LruCache) Cap() int {
	return l.cap
}

func (l *LruCache) Total() int {
	return l.count
}
