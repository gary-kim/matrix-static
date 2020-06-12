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

package mxclient

import "github.com/matrix-org/gomatrix"

// Keeping here in case it becomes used again.
//func ConcatEventsSlices(slices ...[]gomatrix.Event) []gomatrix.Event {
//	var totalLen int
//	for _, s := range slices {
//		totalLen += len(s)
//	}
//	tmp := make([]gomatrix.Event, totalLen)
//	var i int
//	for _, s := range slices {
//		i += copy(tmp[i:], s)
//	}
//	return tmp
//}

// ReverseEventsCopy returns a copy of the input slice with all elements in reverse order.
func ReverseEventsCopy(events []gomatrix.Event) []gomatrix.Event {
	var newEvents []gomatrix.Event
	for i := len(events) - 1; i >= 0; i-- {
		newEvents = append(newEvents, events[i])
	}
	return newEvents
}

// UnwrapRespError takes an error and if it is a HTTPError returns the WrappedError.RespError it contains
func UnwrapRespError(err error) (respErr gomatrix.RespError, respErrOk bool) {
	if err, ok := err.(gomatrix.HTTPError); ok {
		respErr, respErrOk = err.WrappedError.(gomatrix.RespError)
	}
	return
}

var textForRespError = map[string]string{
	"M_GUEST_ACCESS_FORBIDDEN": "This Room does not exist or does not permit guests to access it.",
}

// TextForRespError returns a string representation of the RespError if known, defaulting to the Err field of the RespError.
func TextForRespError(respErr gomatrix.RespError) string {
	if msg, ok := textForRespError[respErr.ErrCode]; ok {
		return msg
	}
	return respErr.Err + " (" + respErr.ErrCode + ")"
}

// ShouldHideEvent returns a bool the event should be ignored in the timeline view, mimicking riot-web
func ShouldHideEvent(ev gomatrix.Event) bool {
	// m.room.create ?

	// we want to hide all unknowns +:
	// m.room.redaction
	// m.room.aliases
	// m.room.canonical_alias

	if ev.StateKey == nil {
		// Message Event
		if ev.Type == "m.room.message" {
			return false
		}
	} else {
		// State Event
		if ev.Type == "m.room.history_visibility" ||
			ev.Type == "m.room.join_rules" ||
			ev.Type == "m.room.member" ||
			ev.Type == "m.room.power_levels" ||
			ev.Type == "m.room.name" ||
			ev.Type == "m.room.topic" ||
			ev.Type == "m.room.avatar" {
			return false
		}

		if ev.Type == "im.vector.modular.widgets" {
			return false
		}
	}

	return true
}
