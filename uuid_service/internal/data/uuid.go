package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	uuid "github.com/satori/go.uuid"
	"time"
	"uuid_service/internal/biz"
)

const (
	// 随机UUID
	UUID_TYPE_DEFAULT = "default"
	// 时间随机值
	UUID_TYPE_TIME_RAND = "time_rand"
)

type uuidRepo struct {
	data *Data
	log  *log.Helper
}

// NewUuidRepo .
func NewUuidRepo(data *Data, logger log.Logger) biz.UuidRepo {
	return &uuidRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *uuidRepo) GenerateUuid(ctx context.Context, g *biz.Uuid) (string, error) {
	var (
		value = ""
		err error
	)

	switch g.GenerateUuid.GetType() {
	case UUID_TYPE_TIME_RAND:
		value = fmt.Sprintf("%v", time.Now().UnixNano() / 1e6)
		break
	default:
		value = uuid.NewV4().String()
	}

	return value, err
}