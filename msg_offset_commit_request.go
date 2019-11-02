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

// OffsetCommitRequest; ApiKey: 8, Versions: 0-8
type OffsetCommitRequest struct {
	GroupId         string                      // The unique group identifier. Versions: 0+
	GenerationId    int32                       // The generation of the group. Versions: 1+
	MemberId        string                      // The member ID assigned by the group coordinator. Versions: 1+
	GroupInstanceId string                      // The unique identifier of the consumer instance provided by end user. Versions: 7+
	RetentionTimeMs int64                       // The time period in ms to retain the offset. Versions: 2-4
	Topics          []OffsetCommitRequestTopic8 // The topics to commit offsets for. Versions: 0+
}

// size of OffsetCommitRequest; Versions: 0-8
func (t OffsetCommitRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.GroupId) // GroupId
	if version >= 1 {
		sz += sizeof.Int32 // GenerationId
	}
	if version >= 1 {
		sz += sizeof.String(t.MemberId) // MemberId
	}
	if version >= 7 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	if version >= 2 && version <= 4 {
		sz += sizeof.Int64 // RetentionTimeMs
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode OffsetCommitRequest; Versions: 0-8
func (t OffsetCommitRequest) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.GroupId) // GroupId
	if version >= 1 {
		e.PutInt32(t.GenerationId) // GenerationId
	}
	if version >= 1 {
		e.PutString(t.MemberId) // MemberId
	}
	if version >= 7 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
	if version >= 2 && version <= 4 {
		e.PutInt64(t.RetentionTimeMs) // RetentionTimeMs
	}
	// Topics
	len5 := len(t.Topics)
	e.PutArrayLength(len5)
	for i := 0; i < len5; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode OffsetCommitRequest; Versions: 0-8
func (t *OffsetCommitRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.GenerationId, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 1 {
		t.MemberId, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 7 {
		t.GroupInstanceId, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 2 && version <= 4 {
		t.RetentionTimeMs, err = d.Int64()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]OffsetCommitRequestTopic8, n)
		for i := 0; i < n; i++ {
			var item OffsetCommitRequestTopic8
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type OffsetCommitRequestTopic8 struct {
	Name       string                          // The topic name. Versions: 0+
	Partitions []OffsetCommitRequestPartition8 // Each partition to commit offsets for. Versions: 0+
}

// size of OffsetCommitRequestTopic8; Versions: 0-8
func (t OffsetCommitRequestTopic8) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode OffsetCommitRequestTopic8; Versions: 0-8
func (t OffsetCommitRequestTopic8) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].encode(e, version)
	}
}

// decode OffsetCommitRequestTopic8; Versions: 0-8
func (t *OffsetCommitRequestTopic8) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]OffsetCommitRequestPartition8, n)
		for i := 0; i < n; i++ {
			var item OffsetCommitRequestPartition8
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type OffsetCommitRequestPartition8 struct {
	PartitionIndex       int32  // The partition index. Versions: 0+
	CommittedOffset      int64  // The message offset to be committed. Versions: 0+
	CommittedLeaderEpoch int32  // The leader epoch of this partition. Versions: 6+
	CommittedMetadata    string // Any associated metadata the client wants to keep. Versions: 0+
}

// size of OffsetCommitRequestPartition8; Versions: 0-8
func (t OffsetCommitRequestPartition8) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int64 // CommittedOffset
	if version >= 6 {
		sz += sizeof.Int32 // CommittedLeaderEpoch
	}
	sz += sizeof.String(t.CommittedMetadata) // CommittedMetadata
	return sz
}

// encode OffsetCommitRequestPartition8; Versions: 0-8
func (t OffsetCommitRequestPartition8) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex)  // PartitionIndex
	e.PutInt64(t.CommittedOffset) // CommittedOffset
	if version >= 6 {
		e.PutInt32(t.CommittedLeaderEpoch) // CommittedLeaderEpoch
	}
	e.PutString(t.CommittedMetadata) // CommittedMetadata
}

// decode OffsetCommitRequestPartition8; Versions: 0-8
func (t *OffsetCommitRequestPartition8) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.CommittedOffset, err = d.Int64()
	if err != nil {
		return err
	}
	if version >= 6 {
		t.CommittedLeaderEpoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.CommittedMetadata, err = d.String()
	if err != nil {
		return err
	}
	return err
}
