package dao

import (
	"fmt"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func GetFloorNumber() int {
	n, err := Get(context.Background(), "floor")
	if err != nil {
		fmt.Println(err)
		fmt.Println(n)
		return -1
	}
	n2 := n[11:]
	n1, err1 := strconv.Atoi(n2)
	if err1 != nil {
		fmt.Println(err1)
		return -1
	}
	return n1
}
func GetUserNumber() int {
	n, err := Get(context.Background(), "user_number")
	if err != nil {
		fmt.Println(err)
		fmt.Println(n)
		return -1
	}
	n2 := n[17:]
	n1, err1 := strconv.Atoi(n2)
	if err1 != nil {
		fmt.Println(err1)
		return -1
	}
	return n1
}
func Set(ctx context.Context, key, value string) error {
	SetKV := Rdb.Set(ctx, key, value, 999*time.Hour)
	return SetKV.Err()
}
func Get(ctx context.Context, key string) (string, error) {
	GetK := Rdb.Get(ctx, key)
	if GetK.Err() != nil {
		return "", GetK.Err()
	}
	return GetK.String(), nil
}
func Del(ctx context.Context, key string) error {
	DelK := Rdb.Del(ctx, key)
	return DelK.Err()
}
func SetNX(ctx context.Context, key string, value int) error {
	SetNxKV := Rdb.SetNX(ctx, key, value, 999*time.Hour)
	return SetNxKV.Err()
}
func SetXX(ctx context.Context, key string, value int) error {
	SetXxKV := Rdb.SetXX(ctx, key, value, 999*time.Hour)
	return SetXxKV.Err()
}
func Incr(ctx context.Context, key string) error {
	IncrK := Rdb.Incr(ctx, key)
	return IncrK.Err()
}
func Decr(ctx context.Context, key string) error {
	DecrK := Rdb.Decr(ctx, key)
	return DecrK.Err()
}
func HSetUser(ctx context.Context, key string, field0, value0, field1, value1, field2, value2, field3, value3, field4, value4, field5, value5 string) error {
	HSetUserKFV := Rdb.HSet(ctx, key, field0, value0, field1, value1, field2, value2, field3, value3, field4, value4, field5, value5)
	return HSetUserKFV.Err()
}
func HSet(ctx context.Context, key string, field, value string) error {
	HSetKFV := Rdb.HSet(ctx, key, field, value)
	return HSetKFV.Err()
}
func HGet(ctx context.Context, key, field string) (string, error) {
	HGetKF := Rdb.HGet(ctx, key, field)
	if HGetKF.Err() != nil {
		return "", HGetKF.Err()
	}
	return HGetKF.String(), nil
}
func HDel(ctx context.Context, key, field string) error {
	HDelKF := Rdb.HDel(ctx, key, field)
	return HDelKF.Err()
}
func SAdd(ctx context.Context, key, value string) (int64, error) {
	SAddKV := Rdb.SAdd(ctx, key, value)
	if SAddKV.Err() != nil {
		return -1, SAddKV.Err()
	}
	return SAddKV.Val(), nil
}
func SRem(ctx context.Context, key, value string) (int64, error) {
	SRemKV := Rdb.SRem(ctx, key, value)
	if SRemKV.Err() != nil {
		return -1, SRemKV.Err()
	}
	return SRemKV.Val(), nil
}
func SCard(ctx context.Context, key string) (int64, error) {
	SCardK := Rdb.SCard(ctx, key)
	if SCardK.Err() != nil {
		return -1, SCardK.Err()
	}
	return SCardK.Val(), nil
}
func FlushAll(ctx context.Context) error {
	FA := Rdb.FlushAll(ctx)
	return FA.Err()
}
