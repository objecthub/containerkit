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
import . "github.com/objecthub/containerkit/impl"
import "github.com/objecthub/containerkit/util"


var HashMap MutableMapClass = HashMapClass(util.UniversalHash, util.UniversalEquality)

var ImmutableHashMap MapClass = ImmutableMap(HashMap)

func HashMapClass(hash Hashfunction, equals Equality) MutableMapClass {
  return &hashMapClass{hash, equals}
}

type hashMapClass struct {
  hash Hashfunction
  equals Equality
}

func (this *hashMapClass) Embed(obj MutableMap) MutableMap {
  res := new(hashMap)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableMapDerived = EmbeddedMutableMap(obj)
  res.table = NewHashTable(17, 80, this.hash, this.equals)
  return res
}

func (this *hashMapClass) New(entries... MapEntry) MutableMap {
  res := this.Embed(nil)
  res.IncludeEntry(entries...)
  return res
}

func (this *hashMapClass) From(coll Container) MutableMap {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

type MutableHashMap interface {
  MutableMap
  Print()
}

type hashMap struct {
  obj MutableMap
  table *HashTable
  MutableMapDerived
}

func (this *hashMap) Size() int {
  return this.table.Size()
}

func (this *hashMap) Get(key interface{}) (value interface{}, exists bool) {
  if entry := this.table.FindEntry(key); entry != nil {
    return entry.Value, true
  }
  return nil, false
}

func (this *hashMap) Elements() Iterator {
  return &hashMapIterator{this.table.Iterator()}
}

func (this *hashMap) Class() MutableMapClass {
  return HashMapClass(this.table.Hash(), this.table.Equality())
}

func (this *hashMap) Include(key, value interface{}) {
  if entry := this.table.FindEntry(key); entry == nil {
    this.table.AddEntry(key, value)
  } else {
    entry.Value = value
  }
}

func (this *hashMap) Exclude(keys ...interface{}) {
  for _, key := range keys {
    this.table.DeleteEntry(key)
  }
}

func (this *hashMap) Clear() {
  this.table.Clear()
}

type hashMapIterator struct {
  hashEntryIter *HashEntryIterator
}

func (this *hashMapIterator) HasNext() bool {
  return this.hashEntryIter.HasNext()
}

func (this *hashMapIterator) Next() interface{} {
  entry := this.hashEntryIter.Next()
  return KV(entry.Key, entry.Value)
}

func (this *hashMap) Print() {
  this.table.Print()
}
