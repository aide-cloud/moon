package build

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
)

type StrategyBuilder struct {
	*bizmodel.Strategy
}

func NewStrategyBuilder(strategy *bizmodel.Strategy) *StrategyBuilder {
	return &StrategyBuilder{
		Strategy: strategy,
	}
}

// ToApi 转换为API层数据
func (b *StrategyBuilder) ToApi(ctx context.Context) *admin.Strategy {
	if types.IsNil(b) || types.IsNil(b.Strategy) {
		return nil
	}
	return &admin.Strategy{
		Id:    b.ID,
		Alert: b.Alert,
		Expr:  b.Expr,
		//Labels: b.Labels,
		Annotations: b.Annotations,
		Datasource: types.SliceTo(b.Datasource, func(datasource *bizmodel.Datasource) *admin.Datasource {
			return NewDatasourceBuilder(datasource).ToApi(ctx)
		}),
		StrategyTemplateId: b.StrategyTemplateID,
		Threshold:          b.Threshold,
		Condition:          b.Condition,
	}
}
