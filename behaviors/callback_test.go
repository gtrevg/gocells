// Tideland Go Cells - Behaviors - Unit Tests - Callback
//
// Copyright (C) 2010-2016 Frank Mueller / Oldenburg / Germany
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

// TestCallbackBehavior tests the callback behavior.
func TestCallbackBehavior(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	env := cells.NewEnvironment("callback-behavior")
	defer env.Stop()

	cbdA := []string{}
	cbfA := func(topic string, payload cells.Payload) error {
		cbdA = append(cbdA, topic)
		return nil
	}
	cbdB := 0
	cbfB := func(topic string, payload cells.Payload) error {
		cbdB++
		return nil
	}
	sigc := audit.MakeSigChan()
	cbfC := func(topic string, payload cells.Payload) error {
		if topic == "baz" {
			sigc <- true
		}
		return nil
	}

	env.StartCell("callback", behaviors.NewCallbackBehavior(cbfA, cbfB, cbfC))

	env.EmitNew("callback", "foo", nil)
	env.EmitNew("callback", "bar", nil)
	env.EmitNew("callback", "baz", nil)

	assert.Wait(sigc, true, time.Second)
	assert.Equal(cbdA, []string{"foo", "bar", "baz"})
	assert.Equal(cbdB, 3)
}

// EOF
