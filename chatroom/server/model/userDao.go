package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{
		pool: pool,
	}
}

// func (me *UserDao) GetByUserId

func (me *UserDao) GetByUserId(userId int) (user *User, err error) {
	user = &User{}
	conn := me.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("redis未找到id:", userId)
		}
		return
	}

	err = json.Unmarshal([]byte(res), user)
	return

}

func (me *UserDao) Insert(user User) (err error) {
	_, err = me.GetByUserId(user.UserId)
	if err == nil {
		err = errors.New("用户已经存在")
		return
	}

	data, err := json.Marshal(user)

	conn := me.pool.Get()
	defer conn.Close()

	_, err = conn.Do("HSet", "users", user.UserId, string(data))

	return

}
