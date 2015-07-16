// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package touch defines an event for touch input.
//
// See the golang.org/x/mobile/event package for details on the event model.
package touch // import "golang.org/x/mobile/event/touch"

// The best source on android input events is the NDK: include/android/input.h
//
// iOS event handling guide:
// https://developer.apple.com/library/ios/documentation/EventHandling/Conceptual/EventHandlingiPhoneOS

import (
	"fmt"

	"golang.org/x/mobile/geom"
)

// Event is a touch event.
//
// The same Sequence is shared by all events in a sequence. A sequence starts
// with a single TypeStart, is followed by zero or more TypeMoves, and ends
// with a single TypeEnd. A Sequence distinguishes concurrent sequences but its
// value is subsequently reused.
type Event struct {
	Sequence Sequence
	Type     Type
	Loc      geom.Point
}

// Sequence identifies a sequence of touch events.
type Sequence int64

// Type describes the type of a touch event.
type Type byte

const (
	// TypeStart is a user first touching the device.
	//
	// On Android, this is a AMOTION_EVENT_ACTION_DOWN.
	// On iOS, this is a call to touchesBegan.
	TypeStart Type = iota

	// TypeMove is a user dragging across the device.
	//
	// A TypeMove is delivered between a TypeStart and TypeEnd.
	//
	// On Android, this is a AMOTION_EVENT_ACTION_MOVE.
	// On iOS, this is a call to touchesMoved.
	TypeMove

	// TypeEnd is a user no longer touching the device.
	//
	// On Android, this is a AMOTION_EVENT_ACTION_UP.
	// On iOS, this is a call to touchesEnded.
	TypeEnd
)

func (t Type) String() string {
	switch t {
	case TypeStart:
		return "start"
	case TypeMove:
		return "move"
	case TypeEnd:
		return "end"
	}
	return fmt.Sprintf("touch.Type(%d)", t)
}