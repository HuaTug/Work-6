package cache //nolint:gofmt
/*
在Go语言中 import声明通常按照以下顺序组织：标准库包，空行，第三方包，空行，项目内部包
*/
import (
	"encoding/json"
	"strconv"

	"HuaTug.com/kitex_gen/comments"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/sirupsen/logrus"
)

func CacheSetAllComment(videoId int64, c []*comments.Comment) {
	vid := strconv.FormatInt(videoId, 10)
	err := CacheHSet("comment:"+vid, vid, c)
	if err != nil {
		logrus.Info("Set Cache error: ", err)
	}
}
func CacheGetListComment(videoId int64) ([]*comments.Comment, error) {
	key := strconv.FormatInt(videoId, 10)
	data, err := CacheHGet("comment:"+key, key)
	var comment []*comments.Comment
	if err != nil {
		return comment, err
	}
	_ = json.Unmarshal(data, &comment)
	return comment, nil
}

func CacheSetCommentVideo(videoId, commentId int64) error {
	key := strconv.FormatInt(commentId, 10)
	err := CacheHSet("convert:"+key, key, videoId)
	if err != nil {
		hlog.Info(err)
		return err
	}
	return err
}

func CacheGetCommentVideo(commentId int64) (videoId int64, err error) {
	key := strconv.FormatInt(commentId, 10)
	data, err := CacheHGet("convert:"+key, key)
	if err != nil {
		return videoId, err
	}
	_ = json.Unmarshal(data, &videoId)
	return videoId, nil
}
