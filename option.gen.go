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
	"crypto/tls"
	"net"
)

const defaultClientID = "github.com/savaki/kafka"

type config struct {
	clientID  string
	dialer    *net.Dialer
	tlsConfig *tls.Config
}

// Option provides connection options
type Option func(*config)

// WithDialer allows dialer properties to be configured
func WithDialer(dialer *net.Dialer) Option {
	return func(c *config) {
		if dialer == nil {
			c.dialer = &net.Dialer{}
			return
		}
		c.dialer = dialer
	}
}

// WithTLS provides TLS configuration
func WithTLS(tlsConfig *tls.Config) Option {
	return func(c *config) {
		c.tlsConfig = tlsConfig
	}
}

func buildConfig(opts []Option) config {
	c := config{
		clientID: defaultClientID,
		dialer:   &net.Dialer{},
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}