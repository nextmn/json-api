// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// N2PduSessionRespMsg is sent to the CP by the gNB as a response of N2PduSessionReqMsg.
type N2PduSessionRespMsg struct {
	UeInfo PduSessionEstabAcceptMsg `json:"ue-info"` // used to identify the PDU Session

	// Downlink FTEID: the CP will use this to configure a downlink GTP Tunnel in Upf-i
	DownlinkFteid jsonapi.Fteid `json:"downlink-fteid"`
}
