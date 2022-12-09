package dao

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "sqy040213",
		DB:       0,
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("redis 链接成功")
}
func AddUser(username, value1, value2, value3 string) {
	id := "id:"
	n2 := GetUserNumber() + 1
	n3 := strconv.Itoa(n2)
	uid := id + n3
	err1 := HSetUser(context.Background(), uid, "username", username, "password", value1, "check_question", value2, "check_answer", value3, "login", "n", "change", "n")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := Incr(context.Background(), "user_number")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func IfLogin() (string, string) {
	n1 := GetUserNumber()
	u := "id:"
	for id := 1; id <= n1; id++ {
		uid := u + strconv.Itoa(id)
		val, err := HGet(context.Background(), uid, "login")
		if err != nil {
			fmt.Println(err)
			fmt.Println(val)
			return "return", "return"
		}
		v := len(val) - 1
		if val[v:] == "y" {
			return "yes", uid
		}
	}
	return "no", "no"
}
func Start() {
	err := SetNX(context.Background(), "user_number", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err1 := SetNX(context.Background(), "floor", 0)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
func SelectUser(username string) (string, string) {
	n1 := GetUserNumber()
	u := "id:"
	for id := 1; id <= n1; id++ {
		uid := u + strconv.Itoa(id)
		val, err := HGet(context.Background(), uid, "username")
		if err != nil {
			fmt.Println(err)
			fmt.Println(val)
			return "return", "return"
		}
		v := len(uid)
		if val[v+16:] == username {
			return "yes", uid
		}
	}
	return "no", "no"
}
func SelectPasswordFromId(uid string) string {
	val, err := HGet(context.Background(), uid, "password")
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return ""
	}
	v := len(uid)
	return val[v+16:]
}
func Login(uid string) {
	val, err := HGet(context.Background(), uid, "username")
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return
	}
	err1 := HDel(context.Background(), uid, "login")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := HSet(context.Background(), uid, "login", "y")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func ChangePassword(uid, newPassword string) {
	err1 := HDel(context.Background(), uid, "password")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := HSet(context.Background(), uid, "password", newPassword)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func Quit(uid string) {
	err1 := HDel(context.Background(), uid, "login")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := HSet(context.Background(), uid, "login", "n")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func CheckQuestion(uid string) string {
	val, err := HGet(context.Background(), uid, "check_question")
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return ""
	}
	v := len(uid)
	return val[v+22:]
}
func CheckAnswer(answer, username string) (bool, string) {
	n := GetUserNumber()
	u := "id:"
	for id := 1; id <= n; id++ {
		uid := u + strconv.Itoa(id)
		val, err := HGet(context.Background(), uid, "username")
		if err != nil {
			fmt.Println(err)
			fmt.Println(val)
			return false, "return"
		}
		v := len(uid)
		if val[v+16:] == username {
			val1, err1 := HGet(context.Background(), uid, "check_answer")
			if err1 != nil {
				fmt.Println(err1)
				fmt.Println(val1)
				return false, "return"
			}
			if val1[v+20:] == answer {
				return true, uid
			}
			return false, "false"
		}
	}
	return false, "return"
}
func IfChange(uid string) {
	err1 := HDel(context.Background(), uid, "change")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := HSet(context.Background(), uid, "change", "y")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func WhoChange() string {
	n := GetUserNumber()
	u := "id:"
	for id := 1; id <= n; id++ {
		uid := u + strconv.Itoa(id)
		val, err := HGet(context.Background(), uid, "change")
		if err != nil {
			fmt.Println(err)
			fmt.Println(val)
			return ""
		}
		v := len(val) - 1
		if val[v:] == "y" {
			return uid
		}
	}
	return ""
}
func ChangeOver(uid string) {
	err1 := HDel(context.Background(), uid, "change")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	err2 := HSet(context.Background(), uid, "change", "n")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func WhoLogin(uid string) string {
	val, err := HGet(context.Background(), uid, "username")
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return ""
	}
	v := len(uid)
	return val[v+16:]
}
func AddComment(comment string) {
	n := GetFloorNumber()
	err := Set(context.Background(), strconv.Itoa(n+1), comment)
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := Incr(context.Background(), "floor")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
func ReadComment(fl int) string {
	floor := strconv.Itoa(fl)
	val, err := Get(context.Background(), floor)
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return ""
	}
	v := len(floor)
	return val[v+5:]
}
func DeleteComment(n int) {
	num := strconv.Itoa(n)
	err := Del(context.Background(), num)
	if err != nil {
		fmt.Println(err)
		return
	}
	j := GetFloorNumber()
	for i := n + 1; i <= j; i++ {
		in := strconv.Itoa(i)
		val, err := Get(context.Background(), in)
		if err != nil {
			fmt.Println(err)
			fmt.Println(val)
			return
		}
		v := len(in)
		comment := val[v+5:]
		err1 := Del(context.Background(), in)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		im := strconv.Itoa(i - 1)
		err2 := Set(context.Background(), im, comment)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
	err1 := Decr(context.Background(), "floor")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
func ClearComments() {
	n := GetFloorNumber()
	for i := 1; i <= n; i++ {
		in := strconv.Itoa(i)
		err := Del(context.Background(), in)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err := SetXX(context.Background(), "floor", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func Unsubscribe(uid string) {
	val, err2 := HGet(context.Background(), uid, "username")
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println(val)
		return
	}
	v := len(uid)
	username := val[v+16:]
	err3 := Del(context.Background(), username)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	err := Del(context.Background(), uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	err1 := HSetUser(context.Background(), uid, "username", "", "password", "", "check_question", "", "check_answer", "", "login", "", "change", "")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
func LikeSomeone(username, like string) int64 {
	val2, err2 := SAdd(context.Background(), like, username)
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println(val2)
		return -1
	}
	return val2
}
func CancelLike(username, dislike string) int64 {
	val, err := SRem(context.Background(), dislike, username)
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return -1
	}
	return val
}
func MyLikes(username string) int64 {
	val, err := SCard(context.Background(), username)
	if err != nil {
		fmt.Println(err)
		fmt.Println(val)
		return -1
	}
	return val
}
func ClearAll() {
	err := FlushAll(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}
