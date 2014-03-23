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


type QueueBase interface {
  FiniteContainerBase
  Enqueue(elem interface{})
  Dequeue() interface{}
  Peek() interface{}
  Class() QueueClass
  Clear()
}

type QueueDerived interface {
  FiniteContainerDerived
  EnqueueFrom(coll Container)
  Copy() Queue
}

type Queue interface {
  QueueBase
  QueueDerived
}

// QueueClass defines the interface for embedding and
// instantiating implementations of the Queue interface.
type QueueClass interface {
  Embed(obj Queue) Queue
  New(elements... interface{}) Queue
  From(coll Container) Queue
}

func EmbeddedQueue(obj Queue) Queue {
  return &queue{obj, obj, EmbeddedFiniteContainer(obj)}
}

type queue struct {
  obj Queue
  QueueBase
  FiniteContainerDerived
}

func (this *queue) EnqueueFrom(coll Container) {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Enqueue(iter.Next())
  }
}

func (this *queue) Copy() Queue {
  return this.obj.Class().From(this.obj)
}

func (this *queue) Force() FiniteContainer {
  return this.obj
}

func (this *queue) String() string {
  return "[" + this.FiniteContainerDerived.String() + "]"
}
