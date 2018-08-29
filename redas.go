package drivers

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	logx "github.com/my0sot1s/godef/log"
)

// RedisCli rd
type RedisCli struct {
	client   *redis.Client
	duration time.Duration
}

// InitRd start rd
func (rc *RedisCli) InitRd(redisHost, redisDb, redisPass string) error {

	rc.client = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})
	rc.duration = 60 * 10 * time.Second // 10 phut

	_, err := rc.client.Ping().Result()
	// fmt.Println(pong, err)
	if err != nil {
		logx.ErrLog(err)
		return err
	}
	logx.Log("+ REDIS CONNECTED RDNAME : ", redisDb)

	return nil
}

// SetValue to rd
func (rc *RedisCli) SetValue(key string, value string, expiration time.Duration) error {
	return rc.client.Set(key, value, expiration).Err()
}

// GetValue to rd
func (rc *RedisCli) GetValue(key string) (string, error) {
	val, err := rc.client.Get(key).Result()
	return val, err
}

// DelKey hard del key
func (rc *RedisCli) DelKey(key []string) (int, error) {
	val, err := rc.client.Del(key...).Result()
	return int(val), err
}

// LPushItem push multiple
func (rc *RedisCli) LPushItem(key string, timeExpired int, values ...interface{}) error {

	// str := make([]string, 0)
	for _, v := range values {
		b, e := json.Marshal(v)
		logx.ErrLog(e)
		_, err := rc.client.LPush(key, string(b)).Result()
		logx.ErrLog(err)
	}
	rc.SetExpired(key, timeExpired)
	return nil
}

// LRangeAll get multiple
func (rc *RedisCli) LRangeAll(key string) ([]map[string]interface{}, error) {
	var raw []string
	raw, err := rc.client.LRange(key, 0, -1).Result()
	desMap := make([]map[string]interface{}, 0)
	for _, v := range raw {
		var d map[string]interface{}
		json.Unmarshal([]byte(v), &d)
		desMap = append(desMap, d)
	}
	return desMap, err
}

// SetExpired set expried
func (rc *RedisCli) SetExpired(key string, min int) bool {
	d := time.Duration(min) * time.Minute

	b, err := rc.client.Expire(key, d).Result()
	logx.ErrLog(err)
	return b
}
