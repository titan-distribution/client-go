// Copyright © 2019 Titan Distribution Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"net/http"
	"time"
)

// DefaultTransport is the optional transport clients use when a NewClient
// recieves a nil parameter.
var DefaultTransport = http.Transport{ResponseHeaderTimeout: time.Second}

// Client is an implementation of http.RoundTripper.
type Client struct {
	Transport http.RoundTripper
}

// NewClient returns a fully initialized Client.
func NewClient(t http.RoundTripper) *Client {
	var client *Client
	if t != nil {
		client.Transport = t
	} else {
		client.Transport = &DefaultTransport
	}
	return client
}

// RoundTrip is the Client implementation of http.RoundTripper. Used to cache
// responses from a registry.
func (c *Client) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return c.Transport.RoundTrip(req)
}
