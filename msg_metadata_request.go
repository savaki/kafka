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

// MetadataRequest; ApiKey: 3, Versions: 0-9
type MetadataRequest struct {
	Topics                             []MetadataRequestTopic3 // The topics to fetch metadata for. Versions: 0+
	AllowAutoTopicCreation             bool                    // If this is true, the broker may auto-create topics that we requested which do not already exist, if it is configured to do so. Versions: 4+
	IncludeClusterAuthorizedOperations bool                    // Whether to include cluster authorized operations. Versions: 8+
	IncludeTopicAuthorizedOperations   bool                    // Whether to include topic authorized operations. Versions: 8+
}

// size of MetadataRequest; Versions: 0-9
func (t MetadataRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	if version >= 4 {
		sz += sizeof.Bool // AllowAutoTopicCreation
	}
	if version >= 8 {
		sz += sizeof.Bool // IncludeClusterAuthorizedOperations
	}
	if version >= 8 {
		sz += sizeof.Bool // IncludeTopicAuthorizedOperations
	}
	return sz
}

// encode MetadataRequest; Versions: 0-9
func (t MetadataRequest) encode(e *protocol.Encoder, version int16) {
	// Topics
	len0 := len(t.Topics)
	e.PutArrayLength(len0)
	for i := 0; i < len0; i++ {
		t.Topics[i].encode(e, version)
	}
	if version >= 4 {
		e.PutBool(t.AllowAutoTopicCreation) // AllowAutoTopicCreation
	}
	if version >= 8 {
		e.PutBool(t.IncludeClusterAuthorizedOperations) // IncludeClusterAuthorizedOperations
	}
	if version >= 8 {
		e.PutBool(t.IncludeTopicAuthorizedOperations) // IncludeTopicAuthorizedOperations
	}
}

// decode MetadataRequest; Versions: 0-9
func (t *MetadataRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]MetadataRequestTopic3, n)
		for i := 0; i < n; i++ {
			var item MetadataRequestTopic3
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	if version >= 4 {
		t.AllowAutoTopicCreation, err = d.Bool()
		if err != nil {
			return err
		}
	}
	if version >= 8 {
		t.IncludeClusterAuthorizedOperations, err = d.Bool()
		if err != nil {
			return err
		}
	}
	if version >= 8 {
		t.IncludeTopicAuthorizedOperations, err = d.Bool()
		if err != nil {
			return err
		}
	}
	return err
}

type MetadataRequestTopic3 struct {
	Name string // The topic name. Versions: 0+
}

// size of MetadataRequestTopic3; Versions: 0-9
func (t MetadataRequestTopic3) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	return sz
}

// encode MetadataRequestTopic3; Versions: 0-9
func (t MetadataRequestTopic3) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name) // Name
}

// decode MetadataRequestTopic3; Versions: 0-9
func (t *MetadataRequestTopic3) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	return err
}
