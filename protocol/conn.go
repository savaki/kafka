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

package protocol

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

type config struct {
	dialFunc func(network, addr string) (net.Conn, error)
}

type Option func(*config)

func WithDialFunc(dialFunc func(network, addr string) (net.Conn, error)) Option {
	return func(c *config) {
		c.dialFunc = dialFunc
		if c.dialFunc == nil {
			c.dialFunc = net.Dial
		}
	}
}

func buildConfig(opts []Option) config {
	c := config{
		dialFunc: net.Dial,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

type Conn struct {
	cancel      context.CancelFunc
	ch          chan request
	doneMessage chan struct{}
	doneRead    chan struct{}
	encoder     *Encoder
	err         error
	id          int32
	raw         net.Conn
	rb          *RingBuffer

	writeLock sync.Mutex
	readLock  sync.Mutex
	requests  map[int32]request
}

type request struct {
	decode func(*Decoder) error
	reply  chan error
}

func Connect(addr string, opts ...Option) (*Conn, error) {
	config := buildConfig(opts)
	raw, err := config.dialFunc("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial broker, %v: %w", addr, err)
	}

	var (
		ctx, cancel = context.WithCancel(context.Background())
		rb          = NewRingBuffer(5e6)
		ch          = make(chan request, 32)
		e           = NewEncoder(bufio.NewWriter(raw))
	)

	c := &Conn{
		cancel:      cancel,
		ch:          ch,
		doneMessage: make(chan struct{}),
		doneRead:    make(chan struct{}),
		encoder:     e,
		raw:         raw,
		rb:          rb,
		requests:    map[int32]request{},
	}
	go c.readLoop(ctx, rb, raw)
	go c.messageLoop(ctx, rb, ch)

	return c, nil
}

func (c *Conn) readLoop(ctx context.Context, rb *RingBuffer, raw net.Conn) {
	defer close(c.doneRead)

	buffer := make([]byte, 2e6) // 2MB
	for {
		n, err := raw.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		rb.WriteN(buffer, n)
	}
}

func (c *Conn) messageLoop(ctx context.Context, rb *RingBuffer, ch <-chan request) {
	defer close(c.doneMessage)

	var (
		buffer = make([]byte, 2e6) // 2MB
		d      = NewDecoder(buffer, len(buffer))
	)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// ok
		}

		// read the header
		rb.ReadN(buffer, 4)
		d.Reset(4)
		v, _ := d.Int32()
		size := int(v)

		// read the payload
		rb.ReadN(buffer, size)
		d.Reset(size)

		var resp ResponseHeader
		if err := (&resp).decode(d, 1); err != nil {
			fmt.Println(err)
			continue
		}

		c.readLock.Lock()
		req, ok := c.requests[resp.CorrelationId]
		delete(c.requests, resp.CorrelationId)
		c.readLock.Unlock()

		if ok {
			req.reply <- req.decode(d)
		}
	}
}

type EncodeFunc func(e *Encoder, correlationID int32)
type DecodeFunc func(d *Decoder) error

func (c *Conn) Write(encode EncodeFunc, decode DecodeFunc) error {
	correlationID := atomic.AddInt32(&c.id, 1)
	req := request{
		decode: decode,
		reply:  make(chan error, 1),
	}

	c.readLock.Lock()
	c.requests[correlationID] = req
	c.readLock.Unlock()

	c.writeLock.Lock()
	encode(c.encoder, correlationID)
	err := c.encoder.Flush()
	c.writeLock.Unlock()

	if err != nil {
		return err
	}

	return <-req.reply
}

func (c *Conn) Close() error {
	c.cancel()
	c.rb.Close()
	c.raw.Close()
	<-c.doneMessage
	<-c.doneRead
	return c.err
}

// ResponseHeader; ApiKey: 0, Versions: 0-1
type ResponseHeader struct {
	CorrelationId int32 // The correlation ID of this response. Versions: 0+
}

// size of ResponseHeader; Versions: 0-1
func (t ResponseHeader) size(version int16) int32 {
	var sz int32
	sz += 4 // CorrelationId
	return sz
}

// encode ResponseHeader; Versions: 0-1
func (t ResponseHeader) encode(e *Encoder, version int16) {
	e.PutInt32(t.CorrelationId) // CorrelationId
}

// decode ResponseHeader; Versions: 0-1
func (t *ResponseHeader) decode(d *Decoder, version int16) error {
	var err error
	t.CorrelationId, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}
