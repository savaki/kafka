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

// LeaveGroupRequest; ApiKey: 13, Versions: 0-4
type LeaveGroupRequest struct {
	GroupId  string             // The ID of the group to leave. Versions: 0+
	MemberId string             // The member ID to remove from the group. Versions: 0-2
	Members  []MemberIdentity13 // List of leaving member identities. Versions: 3+
}

// size of LeaveGroupRequest; Versions: 0-4
func (t LeaveGroupRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.GroupId) // GroupId
	if version >= 0 && version <= 2 {
		sz += sizeof.String(t.MemberId) // MemberId
	}
	if version >= 3 {
		sz += sizeof.ArrayLength // Members
		for i := len(t.Members) - 1; i >= 0; i-- {
			sz += t.Members[i].size(version)
		}
	}
	return sz
}

// encode LeaveGroupRequest; Versions: 0-4
func (t LeaveGroupRequest) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.GroupId) // GroupId
	if version >= 0 && version <= 2 {
		e.PutString(t.MemberId) // MemberId
	}
	if version >= 3 {
		// Members
		len2 := len(t.Members)
		e.PutArrayLength(len2)
		for i := 0; i < len2; i++ {
			t.Members[i].Encode(e, version)
		}
	}
}

// decode LeaveGroupRequest; Versions: 0-4
func (t *LeaveGroupRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	if version >= 0 && version <= 2 {
		t.MemberId, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 3 {
		// Members
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.Members = make([]MemberIdentity13, n)
			for i := 0; i < n; i++ {
				var item MemberIdentity13
				if err := (&item).Decode(d, version); err != nil {
					return err
				}
				t.Members[i] = item
			}
		}
	}
	return err
}

type MemberIdentity13 struct {
	MemberId        string // The member ID to remove from the group. Versions: 3+
	GroupInstanceId string // The group instance ID to remove from the group. Versions: 3+
}

// size of MemberIdentity13; Versions: 0-4
func (t MemberIdentity13) size(version int16) int32 {
	var sz int32
	if version >= 3 {
		sz += sizeof.String(t.MemberId) // MemberId
	}
	if version >= 3 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	return sz
}

// encode MemberIdentity13; Versions: 0-4
func (t MemberIdentity13) Encode(e *protocol.Encoder, version int16) {
	if version >= 3 {
		e.PutString(t.MemberId) // MemberId
	}
	if version >= 3 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
}

// decode MemberIdentity13; Versions: 0-4
func (t *MemberIdentity13) Decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 3 {
		t.MemberId, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 3 {
		t.GroupInstanceId, err = d.String()
		if err != nil {
			return err
		}
	}
	return err
}
