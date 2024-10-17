// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

type Match struct {
	Header  *GtpHeader `json:"gtp,omitempty"`
	Payload *Payload   `json:"payload,omitempty"` // if empty, the rule is considered "default"
}
