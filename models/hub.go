package models

import "log"

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	rooms  map[string]map[*connection]bool
	starts map[string]bool
	keys   map[string]string

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription

	startchan chan subscription
}

var H = hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
	starts:     make(map[string]bool),
	keys:       make(map[string]string),
	startchan:  make(chan subscription),
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.startchan:

			if h.keys[s.room] == s.conn.ip+s.conn.username {

				// start
				h.starts[s.room] = true

				// ok
				for connect := range h.rooms[s.room] {
					h.rooms[s.room][connect] = true
				}
			}

		case s := <-h.register:
			connections := h.rooms[s.room]

			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
				h.starts[s.room] = true
				// h.starts[s.room] = false
				h.keys[s.room] = s.conn.ip + s.conn.username
			}

			// max room = 8
			if len(connections) >= 4 {
				log.Printf("room = %s full \n", s.room)
				close(s.conn.send)
			} else {
				h.rooms[s.room][s.conn] = true
				// h.rooms[s.room][s.conn] = false
			}

		case s := <-h.unregister:

			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}

		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c, ok := range connections {
				if ok {
					select {
					case c.send <- m.data:
					default:
						close(c.send)
						delete(connections, c)
						if len(connections) == 0 {
							delete(h.rooms, m.room)
						}
					}
				}
			}

		}
	}
}
