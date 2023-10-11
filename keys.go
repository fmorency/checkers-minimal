package checkers

import "cosmossdk.io/collections"

const ModuleName = "checkers"

var (
    ParamsKey  = collections.NewPrefix(0)
    StoredGameListKey = collections.NewPrefix(1)
)
