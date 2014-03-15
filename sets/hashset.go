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

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/impl"


var HashSet MutableSetClass = HashSetClass(UniversalHash, UniversalEquality)

var ImmutableHashSet SetClass = ImmutableSet(HashSet)

func HashSetClass(hash Hashfunction, equals Equality) MutableSetClass {
  return &hashSetClass{hash, equals}
}

type hashSetClass struct {
  hash Hashfunction
  equals Equality
}

func (this *hashSetClass) Embed(obj MutableSet) MutableSet {
  res := new(hashSet)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableSetDerived = EmbeddedMutableSet(obj)
  res.table = NewHashTable(17, 80, this.hash, this.equals)
  return res
}

func (this *hashSetClass) New(elements ...interface{}) MutableSet {
  res := this.Embed(nil)
  res.Include(elements...)
  return res
}

func (this *hashSetClass) From(coll Container) MutableSet {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

type hashSet struct {
  obj MutableSet
  table *HashTable
  MutableSetDerived
}

func (this *hashSet) Size() int {
  return this.table.Size()
}

func (this *hashSet) Contains(elem interface{}) bool {
  return this.table.FindEntry(elem) != nil
}

func (this *hashSet) Elements() Iterator {
  return &hashSetIterator{this.table.Iterator()}
}

func (this *hashSet) Class() MutableSetClass {
  return HashSetClass(this.table.Hash(), this.table.Equality())
}

func (this *hashSet) Include(elements ...interface{}) {
  for _, key := range elements {
    if this.table.FindEntry(key) == nil {
      this.table.AddEntry(key, nil)
    }
  }
}

func (this *hashSet) Exclude(elements ...interface{}) {
  for _, key := range elements {
    this.table.DeleteEntry(key)
  }
}

func (this *hashSet) Clear() {
  this.table.Clear()
}

type hashSetIterator struct {
  hashEntryIter *HashEntryIterator
}

func (this *hashSetIterator) HasNext() bool {
  return this.hashEntryIter.HasNext()
}

func (this *hashSetIterator) Next() interface{} {
  return this.hashEntryIter.Next().Key
}

func (this *hashSet) Print() {
  this.table.Print()
}
