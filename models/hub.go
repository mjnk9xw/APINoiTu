package models

import (
	"encoding/json"
	"log"
)

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
	rooms     map[string]map[*connection]bool
	starts    map[string]bool
	keys      map[string]string
	useralive map[string][]string
	messPre   map[string]string

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
	messPre:    make(map[string]string),
	useralive:  make(map[string][]string),
	startchan:  make(chan subscription),
}

func (h *hub) deleteUserAlive(room, name string, conDel *connection) {
	arr := make([]string, 0)
	for _, v := range h.useralive[room] {
		if v != name {
			arr = append(arr, v)
		}
	}
	h.useralive[room] = arr
	h.rooms[room][conDel] = false
}

func (h *hub) searchNextAlive(room, name string) (isWin bool, nameNext string) {
	if len(h.useralive[room]) == 1 {
		return true, name
	}
	checkIndx := -1
	mapIndex := make(map[int]string)
	for i, v := range h.useralive[room] {
		mapIndex[i] = v
		if v == name {
			checkIndx = i
		}
	}
	checkIndx++
	if checkIndx >= len(h.useralive[room]) {
		checkIndx = 0
	}
	return isWin, mapIndex[checkIndx]
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.startchan:

			//start_game_room_wss
			if h.keys[s.room] == s.conn.ip+s.conn.username {

				// start
				h.starts[s.room] = true

				// ok
				h.useralive[s.room] = make([]string, 0)
				for connect := range h.rooms[s.room] {
					h.rooms[s.room][connect] = true
					h.useralive[s.room] = append(h.useralive[s.room], connect.username)
				}
			}

		case s := <-h.register:
			connections := h.rooms[s.room]

			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
				// h.starts[s.room] = true
				h.starts[s.room] = false
				h.keys[s.room] = s.conn.ip + s.conn.username
			}
			if h.starts[s.room] {
				close(s.conn.send)
			} else {
				// max room = 8
				if len(connections) >= 4 {
					log.Printf("room = %s full \n", s.room)
					close(s.conn.send)
				} else {
					// h.rooms[s.room][s.conn] = true
					h.rooms[s.room][s.conn] = false
				}
			}

		case s := <-h.unregister:

			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					h.deleteUserAlive(s.room, s.conn.username, s.conn)
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}

		case m := <-h.broadcast:
			connections := h.rooms[m.room]

			dataSend := &PayLoadWSS{
				MessPre: h.messPre[m.room],
				Mess:    &Mess{},
				AccNext: "",
				ListAcc: make([]*Acc, 0),
			}
			json.Unmarshal(m.data, &dataSend.Mess)

			var cIf interface{}
			for c := range connections {
				if c.username == dataSend.Mess.Username {
					cIf = c
					break
				}
			}

			// check ok k
			if dataSend.MessPre != "" && len(h.useralive[m.room]) > 1 {
				isFalse, _ := GetModels().CheckTu(dataSend.MessPre, dataSend.Mess.MessStr)
				if !isFalse {
					h.deleteUserAlive(m.room, dataSend.Mess.Username, cIf.(*connection))
				} else {
					h.messPre[m.room] = dataSend.Mess.MessStr
				}
			} else {
				h.messPre[m.room] = dataSend.Mess.MessStr
			}
			for c, kill := range connections {
				dataSend.ListAcc = append(dataSend.ListAcc, &Acc{
					Name:    c.username,
					NotKill: kill,
				})
			}
			_, dataSend.AccNext = h.searchNextAlive(m.room, dataSend.Mess.Username)
			if dataSend.AccNext == "" {
				dataSend.AccNext = dataSend.Mess.Username
			}

			dataSendBytes, _ := json.Marshal(dataSend)
			for c, ok := range connections {
				if ok || c.username == dataSend.Mess.Username {
					select {
					// case c.send <- m.data:
					case c.send <- dataSendBytes:
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
