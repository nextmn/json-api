// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6_test

import (
	"net/netip"
	"testing"

	"github.com/nextmn/json-api/jsonapi/n4tosrv6"
)

func TestNewSRH(t *testing.T) {
	if srh, err := n4tosrv6.NewSRH([]string{"::", "fd00::"}); err == nil {
		if len(*srh) != 2 {
			t.Errorf("Length of SRH should me 2, but is %d", len(*srh))
		}
		if (*srh)[0].Compare(netip.MustParseAddr("::")) != 0 {
			t.Errorf("First segment is not correctly created")
		}
		if (*srh)[1].Compare(netip.MustParseAddr("fd00::")) != 0 {
			t.Errorf("Second segment is not correctly created")
		}
	} else {
		t.Errorf("Could not create SRH with '::', ''fd00::'")
	}
}
