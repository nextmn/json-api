// Copyright 2024 Louis Royer and the NextMN-SRv6-ctrl contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi_test

import (
	"testing"

	"github.com/nextmn/json-api/jsonapi"
)

func TestBackboneIP(t *testing.T) {
	b := &jsonapi.BackboneIP{}
	if err := b.UnmarshalText([]byte("192.168.0.0/24")); err == nil {
		t.Errorf("BackboneIP should be an IPv6 Address, not an IPv4 Prefix")
	}
	if err := b.UnmarshalText([]byte("192.168.0.1")); err == nil {
		t.Errorf("Locator should be an IPv6 Address, not an IPv4 Address")
	}
	if err := b.UnmarshalText([]byte("fd00::/80")); err == nil {
		t.Errorf("BackboneIP should be an IPv6 Address, not an IPv6 Prefix")
	}
	if err := b.UnmarshalText([]byte("")); err == nil {
		t.Errorf("Empty BackboneIP should be rejected")
	}
	if err := b.UnmarshalText([]byte("fd00::")); err != nil {
		t.Errorf("Correct IPv6 Address should not be rejected")
	}
}
