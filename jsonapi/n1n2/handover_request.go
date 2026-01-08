// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverRequest is send by the CP to the target gNB during the handover preparation phase
type HandoverRequest struct {
	// Header
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Cp        jsonapi.ControlURI `json:"cp"`
	TargetgNB jsonapi.ControlURI `json:"target-gnb"`

	// Handover Request
	SourcegNB jsonapi.ControlURI `json:"source-gnb"`
	Sessions  []Session          `json:"sessions"` // contains new UL FTeid
}
