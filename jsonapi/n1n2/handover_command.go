// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverCommand is sent by the CP to the source gNB to start the execution of handover, and forwarded to the UE
type HandoverCommand struct {
	// Header
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Cp        jsonapi.ControlURI `json:"cp"`
	SourceGnb jsonapi.ControlURI `json:"source-gnb"`

	// Handover Command
	Sessions  []Session          `json:"sessions"` // contains new ForwardDownlinkFteid
	TargetGnb jsonapi.ControlURI `json:"target-gnb"`
}
