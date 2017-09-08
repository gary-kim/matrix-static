// Copyright 2017 Michael Telatynski <7t3chguy@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/base64"
	"encoding/json"
)

const CookieName = "matrix-static"

type UserSettings struct {
	PageSize *int
}

func (us UserSettings) LoadUserSettings(cookie string, cookieErr error) (UserSettings, error) {
	if cookieErr != nil {
		return us, cookieErr
	}

	data, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return us, err
	}

	err = json.Unmarshal(data, &us)
	return us, err
}

func (us *UserSettings) Serialize() (string, error) {
	data, err := json.Marshal(*us)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
