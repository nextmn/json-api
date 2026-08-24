// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

type GtpHeader struct {
	Ip  IpLayer  `json:"ip,omitzero"`  // Outer IP layer
	Udp UdpLayer `json:"udp,omitzero"` // UDP layer
	Gtp GtpLayer `json:"gtp,omitzero"` // GTP layer
}
