package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var conf = NewConf()
var client *redis.Client

func init() {
	conf.Load("conf.json")
	client = NewRedisClient(conf.Redis.Addr, conf.Redis.Passwd, conf.Redis.Db)
}

func main() {
	//fmt.Printf("%+v\n", conf)

	//keyName := "test*"
	//keys := client.Keys(keyName).Val()
	//for index, name := range keys {
	//	fmt.Println(index, name)
	//}

	//roomId := "2"
	//roomKey := "test"+roomId
	//val, _ := client.Get(roomKey).Result()
	//fmt.Printf("room%+v\n", val)

	roomListKey := "room_list" //list类型存放id
	roomList, _ := client.LRange(roomListKey, 0, -1).Result()
	fmt.Printf("roomList = %+v\n", roomList)

	for _, id := range roomList {
		room := GetRoom(id)
		member := GetRoomMember(id, room["member_num"])

		memberNum, _ := strconv.Atoi(room["member_num"])
		roomStatus := "可加入"
		if len(member) >= memberNum {
			roomStatus = "已满"
		}
		fmt.Printf("房间人数状态: %+v\n", roomStatus)
	}

}

func GetRoom(id string) map[string]string {
	roomKey := "room_" + id //房间id
	room, _ := client.HGetAll(roomKey).Result()
	fmt.Printf("room = %+v\n", room)
	return room
}

func GetRoomMember(id string, memberNum string) []string {
	roomMemberKey := "room_" + id + "_member"
	member, _ := client.LRange(roomMemberKey, 0, -1).Result()
	fmt.Printf("member = %+v\n", member)
	fmt.Printf("len(member) = %+v\n", len(member))
	fmt.Printf("人数 %+v/%+v\n", len(member), memberNum)

	return member
}

func NewRedisClient(addr string, passwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})

	client.Ping().Result()

	return client
}
