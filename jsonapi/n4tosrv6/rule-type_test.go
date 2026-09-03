// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6_test

import (
	"encoding/json/v2"
	"testing"

	"github.com/nextmn/json-api/jsonapi/n4tosrv6"
)

func TestRule(t *testing.T) {
	ruleType := n4tosrv6.RuleTypeUplink

	if ruleTypeJson, err := json.Marshal(ruleType); err == nil {
		if string(ruleTypeJson) != "\"uplink\"" {
			t.Errorf("Marshal of uplink type is wrong: %s", ruleTypeJson)
		}
	} else {
		t.Errorf("Could not Marshal uplink type")
	}

	var out n4tosrv6.RuleType
	if err := json.Unmarshal([]byte("\"downlink\""), &out); err == nil {
		if out != n4tosrv6.RuleTypeDownlink {
			t.Errorf("Unmarshal of downlink type is wrong: %s", out)
		}
	} else {
		t.Errorf("Could not Unmarshal downlink type: %s", err)
	}
}
