package golangredis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB: 0,
})

var ctx = context.Background()
func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}
func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Kareem", time.Second*3)
	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Kareem", result)

	time.Sleep(time.Second * 5)
	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}
func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Kareem")
	client.RPush(ctx, "names", "Abdul")
	client.RPush(ctx, "names", "Melayu")

	assert.Equal(t, "Kareem", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Abdul", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Melayu", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}
func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Kareem")
	client.SAdd(ctx, "students", "Kareem")
	client.SAdd(ctx, "students", "Kareem")
	client.SAdd(ctx, "students", "Abdul")
	client.SAdd(ctx, "students", "Abdul")
	client.SAdd(ctx, "students", "Abdul")
	client.SAdd(ctx, "students", "Abdul")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
}
func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Kareem"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 80, Member: "Abdul"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 90, Member: "Melayu"})

	assert.Equal(t, []string{"Abdul", "Melayu", "Kareem"}, client.ZRange(ctx, "scores", 0, -1).Val())
	assert.Equal(t, "Kareem", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Melayu", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Abdul", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHas(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Kareem")
	client.HSet(ctx, "user:1", "email", "example@email.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Kareem", user["name"])
	assert.Equal(t, "example@email.com", user["email"])

	client.Del(ctx, "user:1")
}
func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name: "Toko A",
		Longitude: 106.822702,
		Latitude: -6.177590,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name: "Toko B",
		Longitude: 106.822702,
		Latitude: -6.177590,
	})
	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
	assert.Equal(t,0.348, distance )
	
	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude: 106.822702,
		Latitude: -6.177590,
		Radius: 5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}


func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "vistiors", "kareem", "Abdul")
	client.PFAdd(ctx, "visitors", "ane", "ano")
	client.PFAdd(ctx, "visitors", "alo", "alu")
	assert.Equal(t, int64(4), client.PFCount(ctx, "visitors").Val())
}
func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func (pipeliner redis.Pipeliner) error  {
		pipeliner.SetEx(ctx, "name", "Kareem", time.Second*5)
		pipeliner.SetEx(ctx, "address", "Indonesia", time.Second*5)
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, "Kareem", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}
func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func (pipeliner redis.Pipeliner) error  {
		pipeliner.SetEx(ctx, "name", "Abdul", time.Second*5)
		pipeliner.SetEx(ctx, "address", "Medan", time.Second*5)
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, "Abdul", client.Get(ctx, "name").Val())
	assert.Equal(t, "Medan", client.Get(ctx, "address").Val())
}
func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name": "Kareem",
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T)  {
	client.XGroupCreate(ctx, "members", "group-1","0")
	client.XGroupCreateConsumer(ctx, "members", "group-1","customer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1","customer-2")
}

func TestConsumerStream(t *testing.T) {
	streams := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group: "group-1",
		Consumer: "consumer-1",
		Streams: []string{"members", ">"},
		Count: 2,
		Block: 5 * time.Second,
	}).Val()

	for _, stream := range streams {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

func TestSubcriberPubSub(t *testing.T) {
	subcriber := client.Subscribe(ctx, "channel-2")
	defer subcriber.Close()
	for i := 0; i < 10; i++ {
		message, err := subcriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}
func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-2", "Hello " + strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}