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


func NewDoubleLinkedList() *DoubleLinkedList {
  return &DoubleLinkedList{nil, nil, 0}
}

type DoubleLinkedList struct {
  front *Element
  back *Element
  size int
}

func (this *DoubleLinkedList) Length() int {
  return this.size
}

func (this *DoubleLinkedList) At(index int) interface{} {
  return this.front.Skip(index).Value
}

func (this *DoubleLinkedList) Iterator() Iterator {
  return this.front.Iterator()
}

func (this *DoubleLinkedList) Find(pred Predicate) *Element {
  return this.front.Find(pred)
}

func (this *DoubleLinkedList) Front() *Element {
  if this.size == 0 {
    panic("DoubleLinkedList.front: list empty")
  }
  return this.front
}

func (this *DoubleLinkedList) Back() *Element {
  if this.size == 0 {
    panic("DoubleLinkedList.back: list empty")
  }
  return this.back
}

func (this *DoubleLinkedList) InsertFront(elem *Element) {
  elem.prev = nil
  elem.next = this.front
  if this.front != nil {
    this.front.prev = elem
  }
  this.front = elem
  if this.size == 0 {
    this.back = this.front
  }
  this.size++
}

func (this *DoubleLinkedList) InsertBack(elem *Element) {
  elem.next = nil
  elem.prev = this.back
  if this.back != nil {
    this.back.next = elem
  }
  this.back = elem
  if this.size == 0 {
    this.front = this.back
  }
  this.size++
}

func (this *DoubleLinkedList) Remove(elem *Element) {
  if elem.next == nil {
    this.back = elem.prev
  } else {
    elem.next.prev = elem.prev
  }
  if elem.prev == nil {
    this.front = elem.next
  } else {
    elem.prev.next = elem.next
  }
  elem.prev = nil
  elem.next = nil
  this.size--
}

func NewElement(value interface{}) *Element {
  return &Element{value, nil, nil}
}

type Element struct {
  Value interface{}
  prev *Element
  next *Element
}

func (this *Element) Find(pred Predicate) *Element {
  for list := this; list != nil; list = list.next {
    if pred(list.Value) {
      return list
    }
  }
  return nil
}

func (this *Element) Skip(index int) *Element {
  list := this
  if (index > 0) {
    for i := 0; i < index; i++ {
      if list == nil {
        return nil
      }
      list = list.next
    }
  } else {
    for i := index; i < 0; i++ {
      if list == nil {
        return nil
      }
      list = list.prev
    }
  }
  return list
}

func (this *Element) Iterator() Iterator {
  return &elementIterator{this}
}

type elementIterator struct {
  next *Element
}

func (this *elementIterator) HasNext() bool {
  return this.next != nil
}

func (this *elementIterator) Next() interface{} {
  if this.next == nil {
    panic("elementIterator: no next element")
  }
  res := this.next.Value
  this.next = this.next.next
  return res
}
