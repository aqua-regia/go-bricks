package Cache_test

import (
	"context"
	cache "github.com/aqua-regia/go-bricks/cache/redis"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

type Object struct {
	Str string
	Num int
}

func TestNonClusterMode(t *testing.T) {

	rdbSimple := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	simpleCache := cache.New(&cache.Options{Redis: rdbSimple})

	err := simpleCache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "key1",
		Value: &Object{Str: "hassan", Num: 21},
		TTL:   time.Hour,
	})

	wanted := new(Object)
	err = simpleCache.Get(context.Background(), "key1", &wanted)
	if err != nil {
		t.Errorf("get failed")
	}
	if wanted.Str != "hassan" {
		t.Errorf("got incorrect value")
	}

	err = simpleCache.Delete(context.Background(), "key1")
	if err != nil {
		t.Errorf("cache delete failed")
	}

	err = simpleCache.Get(context.Background(), "key1", &wanted)
	if err == nil {
		t.Errorf("get should have returned not nil")
	}

}

func TestClusterMode(t *testing.T) {
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":30001", ":30002", ":30003", ":30004", ":30005", ":30006"},
	})

	clusterClient := cache.New(&cache.Options{Redis: cluster})

	err := clusterClient.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "key1",
		Value: &Object{Str: "hassan", Num: 21},
		TTL:   time.Hour,
	})

	wanted := new(Object)
	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err != nil {
		t.Errorf("get failed")
	}
	if wanted.Str != "hassan" {
		t.Errorf("got incorrect value")
	}

	err = clusterClient.Delete(context.Background(), "key1")
	if err != nil {
		t.Errorf("cache delete failed")
	}

	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err == nil {
		t.Errorf("get should have returned not nil")
	}

}

func TestUniversalClient(t *testing.T) {
	universal_client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":30001", ":30002", ":30003", ":30004", ":30005", ":30006"},
	})

	clusterClient := cache.New(&cache.Options{Redis: universal_client})

	err := clusterClient.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "key1",
		Value: &Object{Str: "hassan", Num: 21},
		TTL:   time.Hour,
	})

	wanted := new(Object)
	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err != nil {
		t.Errorf("get failed")
	}
	if wanted.Str != "hassan" {
		t.Errorf("got incorrect value")
	}

	err = clusterClient.Delete(context.Background(), "key1")
	if err != nil {
		t.Errorf("cache delete failed")
	}

	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err == nil {
		t.Errorf("get should have returned not nil")
	}
}

func TestGetDefaultUniversalClient(t *testing.T) {
	clusterClient := cache.GetDefaultUniversalClient([]string{":30001", ":30002", ":30003", ":30004", ":30005", ":30006"})

	err := clusterClient.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "key1",
		Value: &Object{Str: "hassan", Num: 21},
		TTL:   time.Hour,
	})

	wanted := new(Object)
	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err != nil {
		t.Errorf("get failed")
	}
	if wanted.Str != "hassan" {
		t.Errorf("got incorrect value")
	}

	err = clusterClient.Delete(context.Background(), "key1")
	if err != nil {
		t.Errorf("cache delete failed")
	}

	err = clusterClient.Get(context.Background(), "key1", &wanted)
	if err == nil {
		t.Errorf("get should have returned not nil")
	}

}
