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


type StackBase interface {
  FiniteContainerBase
  Push(elem interface{})
  Pop() interface{}
  Peek() interface{}
  Class() StackClass
  Clear()
}

type StackDerived interface {
  FiniteContainerDerived
  PushFrom(coll Container)
  Copy() Stack
}

type Stack interface {
  StackBase
  StackDerived
}

// StackClass defines the interface for embedding and
// instantiating implementations of the Stack interface.
type StackClass interface {
  Embed(obj Stack) Stack
  New(elements... interface{}) Stack
  From(coll Container) Stack
}

func EmbeddedStack(obj Stack) Stack {
  return &stackTrait{obj, obj, EmbeddedFiniteContainer(obj)}
}

type stackTrait struct {
  obj Stack
  StackBase
  FiniteContainerDerived
}

func (this *stackTrait) PushFrom(coll Container) {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Push(iter.Next())
  }
}

func (this *stackTrait) Copy() Stack {
  return this.obj.Class().From(this.obj)
}

func (this *stackTrait) Force() FiniteContainer {
  return this.obj
}

func (this *stackTrait) String() string {
  return "[" + this.FiniteContainerDerived.String() + "]"
}
