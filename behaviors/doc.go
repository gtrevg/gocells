// Tideland Go Cells - Behaviors
//
// Copyright (C) 2010-2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// Package behaviors provides several generic and always useful
// standard behaviors for the Tideland Go Cells. They are simply
// created by calling NewXyzBehavior(). Their configuration
// is done by constructor arguments. Additionally some of them take
// functions or implementations of interfaces to control their
// processing. These behaviors are:
//
// Broadcaster
//
// The broadcaster behavior simply emits all received events to all
// of its subscribers. It is intended to be used as a top level behavior
// to directly rigger multiple handlers instead of emitting an event
// manually to those handlers.
//
// Callback
//
// The callback behavior allows you to provide a number of functions
// which will be called when an event is received. Those functions
// have the topic and the payload of the event as argument.
//
// Collector
//
// The collector behavior collects all received events. They can be
// retrieved and resetted. It also emits all received events to its
// subscribers.
//
// Configurator
//
// After receiving a ReadConfigurationTopic with a filename as
// payload the configuration behavior reads this configuration
// and emits it. If it is started with a validator the configuration
// is validated after the reading.
//
// Counter
//
// The counter behavior is created with a counter function as argument.
// This function is called for each event and returns the IDs of counters
// which are incremented then. The counters are emitted each time and
// also can be resetted.
//
// Filter
//
// The filter behavior is created with a filtering function which is
// called for each event. If this function call returns true the event
// emitted, otherwise it is dropped.
//
// Finite State Machine
//
// The FSM behavior implements a finite state machine. State functions
// process the events and return the following state function.
//
// Logger
//
// The logger behavior logs every event. The used level is INFO.
//
// Mapper
//
// The mapper behavior is created with a mapping. It is called with each
// received event and returns a new mapped one.
//
// Round Robin
//
// The round robin behavior distributes each received event round robin
// to its subscribers. It can be used for load balancing.
//
// Scene
//
// The scene behavior stores a received payload using the event topic as
// key in the event scene. So it can be used later by other behaviors
// or by the external environments, which can wait until the setting.
//
// Simple Processor
//
// The simple behavior is created with a simple event processing function.
// Useful if no state and no complex recovery is needed.
//
// Ticker
//
// The ticker behavior emits a tick event in a defined interval to its
// subscribers. So they can process chronological tasks beside other
// events.
package behaviors

//--------------------
// IMPORTS
//--------------------

import (
	"github.com/tideland/golib/version"
)

//--------------------
// VERSION
//--------------------

// Version returns the version of the behaviors package.
func Version() version.Version {
	return version.New(6, 0, 0, "beta", "2016-10-22")
}

// EOF
