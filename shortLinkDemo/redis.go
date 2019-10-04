package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pilu/go-base62"
	"time"
)

const (
	// URLIDKEY is global counter
	URLIDKEY = "next.url.id"
	// ShortlinkKey mapping the shortlink to the url
	ShortlinkKey = "shortlink:%s:url"
	// URLHashKey mapping the hash of the url to the shortlink
	URLHashKey = "urlhash:%s:url"
	// ShortlinkDetailKey mapping the shortlink to the detail of url
	ShortlinkDetailKey = "shortlink:%s:detail"
)

// RedisCli contains a redis Client
type RedisCli struct {
	Cli *redis.Client
}

// URLDetail contains the detail of the shortlink
type URLDetail struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

// NewRedisCli create a redis Client
func NewRedisCli(addr string, passwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})

	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}

	return &RedisCli{c}
}

//Shorten convert url to shortlink
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {

	//convert url to sh1  hash
	h := toSha1(url)

	//fetch it if the url is cached
	d, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {
		//not existed,noting to do
	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {
			//expiration,noting to do
		} else {
			return d, nil
		}
	}

	//increase the global counter
	err = r.Cli.Incr(URLHashKey).Err()
	if err != nil {
		return "", err
	}

	//encode global counter to base62
	id, err := r.Cli.Get(URLIDKEY).Int64()
	if err != nil {
		return "", err
	}
	eid := base62.Encode(int(id))

	//store the url against this encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url,
		time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	//store the url against the hash of it
	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid,
		time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreateAt:            time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		})
	if err != nil {
		return "", err
	}

	//store the url detail against this encoded id
	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), detail,
		time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	return eid, nil
}

//ShortlinkInfo returns the detail of the shortlink
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	d, err := r.Cli.Get(fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", StausError{404, errors.New("Unkown short URL")}
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}

}

//Unshort converten convert shortlonk to url
func (r *RedisCli) Unshorten(eid string) (string, error) {
	url, err := r.Cli.Get(fmt.Sprintf(ShortlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", StausError{404, err}
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}

}

func toSha1(data string) string {
	sa1 := sha1.New()
	sa1.Write([]byte(data))
	return hex.EncodeToString(sa1.Sum([]byte(nil)))
}
