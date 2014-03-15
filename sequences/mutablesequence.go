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

package sequences

import "sort"
import . "github.com/objecthub/containerkit"


type MutableSequenceBase interface {
  SequenceBase
  MutableSequenceFactoryBase
  Set(index int, element interface{})
  Allocate(index int, n int, element interface{})
  Delete(index int, n int)
}

type MutableSequenceDerived interface {
  SequenceDerived
  MutableSequenceFactoryDerived
  SetFrom(index int, coll Container)
  Insert(index int, element ...interface{})
  Append(element ...interface{})
  Prepend(element ...interface{})
  InsertFrom(index int, elems Container)
  AppendFrom(elems Container)
  PrependFrom(elems Container)
  Trim(nstart, nend int)
  Modify(f func(interface{}) interface{})
  DeleteIf(pred func (interface{}) bool)
  Swap(i int, j int)
  SortWith(comp Comparison)
  Sort()
  Clear()
}

type MutableSequence interface {
  MutableSequenceBase
  MutableSequenceDerived
}

// MutableSequenceClass defines the interface for embedding and
// instantiating implementations of the MutableSequence interface.
type MutableSequenceClass interface {
  Embed(obj MutableSequence) MutableSequence
  New(elements... interface{}) MutableSequence
  From(coll Container) MutableSequence
}

func EmbeddedMutableSequence(obj MutableSequence) MutableSequence {
  return &mutableSequence{obj, obj,
                          EmbeddedSequence(obj),
                          EmbeddedMutableSequenceFactory(obj)}
}

type mutableSequence struct {
  obj MutableSequence
  MutableSequenceBase
  SequenceDerived
  MutableSequenceFactoryDerived
}

func (this *mutableSequence) SetFrom(index int, coll Container) {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Set(index, iter.Next())
    index++
  }
}

func (this *mutableSequence) Insert(index int, elements ...interface{}) {
  n := len(elements)
  this.obj.Allocate(index, n, nil)
  for i := 0; i < n; i++ {
    this.obj.Set(index + i, elements[i])
  }
}

func (this *mutableSequence) Append(elements ...interface{}) {
  this.obj.Insert(this.obj.Size(), elements...)
}

func (this *mutableSequence) Prepend(elements ...interface{}) {
  this.obj.Insert(0, elements...)
}

func (this *mutableSequence) InsertFrom(index int, elems Container) {
  seq := ArraySequence.New()
  n := 0
  for iter := elems.Elements(); iter.HasNext(); {
    seq.Append(iter.Next())
    n++
  }
  this.obj.Allocate(index, n, nil)
  for i := 0; i < n; i++ {
    this.obj.Set(index + i, seq.At(i))
  }
}

func (this *mutableSequence) AppendFrom(elems Container) {
  this.obj.InsertFrom(this.obj.Size(), elems)
}

func (this *mutableSequence) PrependFrom(elems Container) {
  this.obj.InsertFrom(0, elems)
}

func (this *mutableSequence) Trim(nstart, nend int) {
  if nstart > 0 {
    this.obj.Delete(0, nstart)
  }
  if nend > 0 {
    this.obj.Delete(this.obj.Size() - nend, nend)
  }
}

func (this *mutableSequence) Modify(f func(interface{}) interface{}) {
  for i := 0; i < this.obj.Size(); i++ {
    this.obj.Set(i, f(this.obj.At(i)))
  }
}

func (this *mutableSequence) DeleteIf(pred func (interface{}) bool) {
  i := 0
  del := 0
  for i < this.obj.Size() {
    if pred(this.obj.At(i)) {
      del++
      i++
    } else if del > 0 {
      this.obj.Delete(i - del, del)
      del = 0
      i -= del
    } else {
      i++
    }
  }
  if del > 0 {
    this.obj.Trim(0, del)
  }
}

func (this *mutableSequence) Swap(i int, j int) {
  if i != j {
    h := this.obj.At(i)
    this.obj.Set(i, this.obj.At(j))
    this.obj.Set(j, h)
  }
}

func (this *mutableSequence) Clear() {
  this.obj.Delete(0, this.obj.Size())
}


func (this *mutableSequence) SortWith(comp Comparison) {
  sort.Sort(&sortableSeq{comp, this.obj})
}

func (this *mutableSequence) Sort() {
  this.obj.SortWith(UniversalComparison)
}

type sortableSeq struct {
  comp Comparison
  encapsulated MutableSequence
}

func (this *sortableSeq) Len() int {
  return this.encapsulated.Size()
}

func (this *sortableSeq) Less(i, j int) bool {
  return this.comp(this.encapsulated.At(i), this.encapsulated.At(j)) < 0
}

func (this *sortableSeq) Swap(i, j int) {
  this.encapsulated.Swap(i, j)
}
