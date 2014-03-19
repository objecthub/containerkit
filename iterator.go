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


// ============================================================================
// INTERFACE
// ============================================================================

type Iterator interface {
  HasNext() bool
  Next() interface{}
}


// ============================================================================
// IMPLEMENTATION
// ============================================================================

func CountElements(iter Iterator) int {
  res := 0
  for iter.HasNext() {
    res++
    iter.Next()
  }
  return res
}


// Empty iterator

var EmptyIterator Iterator = new(emptyIterator)

type emptyIterator struct {}

func (this emptyIterator) HasNext() bool {
  return false
}

func (this emptyIterator) Next() interface{} {
  panic("EmptyIterator: no next element")
}


// Slice iterators

func NewSlicedIterator(iter Iterator,
                       drop int,
                       dropWhile Predicate,
                       takeWhile Predicate,
                       take int) Iterator {
  // drop initial elements
  for i := 0; i < drop; i++ {
    if iter.HasNext() {
      iter.Next()
    }
  }
  // drop elements as long as they satisfy predicate takeWhile
  hasLookahead := false
  var lookahead interface{} = nil
  for iter.HasNext() && !hasLookahead {
    next := iter.Next()
    if !dropWhile(next) {
      lookahead = next
      hasLookahead = true
    }
  }
  return &slicedIterator{iter, hasLookahead, lookahead, takeWhile, true, take}
}

type slicedIterator struct {
  iter Iterator
  hasLookahead bool
  lookahead interface{}
  takeWhile Predicate
  evalTakeWhile bool
  take int
}

func (this *slicedIterator) HasNext() bool {
  return this.hasLookahead
}

func (this *slicedIterator) Next() interface{} {
  if this.hasLookahead {
    res := this.lookahead
    if this.iter.HasNext() {
      this.lookahead = this.iter.Next()
      if this.evalTakeWhile && !this.takeWhile(this.lookahead) {
        this.evalTakeWhile = false
      }
      if !this.evalTakeWhile {
        if this.take == 0 {
          this.lookahead = nil
          this.hasLookahead = false
        } else {
          this.take--
        }
      }
    } else {
      this.lookahead = nil
      this.hasLookahead = false
    }
    return res
  }
  panic("slicedIterator: no next element")
}


// Filter iterators

func NewFilterIterator(pred Predicate, iter Iterator) Iterator {
  res := &filteredIterator{pred, false, nil, iter}
  res.scan()
  return res
}

type filteredIterator struct {
  pred Predicate
  hasNext bool
  next interface{}
  iter Iterator
}

func (this *filteredIterator) scan() {
  for this.iter.HasNext() {
    this.next = this.iter.Next()
    if this.pred(this.next) {
      this.hasNext = true
      return
    }
  }
  this.next = nil
  this.hasNext = false
}

func (this *filteredIterator) HasNext() bool {
  return this.hasNext
}

func (this *filteredIterator) Next() interface{} {
  if this.hasNext {
    res := this.next
    this.scan()
    return res
  }
  panic("filteredIterator.Next: no next element")
}


// Map iterators

func NewMappedIterator(f Mapping, iter Iterator) Iterator {
  return &mappedIterator{f, iter}
}

type mappedIterator struct {
  f Mapping
  iter Iterator
}

func (this *mappedIterator) HasNext() bool {
  return this.iter.HasNext()
}

func (this *mappedIterator) Next() interface{} {
  if this.iter.HasNext() {
    return this.f(this.iter.Next())
  }
  panic("mappedIterator.Next: no next element")
}


// Flat-map iterators

type flatMappedIterator struct {
  g Generator
  iter Iterator
  current Iterator
}

func (this *flatMappedIterator) scan() {
  for this.iter.HasNext() {
    if next, valid := this.g(this.iter.Next()).(Iterator); valid {
      this.current = next
      if this.current.HasNext() {
        return
      }
    } else {
      panic("flatMappedIterator.scan: generator did not return Iterator")
    }
  }
  this.current = EmptyIterator
}

func (this *flatMappedIterator) HasNext() bool {
  return this.current.HasNext()
}

func (this *flatMappedIterator) Next() interface{} {
  if this.current.HasNext() {
    res := this.current.Next()
    if !this.current.HasNext() {
      this.scan()
    }
    return res
  }
  panic("flatMappedIterator.Next: no next element")
}


// Composite iterators

func NewCompositeIterator(first Iterator, second Iterator) Iterator {
  return &compositeIterator{first, second}
}

type compositeIterator struct {
  first Iterator
  second Iterator
}

func (this *compositeIterator) HasNext() bool {
  return this.first.HasNext() || this.second.HasNext()
}

func (this *compositeIterator) Next() interface{} {
  if this.first.HasNext() {
    return this.first.Next()
  } else if this.second.HasNext() {
    return this.second.Next()
  }
  panic("compositeIterator.Next: no next element")
}


// Combination iterator

type combinedIterator struct {
  first Iterator
  second Iterator
  f Binop
}

func (this *combinedIterator) HasNext() bool {
  return this.first.HasNext() && this.second.HasNext()
}

func (this *combinedIterator) Next() interface{} {
  if this.HasNext() {
    return this.f(this.first.Next(), this.second.Next())
  }
  panic("combinedIterator.Next: no next element")
}
