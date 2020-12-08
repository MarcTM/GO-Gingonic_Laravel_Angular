package Config

import(
	"fmt"

	"github.com/gomodule/redigo/redis"
)


// Redis set value
func RedisSet(name string, value string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer conn.Close()

	_, err = conn.Do(
		"SET",
		name,
		value,
	)
	if err != nil {
		fmt.Println("error")
		return
	}
}


// Redis get value
func RedisGet(name string) string {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer conn.Close()

	value, err := redis.String(conn.Do(
		"GET", name,
	))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}