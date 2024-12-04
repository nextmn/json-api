// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import (
	"encoding/json"
	"fmt"
)

type MessageWithError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func (u *MessageWithError) UnmarshalJSON(data []byte) error {
	a := make(map[string]string)
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	msg, ok := a["message"]
	if !ok {
		return fmt.Errorf("Missing key `message` while unmarshaling MessageWithError")
	}
	u.Message = msg
	e, ok := a["error"]
	if !ok {
		return fmt.Errorf("Missing key `error` while unmarshaling MessageWithError")
	}
	u.Error = fmt.Errorf(e)
	return nil
}

func (u MessageWithError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"message": u.Message,
		"error":   u.Error.Error(),
	})
}
