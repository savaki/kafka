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

// CreateTopicsResponse; ApiKey: 19, Versions: 0-5
type CreateTopicsResponse struct {
	ThrottleTimeMs int32                    // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 2+
	Topics         []CreatableTopicResult19 // Results for each topic we tried to create. Versions: 0+
}

// size of CreateTopicsResponse; Versions: 0-5
func (t CreateTopicsResponse) size(version int16) int32 {
	var sz int32
	if version >= 2 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	sz += sizeof.ArrayLength // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode CreateTopicsResponse; Versions: 0-5
func (t CreateTopicsResponse) encode(e *protocol.Encoder, version int16) {
	if version >= 2 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode CreateTopicsResponse; Versions: 0-5
func (t *CreateTopicsResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 2 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]CreatableTopicResult19, n)
		for i := 0; i < n; i++ {
			var item CreatableTopicResult19
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type CreatableTopicResult19 struct {
	Name                 string                    // The topic name. Versions: 0+
	ErrorCode            int16                     // The error code, or 0 if there was no error. Versions: 0+
	ErrorMessage         string                    // The error message, or null if there was no error. Versions: 1+
	TopicConfigErrorCode int16                     // Optional topic config error returned if configs are not returned in the response. Versions: 5+
	NumPartitions        int32                     // Number of partitions of the topic. Versions: 5+
	ReplicationFactor    int16                     // Replicator factor of the topic. Versions: 5+
	Configs              []CreatableTopicConfigs19 // Configuration of the topic. Versions: 5+
}

// size of CreatableTopicResult19; Versions: 0-5
func (t CreatableTopicResult19) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name) // Name
	sz += sizeof.Int16          // ErrorCode
	if version >= 1 {
		sz += sizeof.String(t.ErrorMessage) // ErrorMessage
	}
	if version >= 5 {
		sz += sizeof.Int16 // TopicConfigErrorCode
	}
	if version >= 5 {
		sz += sizeof.Int32 // NumPartitions
	}
	if version >= 5 {
		sz += sizeof.Int16 // ReplicationFactor
	}
	if version >= 5 {
		sz += sizeof.ArrayLength // Configs
		for i := len(t.Configs) - 1; i >= 0; i-- {
			sz += t.Configs[i].size(version)
		}
	}
	return sz
}

// encode CreatableTopicResult19; Versions: 0-5
func (t CreatableTopicResult19) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name)     // Name
	e.PutInt16(t.ErrorCode) // ErrorCode
	if version >= 1 {
		e.PutString(t.ErrorMessage) // ErrorMessage
	}
	if version >= 5 {
		e.PutInt16(t.TopicConfigErrorCode) // TopicConfigErrorCode
	}
	if version >= 5 {
		e.PutInt32(t.NumPartitions) // NumPartitions
	}
	if version >= 5 {
		e.PutInt16(t.ReplicationFactor) // ReplicationFactor
	}
	if version >= 5 {
		// Configs
		len6 := len(t.Configs)
		e.PutArrayLength(len6)
		for i := 0; i < len6; i++ {
			t.Configs[i].encode(e, version)
		}
	}
}

// decode CreatableTopicResult19; Versions: 0-5
func (t *CreatableTopicResult19) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.ErrorMessage, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.TopicConfigErrorCode, err = d.Int16()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.NumPartitions, err = d.Int32()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.ReplicationFactor, err = d.Int16()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		// Configs
		if n, err := d.ArrayLength(); err != nil {
			return err
		} else if n >= 0 {
			t.Configs = make([]CreatableTopicConfigs19, n)
			for i := 0; i < n; i++ {
				var item CreatableTopicConfigs19
				if err := (&item).decode(d, version); err != nil {
					return err
				}
				t.Configs[i] = item
			}
		}
	}
	return err
}

type CreatableTopicConfigs19 struct {
	Name         string // The configuration name. Versions: 5+
	Value        string // The configuration value. Versions: 5+
	ReadOnly     bool   // True if the configuration is read-only. Versions: 5+
	ConfigSource int8   // The configuration source. Versions: 5+
	IsSensitive  bool   // True if this configuration is sensitive. Versions: 5+
}

// size of CreatableTopicConfigs19; Versions: 0-5
func (t CreatableTopicConfigs19) size(version int16) int32 {
	var sz int32
	if version >= 5 {
		sz += sizeof.String(t.Name) // Name
	}
	if version >= 5 {
		sz += sizeof.String(t.Value) // Value
	}
	if version >= 5 {
		sz += sizeof.Bool // ReadOnly
	}
	if version >= 5 {
		sz += sizeof.Int8 // ConfigSource
	}
	if version >= 5 {
		sz += sizeof.Bool // IsSensitive
	}
	return sz
}

// encode CreatableTopicConfigs19; Versions: 0-5
func (t CreatableTopicConfigs19) encode(e *protocol.Encoder, version int16) {
	if version >= 5 {
		e.PutString(t.Name) // Name
	}
	if version >= 5 {
		e.PutString(t.Value) // Value
	}
	if version >= 5 {
		e.PutBool(t.ReadOnly) // ReadOnly
	}
	if version >= 5 {
		e.PutInt8(t.ConfigSource) // ConfigSource
	}
	if version >= 5 {
		e.PutBool(t.IsSensitive) // IsSensitive
	}
}

// decode CreatableTopicConfigs19; Versions: 0-5
func (t *CreatableTopicConfigs19) decode(d *protocol.Decoder, version int16) error {
	var err error
	if version >= 5 {
		t.Name, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.Value, err = d.String()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.ReadOnly, err = d.Bool()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.ConfigSource, err = d.Int8()
		if err != nil {
			return err
		}
	}
	if version >= 5 {
		t.IsSensitive, err = d.Bool()
		if err != nil {
			return err
		}
	}
	return err
}
