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

// +build integration

package kafka

import (
	"context"
	"testing"

	"github.com/savaki/kafka/message"
)

func TestCreateTopic(t *testing.T) {
	ctx := context.Background()
	client := NewClient([]string{"127.0.0.1:9092"})
	broker, err := client.DialController(ctx)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
	defer broker.Close()

	topicName := "ack"
	resp, err := broker.CreateTopics(message.CreateTopicsRequest{
		Topics: []message.CreatableTopic19{
			{
				Name:              topicName,
				NumPartitions:     1,
				ReplicationFactor: 1,
				Assignments:       []message.CreatableReplicaAssignment19{},
				Configs:           []message.CreateableTopicConfig19{},
			},
		},
	})
	if got, want := len(resp.Topics), 1; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
	for _, topic := range resp.Topics {
		if code := topic.ErrorCode; code != 7 && code != 36 {
			t.Fatalf("got %v; want either 7 or 36", code)
		}
	}

	client.DialLeader()

	leader, err := client.DialLeader(ctx, topicName, 0)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}
	defer leader.Close()
}
