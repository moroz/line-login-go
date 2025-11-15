package config

import (
	"log"
	"os"
)

// MustGetenv 會取得一個環境變數，如果該環境變數沒有值，它就會停止整個程式，因爲沒有值我們不能繼續操作。
func MustGetenv(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", name)
	}
	return val
}

var LineClientId = MustGetenv("LINE_CLIENT_ID")
var LineClientSecret = MustGetenv("LINE_CLIENT_SECRET")

const LineCallbackUri = "http://localhost:3000/oauth/line/callback"
