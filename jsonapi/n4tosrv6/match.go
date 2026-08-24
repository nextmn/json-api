// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

// Match represents a set of rules to be matched.
// When a Match member is zero, we consider it always match.
// For example, if there is no GTP encapsulation, we consider only the "IP" layer field.
type Match struct {
	Ip      IpLayer  `json:"ip,omitzero"`      // Outer IP layer
	Udp     UdpLayer `json:"udp,omitzero"`     // UDP layer
	Gtp     GtpLayer `json:"gtp,omitzero"`     // GTP layer
	Payload IpLayer  `json:"payload,omitzero"` // Inner IP layer
}
