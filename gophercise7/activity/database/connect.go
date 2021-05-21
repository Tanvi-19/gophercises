package database

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var activityBucket = []byte("activities")
var Db *bolt.DB
type Activity struct{
	Num int
	Task string
}




func NewTask(task string)(int, error){
	var id int
	err := Db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket(activityBucket)
		id64, _ := bucket.NextSequence()
		id := int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))
	})
	return id, err
}

func ViewAllTasks() ([]Activity, error) {
	var tasklist []Activity
	err := Db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket(activityBucket)
		point := bucket.Cursor()
		for key, value := point.First(); key != nil; key, value = point.Next() {
			tasklist = append(tasklist, Activity{
				Num:   btoi(key),
				Task: string(value),
			})
		}
		return nil
		
	})
	return tasklist, err
}

func DelTask(num int)error{
	return Db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket(activityBucket)
		return bucket.Delete(itob(num))
	})
}

func itob(val int ) []byte{
	b := make([]byte,8)
	binary.BigEndian.PutUint64(b,uint64(val))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}



func Init(path string)error{
	var err error
	Db, _ = bolt.Open(path,0600,&bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}else{
		return Db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists(activityBucket)
			return err
		})
	}
}