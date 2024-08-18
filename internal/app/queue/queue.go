package queue

import (
	"encoding/json"
	"fmt"
	"framework/internal/app/db"
)

func New() Quer {
	return &Queue{}
}

type Quer interface {
	Construct(db.DBer)
	Dispatch(string, string, map[string]interface{}) error
	Pull(string) (map[string]interface{}, error)
}

type Queue struct {
	db db.DBer
}

func (q *Queue) Construct(d db.DBer) {
	q.db = d
}

func (q *Queue) Dispatch(topic, name string, message map[string]interface{}) error {
	strMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	sql := "insert into jobs (topic, name, message) values (?,?,?)"
	_, err = q.db.Execute(sql, topic, name, string(strMessage))

	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Pull(topic string) (map[string]interface{}, error) {
	// todo add set is_visible false
	sql := "SELECT id, name, message from jobs where topic = ? and is_visible = 1 order by id desc"

	var message map[string]interface{}
	res, err := q.db.QueryOne(sql, topic)
	if err != nil {
		return nil, err
	}

	if mess, ok := res["message"]; ok {
		err = json.Unmarshal([]byte(mess.(string)), &message)
		if err != nil {
			return nil, err
		}

	}

	if id, ok := res["id"]; ok {
		sql := "delete from jobs where id = ?"
		_, err := q.db.Execute(sql, id)
		if err != nil {
			return nil, err
		}
		return message, nil
	}

	return nil, fmt.Errorf("could not get message from queue")
}
