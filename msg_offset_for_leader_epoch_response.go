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

// OffsetForLeaderEpochResponse; ApiKey: 23, Versions: 0-3
type OffsetForLeaderEpochResponse struct {
	ThrottleTimeMs int32                          // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 2+
	Topics         []OffsetForLeaderTopicResult23 // Each topic we fetched offsets for. Versions: 0+
}

// size of OffsetForLeaderEpochResponse; Versions: 0-3
func (t OffsetForLeaderEpochResponse) size(version int16) int32 {
	var sz int32
	if version >= 2 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode OffsetForLeaderEpochResponse; Versions: 0-3
func (t OffsetForLeaderEpochResponse) encode(e *protocol.Encoder, version int16) {
	if version >= 2 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode OffsetForLeaderEpochResponse; Versions: 0-3
func (t *OffsetForLeaderEpochResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 2 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]OffsetForLeaderTopicResult23, n)
		for i := 0; i < n; i++ {
			var item OffsetForLeaderTopicResult23
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type OffsetForLeaderTopicResult23 struct {
	Name       string                             // The topic name. Versions: 0+
	Partitions []OffsetForLeaderPartitionResult23 // Each partition in the topic we fetched offsets for. Versions: 0+
}

// size of OffsetForLeaderTopicResult23; Versions: 0-3
func (t OffsetForLeaderTopicResult23) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode OffsetForLeaderTopicResult23; Versions: 0-3
func (t OffsetForLeaderTopicResult23) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].encode(e, version)
	}
}

// decode OffsetForLeaderTopicResult23; Versions: 0-3
func (t *OffsetForLeaderTopicResult23) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]OffsetForLeaderPartitionResult23, n)
		for i := 0; i < n; i++ {
			var item OffsetForLeaderPartitionResult23
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type OffsetForLeaderPartitionResult23 struct {
	ErrorCode      int16 // The error code 0, or if there was no error. Versions: 0+
	PartitionIndex int32 // The partition index. Versions: 0+
	LeaderEpoch    int32 // The leader epoch of the partition. Versions: 1+
	EndOffset      int64 // The end offset of the epoch. Versions: 0+
}

// size of OffsetForLeaderPartitionResult23; Versions: 0-3
func (t OffsetForLeaderPartitionResult23) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16 // ErrorCode
	sz += sizeof.Int32 // PartitionIndex
	if version >= 1 {
		sz += sizeof.Int32 // LeaderEpoch
	}
	sz += sizeof.Int64 // EndOffset
	return sz
}

// encode OffsetForLeaderPartitionResult23; Versions: 0-3
func (t OffsetForLeaderPartitionResult23) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)      // ErrorCode
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	if version >= 1 {
		e.PutInt32(t.LeaderEpoch) // LeaderEpoch
	}
	e.PutInt64(t.EndOffset) // EndOffset
}

// decode OffsetForLeaderPartitionResult23; Versions: 0-3
func (t *OffsetForLeaderPartitionResult23) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.LeaderEpoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.EndOffset, err = d.Int64()
	if err != nil {
		return err
	}
	return err
}
