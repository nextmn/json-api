// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

// HandoverConfirm is send by the target UE to the target gNB after the UE is synchronized to the new cell
type HandoverConfirm struct {
	Command HandoverCommand `json:"handover-command"`
}