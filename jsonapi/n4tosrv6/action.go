// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import "net/netip"

type Action struct {
	SRH        SRH         `json:"srh"`
	SourceGtp4 *netip.Addr `json:"src-gtp4,omitempty"`
}
