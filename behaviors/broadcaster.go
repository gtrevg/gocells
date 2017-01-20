// Tideland Go Cells - Behaviors - Broadcaster
//
// Copyright (C) 2010-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package behaviors

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/gocells/cells"
)

//--------------------
// BROADCASTER BEHAVIOR
//--------------------

// broadcasterBehavior is a simple repeater.
type broadcasterBehavior struct {
	cell cells.Cell
}

// NewBroadcasterBehavior creates a broadcasting behavior that just emits every
// received event. It's intended to work as an entry point for events, which
// shall be immediately processed by several subscribers.
func NewBroadcasterBehavior() cells.Behavior {
	return &broadcasterBehavior{}
}

// Init the behavior.
func (b *broadcasterBehavior) Init(c cells.Cell) error {
	b.cell = c
	return nil
}

// Terminate the behavior.
func (b *broadcasterBehavior) Terminate() error {
	return nil
}

// ProcessEvent emits the event to all subscribers.
func (b *broadcasterBehavior) ProcessEvent(event cells.Event) error {
	b.cell.Emit(event)
	return nil
}

// Recover from an error.
func (b *broadcasterBehavior) Recover(err interface{}) error {
	return nil
}

// EOF
