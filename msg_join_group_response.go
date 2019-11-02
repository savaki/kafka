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

// JoinGroupResponse; ApiKey: 11, Versions: 0-6
type JoinGroupResponse struct {
	ThrottleTimeMs int32                       // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 2+
	ErrorCode      int16                       // The error code, or 0 if there was no error. Versions: 0+
	GenerationId   int32                       // The generation ID of the group. Versions: 0+
	ProtocolName   string                      // The group protocol selected by the coordinator. Versions: 0+
	Leader         string                      // The leader of the group. Versions: 0+
	MemberId       string                      // The member ID assigned by the group coordinator. Versions: 0+
	Members        []JoinGroupResponseMember11 // Versions: 0+
}

// size of JoinGroupResponse; Versions: 0-6
func (t JoinGroupResponse) size(version int16) int32 {
	var sz int32
	if version >= 2 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.Int32                  // GenerationId
	sz += sizeof.String(t.ProtocolName) // ProtocolName
	sz += sizeof.String(t.Leader)       // Leader
	sz += sizeof.String(t.MemberId)     // MemberId
	sz += sizeof.ArrayLength            // Members
	for i := len(t.Members) - 1; i >= 0; i-- {
		sz += t.Members[i].size(version)
	}
	return sz
}

// encode JoinGroupResponse; Versions: 0-6
func (t JoinGroupResponse) Encode(e *protocol.Encoder, version int16) {
	if version >= 2 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	e.PutInt16(t.ErrorCode)     // ErrorCode
	e.PutInt32(t.GenerationId)  // GenerationId
	e.PutString(t.ProtocolName) // ProtocolName
	e.PutString(t.Leader)       // Leader
	e.PutString(t.MemberId)     // MemberId
	// Members
	len6 := len(t.Members)
	e.PutArrayLength(len6)
	for i := 0; i < len6; i++ {
		t.Members[i].Encode(e, version)
	}
}

// decode JoinGroupResponse; Versions: 0-6
func (t *JoinGroupResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 2 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.GenerationId, err = d.Int32()
	if err != nil {
		return err
	}
	t.ProtocolName, err = d.String()
	if err != nil {
		return err
	}
	t.Leader, err = d.String()
	if err != nil {
		return err
	}
	t.MemberId, err = d.String()
	if err != nil {
		return err
	}
	// Members
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Members = make([]JoinGroupResponseMember11, n)
		for i := 0; i < n; i++ {
			var item JoinGroupResponseMember11
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Members[i] = item
		}
	}
	return err
}

type JoinGroupResponseMember11 struct {
	MemberId        string // The group member ID. Versions: 0+
	GroupInstanceId string // The unique identifier of the consumer instance provided by end user. Versions: 5+
	Metadata        []byte // The group member metadata. Versions: 0+
}

// size of JoinGroupResponseMember11; Versions: 0-6
func (t JoinGroupResponseMember11) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.MemberId) // MemberId
	if version >= 5 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	sz += sizeof.Bytes(t.Metadata) // Metadata
	return sz
}

// encode JoinGroupResponseMember11; Versions: 0-6
func (t JoinGroupResponseMember11) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.MemberId) // MemberId
	if version >= 5 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
	e.PutBytes(t.Metadata) // Metadata
}

// decode JoinGroupResponseMember11; Versions: 0-6
func (t *JoinGroupResponseMember11) Decode(d *protocol.Decoder, version int16) error {
	var err error
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
	t.Metadata, err = d.Bytes()
	if err != nil {
		return err
	}
	return err
}
