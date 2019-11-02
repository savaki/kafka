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

// ControlledShutdownResponse; ApiKey: 7, Versions: 0-3
type ControlledShutdownResponse struct {
	ErrorCode           int16                 // The top-level error code. Versions: 0+
	RemainingPartitions []RemainingPartition7 // The partitions that the broker still leads. Versions: 0+
}

// size of ControlledShutdownResponse; Versions: 0-3
func (t ControlledShutdownResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16       // ErrorCode
	sz += sizeof.ArrayLength // RemainingPartitions
	for i := len(t.RemainingPartitions) - 1; i >= 0; i-- {
		sz += t.RemainingPartitions[i].size(version)
	}
	return sz
}

// encode ControlledShutdownResponse; Versions: 0-3
func (t ControlledShutdownResponse) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode) // ErrorCode
	// RemainingPartitions
	len1 := len(t.RemainingPartitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.RemainingPartitions[i].encode(e, version)
	}
}

// decode ControlledShutdownResponse; Versions: 0-3
func (t *ControlledShutdownResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	// RemainingPartitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.RemainingPartitions = make([]RemainingPartition7, n)
		for i := 0; i < n; i++ {
			var item RemainingPartition7
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.RemainingPartitions[i] = item
		}
	}
	return err
}

type RemainingPartition7 struct {
	TopicName      string // The name of the topic. Versions: 0+
	PartitionIndex int32  // The index of the partition. Versions: 0+
}

// size of RemainingPartition7; Versions: 0-3
func (t RemainingPartition7) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.TopicName) // TopicName
	sz += sizeof.Int32               // PartitionIndex
	return sz
}

// encode RemainingPartition7; Versions: 0-3
func (t RemainingPartition7) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.TopicName)     // TopicName
	e.PutInt32(t.PartitionIndex) // PartitionIndex
}

// decode RemainingPartition7; Versions: 0-3
func (t *RemainingPartition7) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.TopicName, err = d.String()
	if err != nil {
		return err
	}
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}
