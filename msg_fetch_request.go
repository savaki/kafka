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

// FetchRequest; ApiKey: 1, Versions: 0-11
type FetchRequest struct {
	ReplicaId      int32             // The broker ID of the follower, of -1 if this request is from a consumer. Versions: 0+
	MaxWait        int32             // The maximum time in milliseconds to wait for the response. Versions: 0+
	MinBytes       int32             // The minimum bytes to accumulate in the response. Versions: 0+
	MaxBytes       int32             // The maximum bytes to fetch.  See KIP-74 for cases where this limit may not be honored. Versions: 3+
	IsolationLevel int8              // This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records Versions: 4+
	SessionId      int32             // The fetch session ID. Versions: 7+
	Epoch          int32             // The fetch session ID. Versions: 7+
	Topics         []FetchableTopic1 // The topics to fetch. Versions: 0+
	Forgotten      []ForgottenTopic1 // In an incremental fetch request, the partitions to remove. Versions: 7+
	RackId         string            // Rack ID of the consumer making this request Versions: 11+
}

// size of FetchRequest; Versions: 0-11
func (t FetchRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // ReplicaId
	sz += sizeof.Int32 // MaxWait
	sz += sizeof.Int32 // MinBytes
	if version >= 3 {
		sz += sizeof.Int32 // MaxBytes
	}
	if version >= 4 {
		sz += sizeof.Int8 // IsolationLevel
	}
	if version >= 7 {
		sz += sizeof.Int32 // SessionId
	}
	if version >= 7 {
		sz += sizeof.Int32 // Epoch
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	if version >= 7 {
		sz += sizeof.ArrayLength // Forgotten
		for i := len(t.Forgotten) - 1; i >= 0; i-- {
			sz += t.Forgotten[i].size(version)
		}
	}
	if version >= 11 {
		sz += sizeof.String(t.RackId) // RackId
	}
	return sz
}

// encode FetchRequest; Versions: 0-11
func (t FetchRequest) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ReplicaId) // ReplicaId
	e.PutInt32(t.MaxWait)   // MaxWait
	e.PutInt32(t.MinBytes)  // MinBytes
	if version >= 3 {
		e.PutInt32(t.MaxBytes) // MaxBytes
	}
	if version >= 4 {
		e.PutInt8(t.IsolationLevel) // IsolationLevel
	}
	if version >= 7 {
		e.PutInt32(t.SessionId) // SessionId
	}
	if version >= 7 {
		e.PutInt32(t.Epoch) // Epoch
	}
	// Topics
	len7 := len(t.Topics)
	e.PutArrayLength(len7)
	for i := 0; i < len7; i++ {
		t.Topics[i].encode(e, version)
	}
	if version >= 7 {
		// Forgotten
		len8 := len(t.Forgotten)
		e.PutArrayLength(len8)
		for i := 0; i < len8; i++ {
			t.Forgotten[i].encode(e, version)
		}
	}
	if version >= 11 {
		e.PutString(t.RackId) // RackId
	}
}

// decode FetchRequest; Versions: 0-11
func (t *FetchRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ReplicaId, err = d.Int32()
	if err != nil {
		return err
	}
	t.MaxWait, err = d.Int32()
	if err != nil {
		return err
	}
	t.MinBytes, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 3 {
		t.MaxBytes, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 4 {
		t.IsolationLevel, err = d.Int8()
		if err != nil {
			return err
		}
	}
	if version >= 7 {
		t.SessionId, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 7 {
		t.Epoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]FetchableTopic1, n)
		for i := 0; i < n; i++ {
			var item FetchableTopic1
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	if version >= 7 {
		// Forgotten
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.Forgotten = make([]ForgottenTopic1, n)
			for i := 0; i < n; i++ {
				var item ForgottenTopic1
				if err := (&item).decode(d, version); err != nil {
					return err
				}
				t.Forgotten[i] = item
			}
		}
	}
	if version >= 11 {
		t.RackId, err = d.String()
		if err != nil {
			return err
		}
	}
	return err
}

type FetchableTopic1 struct {
	Name            string            // The name of the topic to fetch. Versions: 0+
	FetchPartitions []FetchPartition1 // The partitions to fetch. Versions: 0+
}

// size of FetchableTopic1; Versions: 0-11
func (t FetchableTopic1) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.ArrayLength    // FetchPartitions
	for i := len(t.FetchPartitions) - 1; i >= 0; i-- {
		sz += t.FetchPartitions[i].size(version)
	}
	return sz
}

// encode FetchableTopic1; Versions: 0-11
func (t FetchableTopic1) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
	// FetchPartitions
	len1 := len(t.FetchPartitions)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.FetchPartitions[i].encode(e, version)
	}
}

// decode FetchableTopic1; Versions: 0-11
func (t *FetchableTopic1) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	// FetchPartitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.FetchPartitions = make([]FetchPartition1, n)
		for i := 0; i < n; i++ {
			var item FetchPartition1
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.FetchPartitions[i] = item
		}
	}
	return err
}

type FetchPartition1 struct {
	PartitionIndex     int32 // The partition index. Versions: 0+
	CurrentLeaderEpoch int32 // The current leader epoch of the partition. Versions: 9+
	FetchOffset        int64 // The message offset. Versions: 0+
	LogStartOffset     int64 // The earliest available offset of the follower replica.  The field is only used when the request is sent by the follower. Versions: 5+
	MaxBytes           int32 // The maximum bytes to fetch from this partition.  See KIP-74 for cases where this limit may not be honored. Versions: 0+
}

// size of FetchPartition1; Versions: 0-11
func (t FetchPartition1) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // PartitionIndex
	if version >= 9 {
		sz += sizeof.Int32 // CurrentLeaderEpoch
	}
	sz += sizeof.Int64 // FetchOffset
	if version >= 5 {
		sz += sizeof.Int64 // LogStartOffset
	}
	sz += sizeof.Int32 // MaxBytes
	return sz
}

// encode FetchPartition1; Versions: 0-11
func (t FetchPartition1) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	if version >= 9 {
		e.PutInt32(t.CurrentLeaderEpoch) // CurrentLeaderEpoch
	}
	e.PutInt64(t.FetchOffset) // FetchOffset
	if version >= 5 {
		e.PutInt64(t.LogStartOffset) // LogStartOffset
	}
	e.PutInt32(t.MaxBytes) // MaxBytes
}

// decode FetchPartition1; Versions: 0-11
func (t *FetchPartition1) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 9 {
		t.CurrentLeaderEpoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.FetchOffset, err = d.Int64()
	if err != nil {
		return err
	}
	if version >= 5 {
		t.LogStartOffset, err = d.Int64()
		if err != nil {
			return err
		}
	}
	t.MaxBytes, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}

type ForgottenTopic1 struct {
	Name                      string  // The partition name. Versions: 7+
	ForgottenPartitionIndexes []int32 // The partitions indexes to forget. Versions: 7+
}

// size of ForgottenTopic1; Versions: 0-11
func (t ForgottenTopic1) size(version int16) int32 {
	var sz int32
	if version >= 7 {
		sz += sizeof.String(t.Name) // Name
	}
	if version >= 7 {
		sz += sizeof.Int32Array(t.ForgottenPartitionIndexes) // ForgottenPartitionIndexes
	}
	return sz
}

// encode ForgottenTopic1; Versions: 0-11
func (t ForgottenTopic1) encode(e *protocol.Encoder, version int16) {
	if version >= 7 {
		e.PutString(t.Name) // Name
	}
	if version >= 7 {
		e.PutInt32Array(t.ForgottenPartitionIndexes) // ForgottenPartitionIndexes
	}
}

// decode ForgottenTopic1; Versions: 0-11
func (t *ForgottenTopic1) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 7 {
		t.Name, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 7 {
		t.ForgottenPartitionIndexes, err = d.Int32Array()
		if err != nil {
			return err
		}
	}
	return err
}
