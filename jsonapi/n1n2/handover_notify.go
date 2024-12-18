// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

// HandoverNotify is send by the target gNB to the CP after handover have been confirmed by the UE
type HandoverNotify struct {
	Command HandoverCommand `json:"handover-command"`
}
