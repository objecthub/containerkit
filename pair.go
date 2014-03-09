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

package containerkit

import "github.com/objecthub/containerkit/util"


// A Pair represents a tuple consisting of two components: first and second.
type Pair interface {
  Equals(other interface{}) bool
  Compare(other interface{}) int
  HashCode() int
  String() string
  Get() (first interface{}, second interface{})
  Key() interface{}
  Value() interface{}
  First() interface{}
  Second() interface{}
  Swap() Pair
  Pair() Pair
}

type pair struct {
  first interface{}
  second interface{}
}

func NewPair(first, second interface{}) Pair {
  return &pair{first, second}
}

func PairBinop(first, second interface{}) interface{} {
  return NewPair(first, second)
}

func AsPair(val interface{}) Pair {
  return val.(Pair)
}

func (this *pair) Equals(other interface{}) bool {
  if that, valid := other.(Pair); valid {
    return util.UniversalEquality(this.first, that.First()) &&
           util.UniversalEquality(this.second, that.Second())
  }
  return false
}

func (this *pair) Compare(other interface{}) int {
  if that, valid := other.(Pair); valid {
    fstcmp := util.UniversalComparison(this.first, that.First())
    if (fstcmp == 0) {
      return util.UniversalComparison(this.second, that.Second())
    } else {
      return fstcmp
    }
  }
  panic("pair.Compare: uncomparable values")
}

func (this *pair) HashCode() int {
  return util.UniversalHash(this.first) * 31 + util.UniversalHash(this.second)
}

func (this *pair) String() string {
  return util.NewStringBuilder("(", this.first, ", ", this.second, ")").String()
}

func (this *pair) Get() (first interface{}, second interface{}) {
  return this.first, this.second
}

func (this *pair) Key() interface{} {
  return this.first
}

func (this *pair) Value() interface{} {
  return this.second
}

func (this *pair) First() interface{} {
  return this.first
}

func (this *pair) Second() interface{} {
  return this.second
}

func (this *pair) Swap() Pair {
  return NewPair(this.second, this.first)
}

func (this *pair) Pair() Pair {
  return this
}

