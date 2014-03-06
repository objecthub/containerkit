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

package impl


func NewArray() Array {
  return make(Array, 0, 8)
}

type Array []interface{}

func (this *Array) Length() int {
  return len(*this)
}

func (this *Array) At(i int) interface{} {
  return (*this)[i]
}

func (this *Array) Set(i int, val interface{}) {
  (*this)[i] = val
}

func (this *Array) Append(val interface{}) {
  i := len(*this)
  this.Allocate(i, 1)
  (*this)[i] = val
}

func (this *Array) Allocate(i int, n int) {
  if n < 0 {
    panic("Array.insert: negative number of elements to insert")
  }
  if i < 0 {
    panic("Array.insert: illegal index")
  }
  l := len(*this) + n
  cp := cap(*this)
  if l > cp {
    if l > (cp * 2) {
      cp = l + 1
    } else {
      cp *= 2
    }
    a := make(Array, l, cp)
    copy(a[:i], (*this)[:i])
    copy(a[i + n:l], (*this)[i:])
    *this = a
  } else {
    *this = (*this)[:l]
    copy((*this)[i + n:l], (*this)[i:])
  }
}

func (this *Array) Delete(i int, n int) {
  l := len(*this) - n
  if l < 0 {
    panic("Array.delete: n too large")
  } else if l == 0 {
    *this = (*this)[:0]
  } else {
    copy((*this)[i:l], (*this)[i + n:])
    *this = (*this)[:l]
  }
}

func (this *Array) Copy() Array {
  a := make(Array, len(*this))
	copy(a, *this)
	return a
}

func (this *Array) Iterator() *arrayIterator {
  return &arrayIterator{*this, 0}
}

type arrayIterator struct {
  a Array
  i int
}

func (this *arrayIterator) HasNext() bool {
  return this.i < len(this.a)
}

func (this *arrayIterator) Next() interface{} {
  if this.HasNext() {
    res := this.a[this.i]
    this.i++
    return res
  }
  panic("ArrayIterator.Next: no next element")
}
