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


var NativeMap *nativeMapClass = &nativeMapClass{}

var ImmutableNativeMap MapClass = ImmutableMap(NativeMap)

type nativeMapClass struct {}

func (this *nativeMapClass) Embed(obj MutableMap) MutableMap {
  res := new(nativeMap)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableMapDerived = EmbeddedMutableMap(obj)
  res.nmap = make(map[interface{}] interface{})
  return res
}

func (this *nativeMapClass) New(entries... MapEntry) MutableMap {
  res := this.Embed(nil)
  res.IncludeEntry(entries...)
  return res
}

func (this *nativeMapClass) From(coll Container) MutableMap {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

func (this *nativeMapClass) FromNative(mp map[interface{}] interface{}) MutableMap {
  res := this.Embed(nil)
  res.IncludeFromNative(mp)
  return res
}

type MutableNativeMap interface {
  MutableMap
}

type nativeMap struct {
  obj MutableMap
  nmap map[interface{}] interface{}
  MutableMapDerived
}

func (this *nativeMap) Size() int {
  return len(this.nmap)
}

func (this *nativeMap) Get(key interface{}) (value interface{}, exists bool) {
  if entry, exists := this.nmap[key]; exists {
    return entry, true
  }
  return nil, false
}

func (this *nativeMap) Elements() Iterator {
  var keys []interface{}
  for key := range this.nmap {
    keys = append(keys, key)
  }
  return &nativeMapIterator{this.nmap, keys, 0}
}

func (this *nativeMap) Class() MutableMapClass {
  return NativeMap
}

func (this *nativeMap) Include(key, value interface{}) {
  this.nmap[key] = value
}

func (this *nativeMap) Exclude(keys ...interface{}) {
  for _, key := range keys {
    delete(this.nmap, key)
  }
}

func (this *nativeMap) Clear() {
  this.nmap = make(map[interface{}] interface{})
}

type nativeMapIterator struct {
  nmap map[interface{}] interface{}
  keys []interface{}
  i int
}

func (this *nativeMapIterator) HasNext() bool {
  return this.i < len(this.nmap)
}

func (this *nativeMapIterator) Next() interface{} {
  this.i++
  return this.nmap[this.keys[this.i]]
}
