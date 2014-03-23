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
import "github.com/objecthub/containerkit/impl"


var LruCache CacheClass = LruCacheClass(UniversalHash, UniversalEquality)

func LruCacheClass(hash Hashfunction, equals Equality) CacheClass {
  return &lruCacheClass{hash, equals}
}

type lruCacheClass struct {
  hash Hashfunction
  equals Equality
}

func (this *lruCacheClass) Embed(obj Cache, capacity int, we func (kv MapEntry)) Cache {
  res := new(lruCache)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.CacheDerived = EmbeddedCache(obj)
  res.capacity = capacity
  res.whenEvicted = we
  res.table = impl.NewHashTable(17, 80, this.hash, this.equals)
  res.accessorder = impl.NewDoubleLinkedList()
  return res
}

func (this *lruCacheClass) New(capacity int) Cache {
  return this.Embed(nil, capacity, nil)
}

func (this *lruCacheClass) NewWithCallback(capacity int, we func (kv MapEntry)) Cache {
  return this.Embed(nil, capacity, we)
}

type lruCache struct {
  obj Cache
  capacity int
  whenEvicted func (kv MapEntry)
  table *impl.HashTable
  accessorder *impl.DoubleLinkedList
  CacheDerived
}

func (this *lruCache) Size() int {
  return this.table.Size()
}

func (this *lruCache) Get(key interface{}) (value interface{}, exists bool) {
  if entry := this.table.FindEntry(key); entry != nil {
    elem := entry.Value.(*impl.Element)
    this.accessorder.Remove(elem)
    this.accessorder.InsertFront(elem)
    return elem.Value.(MapEntry).Value(), true
  }
  return nil, false
}

func (this *lruCache) Elements() Iterator {
  return &lruCacheIterator{this.table.Iterator()}
}

func (this *lruCache) Class() CacheClass {
  return LruCacheClass(this.table.Hash(), this.table.Equality())
}

func (this *lruCache) Add(key, value interface{}) {
  kv := KV(key, value)
  if entry := this.table.FindEntry(key); entry == nil {
    elem := impl.NewElement(kv)
    this.table.AddEntry(key, elem)
    this.accessorder.InsertFront(elem)
  } else {
    elem := entry.Value.(*impl.Element)
    elem.Value = kv
    this.accessorder.Remove(elem)
    this.accessorder.InsertFront(elem)
  }
  if this.table.Size() > this.capacity {
    elem := this.accessorder.Back()
    this.accessorder.Remove(elem)
    this.table.DeleteEntry(elem.Value.(MapEntry).Key())
    if this.whenEvicted != nil {
      this.whenEvicted(elem.Value.(MapEntry))
    }
  }
}

func (this *lruCache) Remove(keys ...interface{}) {
  for _, key := range keys {
    if entry := this.table.FindEntry(key); entry != nil {
      elem := entry.Value.(*impl.Element)
      this.accessorder.Remove(elem)
      this.table.DeleteEntry(key)
    }
  }
}

func (this *lruCache) Clear() {
  this.table.Clear()
  this.accessorder = impl.NewDoubleLinkedList()
}

type lruCacheIterator struct {
  hashEntryIter *impl.HashEntryIterator
}

func (this *lruCacheIterator) HasNext() bool {
  return this.hashEntryIter.HasNext()
}

func (this *lruCacheIterator) Next() interface{} {
  entry := this.hashEntryIter.Next()
  return KV(entry.Key, entry.Value)
}
