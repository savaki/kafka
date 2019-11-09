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

package kafka

import (
	"context"
	"fmt"
	"strconv"

	"github.com/savaki/kafka/message"
)

// Client provides a high level Kafka client
type Client struct {
	bootstrap []string
	opts      []Option
}

// NewClient returns a new Kafka Client
func NewClient(bootstrap []string, opts ...Option) *Client {
	return &Client{
		bootstrap: bootstrap,
		opts:      opts,
	}
}

func (c *Client) connectAny(callback func(*Broker) error) error {
	var broker *Broker
	var err error
	for _, addr := range c.bootstrap {
		broker, err = NewBroker(addr, c.opts...)
		if err != nil {
			continue
		}

		return callback(broker)
	}
	return err
}

func (c *Client) dialCoordinator(ctx context.Context, req message.FindCoordinatorRequest) (*Broker, error) {
	var resp message.FindCoordinatorResponse
	callback := func(b *Broker) error {
		v, err := b.FindCoordinator(req)
		if err != nil {
			return err
		}
		resp = v
		return nil
	}
	if err := c.connectAny(callback); err != nil {
		return nil, err
	}

	addr := resp.Host + ":" + strconv.Itoa(int(resp.Port))
	return NewBroker(addr, c.opts...)
}

// DialController connects to the cluster controller.
func (c *Client) DialController(ctx context.Context) (*Broker, error) {
	var meta message.MetadataResponse
	callback := func(b *Broker) error {
		v, err := b.Metadata(message.MetadataRequest{})
		if err != nil {
			return err
		}
		meta = v
		return nil
	}
	if err := c.connectAny(callback); err != nil {
		return nil, err
	}

	for _, broker := range meta.Brokers {
		if broker.NodeId == meta.ControllerId {
			addr := broker.Host + ":" + strconv.Itoa(int(broker.Port))
			return NewBroker(addr, c.opts...)
		}
	}

	return nil, fmt.Errorf("no controller returned by bootstrap brokers")
}

// DialGroupCoordinator returns the group coordinator for the provided consumer group id
func (c *Client) DialGroupCoordinator(ctx context.Context, groupId string) (*Broker, error) {
	return c.dialCoordinator(ctx, message.FindCoordinatorRequest{
		Key:     groupId,
		KeyType: 0,
	})
}
