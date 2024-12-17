// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverRequired is send by the source gNB to the CP to start the handover preparation phase
type HandoverRequired struct {
	// Header
	SourcegNB jsonapi.ControlURI `json:"source-gnb"`
	Cp        jsonapi.ControlURI `json:"cp"`

	// Handover Required content
	Ue        jsonapi.ControlURI `json:"ue"`
	Sessions  []Session          `json:"sessions"` // list of all pdu sessions of the UE to be moved
	TargetgNB jsonapi.ControlURI `json:"target-gnb"`
}
