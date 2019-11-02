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

// MetadataResponse; ApiKey: 3, Versions: 0-9
type MetadataResponse struct {
	ThrottleTimeMs              int32                     // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 3+
	Brokers                     []MetadataResponseBroker3 // Each broker in the response. Versions: 0+
	ClusterId                   string                    // The cluster ID that responding broker belongs to. Versions: 2+
	ControllerId                int32                     // The ID of the controller broker. Versions: 1+
	Topics                      []MetadataResponseTopic3  // Each topic in the response. Versions: 0+
	ClusterAuthorizedOperations int32                     // 32-bit bitfield to represent authorized operations for this cluster. Versions: 8+
}

// size of MetadataResponse; Versions: 0-9
func (t MetadataResponse) size(version int16) int32 {
	var sz int32
	if version >= 3 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.ArrayLength // Brokers
	for i := len(t.Brokers) - 1; i >= 0; i-- {
		sz += t.Brokers[i].size(version)
	}
	if version >= 2 {
		sz += sizeof.String(t.ClusterId) // ClusterId
	}
	if version >= 1 {
		sz += sizeof.Int32 // ControllerId
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	if version >= 8 {
		sz += sizeof.Int32 // ClusterAuthorizedOperations
	}
	return sz
}

// encode MetadataResponse; Versions: 0-9
func (t MetadataResponse) encode(e *protocol.Encoder, version int16) {
	if version >= 3 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	// Brokers
	len1 := len(t.Brokers)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Brokers[i].encode(e, version)
	}
	if version >= 2 {
		e.PutString(t.ClusterId) // ClusterId
	}
	if version >= 1 {
		e.PutInt32(t.ControllerId) // ControllerId
	}
	// Topics
	len4 := len(t.Topics)
	e.PutArrayLength(len4)
	for i := 0; i < len4; i++ {
		t.Topics[i].encode(e, version)
	}
	if version >= 8 {
		e.PutInt32(t.ClusterAuthorizedOperations) // ClusterAuthorizedOperations
	}
}

// decode MetadataResponse; Versions: 0-9
func (t *MetadataResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 3 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Brokers
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Brokers = make([]MetadataResponseBroker3, n)
		for i := 0; i < n; i++ {
			var item MetadataResponseBroker3
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Brokers[i] = item
		}
	}
	if version >= 2 {
		t.ClusterId, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 1 {
		t.ControllerId, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]MetadataResponseTopic3, n)
		for i := 0; i < n; i++ {
			var item MetadataResponseTopic3
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	if version >= 8 {
		t.ClusterAuthorizedOperations, err = d.Int32()
		if err != nil {
			return err
		}
	}
	return err
}

type MetadataResponseBroker3 struct {
	NodeId int32  // The broker ID. Versions: 0+
	Host   string // The broker hostname. Versions: 0+
	Port   int32  // The broker port. Versions: 0+
	Rack   string // The rack of the broker, or null if it has not been assigned to a rack. Versions: 1+
}

// size of MetadataResponseBroker3; Versions: 0-9
func (t MetadataResponseBroker3) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32          // NodeId
	sz += sizeof.String(t.Host) // Host
	sz += sizeof.Int32          // Port
	if version >= 1 {
		sz += sizeof.String(t.Rack) // Rack
	}
	return sz
}

// encode MetadataResponseBroker3; Versions: 0-9
func (t MetadataResponseBroker3) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.NodeId) // NodeId
	e.PutString(t.Host)  // Host
	e.PutInt32(t.Port)   // Port
	if version >= 1 {
		e.PutString(t.Rack) // Rack
	}
}

// decode MetadataResponseBroker3; Versions: 0-9
func (t *MetadataResponseBroker3) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.NodeId, err = d.Int32()
	if err != nil {
		return err
	}
	t.Host, err = d.String()
	if err != nil {
		return err
	}
	t.Port, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.Rack, err = d.String()
		if err != nil {
			return err
		}
	}
	return err
}

type MetadataResponseTopic3 struct {
	ErrorCode                 int16                        // The topic error, or 0 if there was no error. Versions: 0+
	Name                      string                       // The topic name. Versions: 0+
	IsInternal                bool                         // True if the topic is internal. Versions: 1+
	Partitions                []MetadataResponsePartition3 // Each partition in the topic. Versions: 0+
	TopicAuthorizedOperations int32                        // 32-bit bitfield to represent authorized operations for this topic. Versions: 8+
}

// size of MetadataResponseTopic3; Versions: 0-9
func (t MetadataResponseTopic3) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16          // ErrorCode
	sz += sizeof.String(t.Name) // Name
	if version >= 1 {
		sz += sizeof.Bool // IsInternal
	}
	sz += sizeof.ArrayLength // Partitions
	for i := len(t.Partitions) - 1; i >= 0; i-- {
		sz += t.Partitions[i].size(version)
	}
	if version >= 8 {
		sz += sizeof.Int32 // TopicAuthorizedOperations
	}
	return sz
}

// encode MetadataResponseTopic3; Versions: 0-9
func (t MetadataResponseTopic3) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode) // ErrorCode
	e.PutString(t.Name)     // Name
	if version >= 1 {
		e.PutBool(t.IsInternal) // IsInternal
	}
	// Partitions
	len3 := len(t.Partitions)
	e.PutArrayLength(len3)
	for i := 0; i < len3; i++ {
		t.Partitions[i].encode(e, version)
	}
	if version >= 8 {
		e.PutInt32(t.TopicAuthorizedOperations) // TopicAuthorizedOperations
	}
}

// decode MetadataResponseTopic3; Versions: 0-9
func (t *MetadataResponseTopic3) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.IsInternal, err = d.Bool()
		if err != nil {
			return err
		}
	}
	// Partitions
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Partitions = make([]MetadataResponsePartition3, n)
		for i := 0; i < n; i++ {
			var item MetadataResponsePartition3
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Partitions[i] = item
		}
	}
	if version >= 8 {
		t.TopicAuthorizedOperations, err = d.Int32()
		if err != nil {
			return err
		}
	}
	return err
}

type MetadataResponsePartition3 struct {
	ErrorCode       int16   // The partition error, or 0 if there was no error. Versions: 0+
	PartitionIndex  int32   // The partition index. Versions: 0+
	LeaderId        int32   // The ID of the leader broker. Versions: 0+
	LeaderEpoch     int32   // The leader epoch of this partition. Versions: 7+
	ReplicaNodes    []int32 // The set of all nodes that host this partition. Versions: 0+
	IsrNodes        []int32 // The set of nodes that are in sync with the leader for this partition. Versions: 0+
	OfflineReplicas []int32 // The set of offline replicas of this partition. Versions: 5+
}

// size of MetadataResponsePartition3; Versions: 0-9
func (t MetadataResponsePartition3) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16 // ErrorCode
	sz += sizeof.Int32 // PartitionIndex
	sz += sizeof.Int32 // LeaderId
	if version >= 7 {
		sz += sizeof.Int32 // LeaderEpoch
	}
	sz += sizeof.Int32Array(t.ReplicaNodes) // ReplicaNodes
	sz += sizeof.Int32Array(t.IsrNodes)     // IsrNodes
	if version >= 5 {
		sz += sizeof.Int32Array(t.OfflineReplicas) // OfflineReplicas
	}
	return sz
}

// encode MetadataResponsePartition3; Versions: 0-9
func (t MetadataResponsePartition3) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)      // ErrorCode
	e.PutInt32(t.PartitionIndex) // PartitionIndex
	e.PutInt32(t.LeaderId)       // LeaderId
	if version >= 7 {
		e.PutInt32(t.LeaderEpoch) // LeaderEpoch
	}
	e.PutInt32Array(t.ReplicaNodes) // ReplicaNodes
	e.PutInt32Array(t.IsrNodes)     // IsrNodes
	if version >= 5 {
		e.PutInt32Array(t.OfflineReplicas) // OfflineReplicas
	}
}

// decode MetadataResponsePartition3; Versions: 0-9
func (t *MetadataResponsePartition3) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.LeaderId, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 7 {
		t.LeaderEpoch, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.ReplicaNodes, err = d.Int32Array()
	if err != nil {
		return err
	}
	t.IsrNodes, err = d.Int32Array()
	if err != nil {
		return err
	}
	if version >= 5 {
		t.OfflineReplicas, err = d.Int32Array()
		if err != nil {
			return err
		}
	}
	return err
}
