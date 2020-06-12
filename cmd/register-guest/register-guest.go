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
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/matrix-org/gomatrix"
	"github.com/matrix-org/matrix-static/mxclient"
	"io/ioutil"
)

func registerGuest(configPath, homeserverURL, mediaBaseURL string) error {
	m, err := mxclient.NewRawClient(homeserverURL, "", "", "")
	if err != nil {
		return err
	}

	register, inter, err := m.RegisterGuest(&gomatrix.ReqRegister{})

	if err != nil {
		return err
	}
	if inter != nil || register == nil {
		return errors.New("error encountered during guest registration")
	}

	config := mxclient.Config{
		AccessToken:  register.AccessToken,
		DeviceID:     register.DeviceID,
		HomeServer:   homeserverURL,
		RefreshToken: register.RefreshToken,
		UserID:       register.UserID,
		MediaBaseUrl: mediaBaseURL,
	}

	// TODO consider SRV Query on start instead.
	// SRV is primarily for S-S API so not 100% appropriate.
	register.HomeServer = homeserverURL

	configJson, err := json.Marshal(config)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, configJson, 0600)
}

func main() {
	configPath := flag.String("config-file", "./config.json", "The path to the desired config file.")
	homeserverURL := flag.String("homeserver-url", "https://matrix.org", "What Homeserver URL to use when registering a guest.")
	mediaBaseURL := flag.String("media-base-url", "https://matrix.org", "What Homeserver URL to use for Media Repository requests.")
	flag.Parse()

	if *mediaBaseURL == "" || mediaBaseURL == nil { // if media-base-url not provided, default to homeserver-url
		mediaBaseURL = homeserverURL
	}

	if err := registerGuest(*configPath, *homeserverURL, *mediaBaseURL); err != nil {
		fmt.Println("Error encountered when creating guest account: ", err)
	} else {
		fmt.Println("Guest account created successfully!!")
	}
}
