// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2_test

import (
	"encoding/json"
	"testing"

	"github.com/nextmn/json-api/jsonapi/n1n2"
)

func TestSession(t *testing.T) {
	s := &n1n2.Session{}
	if err := json.Unmarshal([]byte("{\"ue-addr\": \"127.0.0.1\", \"uplink-fteid\": {\"addr\": \"127.0.0.2\", \"teid\": 80}}"), s); err != nil {
		t.Errorf("Session with only uplink FTeid could not be unmarshaled")
	}

	if s.DownlinkFteid != nil {
		t.Errorf("Downlink Fteid was not defined but is not nil")
	}
	if s.UplinkFteid == nil {
		t.Errorf("Uplink Fteid was defined but is nil")
	}

}
