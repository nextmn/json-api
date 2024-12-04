// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import "github.com/nextmn/json-api/jsonapi"

// PduSessionEstabReqMessage is sent by the UE to the gNB (then forwarded by the gNB to the CP)
// to start the PDU Session Establishment Procedure.
type PduSessionEstabReqMsg struct {
	Ue  jsonapi.ControlURI `json:"ue"`
	Gnb jsonapi.ControlURI `json:"gnb"`
	Dnn string             `json:"dnn"`
}
