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

import "testing"


func TestArrayStackClass(t *testing.T) {
  s1 := ArrayStack.New()
  checkSize(t, s1, 0, "s1")
  s2 := ArrayStack.New(1)
  checkSize(t, s2, 1, "q2")
  if s2.Peek() != 1 {
    t.Errorf("Expected first element of s2 to be 1")
  }
  s3 := ArrayStack.New(1, 2)
  checkSize(t, s3, 2, "q3")
  if s3.Peek() != 2 {
    t.Errorf("Expected first element of s3 to be 2")
  }
  s4 := ArrayStack.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
  checkSize(t, s4, 17, "q4")
  if s4.Peek() != 17 {
    t.Errorf("Expected first element of s4 to be 17")
  }
  s5 := ArrayStack.New()
  if !s5.IsEmpty() {
    t.Errorf("Expected s5 to be empty")
  }
  s5.Push(1)
  if s5.IsEmpty() {
    t.Errorf("Expected s5 to be not empty")
  }
  s5.Push(2)
  if s5.Peek() != 2 {
    t.Errorf("Expected first element of s5 to be 17")
  }
  s5.Push(3)
  s5.Push(4)
  if s5.Pop() != 4 {
    t.Errorf("Expected first element of s5 to be 4")
  }
  s5.Push(4)
  s5.Push(5)
  s5.Push(6)
  if s5.Pop() != 6 {
    t.Errorf("Expected first element of s5 to be 6")
  }
  s5.Pop()
  s5.Pop()
  s5.Pop()
  if s5.Pop() != 2 {
    t.Errorf("Expected first element of s5 to be 2")
  }
  if s5.IsEmpty() {
    t.Errorf("Expected s5 to be not empty")
  }
  if s5.Pop() != 1 {
    t.Errorf("Expected first element of s5 to be 1")
  }
  if !s5.IsEmpty() {
    t.Errorf("Expected s5 to be empty")
  }
}
