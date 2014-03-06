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


func ImmutableMap(class MutableMapClass) MapClass {
  return &immutableMapClass{class}
}

type immutableMapClass struct {
  class MutableMapClass
}

func (this *immutableMapClass) newImmutableMap() (Map, MutableMap) {
  encapsulated := this.class.New()
  res := new(immutableMap)
  res.fMap = encapsulated
  return res, encapsulated
}

func (this *immutableMapClass) Embed(obj Map) Map {
  res, _ := this.newImmutableMap()
  return res
}

func (this *immutableMapClass) New(entries... MapEntry) Map {
  res, encapsulated := this.newImmutableMap()
  encapsulated.IncludeEntry(entries...)
  return res
}

func (this *immutableMapClass) From(coll Container) Map {
  res, encapsulated := this.newImmutableMap()
  encapsulated.IncludeFrom(coll)
  return res
}

type fMap interface {
  Map
}

type immutableMap struct {
  fMap
}

func (this *immutableMap) ReadOnly() DependentMap {
  return this
}

func (this *immutableMap) Freeze() FiniteContainer {
  return this
}
