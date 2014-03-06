// Copyright 2014 Matthias Zenger. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package impl

import . "github.com/objecthub/containerkit"


func NewLinkedList() *LinkedList {
  return &LinkedList{nil, nil, 0}
}

type LinkedList struct {
  head *Cons
  tail *Cons
  size int
}

func (this *LinkedList) Length() int {
  return this.size
}

func (this *LinkedList) At(index int) interface{} {
  return this.head.Skip(index).Head
}

func (this *LinkedList) Iterator() Iterator {
  return this.head.Iterator()
}

func (this *LinkedList) InsertHead(elem interface{}) {
  this.head = NewCons(elem, this.head)
  if (this.size == 0) {
    this.tail = this.head
  }
  this.size++
}

func (this *LinkedList) InsertTail(elem interface{}) {
  if this.size == 0 {
    this.tail = NewCons(elem, nil)
    this.head = this.tail
  } else {
    this.tail.Tail = NewCons(elem, nil)
    this.tail = this.tail.Tail
  }
  this.size++
}

func (this *LinkedList) RemoveHead() interface{} {
  if this.size == 0 {
    panic("LinkedList.removeHead: list empty")
  }
  res := this.head.Head
  this.head = this.head.Tail
  if this.head == nil {
    this.tail = nil
  }
  this.size--
  return res
}

func (this *LinkedList) GetHead() interface{} {
  if this.size == 0 {
    panic("LinkedList.head: list empty")
  }
  return this.head.Head
}

func (this *LinkedList) GetTail() interface{} {
  if this.size == 0 {
    panic("LinkedList.tail: list empty")
  }
  return this.tail.Head
}

func (this *LinkedList) Find(pred Predicate) *Cons {
  return this.head.Find(pred)
}


func NewCons(head interface{}, tail *Cons) *Cons {
  return &Cons{head, tail}
}

type Cons struct {
  Head interface{}
  Tail *Cons
}

func (this *Cons) Find(pred Predicate) *Cons {
  for list := this; list != nil; list = list.Tail {
    if pred(list.Head) {
      return list
    }
  }
  return nil
}

func (this *Cons) Skip(index int) *Cons {
  list := this
  for i := 0; i < index; i++ {
    if list == nil {
      return nil
    }
    list = list.Tail
  }
  return list
}

func (this *Cons) Iterator() Iterator {
  return &consIterator{this}
}

type consIterator struct {
  next *Cons
}

func (this *consIterator) HasNext() bool {
  return this.next != nil
}

func (this *consIterator) Next() interface{} {
  if this.next == nil {
    panic("consIterator: no next element")
  }
  res := this.next.Head
  this.next = this.next.Tail
  return res
}
