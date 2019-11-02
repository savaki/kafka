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

// UpdateMetadataRequest; ApiKey: 6, Versions: 0-6
type UpdateMetadataRequest struct {
	ControllerId             int32                           // The controller id. Versions: 0+
	ControllerEpoch          int32                           // The controller epoch. Versions: 0+
	BrokerEpoch              int64                           // The broker epoch. Versions: 5+
	UngroupedPartitionStates []UpdateMetadataPartitionState6 // In older versions of this RPC, each partition that we would like to update. Versions: 0-4
	TopicStates              []UpdateMetadataTopicState6     // In newer versions of this RPC, each topic that we would like to update. Versions: 5+
	LiveBrokers              []UpdateMetadataBroker6         // Versions: 0+
}

// size of UpdateMetadataRequest; Versions: 0-6
func (t UpdateMetadataRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // ControllerId
	sz += sizeof.Int32 // ControllerEpoch
	if version >= 5 {
		sz += sizeof.Int64 // BrokerEpoch
	}
	if version >= 0 && version <= 4 {
		sz += sizeof.ArrayLength // UngroupedPartitionStates
		for i := len(t.UngroupedPartitionStates) - 1; i >= 0; i-- {
			sz += t.UngroupedPartitionStates[i].size(version)
		}
	}
	if version >= 5 {
		sz += sizeof.ArrayLength // TopicStates
		for i := len(t.TopicStates) - 1; i >= 0; i-- {
			sz += t.TopicStates[i].size(version)
		}
	}
	sz += sizeof.ArrayLength // LiveBrokers
	for i := len(t.LiveBrokers) - 1; i >= 0; i-- {
		sz += t.LiveBrokers[i].size(version)
	}
	return sz
}

// encode UpdateMetadataRequest; Versions: 0-6
func (t UpdateMetadataRequest) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ControllerId)    // ControllerId
	e.PutInt32(t.ControllerEpoch) // ControllerEpoch
	if version >= 5 {
		e.PutInt64(t.BrokerEpoch) // BrokerEpoch
	}
	if version >= 0 && version <= 4 {
		// UngroupedPartitionStates
		len3 := len(t.UngroupedPartitionStates)
		e.PutArrayLength(len3)
		for i := 0; i < len3; i++ {
			t.UngroupedPartitionStates[i].Encode(e, version)
		}
	}
	if version >= 5 {
		// TopicStates
		len4 := len(t.TopicStates)
		e.PutArrayLength(len4)
		for i := 0; i < len4; i++ {
			t.TopicStates[i].Encode(e, version)
		}
	}
	// LiveBrokers
	len5 := len(t.LiveBrokers)
	e.PutArrayLength(len5)
	for i := 0; i < len5; i++ {
		t.LiveBrokers[i].Encode(e, version)
	}
}

// decode UpdateMetadataRequest; Versions: 0-6
func (t *UpdateMetadataRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ControllerId, err = d.Int32()
	if err != nil {
		return err
	}
	t.ControllerEpoch, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 5 {
		t.BrokerEpoch, err = d.Int64()
		if err != nil {
			return err
		}
	}
	if version >= 0 && version <= 4 {
		// UngroupedPartitionStates
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.UngroupedPartitionStates = make([]UpdateMetadataPartitionState6, n)
			for i := 0; i < n; i++ {
				var item UpdateMetadataPartitionState6
				if err := (&item).Decode(d, version); err != nil {
					return err
				}
				t.UngroupedPartitionStates[i] = item
			}
		}
	}
	if version >= 5 {
		// TopicStates
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.TopicStates = make([]UpdateMetadataTopicState6, n)
			for i := 0; i < n; i++ {
				var item UpdateMetadataTopicState6
				if err := (&item).Decode(d, version); err != nil {
					return err
				}
				t.TopicStates[i] = item
			}
		}
	}
	// LiveBrokers
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.LiveBrokers = make([]UpdateMetadataBroker6, n)
		for i := 0; i < n; i++ {
			var item UpdateMetadataBroker6
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.LiveBrokers[i] = item
		}
	}
	return err
}

type UpdateMetadataTopicState6 struct {
	TopicName       string                          // The topic name. Versions: 5+
	PartitionStates []UpdateMetadataPartitionState6 // The partition that we would like to update. Versions: 5+
}

// size of UpdateMetadataTopicState6; Versions: 0-6
func (t UpdateMetadataTopicState6) size(version int16) int32 {
	var sz int32
	if version >= 5 {
		sz += sizeof.String(t.TopicName) // TopicName
	}
	if version >= 5 {
		sz += sizeof.ArrayLength // PartitionStates
		for i := len(t.PartitionStates) - 1; i >= 0; i-- {
			sz += t.PartitionStates[i].size(version)
		}
	}
	return sz
}

// encode UpdateMetadataTopicState6; Versions: 0-6
func (t UpdateMetadataTopicState6) Encode(e *protocol.Encoder, version int16) {
	if version >= 5 {
		e.PutString(t.TopicName) // TopicName
	}
	if version >= 5 {
		// PartitionStates
		len1 := len(t.PartitionStates)
		e.PutArrayLength(len1)
		for i := 0; i < len1; i++ {
			t.PartitionStates[i].Encode(e, version)
		}
	}
}

// decode UpdateMetadataTopicState6; Versions: 0-6
func (t *UpdateMetadataTopicState6) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 5 {
		t.TopicName, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		// PartitionStates
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.PartitionStates = make([]UpdateMetadataPartitionState6, n)
			for i := 0; i < n; i++ {
				var item UpdateMetadataPartitionState6
				if err := (&item).Decode(d, version); err != nil {
					return err
				}
				t.PartitionStates[i] = item
			}
		}
	}
	return err
}

type UpdateMetadataBroker6 struct {
	Id        int32                     // The broker id. Versions: 0+
	V0Host    string                    // The broker hostname. Versions: 0-0
	V0Port    int32                     // The broker port. Versions: 0-0
	Endpoints []UpdateMetadataEndpoint6 // The broker endpoints. Versions: 1+
	Rack      string                    // The rack which this broker belongs to. Versions: 2+
}

// size of UpdateMetadataBroker6; Versions: 0-6
func (t UpdateMetadataBroker6) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // Id
	if version >= 0 && version <= 0 {
		sz += sizeof.String(t.V0Host) // V0Host
	}
	if version >= 0 && version <= 0 {
		sz += sizeof.Int32 // V0Port
	}
	if version >= 1 {
		sz += sizeof.ArrayLength // Endpoints
		for i := len(t.Endpoints) - 1; i >= 0; i-- {
			sz += t.Endpoints[i].size(version)
		}
	}
	if version >= 2 {
		sz += sizeof.String(t.Rack) // Rack
	}
	return sz
}

// encode UpdateMetadataBroker6; Versions: 0-6
func (t UpdateMetadataBroker6) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.Id) // Id
	if version >= 0 && version <= 0 {
		e.PutString(t.V0Host) // V0Host
	}
	if version >= 0 && version <= 0 {
		e.PutInt32(t.V0Port) // V0Port
	}
	if version >= 1 {
		// Endpoints
		len3 := len(t.Endpoints)
		e.PutArrayLength(len3)
		for i := 0; i < len3; i++ {
			t.Endpoints[i].Encode(e, version)
		}
	}
	if version >= 2 {
		e.PutString(t.Rack) // Rack
	}
}

// decode UpdateMetadataBroker6; Versions: 0-6
func (t *UpdateMetadataBroker6) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Id, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 0 && version <= 0 {
		t.V0Host, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 0 && version <= 0 {
		t.V0Port, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 1 {
		// Endpoints
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.Endpoints = make([]UpdateMetadataEndpoint6, n)
			for i := 0; i < n; i++ {
				var item UpdateMetadataEndpoint6
				if err := (&item).Decode(d, version); err != nil {
					return err
				}
				t.Endpoints[i] = item
			}
		}
	}
	if version >= 2 {
		t.Rack, err = d.String()
		if err != nil {
			return err
		}
	}
	return err
}

type UpdateMetadataEndpoint6 struct {
	Port             int32  // The port of this endpoint Versions: 1+
	Host             string // The hostname of this endpoint Versions: 1+
	Listener         string // The listener name. Versions: 3+
	SecurityProtocol int16  // The security protocol type. Versions: 1+
}

// size of UpdateMetadataEndpoint6; Versions: 0-6
func (t UpdateMetadataEndpoint6) size(version int16) int32 {
	var sz int32
	if version >= 1 {
		sz += sizeof.Int32 // Port
	}
	if version >= 1 {
		sz += sizeof.String(t.Host) // Host
	}
	if version >= 3 {
		sz += sizeof.String(t.Listener) // Listener
	}
	if version >= 1 {
		sz += sizeof.Int16 // SecurityProtocol
	}
	return sz
}

// encode UpdateMetadataEndpoint6; Versions: 0-6
func (t UpdateMetadataEndpoint6) Encode(e *protocol.Encoder, version int16) {
	if version >= 1 {
		e.PutInt32(t.Port) // Port
	}
	if version >= 1 {
		e.PutString(t.Host) // Host
	}
	if version >= 3 {
		e.PutString(t.Listener) // Listener
	}
	if version >= 1 {
		e.PutInt16(t.SecurityProtocol) // SecurityProtocol
	}
}

// decode UpdateMetadataEndpoint6; Versions: 0-6
func (t *UpdateMetadataEndpoint6) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 1 {
		t.Port, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 1 {
		t.Host, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 3 {
		t.Listener, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 1 {
		t.SecurityProtocol, err = d.Int16()
		if err != nil {
			return err
		}
	}
	return err
}

type UpdateMetadataPartitionState6 struct {
	TopicName       string  // In older versions of this RPC, the topic name. Versions: 0-4
	PartitionIndex  int32   // The partition index. Versions: 0+
	ControllerEpoch int32   // The controller epoch. Versions: 0+
	Leader          int32   // The ID of the broker which is the current partition leader. Versions: 0+
	LeaderEpoch     int32   // The leader epoch of this partition. Versions: 0+
	Isr             []int32 // The brokers which are in the ISR for this partition. Versions: 0+
	ZkVersion       int32   // The Zookeeper version. Versions: 0+
	Replicas        []int32 // All the replicas of this partition. Versions: 0+
	OfflineReplicas []int32 // The replicas of this partition which are offline. Versions: 4+
}

// size of UpdateMetadataPartitionState6; Versions: 0-6
func (t UpdateMetadataPartitionState6) size(version int16) int32 {
	var sz int32
	if version >= 0 && version <= 4 {
		sz += sizeof.String(t.TopicName) // TopicName
	}
	sz += sizeof.Int32                  // PartitionIndex
	sz += sizeof.Int32                  // ControllerEpoch
	sz += sizeof.Int32                  // Leader
	sz += sizeof.Int32                  // LeaderEpoch
	sz += sizeof.Int32Array(t.Isr)      // Isr
	sz += sizeof.Int32                  // ZkVersion
	sz += sizeof.Int32Array(t.Replicas) // Replicas
	if version >= 4 {
		sz += sizeof.Int32Array(t.OfflineReplicas) // OfflineReplicas
	}
	return sz
}

// encode UpdateMetadataPartitionState6; Versions: 0-6
func (t UpdateMetadataPartitionState6) Encode(e *protocol.Encoder, version int16) {
	if version >= 0 && version <= 4 {
		e.PutString(t.TopicName) // TopicName
	}
	e.PutInt32(t.PartitionIndex)  // PartitionIndex
	e.PutInt32(t.ControllerEpoch) // ControllerEpoch
	e.PutInt32(t.Leader)          // Leader
	e.PutInt32(t.LeaderEpoch)     // LeaderEpoch
	e.PutInt32Array(t.Isr)        // Isr
	e.PutInt32(t.ZkVersion)       // ZkVersion
	e.PutInt32Array(t.Replicas)   // Replicas
	if version >= 4 {
		e.PutInt32Array(t.OfflineReplicas) // OfflineReplicas
	}
}

// decode UpdateMetadataPartitionState6; Versions: 0-6
func (t *UpdateMetadataPartitionState6) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 0 && version <= 4 {
		t.TopicName, err = d.String()
		if err != nil {
			return err
		}
	}
	t.PartitionIndex, err = d.Int32()
	if err != nil {
		return err
	}
	t.ControllerEpoch, err = d.Int32()
	if err != nil {
		return err
	}
	t.Leader, err = d.Int32()
	if err != nil {
		return err
	}
	t.LeaderEpoch, err = d.Int32()
	if err != nil {
		return err
	}
	t.Isr, err = d.Int32Array()
	if err != nil {
		return err
	}
	t.ZkVersion, err = d.Int32()
	if err != nil {
		return err
	}
	t.Replicas, err = d.Int32Array()
	if err != nil {
		return err
	}
	if version >= 4 {
		t.OfflineReplicas, err = d.Int32Array()
		if err != nil {
			return err
		}
	}
	return err
}
