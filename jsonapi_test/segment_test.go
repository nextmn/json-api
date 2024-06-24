// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

import (
	"encoding/json"
	"testing"

	"github.com/nextmn/json-api/jsonapi"
)

func TestSegment(t *testing.T) {
	s := &jsonapi.Segment{}
	if err := s.UnmarshalText([]byte("192.168.0.0/24")); err == nil {
		t.Errorf("Segment should be an IPv6 Address, not an IPv4 Prefix")
	}
	if err := s.UnmarshalText([]byte("192.168.0.1")); err == nil {
		t.Errorf("Segment should be an IPv6 Address, not an IPv4 Address")
	}
	if err := s.UnmarshalText([]byte("fd00::/80")); err == nil {
		t.Errorf("Segment should be an IPv6 Address, not an IPv6 Prefix")
	}
	if err := s.UnmarshalText([]byte("")); err == nil {
		t.Errorf("Empty Segment should be rejected")
	}
	if err := s.UnmarshalText([]byte("fd00::")); err != nil {
		t.Errorf("Correct Segment should not be rejected")
	}
}

func TestNewSegment(t *testing.T) {
	if seg0, err := jsonapi.NewSegment("::"); err == nil {
		if seg0J, err := json.Marshal(seg0); err == nil {
			if string(seg0J) != "\"::\"" {
				t.Errorf("Marshal of '::' is wrong: %s", seg0J)
			}
		} else {
			t.Errorf("Could not Marshal '::'")
		}
	} else {
		t.Errorf("Could not create segment '::'")
	}
	if seg1, err := jsonapi.NewSegment("fd00::"); err == nil {
		if seg1J, err := json.Marshal(seg1); err == nil {
			if string(seg1J) != "\"fd00::\"" {
				t.Errorf("Marshal of 'fd00::' is wrong: %s", seg1J)
			}
		} else {
			t.Errorf("Could not Marshal 'fd00::'")
		}
	} else {
		t.Errorf("Could not create segment 'fd00::'")
	}
	if _, err := jsonapi.NewSegment("10.0.0.0"); err == nil {
		t.Errorf("NewSegment should error with string '10.0.0.0'")
	}
	if _, err := jsonapi.NewSegment(""); err == nil {
		t.Errorf("NewSegment should error with empty string")
	}
	if _, err := jsonapi.NewSegment("fd00::/128"); err == nil {
		t.Errorf("NewSegment should error with IPv6 Prefix")
	}
	if _, err := jsonapi.NewSegment("10.0.0.0/32"); err == nil {
		t.Errorf("NewSegment should error with IPv4 Prefix")
	}
}
