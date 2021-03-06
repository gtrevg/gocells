// Tideland Go Cells - Behaviors - Simple Processor
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
	"github.com/tideland/golib/logger"

	"github.com/tideland/gocells/cells"
)

//--------------------
// SIMPLE BEHAVIOR
//--------------------

// SimpleProcessor is a function type doing the event processing.
type SimpleProcessor func(cell cells.Cell, event cells.Event) error

// simpleBehavior is a simple event processor using the processor
// function for its own logic.
type simpleBehavior struct {
	cell    cells.Cell
	process SimpleProcessor
}

// NewSimpleProcessorBehavior creates a behavior based on the passed function.
// Instead of an own logic and an own state it uses the passed simple processor
// function for the event processing.
func NewSimpleProcessorBehavior(processor SimpleProcessor) cells.Behavior {
	if processor == nil {
		processor = func(cell cells.Cell, event cells.Event) error {
			logger.Errorf("simple processor %q used without function to handle event %v", cell.ID(), event)
			return nil
		}
	}
	return &simpleBehavior{
		process: processor,
	}
}

// Init the behavior.
func (b *simpleBehavior) Init(c cells.Cell) error {
	b.cell = c
	return nil
}

// Terminate the behavior.
func (b *simpleBehavior) Terminate() error {
	return nil
}

// ProcessEvent calls the simple processor function.
func (b *simpleBehavior) ProcessEvent(event cells.Event) error {
	return b.process(b.cell, event)
}

// Recover from an error.
func (b *simpleBehavior) Recover(err interface{}) error {
	return nil
}

// EOF
