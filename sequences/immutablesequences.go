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


func ImmutableSequence(class MutableSequenceClass) SequenceClass {
  return &immutableSequenceClass{class}
}

type immutableSequenceClass struct {
  class MutableSequenceClass
}

func (this *immutableSequenceClass) newImmutableSequence() (Sequence, MutableSequence) {
  encapsulated := this.class.New()
  res := new(immutableSequence)
  res.Sequence = encapsulated
  return res, encapsulated
}

func (this *immutableSequenceClass) Embed(obj Sequence) Sequence {
  res, _ := this.newImmutableSequence()
  return res
}

func (this *immutableSequenceClass) New(elements... interface{}) Sequence {
  res, encapsulated := this.newImmutableSequence()
  encapsulated.Append(elements...)
  return res
}

func (this *immutableSequenceClass) From(coll Container) Sequence {
  res, encapsulated := this.newImmutableSequence()
  encapsulated.AppendFrom(coll)
  return res
}

type immutableSequence struct {
  Sequence
}

func (this *immutableSequence) ReadOnly() DependentSequence {
  return this
}

func (this *immutableSequence) Freeze() FiniteContainer {
  return this
}
