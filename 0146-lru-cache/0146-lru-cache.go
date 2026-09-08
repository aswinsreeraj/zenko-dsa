type LRUCache struct {
    capacity int
    length int
    items map[int]*node
    head *node
    tail *node
}

type node struct {
    key int
    val int
    prev *node
    next *node
}


func Constructor(capacity int) LRUCache {
    return LRUCache{capacity: capacity, items: make(map[int]*node)}
}


func (this *LRUCache) Get(key int) int {
    n, ok := this.items[key]
    if !ok {
        return -1
    }
    this.moveToFront(n)
    return n.val
}

func (this *LRUCache) moveToFront(n *node) {
    if this.head == n {
        return
    }
    this.remove(n)
    this.pushToFront(n)
}

func (this *LRUCache) pushToFront(n *node) {
    n.prev = nil
    n.next = this.head
    if this.head != nil {
        this.head.prev = n
    }
    this.head = n
    if this.tail == nil {
        this.tail = n
    }
    this.length++
}

func (this *LRUCache) remove(n *node) {
    if n.prev != nil {
        n.prev.next = n.next
    } else {
        this.head = n.next
    }
    if n.next != nil {
        n.next.prev = n.prev
    } else {
        this.tail = n.prev
    }
    n.prev, n.next = nil, nil
    this.length--
}

func (this *LRUCache) Put(key int, value int)  {
    if n, ok := this.items[key]; ok {
        n.val = value
        this.moveToFront(n)
        return
    }
    n := &node{key: key, val: value}
    this.items[key] = n
    this.pushToFront(n)
    if this.length > this.capacity {
        old := this.tail
        this.remove(old)
        delete(this.items, old.key)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */