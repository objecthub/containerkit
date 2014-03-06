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


func SynchronizedStack(class StackClass) StackClass {
  return &synchronizedStackClass{class}
}

type synchronizedStackClass struct {
  class StackClass
}

func (this *synchronizedStackClass) Embed(obj Stack) Stack {
  res := new(synchronizedStack)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedStack = res.unsync
  return res
}

func (this *synchronizedStackClass) New(elements... interface{}) Stack {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Push(elements[i])
  }
  return res
}

func (this *synchronizedStackClass) From(coll Container) Stack {
  res := this.Embed(nil)
  for iter := coll.Elements(); iter.HasNext(); {
    res.Push(iter.Next())
  }
  return res
}

type unsynchronizedStack interface {
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
  Class() StackClass
}

type synchronizedStack struct {
  obj Stack
  class StackClass
  mutex sync.RWMutex
  unsync Stack
  unsynchronizedStack
}

func (this *synchronizedStack) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedStack) Push(elem interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Push(elem)
}

func (this *synchronizedStack) PushFrom(coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.PushFrom(coll)
}

func (this *synchronizedStack) Pop() interface{} {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.Pop()
}

func (this *synchronizedStack) Peek() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Peek()
}

func (this *synchronizedStack) Clear() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Clear()
}

func (this *synchronizedStack) Class() StackClass {
  return this.class
}

func (this *synchronizedStack) Copy() Stack {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedStack) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedStack) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedStack) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedStack) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedStack) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedStack) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedStack) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedStack) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedStack) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
