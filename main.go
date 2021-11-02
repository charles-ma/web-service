package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"web-service/mymath"

	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func handler(w http.ResponseWriter, r *http.Request) {

	rateLmitingStart := time.Now()

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "myredis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := getKey(r.RemoteAddr)

	// fmt.Fprintf(w, "redis key: %s\n", key)

	exist := rdb.Exists(ctx, key)
	if exist == nil {
		fmt.Fprintln(w, "redis error: check exists")
	}

	if exist.Val() == 0 {
		err := rdb.Set(ctx, key, 0, 0).Err()
		if err != nil {
			fmt.Fprintf(w, "redis set error: %s", err)
		}
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Fprintf(w, "redis error: %s", err)
	}
	// else {
	// 	fmt.Fprintf(w, "redis raw result: %s\n", val)
	// }

	c, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "redis value not an int: %s", err)
	}

	c = c / 2
	rdb.IncrBy(ctx, key, 1)

	rateLimitingEnd := time.Now()

	if c > 5 {
		fmt.Fprintln(w, "You are rate limited")
	} else {
		n := mymath.GenerateRandomNumber(100)
		fmt.Fprintf(w, "Look what number you got from Docker!\n%d\n", n)

		if n >= 90 {
			fmt.Fprintf(w, "perfect!\n")
		}

		fmt.Fprintf(w, "\nYour ip address is %s\n", r.RemoteAddr)
	}

	fmt.Fprintln(w, "\n\n------------------")
	fmt.Fprintf(w, "Your have visited %d times\n", c+1)
	fmt.Fprintf(w, "Rate limiting took %d ms\n", (rateLimitingEnd.UnixNano()-rateLmitingStart.UnixNano())/int64(time.Millisecond))
}

func getKey(remoteAddr string) string {
	return remoteAddr + ":" + fmt.Sprint(getCurrentSharpMin())
}

func getCurrentSharpMin() int64 {
	now := time.Now()
	sec := now.Unix()
	return sec - sec%60
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
