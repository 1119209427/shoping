package tool
import (
	"fmt"
	"github.com/go-redis/redis/v8"
)
var Redis *redis.Client
func init() {
	redisCfg := GetCfg().Redis
	RedisConn := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})
	fmt.Println(RedisConn)
}

func GetRedisConn() *redis.Client {
	return Redis
}
