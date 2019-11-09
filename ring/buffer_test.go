// Code generated by kafka-protocol-gen. DO NOT EDIT.
//
// Copyright 2019 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ring

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBuffer_WriteN(t *testing.T) {
	var (
		ring    = New(15)
		content = []byte("hello world")
		length  = len(content)
	)

	ring.WriteN(content, len(content))
	if got, want := string(ring.data[0:length]), string(content); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func BenchmarkAtomic(t *testing.B) {
	var n uint32
	for i := 0; i < t.N; i++ {
		atomic.AddUint32(&n, 1)
	}
	if n == 0 {
		t.Fatalf("got %v; want > 0", n)
	}
}

var digits = []byte("1234567890")

type iterator struct {
	offset int
}

func (iter *iterator) ReadN(data []byte, n int) {
	for i := 0; i < n; i++ {
		data[i] = digits[iter.offset]

		iter.offset++
		if iter.offset == len(digits) {
			iter.offset = 0
		}
	}
}

func TestMutex(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	mutex := newMutex(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			mutex.Wait()
			mutex.Wait()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		mutex.Release()
	}
}

func TestStress(t *testing.T) {
	const size = 1e2

	buffer := New(size)
	go func() {
		data := make([]byte, 10)
		iter := &iterator{}
		max := int(size - 2)
		for i := 0; i < max; i++ {
			for n := 1; n <= i; n++ {
				iter.ReadN(data, n)
				buffer.WriteN(data, n)
			}
			for n := i; n >= 1; n-- {
				iter.ReadN(data, n)
				buffer.WriteN(data, n)
			}
		}
	}()

	data := make([]byte, 10)
	remain := int(size)
	for remain > 0 {
		buffer.ReadN(data, cap(data))
		if got, want := string(data[0:cap(data)]), string(digits); got != want {
			t.Fatalf("got %v; want %v", got, want)
		}
		remain -= cap(data)
	}
}

func Test_newRingBuffer(t *testing.T) {
	var (
		n       = int(1e4)
		content = make([]byte, n)
		buffer  = New(1e6)
		counter int
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//ok
			}
			buffer.WriteN(content, n)
		}
	}()
	go func() {
		defer wg.Done()
		data := make([]byte, n*2)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//ok
			}
			buffer.ReadN(data, n)
			counter++
		}
	}()
	wg.Wait()

	fmt.Println(counter)
}

func BenchmarkRingBuffer(t *testing.B) {
	var (
		n      = int(1e3)
		buffer = make([]byte, n)
		rb     = New(1e6)
	)

	for i := 0; i < t.N; i++ {
		rb.WriteN(buffer, n)
		rb.ReadN(buffer, n)
	}
}
