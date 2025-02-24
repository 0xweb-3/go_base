package main

import "fmt"

type db struct {
	DB  string // 假设为mysql的配置
	Rdb string // 假设为redis配置
}

type DbOption func(*db)

func InfraDbOption(mysqlUrl string) DbOption {
	return func(_db *db) {
		fmt.Println("连接mysqlUrl:", mysqlUrl)
	}
}

func InfraRedisOption(redisID string, redisport int) DbOption {
	return func(_db *db) {
		fmt.Println("连接RedisUrl:", redisID, redisport)
	}
}

func NewDb(options ...DbOption) *db {
	_db := &db{} // 初始化一个默认的车辆对象
	for _, option := range options {
		option(_db) // 应用每一个配置选项
	}
	return _db
}
