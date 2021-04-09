package main

import (
	"flag"
	"github.com/aas-spec/mlog"
	"github.com/aas-spec/redisproxy"
	"net/http"
	"strconv"
)



var redisAddr = flag.String("redis_addr", "localhost", "redis address")
var redisPort = flag.Int("redis_port", 6379, "redis port")
var redisPwd = flag.String("redis_pwd", "pass", "redis password")
var redisDb = flag.Int("redis_db", 0, "redis database")
var redisChannel = flag.String("redis_channel", "/queue", "print tasks channel")
var webHost = flag.String("web_host", "127.0.0.1", "web host")
var webPort = flag.Int("web_port", 12345, "web port")
var logLevel = flag.Int("log_level", 10, "log_level")


func main() {
	flag.Parse()
	mlog.SetLogLevel(mlog.DefLoggerID, *logLevel)
	mlog.LPrint(3, "Print Task Server Started")

	server := goatee.CreateServer()
	server.RegisterAuthFunc(Authenticate)
	webHost := *webHost + ":" + strconv.Itoa(*webPort)
	server.StartServer( goatee.RedisConfig {
		 Host:     *redisAddr,
		 Port:     *redisPort,
		 Password: *redisPwd,
		 Db:       *redisDb,
		 Channel:  *redisChannel,
	}, webHost )
}

func Authenticate(req *http.Request) string {
	vals := req.URL.Query()

	if vals.Get("id") != "" {
		mlog.LPrintf( 5, "Client %s authenticated ",  vals.Get("id"))
		// Возвращаю идентификатор клиента
		return vals.Get("id")
	}
	return ""
}