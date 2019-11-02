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

// SyncGroupRequest; ApiKey: 14, Versions: 0-4
type SyncGroupRequest struct {
	GroupId         string                         // The unique group identifier. Versions: 0+
	GenerationId    int32                          // The generation of the group. Versions: 0+
	MemberId        string                         // The member ID assigned by the group. Versions: 0+
	GroupInstanceId string                         // The unique identifier of the consumer instance provided by end user. Versions: 3+
	Assignments     []SyncGroupRequestAssignment14 // Each assignment. Versions: 0+
}

// size of SyncGroupRequest; Versions: 0-4
func (t SyncGroupRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.GroupId)  // GroupId
	sz += sizeof.Int32              // GenerationId
	sz += sizeof.String(t.MemberId) // MemberId
	if version >= 3 {
		sz += sizeof.String(t.GroupInstanceId) // GroupInstanceId
	}
	sz += sizeof.ArrayLength // Assignments
	for i := len(t.Assignments) - 1; i >= 0; i-- {
		sz += t.Assignments[i].size(version)
	}
	return sz
}

// encode SyncGroupRequest; Versions: 0-4
func (t SyncGroupRequest) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.GroupId)     // GroupId
	e.PutInt32(t.GenerationId) // GenerationId
	e.PutString(t.MemberId)    // MemberId
	if version >= 3 {
		e.PutString(t.GroupInstanceId) // GroupInstanceId
	}
	// Assignments
	len4 := len(t.Assignments)
	e.PutArrayLength(len4)
	for i := 0; i < len4; i++ {
		t.Assignments[i].encode(e, version)
	}
}

// decode SyncGroupRequest; Versions: 0-4
func (t *SyncGroupRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.GroupId, err = d.String()
	if err != nil {
		return err
	}
	t.GenerationId, err = d.Int32()
	if err != nil {
		return err
	}
	t.MemberId, err = d.String()
	if err != nil {
		return err
	}
	if version >= 3 {
		t.GroupInstanceId, err = d.String()
		if err != nil {
			return err
		}
	}
	// Assignments
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Assignments = make([]SyncGroupRequestAssignment14, n)
		for i := 0; i < n; i++ {
			var item SyncGroupRequestAssignment14
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Assignments[i] = item
		}
	}
	return err
}

type SyncGroupRequestAssignment14 struct {
	MemberId   string // The ID of the member to assign. Versions: 0+
	Assignment []byte // The member assignment. Versions: 0+
}

// size of SyncGroupRequestAssignment14; Versions: 0-4
func (t SyncGroupRequestAssignment14) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.MemberId)  // MemberId
	sz += sizeof.Bytes(t.Assignment) // Assignment
	return sz
}

// encode SyncGroupRequestAssignment14; Versions: 0-4
func (t SyncGroupRequestAssignment14) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.MemberId)  // MemberId
	e.PutBytes(t.Assignment) // Assignment
}

// decode SyncGroupRequestAssignment14; Versions: 0-4
func (t *SyncGroupRequestAssignment14) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.MemberId, err = d.String()
	if err != nil {
		return err
	}
	t.Assignment, err = d.Bytes()
	if err != nil {
		return err
	}
	return err
}
