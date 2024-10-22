package db

import (
	"github.com/AnnonaOrg/gokv"
	"github.com/AnnonaOrg/gokv/redis"
	"github.com/AnnonaOrg/osenv"
)

var (
	KVStore gokv.Store
)

func GetRedisOptions() redis.Options {
	options := redis.DefaultOptions
	options.Address = osenv.GetServerDbRedisAddress()
	options.DB = osenv.GetServerDbRedisDB()
	if pw := osenv.GetServerDbRedisPassword(); len(pw) > 0 {
		options.Password = pw
	}
	return options
}

func KVStoreInit() error {
	options := GetRedisOptions()
	kvStore, err := redis.NewClient(options)
	if err != nil {
		return err
	}
	KVStore = kvStore
	return nil
}

func NewKvStoreClient() (gokv.Store, error) {
	options := GetRedisOptions()
	return redis.NewClient(options)
}

// func KvStoreClient() gokv.Store {
// 	return KVStore
// }

func KvStoreClose() error {
	return KVStore.Close()
}
