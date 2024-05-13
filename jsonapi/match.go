// Copyright 2024 Louis Royer and the NextMN-SRv6-ctrl contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

import "net/netip"

type Match struct {
	SrcIpPrefix netip.Prefix
	DstIpPrefix netip.Prefix
	GtpTeid     uint32
}
