package repoimpl

import (
	"context"
	"time"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/data"
)

func NewCaptchaRepository(data *data.Data) repository.Captcha {
	return &captchaRepositoryImpl{
		data: data,
	}
}

type captchaRepositoryImpl struct {
	data *data.Data
}

func (l *captchaRepositoryImpl) CreateCaptcha(ctx context.Context, captcha *bo.ValidateCaptchaItem, duration time.Duration) error {
	bs, err := captcha.MarshalBinary()
	if err != nil {
		return err
	}
	return l.data.GetCacher().Set(ctx, captcha.Id, string(bs), duration)
}

func (l *captchaRepositoryImpl) GetCaptchaById(ctx context.Context, id string) (*bo.ValidateCaptchaItem, error) {
	str, err := l.data.GetCacher().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	var captcha bo.ValidateCaptchaItem
	if err = captcha.UnmarshalBinary([]byte(str)); err != nil {
		return nil, err
	}
	return &captcha, nil
}