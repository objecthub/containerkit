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

package buffers

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/impl"


var PriorityQueue *priorityQueueClass = PriorityQueueClass(UniversalComparison)

var ReversePriorityQueue *priorityQueueClass =
    PriorityQueueClass(InvertComparison(UniversalComparison))

func PriorityQueueClass(comp Comparison) *priorityQueueClass {
  return &priorityQueueClass{comp}
}

type priorityQueueClass struct {
  comp Comparison
}

func (this *priorityQueueClass) Embed(obj Queue) Queue {
  res := new(priorityQueue)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.QueueDerived = EmbeddedQueue(obj)
  res.heap = NewHeap(this.comp)
  return res
}

func (this *priorityQueueClass) New(elements... interface{}) Queue {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Enqueue(elements[i])
  }
  return res
}

func (this *priorityQueueClass) From(coll Container) Queue {
  res := this.Embed(nil)
  res.EnqueueFrom(coll)
  return res
}

type priorityQueue struct {
  obj Queue
  QueueDerived
  heap *Heap
}

func (this *priorityQueue) Size() int {
  return this.heap.Length()
}

func (this *priorityQueue) Elements() Iterator {
  return this.heap.Iterator()
}

func (this *priorityQueue) Enqueue(elem interface{}) {
  this.heap.Add(elem)
}

func (this *priorityQueue) Dequeue() interface{} {
  return this.heap.Next()
}

func (this *priorityQueue) Peek() interface{} {
  return this.heap.First()
}

func (this *priorityQueue) Clear() {
  this.heap = NewHeap(this.heap.Comparison())
}

func (this *priorityQueue) Class() QueueClass {
  return PriorityQueueClass(this.heap.Comparison())
}
