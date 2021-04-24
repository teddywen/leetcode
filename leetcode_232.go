/*
https://leetcode-cn.com/problems/implement-queue-using-stacks/
*/
package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Printf("%d\n", obj.Peek())
	fmt.Printf("%d\n", obj.Pop())
	fmt.Printf("%t\n", obj.Empty())
}

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

func (mq *MyQueue) ReadyIn() {
	for len(mq.stackOut) > 0 {
		mq.stackIn = append(mq.stackIn, mq.stackOut[len(mq.stackOut)-1])
		mq.stackOut = mq.stackOut[:len(mq.stackOut)-1]
	}
}

func (mq *MyQueue) ReadyOut() {
	for len(mq.stackIn) > 0 {
		mq.stackOut = append(mq.stackOut, mq.stackIn[len(mq.stackIn)-1])
		mq.stackIn = mq.stackIn[:len(mq.stackIn)-1]
	}
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int) {
	mq.ReadyIn()
	mq.stackIn = append(mq.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	mq.ReadyOut()
	x := mq.stackOut[len(mq.stackOut)-1]
	mq.stackOut = mq.stackOut[:len(mq.stackOut)-1]
	return x
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	mq.ReadyOut()
	return mq.stackOut[len(mq.stackOut)-1]
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return len(mq.stackOut) == 0 && len(mq.stackIn) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
