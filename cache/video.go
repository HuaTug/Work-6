package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"HuaTug.com/kitex_gen/videos"
	Redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := CacheHSet("video:"+key, key, authorid)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := CacheHGet("video:"+key, key)
	if err != nil {
		return 0, nil
	}
	var uid int64
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// ToDo :实现了批量插入操作 共有两种方法 我使用第二种方法
func Insert(videos []*videos.Video) {
	/*	pipe := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}).Pipeline()
		for _, v := range videos {
			pipe.Set(fmt.Sprintf("VideoId:%d", v.VideoId), fmt.Sprintf("Content:%d", v.FavoriteCount), 24*3600*30)
		}
		cmds, err := pipe.Exec()
		if err != nil {
			fmt.Printf("Pipeline exec error: %s", err)
			return
		} else {
			for _, cmd := range cmds {
				if cmd.Err() != nil {
					logrus.Info("Err: ", cmd.Err())
					return
				}
			}
		}
		logrus.Info("Pipe执行完毕")*/
	client := Redis.NewClient(&Redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	//不要使用for循环进行redis的插入操作，这样会降低效率，可以使用redis的批量插入优化
	data := make(map[string]interface{})
	datas := make(map[string]interface{})
	for _, v := range videos {
		data[fmt.Sprintf("key:%d", v.VideoId)] = v.FavoriteCount
	}

	for _, v := range videos {
		datas[fmt.Sprintf("video:%d", v.VideoId)] = v.AuthorId
	}
	err := client.MSet(context.Background(), data).Err()
	if err != nil {
		fmt.Println(err)
	}
	//设置Video和Author的映射
	err = client.MSet(context.Background(), datas).Err()
	if err != nil {
		fmt.Println(err)
	}
}
