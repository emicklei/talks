import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//START OMIT
type RedisWriter struct {
	connection redis.Conn
}

func NewRedisWriter() (*RedisWriter, error) {
	if len(*redisHost) == 0 {
		return nil, fmt.Errorf("no redis.host given, writer not available")
	}
	connection, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", *redisHost, *redisPort))
	if err != nil {
		return nil, err
	}
	return &RedisWriter{connection}, nil
}
func (w *RedisWriter) Write(payload []byte) (n int, err error) {
	w.connection.Send("RPUSH", *redisKey, payload)
	if err := w.connection.Err(); err != nil {
		if w.tryRecoverFromError(err) {
			w.Write(payload)
			return 0, err
		}
		return 0, err
	}
	w.connection.Flush()
	return len(payload), nil
}

//END OMIT