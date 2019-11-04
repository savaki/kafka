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

// WriteTxnMarkersRequest; ApiKey: 27, Versions: 0
type WriteTxnMarkersRequest struct {
	Markers []WritableTxnMarker27 // The transaction markers to be written. Versions: 0+
}

// size of WriteTxnMarkersRequest; Versions: 0
func (t WriteTxnMarkersRequest) Size(version int16) int32 {
	var sz int32
	sz += sizeof.ArrayLength // Markers
	for i := len(t.Markers) - 1; i >= 0; i-- {
		sz += t.Markers[i].Size(version)
	}
	return sz
}

// encode WriteTxnMarkersRequest; Versions: 0
func (t WriteTxnMarkersRequest) Encode(e *protocol.Encoder, version int16) {
	// Markers
	len0 := len(t.Markers)
	e.PutArrayLength(len0)
	for i := 0; i < len0; i++ {
		t.Markers[i].Encode(e, version)
	}
}

// decode WriteTxnMarkersRequest; Versions: 0
func (t *WriteTxnMarkersRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	// Markers
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Markers = make([]WritableTxnMarker27, n)
		for i := 0; i < n; i++ {
			var item WritableTxnMarker27
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Markers[i] = item
		}
	}
	return err
}

type WritableTxnMarker27 struct {
	ProducerId        int64                      // The current producer ID. Versions: 0+
	ProducerEpoch     int16                      // The current epoch associated with the producer ID. Versions: 0+
	TransactionResult bool                       // The result of the transaction to write to the partitions (false = ABORT, true = COMMIT). Versions: 0+
	Topics            []WritableTxnMarkerTopic27 // Each topic that we want to write transaction marker(s) for. Versions: 0+
	CoordinatorEpoch  int32                      // Epoch associated with the transaction state partition hosted by this transaction coordinator Versions: 0+
}

// size of WritableTxnMarker27; Versions: 0
func (t WritableTxnMarker27) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int64       // ProducerId
	sz += sizeof.Int16       // ProducerEpoch
	sz += sizeof.Bool        // TransactionResult
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].Size(version)
	}
	sz += sizeof.Int32 // CoordinatorEpoch
	return sz
}

// encode WritableTxnMarker27; Versions: 0
func (t WritableTxnMarker27) Encode(e *protocol.Encoder, version int16) {
	e.PutInt64(t.ProducerId)       // ProducerId
	e.PutInt16(t.ProducerEpoch)    // ProducerEpoch
	e.PutBool(t.TransactionResult) // TransactionResult
	// Topics
	len3 := len(t.Topics)
	e.PutArrayLength(len3)
	for i := 0; i < len3; i++ {
		t.Topics[i].Encode(e, version)
	}
	e.PutInt32(t.CoordinatorEpoch) // CoordinatorEpoch
}

// decode WritableTxnMarker27; Versions: 0
func (t *WritableTxnMarker27) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ProducerId, err = d.Int64()
	if err != nil {
		return err
	}
	t.ProducerEpoch, err = d.Int16()
	if err != nil {
		return err
	}
	t.TransactionResult, err = d.Bool()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]WritableTxnMarkerTopic27, n)
		for i := 0; i < n; i++ {
			var item WritableTxnMarkerTopic27
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	t.CoordinatorEpoch, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}

type WritableTxnMarkerTopic27 struct {
	Name             string  // The topic name. Versions: 0+
	PartitionIndexes []int32 // The indexes of the partitions to write transaction markers for. Versions: 0+
}

// size of WritableTxnMarkerTopic27; Versions: 0
func (t WritableTxnMarkerTopic27) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name)                 // Name
	sz += sizeof.Int32Array(t.PartitionIndexes) // PartitionIndexes
	return sz
}

// encode WritableTxnMarkerTopic27; Versions: 0
func (t WritableTxnMarkerTopic27) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name)                 // Name
	e.PutInt32Array(t.PartitionIndexes) // PartitionIndexes
}

// decode WritableTxnMarkerTopic27; Versions: 0
func (t *WritableTxnMarkerTopic27) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	t.PartitionIndexes, err = d.Int32Array()
	if err != nil {
		return err
	}
	return err
}
