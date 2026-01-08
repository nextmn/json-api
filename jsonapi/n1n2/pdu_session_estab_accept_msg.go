// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import "net/netip"

// PduSessionEstabAcceptMsg is sent to the UE by the gNB
// when the PDU Session establishment is accepted.
type PduSessionEstabAcceptMsg struct {
	Header PduSessionEstabReqMsg `json:"header"`  // copy of the PDU Session Establishment Request Message
	Addr   netip.Addr            `json:"address"` // IP Address attributed to the UE for this PDU Session
}
