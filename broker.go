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
	"fmt"
	"net"

	"github.com/savaki/kafka/protocol"
	"github.com/savaki/kafka/protocol/apikey"
)

type config struct {
	brokers  []string
	clientID string
	dialFunc func(network, addr string) (net.Conn, error)
}

// Client provides access to Kafka
type Broker struct {
	apiVersion apiVersion
	config     config
	conn       *protocol.Conn
}

// Option provides functional options to Client
type Option func(*config)

// buildConfig builds config from provided options
func buildConfig(opts []Option) config {
	c := config{
		clientID: "github.com/savaki/kafka",
		dialFunc: net.Dial,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

// New a new Kafka Broker
func New(addr string, opts ...Option) (*Broker, error) {
	config := buildConfig(opts)
	conn, err := protocol.Connect(addr)
	if err != nil {
		return nil, err
	}

	broker := &Broker{
		config: config,
		conn:   conn,
	}

	// negotiate usable api versions prior to returning broker
	//
	resp, err := broker.ApiVersions(ApiVersionsRequest{})
	if err != nil {
		return nil, fmt.Errorf("unable to create broker: %w", err)
	}
	apiVersion, err := negotiateApiVersions(resp.ApiKeys)
	if err != nil {
		return nil, fmt.Errorf("unable to create broker: %w", err)
	}
	broker.apiVersion = apiVersion

	return broker, nil
}

// Close the connection to the broker
func (b *Broker) Close() error {
	return b.conn.Close()
}

// Produce (apiKey: 0)
func (b *Broker) Produce(req ProduceRequest) (ProduceResponse, error) {
	var resp ProduceResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.Produce,
				RequestApiVersion: b.apiVersion.Produce,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.Produce)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.Produce)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.Produce)
		},
	)
	return resp, err
}

// Fetch (apiKey: 1)
func (b *Broker) Fetch(req FetchRequest) (FetchResponse, error) {
	var resp FetchResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.Fetch,
				RequestApiVersion: b.apiVersion.Fetch,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.Fetch)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.Fetch)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.Fetch)
		},
	)
	return resp, err
}

// ListOffset (apiKey: 2)
func (b *Broker) ListOffset(req ListOffsetRequest) (ListOffsetResponse, error) {
	var resp ListOffsetResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ListOffset,
				RequestApiVersion: b.apiVersion.ListOffset,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ListOffset)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ListOffset)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ListOffset)
		},
	)
	return resp, err
}

// Metadata (apiKey: 3)
func (b *Broker) Metadata(req MetadataRequest) (MetadataResponse, error) {
	var resp MetadataResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.Metadata,
				RequestApiVersion: b.apiVersion.Metadata,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.Metadata)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.Metadata)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.Metadata)
		},
	)
	return resp, err
}

// LeaderAndIsr (apiKey: 4)
func (b *Broker) LeaderAndIsr(req LeaderAndIsrRequest) (LeaderAndIsrResponse, error) {
	var resp LeaderAndIsrResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.LeaderAndIsr,
				RequestApiVersion: b.apiVersion.LeaderAndIsr,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.LeaderAndIsr)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.LeaderAndIsr)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.LeaderAndIsr)
		},
	)
	return resp, err
}

// StopReplica (apiKey: 5)
func (b *Broker) StopReplica(req StopReplicaRequest) (StopReplicaResponse, error) {
	var resp StopReplicaResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.StopReplica,
				RequestApiVersion: b.apiVersion.StopReplica,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.StopReplica)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.StopReplica)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.StopReplica)
		},
	)
	return resp, err
}

// UpdateMetadata (apiKey: 6)
func (b *Broker) UpdateMetadata(req UpdateMetadataRequest) (UpdateMetadataResponse, error) {
	var resp UpdateMetadataResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.UpdateMetadata,
				RequestApiVersion: b.apiVersion.UpdateMetadata,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.UpdateMetadata)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.UpdateMetadata)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.UpdateMetadata)
		},
	)
	return resp, err
}

// ControlledShutdown (apiKey: 7)
func (b *Broker) ControlledShutdown(req ControlledShutdownRequest) (ControlledShutdownResponse, error) {
	var resp ControlledShutdownResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ControlledShutdown,
				RequestApiVersion: b.apiVersion.ControlledShutdown,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ControlledShutdown)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ControlledShutdown)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ControlledShutdown)
		},
	)
	return resp, err
}

// OffsetCommit (apiKey: 8)
func (b *Broker) OffsetCommit(req OffsetCommitRequest) (OffsetCommitResponse, error) {
	var resp OffsetCommitResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.OffsetCommit,
				RequestApiVersion: b.apiVersion.OffsetCommit,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.OffsetCommit)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.OffsetCommit)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.OffsetCommit)
		},
	)
	return resp, err
}

// OffsetFetch (apiKey: 9)
func (b *Broker) OffsetFetch(req OffsetFetchRequest) (OffsetFetchResponse, error) {
	var resp OffsetFetchResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.OffsetFetch,
				RequestApiVersion: b.apiVersion.OffsetFetch,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.OffsetFetch)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.OffsetFetch)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.OffsetFetch)
		},
	)
	return resp, err
}

// FindCoordinator (apiKey: 10)
func (b *Broker) FindCoordinator(req FindCoordinatorRequest) (FindCoordinatorResponse, error) {
	var resp FindCoordinatorResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.FindCoordinator,
				RequestApiVersion: b.apiVersion.FindCoordinator,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.FindCoordinator)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.FindCoordinator)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.FindCoordinator)
		},
	)
	return resp, err
}

// JoinGroup (apiKey: 11)
func (b *Broker) JoinGroup(req JoinGroupRequest) (JoinGroupResponse, error) {
	var resp JoinGroupResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.JoinGroup,
				RequestApiVersion: b.apiVersion.JoinGroup,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.JoinGroup)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.JoinGroup)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.JoinGroup)
		},
	)
	return resp, err
}

// Heartbeat (apiKey: 12)
func (b *Broker) Heartbeat(req HeartbeatRequest) (HeartbeatResponse, error) {
	var resp HeartbeatResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.Heartbeat,
				RequestApiVersion: b.apiVersion.Heartbeat,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.Heartbeat)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.Heartbeat)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.Heartbeat)
		},
	)
	return resp, err
}

// LeaveGroup (apiKey: 13)
func (b *Broker) LeaveGroup(req LeaveGroupRequest) (LeaveGroupResponse, error) {
	var resp LeaveGroupResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.LeaveGroup,
				RequestApiVersion: b.apiVersion.LeaveGroup,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.LeaveGroup)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.LeaveGroup)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.LeaveGroup)
		},
	)
	return resp, err
}

// SyncGroup (apiKey: 14)
func (b *Broker) SyncGroup(req SyncGroupRequest) (SyncGroupResponse, error) {
	var resp SyncGroupResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.SyncGroup,
				RequestApiVersion: b.apiVersion.SyncGroup,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.SyncGroup)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.SyncGroup)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.SyncGroup)
		},
	)
	return resp, err
}

// DescribeGroups (apiKey: 15)
func (b *Broker) DescribeGroups(req DescribeGroupsRequest) (DescribeGroupsResponse, error) {
	var resp DescribeGroupsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DescribeGroups,
				RequestApiVersion: b.apiVersion.DescribeGroups,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DescribeGroups)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DescribeGroups)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DescribeGroups)
		},
	)
	return resp, err
}

// ListGroups (apiKey: 16)
func (b *Broker) ListGroups(req ListGroupsRequest) (ListGroupsResponse, error) {
	var resp ListGroupsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ListGroups,
				RequestApiVersion: b.apiVersion.ListGroups,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ListGroups)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ListGroups)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ListGroups)
		},
	)
	return resp, err
}

// SaslHandshake (apiKey: 17)
func (b *Broker) SaslHandshake(req SaslHandshakeRequest) (SaslHandshakeResponse, error) {
	var resp SaslHandshakeResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.SaslHandshake,
				RequestApiVersion: b.apiVersion.SaslHandshake,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.SaslHandshake)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.SaslHandshake)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.SaslHandshake)
		},
	)
	return resp, err
}

// ApiVersions (apiKey: 18)
func (b *Broker) ApiVersions(req ApiVersionsRequest) (ApiVersionsResponse, error) {
	var resp ApiVersionsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ApiVersions,
				RequestApiVersion: b.apiVersion.ApiVersions,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ApiVersions)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ApiVersions)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ApiVersions)
		},
	)
	return resp, err
}

// CreateTopics (apiKey: 19)
func (b *Broker) CreateTopics(req CreateTopicsRequest) (CreateTopicsResponse, error) {
	var resp CreateTopicsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.CreateTopics,
				RequestApiVersion: b.apiVersion.CreateTopics,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.CreateTopics)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.CreateTopics)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.CreateTopics)
		},
	)
	return resp, err
}

// DeleteTopics (apiKey: 20)
func (b *Broker) DeleteTopics(req DeleteTopicsRequest) (DeleteTopicsResponse, error) {
	var resp DeleteTopicsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DeleteTopics,
				RequestApiVersion: b.apiVersion.DeleteTopics,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DeleteTopics)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DeleteTopics)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DeleteTopics)
		},
	)
	return resp, err
}

// DeleteRecords (apiKey: 21)
func (b *Broker) DeleteRecords(req DeleteRecordsRequest) (DeleteRecordsResponse, error) {
	var resp DeleteRecordsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DeleteRecords,
				RequestApiVersion: b.apiVersion.DeleteRecords,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DeleteRecords)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DeleteRecords)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DeleteRecords)
		},
	)
	return resp, err
}

// InitProducerId (apiKey: 22)
func (b *Broker) InitProducerId(req InitProducerIdRequest) (InitProducerIdResponse, error) {
	var resp InitProducerIdResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.InitProducerId,
				RequestApiVersion: b.apiVersion.InitProducerId,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.InitProducerId)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.InitProducerId)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.InitProducerId)
		},
	)
	return resp, err
}

// OffsetForLeaderEpoch (apiKey: 23)
func (b *Broker) OffsetForLeaderEpoch(req OffsetForLeaderEpochRequest) (OffsetForLeaderEpochResponse, error) {
	var resp OffsetForLeaderEpochResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.OffsetForLeaderEpoch,
				RequestApiVersion: b.apiVersion.OffsetForLeaderEpoch,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.OffsetForLeaderEpoch)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.OffsetForLeaderEpoch)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.OffsetForLeaderEpoch)
		},
	)
	return resp, err
}

// AddPartitionsToTxn (apiKey: 24)
func (b *Broker) AddPartitionsToTxn(req AddPartitionsToTxnRequest) (AddPartitionsToTxnResponse, error) {
	var resp AddPartitionsToTxnResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.AddPartitionsToTxn,
				RequestApiVersion: b.apiVersion.AddPartitionsToTxn,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.AddPartitionsToTxn)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.AddPartitionsToTxn)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.AddPartitionsToTxn)
		},
	)
	return resp, err
}

// AddOffsetsToTxn (apiKey: 25)
func (b *Broker) AddOffsetsToTxn(req AddOffsetsToTxnRequest) (AddOffsetsToTxnResponse, error) {
	var resp AddOffsetsToTxnResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.AddOffsetsToTxn,
				RequestApiVersion: b.apiVersion.AddOffsetsToTxn,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.AddOffsetsToTxn)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.AddOffsetsToTxn)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.AddOffsetsToTxn)
		},
	)
	return resp, err
}

// EndTxn (apiKey: 26)
func (b *Broker) EndTxn(req EndTxnRequest) (EndTxnResponse, error) {
	var resp EndTxnResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.EndTxn,
				RequestApiVersion: b.apiVersion.EndTxn,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.EndTxn)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.EndTxn)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.EndTxn)
		},
	)
	return resp, err
}

// WriteTxnMarkers (apiKey: 27)
func (b *Broker) WriteTxnMarkers(req WriteTxnMarkersRequest) (WriteTxnMarkersResponse, error) {
	var resp WriteTxnMarkersResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.WriteTxnMarkers,
				RequestApiVersion: b.apiVersion.WriteTxnMarkers,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.WriteTxnMarkers)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.WriteTxnMarkers)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.WriteTxnMarkers)
		},
	)
	return resp, err
}

// TxnOffsetCommit (apiKey: 28)
func (b *Broker) TxnOffsetCommit(req TxnOffsetCommitRequest) (TxnOffsetCommitResponse, error) {
	var resp TxnOffsetCommitResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.TxnOffsetCommit,
				RequestApiVersion: b.apiVersion.TxnOffsetCommit,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.TxnOffsetCommit)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.TxnOffsetCommit)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.TxnOffsetCommit)
		},
	)
	return resp, err
}

// DescribeAcls (apiKey: 29)
func (b *Broker) DescribeAcls(req DescribeAclsRequest) (DescribeAclsResponse, error) {
	var resp DescribeAclsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DescribeAcls,
				RequestApiVersion: b.apiVersion.DescribeAcls,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DescribeAcls)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DescribeAcls)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DescribeAcls)
		},
	)
	return resp, err
}

// CreateAcls (apiKey: 30)
func (b *Broker) CreateAcls(req CreateAclsRequest) (CreateAclsResponse, error) {
	var resp CreateAclsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.CreateAcls,
				RequestApiVersion: b.apiVersion.CreateAcls,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.CreateAcls)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.CreateAcls)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.CreateAcls)
		},
	)
	return resp, err
}

// DeleteAcls (apiKey: 31)
func (b *Broker) DeleteAcls(req DeleteAclsRequest) (DeleteAclsResponse, error) {
	var resp DeleteAclsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DeleteAcls,
				RequestApiVersion: b.apiVersion.DeleteAcls,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DeleteAcls)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DeleteAcls)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DeleteAcls)
		},
	)
	return resp, err
}

// DescribeConfigs (apiKey: 32)
func (b *Broker) DescribeConfigs(req DescribeConfigsRequest) (DescribeConfigsResponse, error) {
	var resp DescribeConfigsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DescribeConfigs,
				RequestApiVersion: b.apiVersion.DescribeConfigs,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DescribeConfigs)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DescribeConfigs)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DescribeConfigs)
		},
	)
	return resp, err
}

// AlterConfigs (apiKey: 33)
func (b *Broker) AlterConfigs(req AlterConfigsRequest) (AlterConfigsResponse, error) {
	var resp AlterConfigsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.AlterConfigs,
				RequestApiVersion: b.apiVersion.AlterConfigs,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.AlterConfigs)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.AlterConfigs)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.AlterConfigs)
		},
	)
	return resp, err
}

// AlterReplicaLogDirs (apiKey: 34)
func (b *Broker) AlterReplicaLogDirs(req AlterReplicaLogDirsRequest) (AlterReplicaLogDirsResponse, error) {
	var resp AlterReplicaLogDirsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.AlterReplicaLogDirs,
				RequestApiVersion: b.apiVersion.AlterReplicaLogDirs,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.AlterReplicaLogDirs)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.AlterReplicaLogDirs)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.AlterReplicaLogDirs)
		},
	)
	return resp, err
}

// DescribeLogDirs (apiKey: 35)
func (b *Broker) DescribeLogDirs(req DescribeLogDirsRequest) (DescribeLogDirsResponse, error) {
	var resp DescribeLogDirsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DescribeLogDirs,
				RequestApiVersion: b.apiVersion.DescribeLogDirs,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DescribeLogDirs)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DescribeLogDirs)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DescribeLogDirs)
		},
	)
	return resp, err
}

// SaslAuthenticate (apiKey: 36)
func (b *Broker) SaslAuthenticate(req SaslAuthenticateRequest) (SaslAuthenticateResponse, error) {
	var resp SaslAuthenticateResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.SaslAuthenticate,
				RequestApiVersion: b.apiVersion.SaslAuthenticate,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.SaslAuthenticate)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.SaslAuthenticate)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.SaslAuthenticate)
		},
	)
	return resp, err
}

// CreatePartitions (apiKey: 37)
func (b *Broker) CreatePartitions(req CreatePartitionsRequest) (CreatePartitionsResponse, error) {
	var resp CreatePartitionsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.CreatePartitions,
				RequestApiVersion: b.apiVersion.CreatePartitions,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.CreatePartitions)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.CreatePartitions)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.CreatePartitions)
		},
	)
	return resp, err
}

// CreateDelegationToken (apiKey: 38)
func (b *Broker) CreateDelegationToken(req CreateDelegationTokenRequest) (CreateDelegationTokenResponse, error) {
	var resp CreateDelegationTokenResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.CreateDelegationToken,
				RequestApiVersion: b.apiVersion.CreateDelegationToken,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.CreateDelegationToken)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.CreateDelegationToken)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.CreateDelegationToken)
		},
	)
	return resp, err
}

// RenewDelegationToken (apiKey: 39)
func (b *Broker) RenewDelegationToken(req RenewDelegationTokenRequest) (RenewDelegationTokenResponse, error) {
	var resp RenewDelegationTokenResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.RenewDelegationToken,
				RequestApiVersion: b.apiVersion.RenewDelegationToken,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.RenewDelegationToken)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.RenewDelegationToken)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.RenewDelegationToken)
		},
	)
	return resp, err
}

// ExpireDelegationToken (apiKey: 40)
func (b *Broker) ExpireDelegationToken(req ExpireDelegationTokenRequest) (ExpireDelegationTokenResponse, error) {
	var resp ExpireDelegationTokenResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ExpireDelegationToken,
				RequestApiVersion: b.apiVersion.ExpireDelegationToken,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ExpireDelegationToken)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ExpireDelegationToken)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ExpireDelegationToken)
		},
	)
	return resp, err
}

// DescribeDelegationToken (apiKey: 41)
func (b *Broker) DescribeDelegationToken(req DescribeDelegationTokenRequest) (DescribeDelegationTokenResponse, error) {
	var resp DescribeDelegationTokenResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DescribeDelegationToken,
				RequestApiVersion: b.apiVersion.DescribeDelegationToken,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DescribeDelegationToken)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DescribeDelegationToken)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DescribeDelegationToken)
		},
	)
	return resp, err
}

// DeleteGroups (apiKey: 42)
func (b *Broker) DeleteGroups(req DeleteGroupsRequest) (DeleteGroupsResponse, error) {
	var resp DeleteGroupsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.DeleteGroups,
				RequestApiVersion: b.apiVersion.DeleteGroups,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.DeleteGroups)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.DeleteGroups)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.DeleteGroups)
		},
	)
	return resp, err
}

// ElectLeaders (apiKey: 43)
func (b *Broker) ElectLeaders(req ElectLeadersRequest) (ElectLeadersResponse, error) {
	var resp ElectLeadersResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ElectLeaders,
				RequestApiVersion: b.apiVersion.ElectLeaders,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ElectLeaders)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ElectLeaders)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ElectLeaders)
		},
	)
	return resp, err
}

// IncrementalAlterConfigs (apiKey: 44)
func (b *Broker) IncrementalAlterConfigs(req IncrementalAlterConfigsRequest) (IncrementalAlterConfigsResponse, error) {
	var resp IncrementalAlterConfigsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.IncrementalAlterConfigs,
				RequestApiVersion: b.apiVersion.IncrementalAlterConfigs,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.IncrementalAlterConfigs)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.IncrementalAlterConfigs)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.IncrementalAlterConfigs)
		},
	)
	return resp, err
}

// AlterPartitionReassignments (apiKey: 45)
func (b *Broker) AlterPartitionReassignments(req AlterPartitionReassignmentsRequest) (AlterPartitionReassignmentsResponse, error) {
	var resp AlterPartitionReassignmentsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.AlterPartitionReassignments,
				RequestApiVersion: b.apiVersion.AlterPartitionReassignments,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.AlterPartitionReassignments)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.AlterPartitionReassignments)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.AlterPartitionReassignments)
		},
	)
	return resp, err
}

// ListPartitionReassignments (apiKey: 46)
func (b *Broker) ListPartitionReassignments(req ListPartitionReassignmentsRequest) (ListPartitionReassignmentsResponse, error) {
	var resp ListPartitionReassignmentsResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.ListPartitionReassignments,
				RequestApiVersion: b.apiVersion.ListPartitionReassignments,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.ListPartitionReassignments)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.ListPartitionReassignments)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.ListPartitionReassignments)
		},
	)
	return resp, err
}

// OffsetDelete (apiKey: 47)
func (b *Broker) OffsetDelete(req OffsetDeleteRequest) (OffsetDeleteResponse, error) {
	var resp OffsetDeleteResponse
	err := b.conn.Write(
		// encode request
		func(e *protocol.Encoder, correlationID int32) {
			hdr := RequestHeader{
				RequestApiKey:     apikey.OffsetDelete,
				RequestApiVersion: b.apiVersion.OffsetDelete,
				CorrelationId:     correlationID,
				ClientId:          b.config.clientID,
			}
			size := hdr.size(2) + req.size(b.apiVersion.OffsetDelete)
			e.PutInt32(size)
			hdr.Encode(e, 2)
			req.Encode(e, b.apiVersion.OffsetDelete)
		},
		// decode response
		func(d *protocol.Decoder) error {
			return (&resp).Decode(d, b.apiVersion.OffsetDelete)
		},
	)
	return resp, err
}

// Message defines the standard consumer and producer message
type Message struct {
	Topic string
	Key   []byte
	Value []byte
}

type MessageHandler func(*Message) error

type Subscription struct {
}

func (b *Broker) Subscribe(topics string, handler MessageHandler) (*Subscription, error) {
	return nil, nil
}
