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
	obj.Add(1)
	obj.Add(2)
	fmt.Printf("%t\n", obj.Contains(1))
	fmt.Printf("%t\n", obj.Contains(3))
	obj.Add(2)
	fmt.Printf("%t\n", obj.Contains(2))
	obj.Remove(2)
	fmt.Printf("%t\n", obj.Contains(2))
}

type MyHashSet struct {
	table []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{table: make([]list.List, 769)}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	slot := key % len(this.table)
	this.table[slot].PushBack(key)
}

func (this *MyHashSet) Remove(key int) {
	slot := key % len(this.table)
	for e := this.table[slot].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.table[slot].Remove(e)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for e := this.table[key % len(this.table)].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
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
