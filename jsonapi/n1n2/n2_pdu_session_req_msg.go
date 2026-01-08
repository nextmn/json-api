// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// N2PduSessionReqMsg is sent by the CP to the gNB to establish a PDU Session.
type N2PduSessionReqMsg struct {
	Cp     jsonapi.ControlURI       `json:"cp"`
	UeInfo PduSessionEstabAcceptMsg `json:"ue-info"` // information to forward to the UE

	// Uplink FTEID: the gNB will establish an Uplink GTP Tunnel using the following
	UplinkFteid jsonapi.Fteid `json:"uplink-fteid"`
}
