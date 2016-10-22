// Tideland Go Cells - Behaviors - Unit Tests - Broadcaster
//
// Copyright (C) 2010-2016 Frank Mueller / Tideland / Oldenburg / Germany
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
	env := cells.NewEnvironment("broadcaster-behavior")
	defer env.Stop()

	env.StartCell("broadcast", behaviors.NewBroadcasterBehavior())
	env.StartCell("test-1", behaviors.NewCollectorBehavior(10))
	env.StartCell("test-2", behaviors.NewCollectorBehavior(10))
	env.Subscribe("broadcast", "test-1", "test-2")

	env.EmitNew("broadcast", "test", "a")
	env.EmitNew("broadcast", "test", "b")
	env.EmitNew("broadcast", "test", "c")

	time.Sleep(100 * time.Millisecond)

	collected, err := env.Request("test-1", cells.CollectedTopic, nil, cells.DefaultTimeout, nil)
	assert.Nil(err)
	assert.Length(collected, 3)

	collected, err = env.Request("test-2", cells.CollectedTopic, nil, cells.DefaultTimeout, nil)
	assert.Nil(err)
	assert.Length(collected, 3)
}

// EOF
