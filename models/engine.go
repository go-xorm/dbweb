package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

var (
	ListEngine  = []byte("list:engine")
	IndexEngine = []byte("index:engine")
)

type Engine struct {
	Id         int64
	Name       string `binding:"Required"`
	Driver     string `binding:"Required"`
	DataSource string `binding:"Required"`
}

func AddEngine(engine *Engine) error {
	id, err := Db.Incr(IndexEngine)
	if err != nil {
		return err
	}

	engine.Id = id
	bs, err := json.Marshal(engine)
	if err != nil {
		return err
	}

	var buf = make([]byte, 8)
	binary.PutVarint(buf, id)

	Db.Set([]byte(fmt.Sprintf("engine:name:%s", engine.Name)), buf)
	Db.Set([]byte(fmt.Sprintf("engine:%d", id)), bs)
	Db.SAdd(ListEngine, buf)

	return err
}

func GetEngineById(id int64) (*Engine, error) {
	bs, err := Db.Get([]byte(fmt.Sprintf("engine:%d", id)))
	if err != nil {
		return nil, err
	}

	var engine Engine
	err = json.Unmarshal(bs, &engine)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func GetEngineByName(name string) (*Engine, error) {
	bs, err := Db.Get([]byte(fmt.Sprintf("engine:name:%s", name)))
	if err != nil {
		return nil, err
	}

	id, _ := binary.Varint(bs)

	return GetEngineById(id)
}

func DelEngineByName(name string) error {
	bs, err := Db.Get([]byte(fmt.Sprintf("engine:name:%s", name)))
	if err != nil {
		return err
	}

	id, _ := binary.Varint(bs)

	_, err = Db.Del([]byte(fmt.Sprintf("engine:%d", id)))
	if err != nil {
		return err
	}

	_, err = Db.Del([]byte(fmt.Sprintf("engine:name:%s", name)))
	if err != nil {
		return err
	}

	_, err = Db.SRem(ListEngine, bs)

	return err
}

func DelEngineById(id int64) error {
	engine, err := GetEngineById(id)
	if err != nil {
		return err
	}

	return DelEngineByName(engine.Name)
}

func FindEngines() ([]*Engine, error) {
	engines := make([]*Engine, 0)

	members, err := Db.SMembers(ListEngine)
	if err != nil {
		return nil, err
	}
	for _, bs := range members {
		id, _ := binary.Varint(bs)

		engine, err := GetEngineById(id)
		if err != nil {
			return nil, err
		}
		engines = append(engines, engine)
	}
	return engines, nil
}
