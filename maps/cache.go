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


type Cache interface {
  CacheBase
  CacheDerived
}

type CacheBase interface {
  FiniteContainerBase
  MapperBase
  Add(key, value interface{})
  Remove(key ...interface{})
}

type CacheDerived interface {
  FiniteContainerDerived
  MapEntryContainerDerived
  MapperDerived
}

// CacheClass defines the functionality of Map implementations,
// ie. records that act as Map factories, providing an Embed, New,
// and From method.
type CacheClass interface {
  Embed(obj Cache, capacity int, whenEvicted func (kv MapEntry)) Cache
  New(capacity int) Cache
  NewWithCallback(capacity int, whenEvicted func (kv MapEntry)) Cache
}

func EmbeddedCache(obj Cache) Cache {
  return &cache{obj,
                obj,
                EmbeddedFiniteContainer(obj),
                EmbeddedMapEntryContainer(obj),
                EmbeddedMapper(obj)}
}
 
type cache struct {
  obj Cache
  CacheBase
  FiniteContainerDerived
  MapEntryContainerDerived
  MapperDerived
}

func (this *cache) String() string {
  return "«" + this.FiniteContainerDerived.String() + "»"
}
