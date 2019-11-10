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

// dialBootstrap dials one of the initial bootstrap servers
func (c *Client) dialBootstrap(callback func(*Broker) error) error {
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

// dial a specific node within the kafka cluster
func (c *Client) dialNode(meta message.MetadataResponse, nodeID int32) (*Broker, error) {
	for _, broker := range meta.Brokers {
		if broker.NodeId == nodeID {
			addr := broker.Host + ":" + strconv.Itoa(int(broker.Port))
			return NewBroker(addr, c.opts...)
		}
	}
	return nil, fmt.Errorf("no such broker node, %v", nodeID)
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
	if err := c.dialBootstrap(callback); err != nil {
		return nil, err
	}

	addr := resp.Host + ":" + strconv.Itoa(int(resp.Port))
	return NewBroker(addr, c.opts...)
}

// DialController connects to the cluster controller.
func (c *Client) DialController(ctx context.Context) (*Broker, error) {
	meta, err := c.Metadata(ctx)
	if err != nil {
		return nil, err
	}

	return c.dialNode(meta, meta.ControllerId)
}

// DialGroupCoordinator returns the group coordinator for the provided consumer group id
func (c *Client) DialGroupCoordinator(ctx context.Context, groupId string) (*Broker, error) {
	return c.dialCoordinator(ctx, message.FindCoordinatorRequest{
		Key:     groupId,
		KeyType: 0,
	})
}

// DialLeader dials the leads for the specified partition
func (c *Client) DialLeader(ctx context.Context, topic string, partition int32) (*Broker, error) {
	meta, err := c.Metadata(ctx, topic)
	if err != nil {
		return nil, err
	}

	t, err := findTopic(meta, topic)
	if err != nil {
		return nil, fmt.Errorf("unable to dial leader: %w", err)
	}

	p, err := findPartition(t, partition)
	if err != nil {
		return nil, fmt.Errorf("unable to dial leader: %w", err)
	}

	return c.dialNode(meta, p.LeaderId)
}

// Metadata returns cluster metadata with optional topic info
func (c *Client) Metadata(ctx context.Context, topics ...string) (message.MetadataResponse, error) {
	var meta message.MetadataResponse
	callback := func(b *Broker) error {
		req := message.MetadataRequest{}
		for _, topic := range topics {
			req.Topics = append(req.Topics, message.MetadataRequestTopic3{Name: topic})
		}
		v, err := b.Metadata(req)
		if err != nil {
			return err
		}
		meta = v
		return nil
	}
	if err := c.dialBootstrap(callback); err != nil {
		return message.MetadataResponse{}, err
	}
	return meta, nil
}

func findPartition(topic message.MetadataResponseTopic3, partition int32) (message.MetadataResponsePartition3, error) {
	for _, p := range topic.Partitions {
		if p.PartitionIndex == partition {
			return p, nil
		}
	}
	return message.MetadataResponsePartition3{}, fmt.Errorf("unable to find partition, %v, for topic, %v", partition, topic.Name)
}

func findTopic(meta message.MetadataResponse, topicName string) (message.MetadataResponseTopic3, error) {
	for _, topic := range meta.Topics {
		if topic.Name == topicName {
			return topic, nil
		}
	}
	return message.MetadataResponseTopic3{}, fmt.Errorf("no such topic, %v", topicName)
}
