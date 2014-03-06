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

package sequences

import "sync"
import . "github.com/objecthub/containerkit"


func SynchronizedSequence(class MutableSequenceClass) MutableSequenceClass {
  return &synchronizedSequenceClass{class}
}

type synchronizedSequenceClass struct {
  class MutableSequenceClass
}

func (this *synchronizedSequenceClass) Embed(obj MutableSequence) MutableSequence {
  res := new(synchronizedSequence)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedSequence = res.unsync
  return res
}

func (this *synchronizedSequenceClass) New(elements ...interface{}) MutableSequence {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *synchronizedSequenceClass) From(coll Container) MutableSequence {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type unsynchronizedSequence interface {
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
  Immutable() Sequence
  Class() MutableSequenceClass
  Func() func (int) interface{}
  ReadOnly() DependentSequence
  Subsequence(start int, maxSize int) DependentSequence
  MapValues(f Mapping) DependentSequence
  Reverse() DependentSequence
  Join(other Sequence) DependentSequence
}

type synchronizedSequence struct {
  obj MutableSequence
  class MutableSequenceClass
  mutex sync.RWMutex
  unsync MutableSequence
  unsynchronizedSequence
}

func (this *synchronizedSequence) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedSequence) Set(index int, element interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Set(index, element)
}

func (this *synchronizedSequence) Allocate(index int, n int, element interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Allocate(index, n, element)
}

func (this *synchronizedSequence) Delete(index int, n int) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Delete(index, n)
}

func (this *synchronizedSequence) SetFrom(index int, coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.SetFrom(index, coll)
}

func (this *synchronizedSequence) Insert(index int, element... interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Insert(index, element...)
}

func (this *synchronizedSequence) Append(element... interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Append(element...)
}

func (this *synchronizedSequence) Prepend(element... interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Prepend(element...)
}

func (this *synchronizedSequence) InsertFrom(index int, elems Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.InsertFrom(index, elems)
}

func (this *synchronizedSequence) AppendFrom(elems Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.AppendFrom(elems)
}

func (this *synchronizedSequence) PrependFrom(elems Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.PrependFrom(elems)
}

func (this *synchronizedSequence) Trim(nstart, nend int) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Trim(nstart, nend)
}

func (this *synchronizedSequence) Modify(f func(interface{}) interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Modify(f)
}

func (this *synchronizedSequence) DeleteIf(pred func (interface{}) bool) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.DeleteIf(pred)
}

func (this *synchronizedSequence) Swap(i int, j int) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Swap(i, j)
}

func (this *synchronizedSequence) SortWith(comp Comparison) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.SortWith(comp)
}

func (this *synchronizedSequence) Sort() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Sort()
}

func (this *synchronizedSequence) At(index int) interface{} {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  return this.unsync.At(index)
}

func (this *synchronizedSequence) First() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.First()
}

func (this *synchronizedSequence) Last() interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Last()
}

func (this *synchronizedSequence) NextIndexEq(start int, val interface{}) int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NextIndexEq(start, val)
}

func (this *synchronizedSequence) NextIndex(start int, pred func (interface{}) bool) int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NextIndex(start, pred)
}

func (this *synchronizedSequence) NativeMap() map[int] interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.NativeMap()
}

func (this *synchronizedSequence) Array() []interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Array()
}

func (this *synchronizedSequence) Class() MutableSequenceClass {
  return this.class
}

func (this *synchronizedSequence) Copy() MutableSequence {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedSequence) Project(f func (interface{}) interface{}) MutableSequence {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Project(f)
}

func (this *synchronizedSequence) DropIf(pred func (interface{}) bool) MutableSequence {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.DropIf(pred)
}

func (this *synchronizedSequence) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedSequence) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedSequence) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedSequence) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedSequence) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedSequence) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedSequence) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedSequence) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedSequence) Clear() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Clear()
}

func (this *synchronizedSequence) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
