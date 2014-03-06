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
import . "github.com/objecthub/containerkit/sequences"


func SynchronizedBuffer(class BufferClass) BufferClass {
  return &synchronizedBufferClass{class}
}

type synchronizedBufferClass struct {
  class BufferClass
}

func (this *synchronizedBufferClass) Embed(obj Buffer) Buffer {
  res := new(synchronizedBuffer)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedBuffer = res.unsync
  return res
}

func (this *synchronizedBufferClass) New(elements ...interface{}) Buffer {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *synchronizedBufferClass) From(coll Container) Buffer {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type unsynchronizedBuffer interface {
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
  Class() BufferClass
  ReadOnly() DependentSequence
  Subsequence(start int, maxSize int) DependentSequence
  MapValues(f Mapping) DependentSequence
  Reverse() DependentSequence
  Join(other Sequence) DependentSequence
  Func() func (int) interface{}
}

type synchronizedBuffer struct {
  obj Buffer
  class BufferClass
  mutex sync.RWMutex
  unsync Buffer
  unsynchronizedBuffer
}

func (this *synchronizedBuffer) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedBuffer) At(index int) interface{} {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.At(index)
}

func (this *synchronizedBuffer) First() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.First()
}

func (this *synchronizedBuffer) Last() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Last()
}

func (this *synchronizedBuffer) NextIndexEq(start int, val interface{}) int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NextIndexEq(start, val)
}

func (this *synchronizedBuffer) NextIndex(start int, pred func (interface{}) bool) int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NextIndex(start, pred)
}

func (this *synchronizedBuffer) NativeMap() map[int] interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NativeMap()
}

func (this *synchronizedBuffer) Array() []interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Array()
}

func (this *synchronizedBuffer) Append(elem... interface{}) Buffer {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.Append(elem...)
}

func (this *synchronizedBuffer) AppendFrom(coll Container) Buffer {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.AppendFrom(coll)
}

func (this *synchronizedBuffer) Clear() Buffer {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.Clear()
}

func (this *synchronizedBuffer) Class() BufferClass {
  return this.class
}

func (this *synchronizedBuffer) Copy() Buffer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedBuffer) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedBuffer) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedBuffer) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedBuffer) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedBuffer) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedBuffer) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedBuffer) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedBuffer) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedBuffer) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
