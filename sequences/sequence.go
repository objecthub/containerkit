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

import . "github.com/objecthub/containerkit"


type SequenceBase interface {
  IndexedBase
}

type SequenceDerived interface {
  IndexedDerived
  Container
  ReadOnly() DependentSequence
  Subsequence(start int, maxSize int) DependentSequence
  MapValues(f Mapping) DependentSequence
  Reverse() DependentSequence
  Join(other Sequence) DependentSequence
}

type Sequence interface {
  SequenceBase
  SequenceDerived
}

// SequenceClass defines the interface for embedding and
// instantiating implementations of the Sequence interface.
type SequenceClass interface {
  Embed(obj Sequence) Sequence
  New(elements... interface{}) Sequence
  From(coll Container) Sequence
}

func EmbeddedSequence(obj Sequence) Sequence {
  return &sequenceTrait{obj,
                        obj,
                        EmbeddedIndexed(obj),
                        EmbeddedFiniteContainer(obj)}
}

type sequenceTrait struct {
  obj Sequence
  SequenceBase
  IndexedDerived
  ContainerDerived
}

func (this *sequenceTrait) Elements() Iterator {
  return &sequenceIterator{this.obj, 0}
}

func (this *sequenceTrait) Reverse() DependentSequence {
  return newReversedSequence(this.obj)
}

func (this *sequenceTrait) Subsequence(start int, maxSize int) DependentSequence {
  return newSubsequence(this.obj, start, maxSize)
}

func (this *sequenceTrait) MapValues(f Mapping) DependentSequence {
  return newMappedSequence(this.obj, f)
}

func (this *sequenceTrait) ReadOnly() DependentSequence {
  return wrappedSequence(this.obj, false)
}

func (this *sequenceTrait) Join(other Sequence) DependentSequence {
  return newAppendedSequence(this.obj, other)
}

func (this *sequenceTrait) String() string {
  return "[" + this.ContainerDerived.String() + "]"
}

type sequenceIterator struct {
  data Sequence
  index int
}

func (this *sequenceIterator) HasNext() bool {
  return this.index < this.data.Size()
}

func (this *sequenceIterator) Next() interface{} {
  if !this.HasNext() {
    panic("sequenceIterator.Next: no next value")
  }
  res := this.data.At(this.index)
  this.index++
  return res
}
