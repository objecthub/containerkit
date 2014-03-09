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


func TestArrayQueueClass(t *testing.T) {
  q1 := ArrayQueue.New()
  checkQueueSize(t, q1, 0, "q1")
  q2 := ArrayQueue.New(1)
  checkQueueSize(t, q2, 1, "q2")
  if q2.Peek() != 1 {
    t.Errorf("Expected first element of q2 to be 1")
  }
  q3 := ArrayQueue.New(2, 2)
  checkQueueSize(t, q3, 2, "q3")
  if q3.Peek() != 2 {
    t.Errorf("Expected first element of q3 to be 2")
  }
  q4 := ArrayQueue.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
  checkQueueSize(t, q4, 17, "q4")
  if q4.Peek() != 1 {
    t.Errorf("Expected first element of q4 to be 1")
  }
}
