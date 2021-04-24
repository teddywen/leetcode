/*
https://leetcode-cn.com/problems/design-hashset/
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Put(1,1)
	obj.Put(2,2)
	fmt.Printf("%d\n", obj.Get(1))
	fmt.Printf("%d\n", obj.Get(3))
	obj.Put(2,1)
	fmt.Printf("%d\n", obj.Get(2))
	obj.Remove(2)
	fmt.Printf("%d\n", obj.Get(2))
}

type MyHashMap struct {
	table []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{table: make([]list.List, 769)}
}

func (this *MyHashMap) Put(key int, value int) {
	for e := this.table[key % len(this.table)].Front(); e != nil; e = e.Next() {
		if e.Value.(*[2]int)[0] == key {
			e.Value.(*[2]int)[1] = value
			return
		}
	}
	slot := key % len(this.table)
	this.table[slot].PushBack(&[2]int{key, value})
}

func (this *MyHashMap) Remove(key int) {
	slot := key % len(this.table)
	for e := this.table[slot].Front(); e != nil; e = e.Next() {
		if e.Value.(*[2]int)[0] == key {
			this.table[slot].Remove(e)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashMap) Get(key int) int {
	for e := this.table[key % len(this.table)].Front(); e != nil; e = e.Next() {
		if e.Value.(*[2]int)[0] == key {
			return e.Value.(*[2]int)[1]
		}
	}
	return -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
