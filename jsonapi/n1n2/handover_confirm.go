// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"github.com/nextmn/json-api/jsonapi"
)

// HandoverConfirm is send by the target UE to the target gNB after the UE is synchronized to the new cell
type HandoverConfirm struct {
	// Header
	UeCtrl jsonapi.ControlURI `json:"ue-ctrl"`
	Cp     jsonapi.ControlURI `json:"cp"`

	// Handover Confirm
	Sessions  []Session          `json:"sessions"`
	SourceGnb jsonapi.ControlURI `json:"source-gnb"`
	TargetGnb jsonapi.ControlURI `json:"target-gnb"`
}
