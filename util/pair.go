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

package util


func NewPair(first, second interface{}) *Pair {
  return &Pair{first, second}
}

func PairBinop(first, second interface{}) interface{} {
  return NewPair(first, second)
}

func AsPair(val interface{}) *Pair {
  return val.(*Pair)
}

type Pair struct {
  first interface{}
  second interface{}
}

func (this *Pair) Equals(other interface{}) bool {
  if that, valid := other.(Pair); valid {
    return UniversalEquality(this.first, that.first) &&
           UniversalEquality(this.second, that.second)
  }
  return false
}

func (this *Pair) Compare(other interface{}) int {
  if that, valid := other.(Pair); valid {
    fstcmp := UniversalComparison(this.first, that.first)
    if (fstcmp == 0) {
      return UniversalComparison(this.second, that.second)
    } else {
      return fstcmp
    }
  }
  panic("Pair.Compare: uncomparable values")
}

func (this *Pair) HashCode() int {
  return UniversalHash(this.first) * 31 + UniversalHash(this.second)
}

func (this *Pair) String() string {
  return NewStringBuilder("(", this.first, ", ", this.second, ")").String()
}

func (this *Pair) Get() (first interface{}, second interface{}) {
  return this.first, this.second
}

func (this *Pair) Key() interface{} {
  return this.first
}

func (this *Pair) Value() interface{} {
  return this.second
}

func (this *Pair) First() interface{} {
  return this.first
}

func (this *Pair) Second() interface{} {
  return this.second
}

func (this *Pair) Swap() *Pair {
  return NewPair(this.second, this.first)
}

func (this *Pair) Pair() *Pair {
  return this
}

