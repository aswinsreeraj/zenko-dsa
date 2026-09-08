type MyHashSet struct {
    buckets []*Node
    bucketSize int
}

type Node struct {
    Val int
    Next *Node
}

func Constructor() MyHashSet {
    bucketSize := 1000
    return MyHashSet{buckets: make([]*Node, bucketSize), bucketSize: bucketSize}
}

func (this *MyHashSet) Add(key int)  {
    idx := key % this.bucketSize
    if this.buckets[idx] == nil {
        this.buckets[idx] = &Node{Val: key}
    } else {
        temp := this.buckets[idx]
        if temp.Val == key {
                return
        }
        for temp.Next != nil {
            if temp.Next.Val == key {
                return
            }
            temp = temp.Next
        }
        temp.Next = &Node{Val: key}
    }
}


func (this *MyHashSet) Remove(key int)  {
    idx := key % this.bucketSize
    if this.buckets[idx] == nil {
        return
    }
    if this.buckets[idx].Val == key {
        this.buckets[idx] = this.buckets[idx].Next
        return
    }
    temp := this.buckets[idx]
    for temp != nil && temp.Next != nil {
        if temp.Next.Val == key {
            temp.Next = temp.Next.Next
            return
        }
        temp = temp.Next
    }
}


func (this *MyHashSet) Contains(key int) bool {
    idx := key % this.bucketSize
    if this.buckets[idx] == nil {
        return false
    }
    temp := this.buckets[idx]
    for temp != nil {
        if temp.Val == key {
            return true
        }
        temp = temp.Next
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */