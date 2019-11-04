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

// JoinGroupRequest; ApiKey: 11, Versions: 0-6
type JoinGroupRequest struct {
	GroupId            string                       // The group identifier. Versions: 0+
	SessionTimeoutMs   int32                        // The coordinator considers the consumer dead if it receives no heartbeat after this timeout in milliseconds. Versions: 0+
	RebalanceTimeoutMs int32                        // The maximum time in milliseconds that the coordinator will wait for each member to rejoin when rebalancing the group. Versions: 1+
	MemberId           string                       // The member id assigned by the group coordinator. Versions: 0+
	GroupInstanceId    string                       // The unique identifier of the consumer instance provided by end user. Versions: 5+
	ProtocolType       string                       // The unique name the for class of protocols implemented by the group we want to join. Versions: 0+
	Protocols          []JoinGroupRequestProtocol11 // The list of protocols that the member supports. Versions: 0+
}

// size of JoinGroupRequest; Versions: 0-6
func (t JoinGroupRequest) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.GroupId) // GroupId
	sz += sizeof.Int32             // SessionTimeoutMs
	if version >= 1 {
		sz += sizeof.Int32 // RebalanceTimeoutMs
	}
	sz += sizeof.String(t.MemberId) // MemberId
	if version >= 5 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	sz += sizeof.String(t.ProtocolType) // ProtocolType
	sz += sizeof.ArrayLength            // Protocols
	for i := len(t.Protocols) - 1; i >= 0; i-- {
		sz += t.Protocols[i].Size(version)
	}
	return sz
}

// encode JoinGroupRequest; Versions: 0-6
func (t JoinGroupRequest) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.GroupId)         // GroupId
	e.PutInt32(t.SessionTimeoutMs) // SessionTimeoutMs
	if version >= 1 {
		e.PutInt32(t.RebalanceTimeoutMs) // RebalanceTimeoutMs
	}
	e.PutString(t.MemberId) // MemberId
	if version >= 5 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
	e.PutString(t.ProtocolType) // ProtocolType
	// Protocols
	len6 := len(t.Protocols)
	e.PutArrayLength(len6)
	for i := 0; i < len6; i++ {
		t.Protocols[i].Encode(e, version)
	}
}

// decode JoinGroupRequest; Versions: 0-6
func (t *JoinGroupRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	t.SessionTimeoutMs, err = d.Int32()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.RebalanceTimeoutMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.MemberId, err = d.String()
	if err != nil {
		return err
	}
	if version >= 5 {
		t.GroupInstanceId, err = d.String()
		if err != nil {
			return err
		}
	}
	t.ProtocolType, err = d.String()
	if err != nil {
		return err
	}
	// Protocols
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Protocols = make([]JoinGroupRequestProtocol11, n)
		for i := 0; i < n; i++ {
			var item JoinGroupRequestProtocol11
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Protocols[i] = item
		}
	}
	return err
}

type JoinGroupRequestProtocol11 struct {
	Name     string // The protocol name. Versions: 0+
	Metadata []byte // The protocol metadata. Versions: 0+
}

// size of JoinGroupRequestProtocol11; Versions: 0-6
func (t JoinGroupRequestProtocol11) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name)    // Name
	sz += sizeof.Bytes(t.Metadata) // Metadata
	return sz
}

// encode JoinGroupRequestProtocol11; Versions: 0-6
func (t JoinGroupRequestProtocol11) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name)    // Name
	e.PutBytes(t.Metadata) // Metadata
}

// decode JoinGroupRequestProtocol11; Versions: 0-6
func (t *JoinGroupRequestProtocol11) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	t.Metadata, err = d.Bytes()
	if err != nil {
		return err
	}
	return err
}
