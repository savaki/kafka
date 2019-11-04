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

// OffsetFetchResponse; ApiKey: 9, Versions: 0-6
type OffsetFetchResponse struct {
	ThrottleTimeMs int32                       // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 3+
	Topics         []OffsetFetchResponseTopic9 // The responses per topic. Versions: 0+
	ErrorCode      int16                       // The top-level error code, or 0 if there was no error. Versions: 2+
}

// size of OffsetFetchResponse; Versions: 0-6
func (t OffsetFetchResponse) Size(version int16) int32 {
	var sz int32
	if version >= 3 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].Size(version)
	}
	if version >= 2 {
		sz += sizeof.Int16 // ErrorCode
	}
	return sz
}

// encode OffsetFetchResponse; Versions: 0-6
func (t OffsetFetchResponse) Encode(e *protocol.Encoder, version int16) {
	if version >= 3 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].Encode(e, version)
	}
	if version >= 2 {
		e.PutInt16(t.ErrorCode) // ErrorCode
	}
}

// decode OffsetFetchResponse; Versions: 0-6
func (t *OffsetFetchResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 3 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]OffsetFetchResponseTopic9, n)
		for i := 0; i < n; i++ {
			var item OffsetFetchResponseTopic9
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	if version >= 2 {
		t.ErrorCode, err = d.Int16()
		if err != nil {
			return err
		}
	}
	return err
}

type OffsetFetchResponseTopic9 struct {
	Name       string                          // The topic name. Versions: 0+
	Partitions []OffsetFetchResponsePartition9 // The responses per partition Versions: 0+
}

// size of OffsetFetchResponseTopic9; Versions: 0-6
func (t OffsetFetchResponseTopic9) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].Size(version)
	}
	return sz
}

// encode OffsetFetchResponseTopic9; Versions: 0-6
func (t OffsetFetchResponseTopic9) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].Encode(e, version)
	}
}

// decode OffsetFetchResponseTopic9; Versions: 0-6
func (t *OffsetFetchResponseTopic9) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]OffsetFetchResponsePartition9, n)
		for i := 0; i < n; i++ {
			var item OffsetFetchResponsePartition9
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type OffsetFetchResponsePartition9 struct {
	PartitionIndex       int32  // The partition index. Versions: 0+
	CommittedOffset      int64  // The committed message offset. Versions: 0+
	CommittedLeaderEpoch int32  // The leader epoch. Versions: 5+
	Metadata             string // The partition metadata. Versions: 0+
	ErrorCode            int16  // The error code, or 0 if there was no error. Versions: 0+
}

// size of OffsetFetchResponsePartition9; Versions: 0-6
func (t OffsetFetchResponsePartition9) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int64 // CommittedOffset
	if version >= 5 {
		sz += sizeof.Int32 // CommittedLeaderEpoch
	}
	sz += sizeof.String(t.Metadata) // Metadata
	sz += sizeof.Int16              // ErrorCode
	return sz
}

// encode OffsetFetchResponsePartition9; Versions: 0-6
func (t OffsetFetchResponsePartition9) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex)  // PartitionIndex
	e.PutInt64(t.CommittedOffset) // CommittedOffset
	if version >= 5 {
		e.PutInt32(t.CommittedLeaderEpoch) // CommittedLeaderEpoch
	}
	e.PutString(t.Metadata) // Metadata
	e.PutInt16(t.ErrorCode) // ErrorCode
}

// decode OffsetFetchResponsePartition9; Versions: 0-6
func (t *OffsetFetchResponsePartition9) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.CommittedOffset, err = d.Int64()
	if err != nil {
		return err
	}
	if version >= 5 {
		t.CommittedLeaderEpoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.Metadata, err = d.String()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	return err
}
