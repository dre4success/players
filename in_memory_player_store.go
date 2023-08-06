package main

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type InMemoryPlayerStore struct {
	store  map[string]int
	client *redis.Client
}

var ctx = context.Background()

const table = "player"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	return &InMemoryPlayerStore{map[string]int{}, client}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++

	i.client.HSet(ctx, table, map[string]interface{}{name: i.store[name]})

}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.GetValue(name)
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	redisValue := i.client.HGetAll(ctx, table).Val()

	for name, wins := range redisValue {
		winsint, _ := strconv.Atoi(wins)
		league = append(league, Player{name, winsint})
	}
	return league
}

func (i *InMemoryPlayerStore) GetValue(name string) int {
	result, err := i.client.HGet(ctx, table, name).Result()
	log.Println(result)
	if err != nil {
		if err == redis.Nil {
			log.Println("key does not exist")
		}
		panic(err)
	}
	value, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return value
}
