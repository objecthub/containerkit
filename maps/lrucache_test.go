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

import "testing"
import "github.com/objecthub/containerkit/util"
import . "github.com/objecthub/containerkit"


func checkCacheEntry(t *testing.T,
                     cache Cache,
                     key interface{},
                     value interface{},
                     size int) {
  if cache.Size() != size {
    t.Errorf("Expected size of cache: %d; was %d", size, cache.Size())
  }
  if val, exists := cache.Get(key); exists {
    if value != nil {
      if !UniversalEquality(val, value) {
        t.Errorf("Expected value '%s' for key '%s', but found '%s' in cache",
                 util.ToString(value),
                 util.ToString(key),
                 util.ToString(val))
      }
    } else {
      t.Errorf("Did not expect key '%s' in cache; found value '%s'",
               util.ToString(key),
               util.ToString(val))
    }
  } else {
    if value != nil {
      t.Errorf("Key '%s' does not exist in cache; is suppoesed to map to '%s'",
               util.ToString(key),
               util.ToString(value))
    }
  }
}

func TestLruCacheClass(t *testing.T) {
  count := 0
  cache := LruCache.NewWithCallback(8, func (kv MapEntry) { count++ })
  checkCacheEntry(t, cache, "k1", nil, 0)
  cache.Add("k1", "v1")
  checkCacheEntry(t, cache, "k1", "v1", 1)
  cache.Add("k2", "v2")
  cache.Add("k3", "v3")
  checkCacheEntry(t, cache, "k1", "v1", 3)
  checkCacheEntry(t, cache, "k2", "v2", 3)
  checkCacheEntry(t, cache, "k4", nil, 3)
  cache.Add("k4", "v4")
  cache.Add("k5", "v5")
  checkCacheEntry(t, cache, "k4", "v4", 5)
  checkCacheEntry(t, cache, "k5", "v5", 5)
  checkCacheEntry(t, cache, "k6", nil, 5)
  cache.Add("k6", "v6")
  checkCacheEntry(t, cache, "k6", "v6", 6)
  cache.Remove("k3")
  checkCacheEntry(t, cache, "k3", nil, 5)
  if count != 0 {
    t.Errorf("No eviction exepcted yet (73)")
  }
  cache.Add("k7", "v7")
  cache.Add("k8", "v8")
  checkCacheEntry(t, cache, "k1", "v1", 7)
  cache.Add("k9", "v9")
  checkCacheEntry(t, cache, "k2", "v2", 8)
  if count != 0 {
    t.Errorf("No eviction exepcted yet (81)")
  }
  cache.Add("k10", "v10")
  if count != 1 {
    t.Errorf("1 eviction expected")
  }
  checkCacheEntry(t, cache, "k8", "v8", 8)
  checkCacheEntry(t, cache, "k9", "v9", 8)
  checkCacheEntry(t, cache, "k10", "v10", 8)
  checkCacheEntry(t, cache, "k1", "v1", 8)
  checkCacheEntry(t, cache, "k2", "v2", 8)
  checkCacheEntry(t, cache, "k3", nil, 8)
  cache.Add("k11", "v11")
  cache.Add("k12", "v12")
  if count != 3 {
    t.Errorf("3 evictions expected (96)")
  }
  checkCacheEntry(t, cache, "k11", "v11", 8)
  cache.Remove("k10")
  checkCacheEntry(t, cache, "k12", "v12", 7)
  if count != 3 {
    t.Errorf("3 evictions expected (102)")
  }
}

