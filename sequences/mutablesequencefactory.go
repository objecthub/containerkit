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


type MutableSequenceFactoryBase interface {
  Class() MutableSequenceClass
}

type MutableSequenceFactoryContext interface {
  Sequence
  MutableSequenceFactoryBase
}

type MutableSequenceFactoryDerived interface {
  Copy() MutableSequence
  Immutable() Sequence
  Project(f func (interface{}) interface{}) MutableSequence
  DropIf(pred func (interface{}) bool) MutableSequence
}

type MutableSequenceFactory interface {
  MutableSequenceFactoryContext
  MutableSequenceFactoryDerived
}

func EmbeddedMutableSequenceFactory(obj MutableSequenceFactory) MutableSequenceFactory {
  return &sequenceFactory{obj, obj}
}

type sequenceFactory struct {
  obj MutableSequenceFactory
  MutableSequenceFactoryContext
}

func (this *sequenceFactory) Copy() MutableSequence {
  n := this.obj.Size()
  seq := this.obj.Class().New()
  seq.Allocate(0, n, nil)
  for i := 0; i < n; i++ {
    seq.Set(i, this.obj.At(i))
  }
  return seq
}

func (this *sequenceFactory) Immutable() Sequence {
  return wrappedSequence(this.obj.Copy(), true)
}

func (this *sequenceFactory) Project(f func (interface{}) interface{}) MutableSequence {
  n := this.obj.Size()
  seq := this.obj.Class().New()
  seq.Allocate(0, n, nil)
  for i := 0; i < n; i++ {
    seq.Set(i, f(this.obj.At(i)))
  }
  return seq
}

func (this *sequenceFactory) DropIf(pred func (interface{}) bool) MutableSequence {
  n := this.obj.Size()
  seq := this.obj.Class().New()
  for i := 0; i < n; i++ {
    if !pred(this.obj.At(i)) {
      seq.Allocate(seq.Size(), 1, this.obj.At(i))
    }
  }
  return seq
}
