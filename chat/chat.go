package chat

import (

)

var (

)

type Room struct {
  onlineUsers map[string]*user
  broadcast chan Message
}

type Message struct {
  user Client
  txt string
}

