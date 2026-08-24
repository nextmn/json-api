// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6_test

import (
	"testing"

	"github.com/nextmn/json-api/jsonapi/n4tosrv6"
)

func TestLocator(t *testing.T) {
	l := &n4tosrv6.Locator{}
	if err := l.UnmarshalText([]byte("192.168.0.0/24")); err == nil {
		t.Errorf("Locator should be an IPv6 Prefix, not an IPv4 Prefix")
	}
	if err := l.UnmarshalText([]byte("192.168.0.1")); err == nil {
		t.Errorf("Locator should be an IPv6 Prefix, not an IPv4 Address")
	}
	if err := l.UnmarshalText([]byte("fd00::")); err == nil {
		t.Errorf("Locator should be an IPv6 Prefix, not an IPv6 Address")
	}
	if err := l.UnmarshalText([]byte("")); err == nil {
		t.Errorf("Empty Locator should be rejected")
	}
	if err := l.UnmarshalText([]byte("fd00::/80")); err != nil {
		t.Errorf("Correct Locator should not be rejected")
	}
}
