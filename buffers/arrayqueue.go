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


var ArrayQueue *arrayQueueClass = &arrayQueueClass{}

type arrayQueueClass struct {}

func (this *arrayQueueClass) Embed(obj Queue) Queue {
  res := new(arrayQueue)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.QueueDerived = EmbeddedQueue(obj)
  res.array = NewArray()
  res.array.Extend(res.array.Capacity())
  res.first = 0
  res.next = 0
  return res
}

func (this *arrayQueueClass) New(elements... interface{}) Queue {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Enqueue(elements[i])
  }
  return res
}

func (this *arrayQueueClass) From(coll Container) Queue {
  res := this.Embed(nil)
  res.EnqueueFrom(coll)
  return res
}

type arrayQueue struct {
  obj Queue
  QueueDerived
  array Array
  first int
  next int
}

func (this *arrayQueue) Size() int {
  return (this.array.Length() + this.next - this.first) % this.array.Length()
}

func (this *arrayQueue) Elements() Iterator {
  return this.array.ArrayIterator(
      this.first,
      this.next,
      func (i int) int { return (i + 1) % this.array.Length() })
}

func (this *arrayQueue) Enqueue(elem interface{}) {
  this.array.Set(this.next, elem)
  this.next = (this.next + 1) % this.array.Length()
  if this.next == this.first {
    l := this.array.Length()
    this.array.Extend(l * 2)
    this.array.Move(0, l, this.first)
    this.next += l
  }
}

func (this *arrayQueue) Dequeue() interface{} {
  res := this.Peek()
  this.first = (this.first + 1) % this.array.Length()
  return res
}

func (this *arrayQueue) Peek() interface{} {
  if this.next == this.first {
    panic("arrayQueue: queue empty")
  }
  return this.array.At(this.first)
}

func (this *arrayQueue) Clear() {
  this.array.Clear()
  this.first = 0
  this.next = 0
}

func (this *arrayQueue) Class() QueueClass {
  return ArrayQueue
}
