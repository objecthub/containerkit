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

import "sync"
import . "github.com/objecthub/containerkit"


func SynchronizedQueue(class QueueClass) QueueClass {
  return &synchronizedQueueClass{class}
}

type synchronizedQueueClass struct {
  class QueueClass
}

func (this *synchronizedQueueClass) Embed(obj Queue) Queue {
  res := new(synchronizedQueue)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedQueue = res.unsync
  return res
}

func (this *synchronizedQueueClass) New(elements... interface{}) Queue {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Enqueue(elements[i])
  }
  return res
}

func (this *synchronizedQueueClass) From(coll Container) Queue {
  res := this.Embed(nil)
  res.EnqueueFrom(coll)
  return res
}

type unsynchronizedQueue interface {
  Elements() Iterator
  Take(n int) DependentContainer
  TakeWhile(pred Predicate) DependentContainer
  Drop(n int) DependentContainer
  DropWhile(pred Predicate) DependentContainer
  Filter(pred Predicate) DependentContainer
  Map(f Mapping) DependentContainer
  FlatMap(g Generator) DependentContainer
  Flatten() DependentContainer
  Concat(other Container) DependentContainer
  Combine(f Binop, other Container) DependentContainer
  Zip(other Container) DependentContainer
  Class() QueueClass
}

type synchronizedQueue struct {
  obj Queue
  class QueueClass
  mutex sync.RWMutex
  unsync Queue
  unsynchronizedQueue
}

func (this *synchronizedQueue) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedQueue) Enqueue(elem interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Enqueue(elem)
}

func (this *synchronizedQueue) Dequeue() interface{} {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.Dequeue()
}

func (this *synchronizedQueue) Peek() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Peek()
}

func (this *synchronizedQueue) EnqueueFrom(coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.EnqueueFrom(coll)
}

func (this *synchronizedQueue) Clear() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Clear()
}

func (this *synchronizedQueue) Class() QueueClass {
  return this.class
}

func (this *synchronizedQueue) Copy() Queue {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedQueue) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedQueue) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedQueue) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedQueue) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedQueue) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedQueue) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedQueue) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedQueue) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedQueue) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
