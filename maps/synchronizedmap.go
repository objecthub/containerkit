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

package maps

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/sets"
import "sync"


func SynchronizedMap(class MutableMapClass) MutableMapClass {
  return &synchronizedMapClass{class}
}

type synchronizedMapClass struct {
  class MutableMapClass
}

func (this *synchronizedMapClass) Embed(obj MutableMap) MutableMap {
  res := new(synchronizedMap)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.class = this
  res.unsync = this.class.Embed(obj)
  res.unsynchronizedMap = res.unsync
  return res
}

func (this *synchronizedMapClass) New(entries... MapEntry) MutableMap {
  res := this.Embed(nil)
  res.IncludeEntry(entries...)
  return res
}

func (this *synchronizedMapClass) From(coll Container) MutableMap {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

func (this *synchronizedMapClass) FromNative(mp map[interface{}] interface{}) MutableMap {
  res := this.Embed(nil)
  res.IncludeFromNative(mp)
  return res
}

type unsynchronizedMap interface {
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
  Immutable() Map
  Class() MutableMapClass
  Func() func (interface{}) interface{}
  GetString(key interface{}) string
  ReadOnly() DependentMap
  Keys() DependentContainer
  Values() DependentContainer
  KeySet() DependentSet
  RestrictTo(domain Set) DependentMap
  MapValues(f Mapping) DependentMap
  Override(base Map) DependentMap
}

type synchronizedMap struct {
  obj MutableMap
  class MutableMapClass
  mutex sync.RWMutex
  unsync MutableMap
  unsynchronizedMap
}

func (this *synchronizedMap) Size() int {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Size()
}

func (this *synchronizedMap) Get(key interface{}) (value interface{}, exists bool) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Get(key)
}

func (this *synchronizedMap) HasKey(key interface{}) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.HasKey(key)
}

func (this *synchronizedMap) GetValue(key interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.GetValue(key)
}

func (this *synchronizedMap) Include(key, value interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Include(key, value)
}

func (this *synchronizedMap) Exclude(key... interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Exclude(key...)
}

func (this *synchronizedMap) IncludeEntry(entries... MapEntry) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.IncludeEntry(entries...)
}

func (this *synchronizedMap) IncludeFrom(entries Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.IncludeFrom(entries)
}

func (this *synchronizedMap) IncludeFromNative(mp map[interface{}] interface{}) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.IncludeFromNative(mp)
}

func (this *synchronizedMap) ExcludeKeys(keys Container) {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.ExcludeKeys(keys)
}

func (this *synchronizedMap) ProjectValues(proj Mapping) MutableMap {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ProjectValues(proj)
}

func (this *synchronizedMap) FilterKeys(pred Predicate) MutableMap {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FilterKeys(pred)
}

func (this *synchronizedMap) Class() MutableMapClass {
  return this.class
}

func (this *synchronizedMap) Copy() MutableMap {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Copy()
}

func (this *synchronizedMap) IsEmpty() bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.IsEmpty()
}

func (this *synchronizedMap) Exists(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Exists(pred)
}

func (this *synchronizedMap) ForAll(pred Predicate) bool {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.ForAll(pred)
}

func (this *synchronizedMap) ForEach(proc Procedure) {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  this.unsync.ForEach(proc)
}

func (this *synchronizedMap) FoldLeft(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldLeft(f, z)
}

func (this *synchronizedMap) FoldRight(f Binop, z interface{}) interface{} {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.FoldRight(f, z)
}

func (this *synchronizedMap) Force() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Force()
}

func (this *synchronizedMap) Freeze() FiniteContainer {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.Freeze()
}

func (this *synchronizedMap) Clear() {
  this.mutex.Lock()
  defer this.mutex.Unlock()
  this.unsync.Clear()
}

func (this *synchronizedMap) String() string {
  this.mutex.RLock()
  defer this.mutex.RUnlock()
  return this.unsync.String()
}
