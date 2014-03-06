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

package impl

import "fmt"
import . "github.com/objecthub/containerkit"
import "github.com/objecthub/containerkit/util"


type HashEntry struct {
  Key interface{}
  Value interface{}
  Next *HashEntry
}

func NewHashEntry(key, value interface{}, next *HashEntry) *HashEntry {
  return &HashEntry{key, value, next}
}

type HashTable struct {
  table [](*HashEntry)
  entries int
  maxLoadFactor int
  hash Hashfunction
  equals Equality
}

func NewHashTable(size int, maxLoadFactor int,
                  hash Hashfunction, equals Equality) *HashTable {
  if (size < 11) {
    size = 11
  }
  return &HashTable{make([](*HashEntry), size),
                    0,
                    maxLoadFactor,
                    hash,
                    equals}
}

func (this *HashTable) Size() int {
  return this.entries
}

func (this *HashTable) Hash() Hashfunction {
  return this.hash
}

func (this *HashTable) Equality() Equality {
  return this.equals
}

func (this *HashTable) tableFull() bool {
  return this.entries * 100 > len(this.table) * this.maxLoadFactor
}

func (this *HashTable) resize(newSize int) {
  newTable := make([](*HashEntry), newSize)
  for i := len(this.table) - 1; i >= 0; i-- {
    for entry := this.table[i]; entry != nil; {
      next := entry.Next
      b := this.bucketInTable(entry.Key, newSize)
      entry.Next = newTable[b]
      newTable[b] = entry
      entry = next
    }
  }
  this.table = newTable
}

func (this *HashTable) resizeIfNeeded() {
  if this.tableFull() {
    this.resize(len(this.table) * 2 + 1)
  }
}

func (this *HashTable) Clear() {
  for i := len(this.table) - 1; i >= 0; i-- {
    this.table[i] = nil
  }
  this.entries = 0
}

func (this *HashTable) bucket(key interface{}) int {
  return this.bucketInTable(key, len(this.table))
}

func (this *HashTable) bucketInTable(key interface{}, tableSize int) int {
  res := this.hash(key) % tableSize
  if res < 0 {
    return -res
  }
  return res
}

func (this *HashTable) FindEntry(key interface{}) *HashEntry {
  entry := this.table[this.bucket(key)]
  for ; entry != nil && !this.equals(key, entry.Key); entry = entry.Next {}
  return entry
}

func (this *HashTable) AddEntry(key, value interface{}) {
  b := this.bucket(key)
  this.table[b] = NewHashEntry(key, value, this.table[b])
  this.entries++
  this.resizeIfNeeded()
}

func (this *HashTable) DeleteEntry(key interface{}) {
  b := this.bucket(key)
  if entry := this.table[b]; entry != nil {
    if this.equals(key, entry.Key) {
      this.table[b] = entry.Next
      this.entries--
    } else {
      for ; entry.Next != nil; entry = entry.Next {
        if this.equals(key, entry.Next.Key) {
          entry.Next = entry.Next.Next
          this.entries--
          return
        }
      }
    }
  }
}

func (this *HashTable) Iterator() *HashEntryIterator {
  firstBucket := len(this.table) - 1
  return &HashEntryIterator{this.table, firstBucket, this.table[firstBucket]}
}

type HashEntryIterator struct {
  table [](*HashEntry)
  currentBucket int
  nextEntry *HashEntry
}

func (this *HashEntryIterator) HasNext() bool {
  if this.nextEntry != nil {
    return true
  }
  this.currentBucket--
  for ; this.currentBucket >= 0; this.currentBucket-- {
    this.nextEntry = this.table[this.currentBucket]
    if this.nextEntry != nil {
      return true
    }
  }
  return false
}

func (this *HashEntryIterator) Next() *HashEntry {
  if this.HasNext() {
    entry := this.nextEntry
    this.nextEntry = entry.Next
    return entry
  }
  panic("HashEntryIterator.next: No next entry")
}

func (this *HashTable) Print() {
  builder := util.NewStringBuilder("<<HashTable entries = ")
  builder.Append(this.entries,
                 ", maxLoadFactor = ",
                 this.maxLoadFactor)
  fmt.Println(builder.String())
  for i := 0; i < len(this.table); i++ {
    builder.Clear()
    builder.Append("  table[", i, "] = [");
    entry := this.table[i]
    if (entry != nil) {
      builder.Append(entry.Key, " -> ", entry.Value)
      entry = entry.Next
    }
    for ; entry != nil; entry = entry.Next {
      builder.Append(", ", entry.Key, "->", entry.Value)
    }
    builder.Append("]")
    fmt.Println(builder.String())
  }
  fmt.Println(">>")
}
