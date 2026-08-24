// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi_test

import (
	"net/url"
	"testing"

	"github.com/nextmn/json-api/jsonapi"
)

func TestControlURI(t *testing.T) {
	u := &jsonapi.ControlURI{}
	if err := u.UnmarshalText([]byte("http://example.org/")); err == nil {
		t.Errorf("URI with trailing slash should be rejected")
	}
	if err := u.UnmarshalText([]byte("example.org")); err == nil {
		t.Errorf("URI without scheme should be rejected")
	}
	if err := u.UnmarshalText([]byte("http://example.org")); err != nil {
		t.Errorf("URI with a domain should be accepted")
	}
	if err := u.UnmarshalText([]byte("http://example.org:8000")); err != nil {
		t.Errorf("URI with a domain and a port should be accepted")
	}
	if err := u.UnmarshalText([]byte("http://10.0.0.1:8000")); err != nil {
		t.Errorf("URI with an IPv4 address and a port should be accepted")
	}
	if err := u.UnmarshalText([]byte("http://[fd00::1]:8000")); err != nil {
		t.Errorf("URI with an IPv6 address and a port should be accepted")
	}

	u.UnmarshalText([]byte("http://example.org"))
	cmp, _ := url.ParseRequestURI("http://example.org")
	if u.URL != *cmp {
		t.Errorf("Valid ControlURI should be unmarshaled the same as ParseRequestURI does")
	}
}
