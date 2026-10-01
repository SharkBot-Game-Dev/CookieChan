package cache

import "github.com/disgoorg/snowflake/v2"

// {"ユーザーid": "チャンネルid"}
var JankenSessions = make(map[snowflake.ID]snowflake.ID)
var JankenAnswer = make(map[snowflake.ID]string)
