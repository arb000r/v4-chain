package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/indexer/indexer_manager"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/x/vault/types"
)

type (
	Keeper struct {
		cdc                 codec.BinaryCodec
		storeKey            storetypes.StoreKey
		clobKeeper          types.ClobKeeper
		perpetualsKeeper    types.PerpetualsKeeper
		pricesKeeper        types.PricesKeeper
		subaccountsKeeper   types.SubaccountsKeeper
		indexerEventManager indexer_manager.IndexerEventManager
		authorities         map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	clobKeeper types.ClobKeeper,
	perpetualsKeeper types.PerpetualsKeeper,
	pricesKeeper types.PricesKeeper,
	subaccountsKeeper types.SubaccountsKeeper,
	indexerEventManager indexer_manager.IndexerEventManager,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		clobKeeper:          clobKeeper,
		perpetualsKeeper:    perpetualsKeeper,
		pricesKeeper:        pricesKeeper,
		subaccountsKeeper:   subaccountsKeeper,
		indexerEventManager: indexerEventManager,
		authorities:         lib.UniqueSliceToSet(authorities),
	}
}

func (k Keeper) GetIndexerEventManager() indexer_manager.IndexerEventManager {
	return k.indexerEventManager
}

func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {}
