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

// IncrementalAlterConfigsRequest; ApiKey: 44, Versions: 0-1
type IncrementalAlterConfigsRequest struct {
	Resources    []AlterConfigsResource44 // The incremental updates for each resource. Versions: 0+
	ValidateOnly bool                     // True if we should validate the request, but not change the configurations. Versions: 0+
}

// size of IncrementalAlterConfigsRequest; Versions: 0-1
func (t IncrementalAlterConfigsRequest) Size(version int16) int32 {
	var sz int32
	sz += sizeof.ArrayLength // Resources
	for i := len(t.Resources) - 1; i >= 0; i-- {
		sz += t.Resources[i].Size(version)
	}
	sz += sizeof.Bool // ValidateOnly
	return sz
}

// encode IncrementalAlterConfigsRequest; Versions: 0-1
func (t IncrementalAlterConfigsRequest) Encode(e *protocol.Encoder, version int16) {
	// Resources
	len0 := len(t.Resources)
	e.PutArrayLength(len0)
	for i := 0; i < len0; i++ {
		t.Resources[i].Encode(e, version)
	}
	e.PutBool(t.ValidateOnly) // ValidateOnly
}

// decode IncrementalAlterConfigsRequest; Versions: 0-1
func (t *IncrementalAlterConfigsRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	// Resources
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Resources = make([]AlterConfigsResource44, n)
		for i := 0; i < n; i++ {
			var item AlterConfigsResource44
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Resources[i] = item
		}
	}
	t.ValidateOnly, err = d.Bool()
	if err != nil {
		return err
	}
	return err
}

type AlterConfigsResource44 struct {
	ResourceType int8                // The resource type. Versions: 0+
	ResourceName string              // The resource name. Versions: 0+
	Configs      []AlterableConfig44 // The configurations. Versions: 0+
}

// size of AlterConfigsResource44; Versions: 0-1
func (t AlterConfigsResource44) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int8                   // ResourceType
	sz += sizeof.String(t.ResourceName) // ResourceName
	sz += sizeof.ArrayLength            // Configs
	for i := len(t.Configs) - 1; i >= 0; i-- {
		sz += t.Configs[i].Size(version)
	}
	return sz
}

// encode AlterConfigsResource44; Versions: 0-1
func (t AlterConfigsResource44) Encode(e *protocol.Encoder, version int16) {
	e.PutInt8(t.ResourceType)   // ResourceType
	e.PutString(t.ResourceName) // ResourceName
	// Configs
	len2 := len(t.Configs)
	e.PutArrayLength(len2)
	for i := 0; i < len2; i++ {
		t.Configs[i].Encode(e, version)
	}
}

// decode AlterConfigsResource44; Versions: 0-1
func (t *AlterConfigsResource44) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ResourceType, err = d.Int8()
	if err != nil {
		return err
	}
	t.ResourceName, err = d.String()
	if err != nil {
		return err
	}
	// Configs
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Configs = make([]AlterableConfig44, n)
		for i := 0; i < n; i++ {
			var item AlterableConfig44
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Configs[i] = item
		}
	}
	return err
}

type AlterableConfig44 struct {
	Name            string // The configuration key name. Versions: 0+
	ConfigOperation int8   // The type (Set, Delete, Append, Subtract) of operation. Versions: 0+
	Value           string // The value to set for the configuration key. Versions: 0+
}

// size of AlterableConfig44; Versions: 0-1
func (t AlterableConfig44) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name)  // Name
	sz += sizeof.Int8            // ConfigOperation
	sz += sizeof.String(t.Value) // Value
	return sz
}

// encode AlterableConfig44; Versions: 0-1
func (t AlterableConfig44) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name)          // Name
	e.PutInt8(t.ConfigOperation) // ConfigOperation
	e.PutString(t.Value)         // Value
}

// decode AlterableConfig44; Versions: 0-1
func (t *AlterableConfig44) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	t.ConfigOperation, err = d.Int8()
	if err != nil {
		return err
	}
	t.Value, err = d.String()
	if err != nil {
		return err
	}
	return err
}
