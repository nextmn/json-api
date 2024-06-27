// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

type Rule struct {
	Enabled bool
	Type    string // uplink or downlink
	Match   Match
	Action  Action
}
