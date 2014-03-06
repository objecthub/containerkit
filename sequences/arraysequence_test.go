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

import "testing"


func checkSize(t *testing.T, q Sequence, size int, name string) {
  qsize := q.Size()
  if qsize != size {
    t.Errorf("Expected size of sequence %s to be %d; was %d", name, size, qsize)
  }
}

func TestArraySequenceClass(t *testing.T) {
  s1 := ArraySequence.New()
  checkSize(t, s1, 0, "s1")
  s2 := ArraySequence.New(1)
  checkSize(t, s2, 1, "s2")
  if s2.First() != 1 {
    t.Errorf("Expected first element of s2 to be 1")
  }
  s3 := ArraySequence.New(1, 2)
  checkSize(t, s3, 2, "s3")
  if s3.First() != 1 {
    t.Errorf("Expected first element of s3 to be 1")
  }
  s4 := ArraySequence.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
  checkSize(t, s4, 17, "s4")
  if s4.First() != 1 {
    t.Errorf("Expected first element of s4 to be 1")
  }
}
