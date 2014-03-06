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


var ListQueue QueueClass = &listQueueClass{}

type listQueueClass struct {
}

func (this *listQueueClass) Embed(obj Queue) Queue {
  res := new(listQueue)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.QueueDerived = EmbeddedQueue(obj)
  res.list = NewLinkedList()
  return res
}

func (this *listQueueClass) New(elements... interface{}) Queue {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Enqueue(elements[i])
  }
  return res
}

func (this *listQueueClass) From(coll Container) Queue {
  res := this.Embed(nil)
  res.EnqueueFrom(coll)
  return res
}

type listQueue struct {
  obj Queue
  QueueDerived
  list *LinkedList
}

func (this *listQueue) Size() int {
  return this.list.Length()
}

func (this *listQueue) Elements() Iterator {
  return this.list.Iterator()
}

func (this *listQueue) Enqueue(elem interface{}) {
  this.list.InsertTail(elem)
}

func (this *listQueue) Dequeue() interface{} {
  return this.list.RemoveHead()
}

func (this *listQueue) Peek() interface{} {
  return this.list.GetHead()
}

func (this *listQueue) Clear() {
  this.list = NewLinkedList()
}

func (this *listQueue) Class() QueueClass {
  return ListQueue
}
