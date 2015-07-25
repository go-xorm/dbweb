package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

var (
	ListUser  = []byte("list:user")
	IndexUser = []byte("index:user")
)

type User struct {
	Id       int64
	Name     string
	Password string
}

func EncodePassword(unencoded string) string {
	newPasswd, _ := scrypt.Key([]byte(unencoded), []byte("!#@FDEWREWR&*("), 16384, 8, 1, 64)
	return fmt.Sprintf("%x", newPasswd)
}

func (u *User) EncodePassword() {
	u.Password = EncodePassword(u.Password)
}

func AddUser(user *User) error {
	user.EncodePassword()

	id, err := Db.Incr(IndexUser)
	if err != nil {
		return err
	}

	user.Id = id
	bs, err := json.Marshal(user)
	if err != nil {
		return err
	}

	var buf = make([]byte, 8)
	binary.PutVarint(buf, id)

	Db.Set([]byte(fmt.Sprintf("user:name:%s", user.Name)), buf)
	Db.Set([]byte(fmt.Sprintf("user:%d", id)), bs)
	Db.SAdd(ListUser, buf)

	return err
}

func GetUserById(id int64) (*User, error) {
	bs, err := Db.Get([]byte(fmt.Sprintf("user:%d", id)))
	if err != nil {
		return nil, err
	}

	if len(bs) <= 0 {
		return nil, ErrNotExist
	}

	var user User
	err = json.Unmarshal(bs, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(name string) (*User, error) {
	bs, err := Db.Get([]byte(fmt.Sprintf("user:name:%s", name)))
	if err != nil {
		return nil, err
	}

	if len(bs) <= 0 {
		return nil, ErrNotExist
	}

	id, _ := binary.Varint(bs)

	return GetUserById(id)
}

func UpdateUser(user *User) error {
	if user.Id == 0 {
		return ErrParamError
	}

	user.EncodePassword()

	bs, err := json.Marshal(user)
	if err != nil {
		return err
	}

	var buf = make([]byte, 8)
	binary.PutVarint(buf, user.Id)

	oldUser, err := GetUserById(user.Id)
	if err != nil {
		return err
	}
	Db.Del([]byte(fmt.Sprintf("user:name:%s", oldUser.Name)))

	Db.Set([]byte(fmt.Sprintf("user:name:%s", user.Name)), buf)
	Db.Set([]byte(fmt.Sprintf("user:%d", user.Id)), bs)

	return err
}
