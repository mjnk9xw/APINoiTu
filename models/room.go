package models

const (
	maxUserOnRoom = 8
)

// func (m *Models) CreateRoom(usname string, ip string, nameRoom string) string {
// 	key := usname + ip
// 	listUser := []string{nameRoom}
// 	m.cacheRoom.Add(nameRoom+"_key", key)
// 	m.cacheRoom.Add(nameRoom, listUser)
// 	return nameRoom
// }

// func (m *Models) startGameRoom(nameRoom string, usname string, ip string) bool {
// 	key := usname + ip

// }
// func (m *Models) outRoom(usname string, ip string, nameRoom string) bool {

// 	listUser, _ := m.cacheRoom.Get(nameRoom)

// 	return true
// }

// func (m *Models) joinRoom(usname string) int64 {

// 	m.cacheRoom.Add(nameRoom+"_number", 1)
// 	return
// }
