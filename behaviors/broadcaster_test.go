// Tideland Go Cells - Behaviors - Unit Tests - Broadcaster
//
// Copyright (C) 2010-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package behaviors_test

//--------------------
// IMPORTS
//--------------------

import (
	"testing"
	"time"

	"github.com/tideland/golib/audit"

	"github.com/tideland/gocells/behaviors"
	"github.com/tideland/gocells/cells"
)

//--------------------
// TESTS
//--------------------

// TestBroadcasterBehavior tests the broadcast behavior.
func TestBroadcasterBehavior(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	sigc := audit.MakeSigChan()
	env := cells.NewEnvironment("broadcaster-behavior")
	defer env.Stop()

	mktester := func() behaviors.ConditionTester {
		counter := 0
		return func(event cells.Event) bool {
			counter++
			return counter == 3
		}
	}
	processor := func(cell cells.Cell, event cells.Event) error {
		sigc <- true
		return nil
	}

	env.StartCell("broadcast", behaviors.NewBroadcasterBehavior())
	env.StartCell("test-a", behaviors.NewConditionBehavior(mktester(), processor))
	env.StartCell("test-b", behaviors.NewConditionBehavior(mktester(), processor))
	env.Subscribe("broadcast", "test-a", "test-b")

	env.EmitNew("broadcast", "test", nil)
	env.EmitNew("broadcast", "test", nil)
	env.EmitNew("broadcast", "test", nil)

	assert.Wait(sigc, true, time.Second)
	assert.Wait(sigc, true, time.Second)
}

// EOF
