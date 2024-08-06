package repoimpl

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/util/types"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

// NewAlarmGroupRepository 创建策略分组仓库
func NewAlarmGroupRepository(data *data.Data) repository.AlarmGroup {
	return &alarmGroupRepositoryImpl{
		data: data,
	}
}

type (
	alarmGroupRepositoryImpl struct {
		data *data.Data
	}
)

func (a *alarmGroupRepositoryImpl) CreateAlarmGroup(ctx context.Context, params *bo.CreateAlarmGroupParams) (*bizmodel.AlarmGroup, error) {
	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return nil, err
	}

	alarmGroupModel := createAlarmGroupParamsToModel(ctx, params)

	err = bizQuery.Transaction(func(tx *bizquery.Query) error {
		if err := tx.AlarmGroup.WithContext(ctx).Create(alarmGroupModel); err != nil {
			return err
		}
		noticeUsers := createAlarmNoticeUsersToModel(ctx, params.NoticeUsers, alarmGroupModel.ID)
		if err := tx.AlarmNoticeUser.WithContext(ctx).Create(noticeUsers...); err != nil {
			return err
		}
		return nil
	})
	if !types.IsNil(err) {
		return nil, err
	}
	return alarmGroupModel, nil

}

func (a *alarmGroupRepositoryImpl) UpdateAlarmGroup(ctx context.Context, params *bo.UpdateAlarmGroupParams) error {
	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return err
	}
	noticeUsers := createAlarmNoticeUsersToModel(ctx, params.UpdateParam.NoticeUsers, params.ID)
	return bizQuery.Transaction(func(tx *bizquery.Query) error {
		if !types.IsNil(params.UpdateParam.NoticeUsers) {
			// 更新告警分组
			if _, err = tx.AlarmGroup.WithContext(ctx).Where(tx.AlarmGroup.ID.Eq(params.ID)).UpdateSimple(
				tx.AlarmGroup.Name.Value(params.UpdateParam.Name),
				tx.AlarmGroup.Remark.Value(params.UpdateParam.Remark),
			); !types.IsNil(err) {
				return err
			}

			if len(params.UpdateParam.NoticeUsers) > 0 {
				// 清除通知人员关联信息
				groupModel := &bizmodel.AlarmGroup{AllFieldModel: model.AllFieldModel{ID: params.ID}}
				if err := tx.AlarmGroup.NoticeUsers.Model(groupModel).Clear(); err != nil {
					return err
				}
				if err := tx.AlarmNoticeUser.WithContext(ctx).Create(noticeUsers...); err != nil {
					return err
				}
			}
		}
		//TODO 更新hook
		return nil
	})
}

func (a *alarmGroupRepositoryImpl) DeleteAlarmGroup(ctx context.Context, alarmId uint32) error {
	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return err
	}
	groupModel := &bizmodel.AlarmGroup{AllFieldModel: model.AllFieldModel{ID: alarmId}}
	return bizQuery.Transaction(func(tx *bizquery.Query) error {
		// 清除通知人员关联信息
		if err := tx.AlarmGroup.NoticeUsers.Model(groupModel).Clear(); err != nil {
			return err
		}

		// TODO 清空hook关联信息

		if _, err = tx.AlarmGroup.WithContext(ctx).Where(bizQuery.AlarmGroup.ID.Eq(alarmId)).Delete(); !types.IsNil(err) {
			return err
		}
		return nil
	})
}

func (a *alarmGroupRepositoryImpl) GetAlarmGroup(ctx context.Context, alarmId uint32) (*bizmodel.AlarmGroup, error) {
	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return nil, err
	}
	return bizQuery.AlarmGroup.WithContext(ctx).Where(bizQuery.AlarmGroup.ID.Eq(alarmId)).Preload(field.Associations).First()
}

func (a *alarmGroupRepositoryImpl) AlarmGroupPage(ctx context.Context, params *bo.QueryAlarmGroupListParams) ([]*bizmodel.AlarmGroup, error) {
	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return nil, err
	}
	bizWrapper := bizQuery.AlarmGroup.WithContext(ctx)
	var wheres []gen.Condition
	if !types.TextIsNull(params.Name) {
		wheres = append(wheres, bizQuery.AlarmGroup.Name.Like(params.Name))
	}

	if !params.Status.IsUnknown() {
		wheres = append(wheres, bizQuery.AlarmGroup.Status.Eq(params.Status.GetValue()))
	}
	if !types.TextIsNull(params.Keyword) {
		bizWrapper = bizWrapper.Or(bizQuery.AlarmGroup.Name.Like(params.Keyword))
		bizWrapper = bizWrapper.Or(bizQuery.AlarmGroup.Remark.Like(params.Keyword))
	}
	bizWrapper = bizWrapper.Where(wheres...)

	if err := types.WithPageQuery[bizquery.IAlarmGroupDo](bizWrapper, params.Page); err != nil {
		return nil, err
	}
	return bizWrapper.Order(bizQuery.AlarmGroup.ID.Desc()).Find()
}

func (a *alarmGroupRepositoryImpl) UpdateStatus(ctx context.Context, params *bo.UpdateAlarmGroupStatusParams) error {
	if len(params.IDs) == 0 {
		return nil
	}

	bizQuery, err := getBizQuery(ctx, a.data)
	if !types.IsNil(err) {
		return err
	}

	if _, err = bizQuery.AlarmGroup.WithContext(ctx).Where(bizQuery.AlarmGroup.ID.In(params.IDs...)).Update(bizQuery.AlarmGroup.Status, params.Status); !types.IsNil(err) {
		return err
	}
	return nil
}

// convert bo params to model
func createAlarmGroupParamsToModel(ctx context.Context, params *bo.CreateAlarmGroupParams) *bizmodel.AlarmGroup {
	alarmGroup := &bizmodel.AlarmGroup{
		Name:   params.Name,
		Status: params.Status,
		Remark: params.Remark,
		// TODO 添加hook关联信息
	}
	alarmGroup.WithContext(ctx)
	return alarmGroup
}

func createAlarmNoticeUsersToModel(ctx context.Context, params []*bo.CreateNoticeUserParams, alarmGroupID uint32) []*bizmodel.AlarmNoticeUser {
	return types.SliceToWithFilter(params, func(noticeUser *bo.CreateNoticeUserParams) (*bizmodel.AlarmNoticeUser, bool) {
		if noticeUser.UserId <= 0 {
			return nil, false
		}
		resUser := &bizmodel.AlarmNoticeUser{
			AlarmNoticeType: noticeUser.NotifyType,
			UserId:          noticeUser.UserId,
			AlarmGroupID:    alarmGroupID,
		}
		resUser.WithContext(ctx)
		return resUser, true
	})
}
