package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	gorma "github.com/hhxsv5/go-redis-memory-analysis"
)

var ctx context.Context

var client *redis.Client

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

func main() {
	write(10000, "len10_10k", getValue(10))
	write(50000, "len10_50k", getValue(10))
	write(500000, "len10_500k", getValue(10))

	write(10000, "len1000_10k", getValue(1000))
	write(50000, "len1000_50k", getValue(1000))
	write(500000, "len1000_500k", getValue(1000))

	write(10000, "len5000_10k", getValue(5000))
	write(50000, "len5000_50k", getValue(5000))
	write(500000, "len5000_500k", getValue(5000))

	analysis()

}

func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)

		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("/Users/redisreport/")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func getValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 't'
	}
	return string(arr)
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})
	ctx = context.Background()
}
