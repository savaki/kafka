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

package kafka

import (
	"github.com/savaki/kafka/protocol"
	"github.com/savaki/kafka/sizeof"
)

// AlterPartitionReassignmentsRequest; ApiKey: 45, Versions: 0
type AlterPartitionReassignmentsRequest struct {
	TimeoutMs int32                 // The time in ms to wait for the request to complete. Versions: 0+
	Topics    []ReassignableTopic45 // The topics to reassign. Versions: 0+
}

// size of AlterPartitionReassignmentsRequest; Versions: 0
func (t AlterPartitionReassignmentsRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32       // TimeoutMs
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode AlterPartitionReassignmentsRequest; Versions: 0
func (t AlterPartitionReassignmentsRequest) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.TimeoutMs) // TimeoutMs
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].Encode(e, version)
	}
}

// decode AlterPartitionReassignmentsRequest; Versions: 0
func (t *AlterPartitionReassignmentsRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.TimeoutMs, err = d.Int32()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]ReassignableTopic45, n)
		for i := 0; i < n; i++ {
			var item ReassignableTopic45
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type ReassignableTopic45 struct {
	Name       string                    // The topic name. Versions: 0+
	Partitions []ReassignablePartition45 // The partitions to reassign. Versions: 0+
}

// size of ReassignableTopic45; Versions: 0
func (t ReassignableTopic45) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode ReassignableTopic45; Versions: 0
func (t ReassignableTopic45) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].Encode(e, version)
	}
}

// decode ReassignableTopic45; Versions: 0
func (t *ReassignableTopic45) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]ReassignablePartition45, n)
		for i := 0; i < n; i++ {
			var item ReassignablePartition45
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type ReassignablePartition45 struct {
	PartitionIndex int32   // The partition index. Versions: 0+
	Replicas       []int32 // The replicas to place the partitions on, or null to cancel a pending reassignment for this partition. Versions: 0+
}

// size of ReassignablePartition45; Versions: 0
func (t ReassignablePartition45) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32                  // PartitionIndex
	sz += sizeof.Int32Array(t.Replicas) // Replicas
	return sz
}

// encode ReassignablePartition45; Versions: 0
func (t ReassignablePartition45) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	e.PutInt32Array(t.Replicas)  // Replicas
}

// decode ReassignablePartition45; Versions: 0
func (t *ReassignablePartition45) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.Replicas, err = d.Int32Array()
	if err != nil {
		return err
	}
	return err
}
