// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverRequestAck is send by the target gNB to the CP in response to an HandoverRequest
type HandoverRequestAck struct {
	// Header
	Cp        jsonapi.ControlURI `json:"cp"`
	TargetgNB jsonapi.ControlURI `json:"target-gnb"`

	// Handover Request Ack
	SourcegNB jsonapi.ControlURI `json:"source-gnb"`
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Sessions  []Session          `json:"sessions"` // contains new DL FTeid
}
