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

package sets

import "sync"
import . "github.com/objecthub/containerkit"


func SynchronizedSet(class MutableSetClass) MutableSetClass {
  return &synchronizedSetClass{class}
}

type synchronizedSetClass struct {
  class MutableSetClass
}

func (this *synchronizedSetClass) Embed(obj MutableSet) MutableSet {
  res := new(synchronizedSet)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedSet = res.unsync
  return res
}

func (this *synchronizedSetClass) New(elements ...interface{}) MutableSet {
  res := this.Embed(nil)
  res.Include(elements...)
  return res
}

func (this *synchronizedSetClass) From(coll Container) MutableSet {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

type unsynchronizedSet interface {
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
  Immutable() Set
  Class() MutableSetClass
  Func() Predicate
  ReadOnly() DependentSet
  Union(set Set) DependentSet
  Intersection(set Set) DependentSet
  Difference(set Set) DependentSet
}

type synchronizedSet struct {
  obj MutableSet
  class MutableSetClass
  mutex sync.RWMutex
  unsync MutableSet
  unsynchronizedSet
}

func (this *synchronizedSet) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedSet) Contains(elem interface{}) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Contains(elem)
}

func (this *synchronizedSet) ContainsAll(elements ...interface{}) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsAll(elements...)
}

func (this *synchronizedSet) ContainsNone(elements ...interface{}) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsNone(elements...)
}

func (this *synchronizedSet) ContainsSome(elements ...interface{}) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsSome(elements...)
}

func (this *synchronizedSet) ContainsAllFrom(elements Container) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsAllFrom(elements)
}

func (this *synchronizedSet) ContainsNoneFrom(elements Container) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsNoneFrom(elements)
}

func (this *synchronizedSet) ContainsSomeFrom(elements Container) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ContainsSomeFrom(elements)
}

func (this *synchronizedSet) Include(elements ...interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Include(elements...)
}

func (this *synchronizedSet) Exclude(elements ...interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Exclude(elements...)
}

func (this *synchronizedSet) IncludeFrom(coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.IncludeFrom(coll)
}

func (this *synchronizedSet) ExcludeFrom(coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.ExcludeFrom(coll)
}

func (this *synchronizedSet) IntersectWith(coll Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.IntersectWith(coll)
}

func (this *synchronizedSet) ExcludeIf(pred Predicate) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.ExcludeIf(pred)
}

func (this *synchronizedSet) Class() MutableSetClass {
  return this.class
}

func (this *synchronizedSet) Copy() MutableSet {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedSet) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedSet) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedSet) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedSet) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedSet) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedSet) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedSet) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedSet) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedSet) Clear() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Clear()
}

func (this *synchronizedSet) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
