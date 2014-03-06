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


func NewHeap(comparator Comparison) *Heap {
  return &Heap{NewArray(), comparator}
}

type Heap struct {
  as Array
  comparator Comparison
}

func (this *Heap) Comparison() Comparison {
  return this.comparator
}

func (this *Heap) Length() int {
  return this.as.Length()
}

func (this *Heap) First() interface{} {
  return this.as.At(0)
}

func (this *Heap) Copy() Heap {
  return Heap{this.as.Copy(), this.comparator}
}

func (this *Heap) Iterator() *heapIterator {
  return &heapIterator{this.Copy()}
}

func (this *Heap) Add(elem interface{}) {
  this.as.Append(elem)
  this.fixUp(this.as.Length() - 1)
}

func (this *Heap) Next() interface{} {
  n := this.as.Length() - 1
  this.swap(0, n)
  res := this.as.At(n)
  this.as.Delete(n, 1)
  this.fixDown(0)
  return res
}

func (this *Heap) swap(i, j int) {
  if i != j {
    h := this.as[i]
    this.as[i] = this.as[j]
    this.as[j] = h
  }
}

func (this *Heap) fixUp(i int) {
  for i > 0 && this.isLess(parent(i), i) {
    this.swap(i, parent(i))
    i = parent(i)
  }
}

func (this *Heap) fixDown(i int) {
  for this.isDefined(left(i)) {
    max := this.max(left(i), right(i))
    if this.isLess(i, max) {
      this.swap(i, max)
      i = max
    } else {
      return
    }
  }
}

func (this *Heap) isLess(i, j int) bool {
  return this.comparator(this.as[i], this.as[j]) < 0
}

func (this *Heap) isDefined(i int) bool {
  return i < this.as.Length()
}

func (this *Heap) max(i, j int) int {
  if this.isDefined(i) {
    if this.isDefined(j) && this.isLess(i, j) {
      return j
    }
    return i
  }
  return j
}

func parent(i int) int {
  return (i - 1) / 2
}

func left(i int) int {
  return 2 * i + 1
}

func right(i int) int {
  return 2 * i + 2
}

type heapIterator struct {
  heap Heap
}

func (this *heapIterator) HasNext() bool {
  return this.heap.Length() > 0
}

func (this *heapIterator) Next() interface{} {
  return this.heap.Next()
}
