// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverNotify is send by the target gNB to the CP after handover have been confirmed by the UE
type HandoverNotify struct {
	// Header
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Cp        jsonapi.ControlURI `json:"cp"`
	TargetGnb jsonapi.ControlURI `json:"target-gnb"`

	// Handover Notify
	Sessions  []Session          `json:"sessions"`
	SourceGnb jsonapi.ControlURI `json:"source-gnb"`
}
