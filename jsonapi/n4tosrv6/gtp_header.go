// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import "net/netip"

type GtpHeader struct {
	OuterIpSrc netip.Addr  `json:"outer-ip-src"` // i.e. gNB ip
	Teid       uint32      `json:"teid"`
	InnerIpSrc *netip.Addr `json:"inner-ip-src,omitempty"` // i.e. UE ip; only useful for uplink when multiple PDU Sessions are aggregated on a single TEID
}
