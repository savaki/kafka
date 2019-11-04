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

// WriteTxnMarkersResponse; ApiKey: 27, Versions: 0
type WriteTxnMarkersResponse struct {
	Markers []WritableTxnMarkerResult27 // The results for writing makers. Versions: 0+
}

// size of WriteTxnMarkersResponse; Versions: 0
func (t WriteTxnMarkersResponse) Size(version int16) int32 {
	var sz int32
	sz += sizeof.ArrayLength // Markers
	for i := len(t.Markers) - 1; i >= 0; i-- {
		sz += t.Markers[i].Size(version)
	}
	return sz
}

// encode WriteTxnMarkersResponse; Versions: 0
func (t WriteTxnMarkersResponse) Encode(e *protocol.Encoder, version int16) {
	// Markers
	len0 := len(t.Markers)
	e.PutArrayLength(len0)
	for i := 0; i < len0; i++ {
		t.Markers[i].Encode(e, version)
	}
}

// decode WriteTxnMarkersResponse; Versions: 0
func (t *WriteTxnMarkersResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	// Markers
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Markers = make([]WritableTxnMarkerResult27, n)
		for i := 0; i < n; i++ {
			var item WritableTxnMarkerResult27
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Markers[i] = item
		}
	}
	return err
}

type WritableTxnMarkerResult27 struct {
	ProducerId int64                            // The current producer ID in use by the transactional ID. Versions: 0+
	Topics     []WritableTxnMarkerTopicResult27 // The results by topic. Versions: 0+
}

// size of WritableTxnMarkerResult27; Versions: 0
func (t WritableTxnMarkerResult27) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int64       // ProducerId
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].Size(version)
	}
	return sz
}

// encode WritableTxnMarkerResult27; Versions: 0
func (t WritableTxnMarkerResult27) Encode(e *protocol.Encoder, version int16) {
	e.PutInt64(t.ProducerId) // ProducerId
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].Encode(e, version)
	}
}

// decode WritableTxnMarkerResult27; Versions: 0
func (t *WritableTxnMarkerResult27) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ProducerId, err = d.Int64()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]WritableTxnMarkerTopicResult27, n)
		for i := 0; i < n; i++ {
			var item WritableTxnMarkerTopicResult27
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type WritableTxnMarkerTopicResult27 struct {
	Name       string                               // The topic name. Versions: 0+
	Partitions []WritableTxnMarkerPartitionResult27 // The results by partition. Versions: 0+
}

// size of WritableTxnMarkerTopicResult27; Versions: 0
func (t WritableTxnMarkerTopicResult27) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].Size(version)
	}
	return sz
}

// encode WritableTxnMarkerTopicResult27; Versions: 0
func (t WritableTxnMarkerTopicResult27) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// Partitions
	len1 := len(t.Partitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Partitions[i].Encode(e, version)
	}
}

// decode WritableTxnMarkerTopicResult27; Versions: 0
func (t *WritableTxnMarkerTopicResult27) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]WritableTxnMarkerPartitionResult27, n)
		for i := 0; i < n; i++ {
			var item WritableTxnMarkerPartitionResult27
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	return err
}

type WritableTxnMarkerPartitionResult27 struct {
	PartitionIndex int32 // The partition index. Versions: 0+
	ErrorCode      int16 // The error code, or 0 if there was no error. Versions: 0+
}

// size of WritableTxnMarkerPartitionResult27; Versions: 0
func (t WritableTxnMarkerPartitionResult27) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int16 // ErrorCode
	return sz
}

// encode WritableTxnMarkerPartitionResult27; Versions: 0
func (t WritableTxnMarkerPartitionResult27) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	e.PutInt16(t.ErrorCode)      // ErrorCode
}

// decode WritableTxnMarkerPartitionResult27; Versions: 0
func (t *WritableTxnMarkerPartitionResult27) Decode(d *protocol.Decoder, version int16) error {
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
