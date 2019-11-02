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

// DeleteRecordsResponse; ApiKey: 21, Versions: 0-1
type DeleteRecordsResponse struct {
	ThrottleTimeMs int32                        // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	Topics         []DeleteRecordsTopicResult21 // Each topic that we wanted to delete records from. Versions: 0+
}

// size of DeleteRecordsResponse; Versions: 0-1
func (t DeleteRecordsResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32       // ThrottleTimeMs
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode DeleteRecordsResponse; Versions: 0-1
func (t DeleteRecordsResponse) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode DeleteRecordsResponse; Versions: 0-1
func (t *DeleteRecordsResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]DeleteRecordsTopicResult21, n)
		for i := 0; i < n; i++ {
			var item DeleteRecordsTopicResult21
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type DeleteRecordsTopicResult21 struct {
	Name       string                           // The topic name. Versions: 0+
	Partitions []DeleteRecordsPartitionResult21 // Each partition that we wanted to delete records from. Versions: 0+
}

// size of DeleteRecordsTopicResult21; Versions: 0-1
func (t DeleteRecordsTopicResult21) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode DeleteRecordsTopicResult21; Versions: 0-1
func (t DeleteRecordsTopicResult21) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].encode(e, version)
	}
}

// decode DeleteRecordsTopicResult21; Versions: 0-1
func (t *DeleteRecordsTopicResult21) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]DeleteRecordsPartitionResult21, n)
		for i := 0; i < n; i++ {
			var item DeleteRecordsPartitionResult21
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type DeleteRecordsPartitionResult21 struct {
	PartitionIndex int32 // The partition index. Versions: 0+
	LowWatermark   int64 // The partition low water mark. Versions: 0+
	ErrorCode      int16 // The deletion error code, or 0 if the deletion succeeded. Versions: 0+
}

// size of DeleteRecordsPartitionResult21; Versions: 0-1
func (t DeleteRecordsPartitionResult21) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int64 // LowWatermark
	sz += sizeof.Int16 // ErrorCode
	return sz
}

// encode DeleteRecordsPartitionResult21; Versions: 0-1
func (t DeleteRecordsPartitionResult21) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	e.PutInt64(t.LowWatermark)   // LowWatermark
	e.PutInt16(t.ErrorCode)      // ErrorCode
}

// decode DeleteRecordsPartitionResult21; Versions: 0-1
func (t *DeleteRecordsPartitionResult21) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.LowWatermark, err = d.Int64()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	return err
}
