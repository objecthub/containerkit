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


// MutableMapBase defines minimal functionality that is required for
// supporting the full MutableMap interface. Embedding the record defined
// by the EmbeddedMutableMap function allows one to construct a full
// MutableMap implementation from the methods given by the MutableMapBase
// interface.
type MutableMapBase interface {
  MapBase
  MutableMapFactoryBase
  Include(key, value interface{})
  Exclude(key ...interface{})
  Clear()
}

// MutableMapDerived defines methods whose implementation can be generically
// derived from the methods defined by the MutableMapBase interface.
type MutableMapDerived interface {
  MapDerived
  MutableMapFactoryDerived
  IncludeEntry(entries ...MapEntry)
  IncludeFrom(entries Container)
  IncludeFromNative(mp map[interface{}] interface{})
  ExcludeKeys(keys Container)
}

// A MutableMap is a Map that provides functionality for changing the state
// of the map by either including or excluding mappings/map entries.
type MutableMap interface {
  MutableMapBase
  MutableMapDerived
}

// MutableMapClass defines the functionality of MutableMap implementations,
// ie. records that act as MutableMap factories, providing an Embed, New,
// and From method.
type MutableMapClass interface {
  Embed(obj MutableMap) MutableMap
  New(entries... MapEntry) MutableMap
  From(coll Container) MutableMap
  FromNative(mp map[interface{}] interface{}) MutableMap
}

// EmbeddedMutableMap returns a new record that implements the MutableMap
// interface. obj needs to refer to a record which embedds the result of
// this function.
func EmbeddedMutableMap(obj MutableMap) MutableMap {
  return &mutableMap{obj,
                     obj,
                     EmbeddedMap(obj),
                     EmbeddedMutableMapFactory(obj)}
}

type mutableMap struct {
  obj MutableMap
  MutableMapBase
  MapDerived
  MutableMapFactoryDerived
}

func (this *mutableMap) IncludeEntry(entries ...MapEntry) {
  for _, entry := range entries {
    this.obj.Include(entry.Key(), entry.Value())
  }
}

func (this *mutableMap) IncludeFrom(entries Container) {
  for iter := entries.Elements(); iter.HasNext(); {
    if entry, valid := iter.Next().(MapEntry); valid {
      this.obj.IncludeEntry(entry)
    } else {
      panic("mutableMap.IncludeEntry: not a MapEntry value")
    }
  }
}

func (this *mutableMap) IncludeFromNative(mp map[interface{}] (interface{})) {
  for key, value := range mp {
    this.obj.Include(key, value)
  }
}

func (this *mutableMap) ExcludeKeys(keys Container) {
  for iter := keys.Elements(); iter.HasNext(); {
    this.obj.Exclude(iter.Next())
  }
}
