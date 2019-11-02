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

// OffsetDeleteRequest; ApiKey: 47, Versions: 0
type OffsetDeleteRequest struct {
	GroupId string                       // The unique group identifier. Versions: 0+
	Topics  []OffsetDeleteRequestTopic47 // The topics to delete offsets for Versions: 0+
}

// size of OffsetDeleteRequest; Versions: 0
func (t OffsetDeleteRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.GroupId) // GroupId
	sz += sizeof.ArrayLength       // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode OffsetDeleteRequest; Versions: 0
func (t OffsetDeleteRequest) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.GroupId) // GroupId
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode OffsetDeleteRequest; Versions: 0
func (t *OffsetDeleteRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]OffsetDeleteRequestTopic47, n)
		for i := 0; i < n; i++ {
			var item OffsetDeleteRequestTopic47
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type OffsetDeleteRequestTopic47 struct {
	Name       string                           // The topic name. Versions: 0+
	Partitions []OffsetDeleteRequestPartition47 // Each partition to delete offsets for. Versions: 0+
}

// size of OffsetDeleteRequestTopic47; Versions: 0
func (t OffsetDeleteRequestTopic47) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	return sz
}

// encode OffsetDeleteRequestTopic47; Versions: 0
func (t OffsetDeleteRequestTopic47) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].encode(e, version)
	}
}

// decode OffsetDeleteRequestTopic47; Versions: 0
func (t *OffsetDeleteRequestTopic47) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]OffsetDeleteRequestPartition47, n)
		for i := 0; i < n; i++ {
			var item OffsetDeleteRequestPartition47
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type OffsetDeleteRequestPartition47 struct {
	PartitionIndex int32 // The partition index. Versions: 0+
}

// size of OffsetDeleteRequestPartition47; Versions: 0
func (t OffsetDeleteRequestPartition47) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	return sz
}

// encode OffsetDeleteRequestPartition47; Versions: 0
func (t OffsetDeleteRequestPartition47) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
}

// decode OffsetDeleteRequestPartition47; Versions: 0
func (t *OffsetDeleteRequestPartition47) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}
