// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package bizquery

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	CasbinRule           *casbinRule
	Datasource           *datasource
	DatasourceLabelValue *datasourceLabelValue
	DatasourceMetric     *datasourceMetric
	MetricLabel          *metricLabel
	SysTeamAPI           *sysTeamAPI
	SysTeamMember        *sysTeamMember
	SysTeamMemberRole    *sysTeamMemberRole
	SysTeamMenu          *sysTeamMenu
	SysTeamRole          *sysTeamRole
	SysTeamRoleAPI       *sysTeamRoleAPI
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CasbinRule = &Q.CasbinRule
	Datasource = &Q.Datasource
	DatasourceLabelValue = &Q.DatasourceLabelValue
	DatasourceMetric = &Q.DatasourceMetric
	MetricLabel = &Q.MetricLabel
	SysTeamAPI = &Q.SysTeamAPI
	SysTeamMember = &Q.SysTeamMember
	SysTeamMemberRole = &Q.SysTeamMemberRole
	SysTeamMenu = &Q.SysTeamMenu
	SysTeamRole = &Q.SysTeamRole
	SysTeamRoleAPI = &Q.SysTeamRoleAPI
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		CasbinRule:           newCasbinRule(db, opts...),
		Datasource:           newDatasource(db, opts...),
		DatasourceLabelValue: newDatasourceLabelValue(db, opts...),
		DatasourceMetric:     newDatasourceMetric(db, opts...),
		MetricLabel:          newMetricLabel(db, opts...),
		SysTeamAPI:           newSysTeamAPI(db, opts...),
		SysTeamMember:        newSysTeamMember(db, opts...),
		SysTeamMemberRole:    newSysTeamMemberRole(db, opts...),
		SysTeamMenu:          newSysTeamMenu(db, opts...),
		SysTeamRole:          newSysTeamRole(db, opts...),
		SysTeamRoleAPI:       newSysTeamRoleAPI(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CasbinRule           casbinRule
	Datasource           datasource
	DatasourceLabelValue datasourceLabelValue
	DatasourceMetric     datasourceMetric
	MetricLabel          metricLabel
	SysTeamAPI           sysTeamAPI
	SysTeamMember        sysTeamMember
	SysTeamMemberRole    sysTeamMemberRole
	SysTeamMenu          sysTeamMenu
	SysTeamRole          sysTeamRole
	SysTeamRoleAPI       sysTeamRoleAPI
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		CasbinRule:           q.CasbinRule.clone(db),
		Datasource:           q.Datasource.clone(db),
		DatasourceLabelValue: q.DatasourceLabelValue.clone(db),
		DatasourceMetric:     q.DatasourceMetric.clone(db),
		MetricLabel:          q.MetricLabel.clone(db),
		SysTeamAPI:           q.SysTeamAPI.clone(db),
		SysTeamMember:        q.SysTeamMember.clone(db),
		SysTeamMemberRole:    q.SysTeamMemberRole.clone(db),
		SysTeamMenu:          q.SysTeamMenu.clone(db),
		SysTeamRole:          q.SysTeamRole.clone(db),
		SysTeamRoleAPI:       q.SysTeamRoleAPI.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		CasbinRule:           q.CasbinRule.replaceDB(db),
		Datasource:           q.Datasource.replaceDB(db),
		DatasourceLabelValue: q.DatasourceLabelValue.replaceDB(db),
		DatasourceMetric:     q.DatasourceMetric.replaceDB(db),
		MetricLabel:          q.MetricLabel.replaceDB(db),
		SysTeamAPI:           q.SysTeamAPI.replaceDB(db),
		SysTeamMember:        q.SysTeamMember.replaceDB(db),
		SysTeamMemberRole:    q.SysTeamMemberRole.replaceDB(db),
		SysTeamMenu:          q.SysTeamMenu.replaceDB(db),
		SysTeamRole:          q.SysTeamRole.replaceDB(db),
		SysTeamRoleAPI:       q.SysTeamRoleAPI.replaceDB(db),
	}
}

type queryCtx struct {
	CasbinRule           ICasbinRuleDo
	Datasource           IDatasourceDo
	DatasourceLabelValue IDatasourceLabelValueDo
	DatasourceMetric     IDatasourceMetricDo
	MetricLabel          IMetricLabelDo
	SysTeamAPI           ISysTeamAPIDo
	SysTeamMember        ISysTeamMemberDo
	SysTeamMemberRole    ISysTeamMemberRoleDo
	SysTeamMenu          ISysTeamMenuDo
	SysTeamRole          ISysTeamRoleDo
	SysTeamRoleAPI       ISysTeamRoleAPIDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CasbinRule:           q.CasbinRule.WithContext(ctx),
		Datasource:           q.Datasource.WithContext(ctx),
		DatasourceLabelValue: q.DatasourceLabelValue.WithContext(ctx),
		DatasourceMetric:     q.DatasourceMetric.WithContext(ctx),
		MetricLabel:          q.MetricLabel.WithContext(ctx),
		SysTeamAPI:           q.SysTeamAPI.WithContext(ctx),
		SysTeamMember:        q.SysTeamMember.WithContext(ctx),
		SysTeamMemberRole:    q.SysTeamMemberRole.WithContext(ctx),
		SysTeamMenu:          q.SysTeamMenu.WithContext(ctx),
		SysTeamRole:          q.SysTeamRole.WithContext(ctx),
		SysTeamRoleAPI:       q.SysTeamRoleAPI.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
