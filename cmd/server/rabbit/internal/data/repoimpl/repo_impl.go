package repoimpl

import (
	"github.com/google/wire"
)

var ProviderSetRepoImpl = wire.NewSet(
	NewTransactionRepo,
	NewCacheRepo,
)
