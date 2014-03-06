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


type DependentSequence interface {
  Sequence
}

func EmbeddedDependentSequence(obj DependentSequence) SequenceDerived {
  return &dependentSequenceTrait{obj, EmbeddedSequence(obj)}
}

type dependentSequenceTrait struct {
  obj DependentSequence
  SequenceDerived
}

func (this *dependentSequenceTrait) ReadOnly() DependentSequence {
  return this.obj
}

func (this *dependentSequenceTrait) String() string {
  return "<" + this.SequenceDerived.String() + ">"
}

// Reversed sequences

func newReversedSequence(sequence Sequence) DependentSequence {
  res := new(reversedSequence)
  res.SequenceDerived = EmbeddedDependentSequence(res)
  res.fst = sequence
  return res
}

type reversedSequence struct {
  SequenceDerived
  fst Sequence
}

func (this *reversedSequence) Size() int {
  return this.fst.Size()
}

func (this *reversedSequence) At(i int) interface{} {
  return this.fst.At(this.fst.Size() - i - 1)
}

// Subsequences

func newSubsequence(sequence Sequence,
                    start int,
                    maxSize int) DependentSequence {
  res := new(subsequence)
  res.SequenceDerived = EmbeddedDependentSequence(res)
  res.fst = sequence
  res.start = start
  res.maxSize = maxSize
  return res
}

type subsequence struct {
  SequenceDerived
  fst Sequence
  start int
  maxSize int
}

func (this *subsequence) Size() int {
  switch size := this.fst.Size() - this.start; {
    case size > this.maxSize:
      return this.maxSize
    case size >= 0:
      return size
  }
  return 0
}

func (this *subsequence) At(i int) interface{} {
  switch {
    case i < 0:
      panic("subsequence.At: index below 0")
    case i - this.start < this.maxSize:
      return this.fst.At(i - this.start)
  }
  panic("subsequence.At: index above size")
}

// Mapped sequences

func newMappedSequence(sequence Sequence, f Mapping) DependentSequence {
  res := new(mappedSequence)
  res.SequenceDerived = EmbeddedDependentSequence(res)
  res.fst = sequence
  res.f = f
  return res
}

type mappedSequence struct {
  SequenceDerived
  fst Sequence
  f Mapping
}

func (this *mappedSequence) Size() int {
  return this.fst.Size()
}

func (this *mappedSequence) At(i int) interface{} {
  return this.f(this.fst.At(i))
}

// Appended sequences

func newAppendedSequence(fst Sequence, snd Sequence) DependentSequence {
  res := new(appendedSequence)
  res.SequenceDerived = EmbeddedDependentSequence(res)
  res.fst = fst
  res.snd = snd
  return res
}

type appendedSequence struct {
  SequenceDerived
  fst Sequence
  snd Sequence
}

func (this *appendedSequence) Size() int {
  return this.fst.Size() + this.snd.Size()
}

func (this *appendedSequence) At(i int) interface{} {
  if i >= this.fst.Size() {
    return this.snd.At(i - this.fst.Size())
  }
  return this.fst.At(i)
}

// Wrapped sequences

func wrappedSequence(encapsulated Sequence, immutable bool) DependentSequence {
  res := new(sequenceWrapper)
  res.SequenceDerived = EmbeddedDependentSequence(res)
  res.encapsulated = encapsulated
  res.immutable = immutable
  return res
}

type sequenceWrapper struct {
  SequenceDerived
  encapsulated Sequence
  immutable bool
}

func (this *sequenceWrapper) Size() int {
  return this.encapsulated.Size()
}

func (this *sequenceWrapper) At(index int) interface{} {
  return this.encapsulated.At(index)
}

func (this *sequenceWrapper) Force() FiniteContainer {
  return this
}

func (this *sequenceWrapper) Freeze() FiniteContainer {
  if this.immutable {
    return this
  }
  return this.SequenceDerived.Freeze()
}
