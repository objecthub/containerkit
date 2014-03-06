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

package sets

import . "github.com/objecthub/containerkit"
import "github.com/objecthub/containerkit/util"


// MutableSetBase defines minimal functionality that is required for
// supporting the full MutableSet interface. Embedding the record defined
// by the EmbeddedMutableSet function allows one to construct a full
// MutableSet implementation from the methods given by the MutableSetBase
// interface.
type MutableSetBase interface {
  SetBase
  MutableSetFactoryBase
  Include(elements ...interface{})
  Exclude(elements ...interface{})
  Clear()
}

// MutableSetDerived defines methods whose implementation can be generically
// derived from the methods defined by the MutableSetBase interface.
type MutableSetDerived interface {
  SetDerived
  MutableSetFactoryDerived
  IncludeFrom(coll Container)
  ExcludeFrom(coll Container)
  IntersectWith(coll Container)
  ExcludeIf(pred Predicate)
}

// A MutableSet is a Set that provides functionality for changing the state
// of the set by either including or excluding elements.
type MutableSet interface {
  MutableSetBase
  MutableSetDerived
}

// MutableSetClass defines the functionality of MutableSet implementations,
// ie. records that act as MutableSet factories, providing an Embed, New,
// and From method.
type MutableSetClass interface {
  Embed(obj MutableSet) MutableSet
  New(elements ...interface{}) MutableSet
  From(coll Container) MutableSet
}

// EmbeddedMutableSet returns a new record that implements the MutableSet
// interface. obj needs to refer to a record which embedds the result of
// this function.
func EmbeddedMutableSet(obj MutableSet) MutableSet {
  return &mutableSetTrait{obj,
                          obj,
                          EmbeddedSet(obj),
                          EmbeddedMutableSetFactory(obj)}
}

type mutableSetTrait struct {
  obj MutableSet
  MutableSetBase
  SetDerived
  MutableSetFactoryDerived
}

func (this *mutableSetTrait) IncludeFrom(coll Container) {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Include(iter.Next())
  }
}

func (this *mutableSetTrait) ExcludeFrom(coll Container) {
  for iter := coll.Elements(); iter.HasNext(); {
    this.obj.Exclude(iter.Next())
  }
}

func (this *mutableSetTrait) IntersectWith(coll Container) {
  this.obj.ExcludeIf(util.Negate(this.obj.Class().From(coll).Func()))
}

func (this *mutableSetTrait) ExcludeIf(pred Predicate) {
  this.obj.ExcludeFrom(this.obj.Filter(pred).Force())
}
