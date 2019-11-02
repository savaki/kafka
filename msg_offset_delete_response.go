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

// OffsetDeleteResponse; ApiKey: 47, Versions: 0
type OffsetDeleteResponse struct {
	ErrorCode      int16                         // The top-level error code, or 0 if there was no error. Versions: 0+
	ThrottleTimeMs int32                         // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	Topics         []OffsetDeleteResponseTopic47 // The responses for each topic. Versions: 0+
}

// size of OffsetDeleteResponse; Versions: 0
func (t OffsetDeleteResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16       // ErrorCode
	sz += sizeof.Int32       // ThrottleTimeMs
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode OffsetDeleteResponse; Versions: 0
func (t OffsetDeleteResponse) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)      // ErrorCode
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	// Topics
	len2 := len(t.Topics)
	e.PutArrayLength(len2)
	for i := 0; i < len2; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode OffsetDeleteResponse; Versions: 0
func (t *OffsetDeleteResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]OffsetDeleteResponseTopic47, n)
		for i := 0; i < n; i++ {
			var item OffsetDeleteResponseTopic47
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type OffsetDeleteResponseTopic47 struct {
	Name       string                            // The topic name. Versions: 0+
	Partitions []OffsetDeleteResponsePartition47 // The responses for each partition in the topic. Versions: 0+
}

// size of OffsetDeleteResponseTopic47; Versions: 0
func (t OffsetDeleteResponseTopic47) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode OffsetDeleteResponseTopic47; Versions: 0
func (t OffsetDeleteResponseTopic47) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].encode(e, version)
	}
}

// decode OffsetDeleteResponseTopic47; Versions: 0
func (t *OffsetDeleteResponseTopic47) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]OffsetDeleteResponsePartition47, n)
		for i := 0; i < n; i++ {
			var item OffsetDeleteResponsePartition47
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type OffsetDeleteResponsePartition47 struct {
	PartitionIndex int32 // The partition index. Versions: 0+
	ErrorCode      int16 // The error code, or 0 if there was no error. Versions: 0+
}

// size of OffsetDeleteResponsePartition47; Versions: 0
func (t OffsetDeleteResponsePartition47) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int16 // ErrorCode
	return sz
}

// encode OffsetDeleteResponsePartition47; Versions: 0
func (t OffsetDeleteResponsePartition47) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	e.PutInt16(t.ErrorCode)      // ErrorCode
}

// decode OffsetDeleteResponsePartition47; Versions: 0
func (t *OffsetDeleteResponsePartition47) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	return err
}
