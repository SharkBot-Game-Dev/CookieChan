package cache

import (
	"github.com/disgoorg/snowflake/v2"
	"sync"
)

type jankenSession struct {
	channel snowflake.ID
	answer  string
}

var jankenMu sync.Mutex
var jankenSessions = make(map[snowflake.ID]jankenSession)

func StartJanken(user, channel snowflake.ID, answer string) bool {
	jankenMu.Lock()
	defer jankenMu.Unlock()
	if session, ok := jankenSessions[user]; ok && session.channel == channel {
		return false
	}
	jankenSessions[user] = jankenSession{channel: channel, answer: answer}
	return true
}

// Consume a valid move atomically so only one message can finish a game.
func PlayJanken(user, channel snowflake.ID, valid bool) (string, bool) {
	jankenMu.Lock()
	defer jankenMu.Unlock()
	session, ok := jankenSessions[user]
	if !ok || session.channel != channel {
		return "", false
	}
	if valid {
		delete(jankenSessions, user)
	}
	return session.answer, true
}
