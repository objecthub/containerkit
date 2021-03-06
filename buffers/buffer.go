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

package buffers

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/sequences"


type Buffer interface {
  BufferBase
  BufferDerived
}

type BufferBase interface {
  SequenceBase
  Append(elem ...interface{}) Buffer
  Class() BufferClass
  Clear() Buffer
}

type BufferDerived interface {
  SequenceDerived
  AppendFrom(coll Container) Buffer
  Copy() Buffer
}

// BufferClass defines the interface for embedding and
// instantiating implementations of the Buffer interface.
type BufferClass interface {
  Embed(obj Buffer) Buffer
  New(elements... interface{}) Buffer
  From(coll Container) Buffer
}

func EmbeddedBuffer(obj Buffer) Buffer {
  return &buffer{obj, obj, EmbeddedSequence(obj)}
}

type buffer struct {
  obj Buffer
  BufferBase
  SequenceDerived
}

func (this *buffer) AppendFrom(coll Container) Buffer {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Append(iter.Next())
  }
  return this.obj
}

func (this *buffer) Copy() Buffer {
  return this.obj.Class().From(this.obj)
}
