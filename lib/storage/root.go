package storage

import (
	"fmt"
	"log"
	"net"

	bolt "go.etcd.io/bbolt"
)

var Instance *bolt.DB

func Init(path string) {
	_db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	_db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ips"))
		if err != nil {
			return err
		}
		return nil
	})
	Instance = _db
}

func IpExists(ip *net.IP) bool {
	var v []byte
	err := Instance.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		k := []byte(ip.String())
		v = b.Get(k)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return v != nil
}

func StoreIpStatus(ip *net.IP, isAccessible bool) error {
	return Instance.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ips"))
		return b.Put([]byte(ip.String()), []byte(fmt.Sprintf("%t", isAccessible)))
	})
}
