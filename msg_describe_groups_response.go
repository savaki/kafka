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

// DescribeGroupsResponse; ApiKey: 15, Versions: 0-5
type DescribeGroupsResponse struct {
	ThrottleTimeMs int32              // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 1+
	Groups         []DescribedGroup15 // Each described group. Versions: 0+
}

// size of DescribeGroupsResponse; Versions: 0-5
func (t DescribeGroupsResponse) size(version int16) int32 {
	var sz int32
	if version >= 1 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.ArrayLength // Groups
	for i := len(t.Groups) - 1; i >= 0; i-- {
		sz += t.Groups[i].size(version)
	}
	return sz
}

// encode DescribeGroupsResponse; Versions: 0-5
func (t DescribeGroupsResponse) encode(e *protocol.Encoder, version int16) {
	if version >= 1 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	// Groups
	len1 := len(t.Groups)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Groups[i].encode(e, version)
	}
}

// decode DescribeGroupsResponse; Versions: 0-5
func (t *DescribeGroupsResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 1 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Groups
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Groups = make([]DescribedGroup15, n)
		for i := 0; i < n; i++ {
			var item DescribedGroup15
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Groups[i] = item
		}
	}
	return err
}

type DescribedGroup15 struct {
	ErrorCode            int16                    // The describe error, or 0 if there was no error. Versions: 0+
	GroupId              string                   // The group ID string. Versions: 0+
	GroupState           string                   // The group state string, or the empty string. Versions: 0+
	ProtocolType         string                   // The group protocol type, or the empty string. Versions: 0+
	ProtocolData         string                   // The group protocol data, or the empty string. Versions: 0+
	Members              []DescribedGroupMember15 // The group members. Versions: 0+
	AuthorizedOperations int32                    // 32-bit bitfield to represent authorized operations for this group. Versions: 3+
}

// size of DescribedGroup15; Versions: 0-5
func (t DescribedGroup15) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.String(t.GroupId)      // GroupId
	sz += sizeof.String(t.GroupState)   // GroupState
	sz += sizeof.String(t.ProtocolType) // ProtocolType
	sz += sizeof.String(t.ProtocolData) // ProtocolData
	sz += sizeof.ArrayLength            // Members
	for i := len(t.Members) - 1; i >= 0; i-- {
		sz += t.Members[i].size(version)
	}
	if version >= 3 {
		sz += sizeof.Int32 // AuthorizedOperations
	}
	return sz
}

// encode DescribedGroup15; Versions: 0-5
func (t DescribedGroup15) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)     // ErrorCode
	e.PutString(t.GroupId)      // GroupId
	e.PutString(t.GroupState)   // GroupState
	e.PutString(t.ProtocolType) // ProtocolType
	e.PutString(t.ProtocolData) // ProtocolData
	// Members
	len5 := len(t.Members)
	e.PutArrayLength(len5)
	for i := 0; i < len5; i++ {
		t.Members[i].encode(e, version)
	}
	if version >= 3 {
		e.PutInt32(t.AuthorizedOperations) // AuthorizedOperations
	}
}

// decode DescribedGroup15; Versions: 0-5
func (t *DescribedGroup15) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	t.GroupState, err = d.String()
	if err != nil {
		return err
	}
	t.ProtocolType, err = d.String()
	if err != nil {
		return err
	}
	t.ProtocolData, err = d.String()
	if err != nil {
		return err
	}
	// Members
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Members = make([]DescribedGroupMember15, n)
		for i := 0; i < n; i++ {
			var item DescribedGroupMember15
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Members[i] = item
		}
	}
	if version >= 3 {
		t.AuthorizedOperations, err = d.Int32()
		if err != nil {
			return err
		}
	}
	return err
}

type DescribedGroupMember15 struct {
	MemberId         string // The member ID assigned by the group coordinator. Versions: 0+
	GroupInstanceId  string // The unique identifier of the consumer instance provided by end user. Versions: 4+
	ClientId         string // The client ID used in the member's latest join group request. Versions: 0+
	ClientHost       string // The client host. Versions: 0+
	MemberMetadata   []byte // The metadata corresponding to the current group protocol in use. Versions: 0+
	MemberAssignment []byte // The current assignment provided by the group leader. Versions: 0+
}

// size of DescribedGroupMember15; Versions: 0-5
func (t DescribedGroupMember15) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.MemberId) // MemberId
	if version >= 4 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	sz += sizeof.String(t.ClientId)        // ClientId
	sz += sizeof.String(t.ClientHost)      // ClientHost
	sz += sizeof.Bytes(t.MemberMetadata)   // MemberMetadata
	sz += sizeof.Bytes(t.MemberAssignment) // MemberAssignment
	return sz
}

// encode DescribedGroupMember15; Versions: 0-5
func (t DescribedGroupMember15) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.MemberId) // MemberId
	if version >= 4 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
	e.PutString(t.ClientId)        // ClientId
	e.PutString(t.ClientHost)      // ClientHost
	e.PutBytes(t.MemberMetadata)   // MemberMetadata
	e.PutBytes(t.MemberAssignment) // MemberAssignment
}

// decode DescribedGroupMember15; Versions: 0-5
func (t *DescribedGroupMember15) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.MemberId, err = d.String()
	if err != nil {
		return err
	}
	if version >= 4 {
		t.GroupInstanceId, err = d.String()
		if err != nil {
			return err
		}
	}
	t.ClientId, err = d.String()
	if err != nil {
		return err
	}
	t.ClientHost, err = d.String()
	if err != nil {
		return err
	}
	t.MemberMetadata, err = d.Bytes()
	if err != nil {
		return err
	}
	t.MemberAssignment, err = d.Bytes()
	if err != nil {
		return err
	}
	return err
}
