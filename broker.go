package dashing

import "github.com/znly/go-dashing/dashingtypes"

//Inspired by https://github.com/gigablah/dashing-go

// An eventCache stores the latest event for each key, so that new clients can
// catch up.
type eventCache map[string]*dashingtypes.Event

type broker struct {
	// Create a map of clients, the keys of the map are the channels
	// over which we can push messages to attached clients. (The values
	// are just booleans and are meaningless)
	clients map[chan *dashingtypes.Event]bool

	// Channel into which new clients can be pushed
	newClients chan chan *dashingtypes.Event

	// Channel into which disconnected clients should be pushed
	defunctClients chan chan *dashingtypes.Event

	// Channel into which events are pushed to be broadcast out
	// to attached clients
	events chan *dashingtypes.Event

	// Cache for most recent events with a certain ID
	cache eventCache
}

func (b *broker) start() {
	go func() {
		for {
			// Block until we receive from one of the
			// three following channels.
			select {
			case s := <-b.newClients:
				// There is a new client attached and we
				// want to start sending them events.
				b.clients[s] = true
				// Send all the cached events so that when a new client connects, it
				// doesn't miss previous events
				for _, e := range b.cache {
					s <- e
				}
				// log.Println("Added new client")
			case s := <-b.defunctClients:
				// A client has detached and we want to
				// stop sending them events.
				delete(b.clients, s)
				// log.Println("Removed client")
			case event := <-b.events:
				if event.Target != "dashboards" {
					b.cache[event.ID] = event
				}
				// There is a new event to send. For each
				// attached client, push the new event
				// into the client's channel.
				for s := range b.clients {
					s <- event
				}
				// log.Printf("Broadcast event to %d clients", len(b.clients))
			}
		}
	}()
}

func newBroker() *broker {
	return &broker{
		make(map[chan *dashingtypes.Event]bool),
		make(chan (chan *dashingtypes.Event)),
		make(chan (chan *dashingtypes.Event)),
		make(chan *dashingtypes.Event),
		map[string]*dashingtypes.Event{},
	}
}
