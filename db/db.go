package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/bandprotocol/bandchain/chain/x/zoracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type BandDB struct {
	db  *gorm.DB
	tx  *gorm.DB
	ctx sdk.Context

	ZoracleKeeper zoracle.Keeper
}

func NewDB(dialect, path string, metadata map[string]string) (*BandDB, error) {
	db, err := gorm.Open(dialect, path)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&Metadata{},
		&Event{},
		&Validator{},
		&ValidatorVote{},
		&DataSource{},
		&DataSourceRevision{},
		&Transaction{},
	)

	db.Model(&ValidatorVote{}).AddForeignKey(
		"consensus_address",
		"validators(consensus_address)",
		"RESTRICT",
		"RESTRICT",
	)

	db.Model(&DataSourceRevision{}).AddForeignKey(
		"data_source_id",
		"data_sources(id)",
		"CASCADE",
		"CASCADE",
	)

	for key, value := range metadata {
		err := db.Where(Metadata{Key: key}).
			Assign(Metadata{Value: value}).
			FirstOrCreate(&Metadata{}).Error
		if err != nil {
			panic(err)
		}
	}

	return &BandDB{db: db}, nil
}

func (b *BandDB) BeginTransaction() {
	if b.tx != nil {
		panic("BeginTransaction: Cannot begin a new transaction without closing the pending one.")
	}
	b.tx = b.db.Begin()
	if b.tx.Error != nil {
		panic(b.tx.Error)
	}
}

func (b *BandDB) Commit() {
	err := b.tx.Commit().Error
	if err != nil {
		panic(err)
	}
	b.tx = nil
}

func (b *BandDB) RollBack() {
	err := b.tx.Rollback()
	if err != nil {
		panic(err)
	}
	b.tx = nil
}

func (b *BandDB) SetContext(ctx sdk.Context) {
	b.ctx = ctx
}

func (b *BandDB) HandleTransaction(tx auth.StdTx, txHash []byte, logs sdk.ABCIMessageLogs) {
	msgs := tx.GetMsgs()

	err := b.AddTransaction(
		txHash,
		b.ctx.BlockTime(),
		b.ctx.BlockGasMeter().GasConsumed(),
		b.ctx.BlockGasMeter().GasConsumedToLimit(),
		string(tx.Fee.Bytes()),
		tx.GetSigners()[0],
		true,
		b.ctx.BlockHeight(),
	)
	if err != nil {
		panic(err)
	}

	if len(msgs) != len(logs) {
		panic("Inconsistent size of msgs and logs.")
	}

	for idx, msg := range msgs {
		events := logs[idx].Events
		kvMap := make(map[string]string)
		for _, event := range events {
			for _, kv := range event.Attributes {
				kvMap[event.Type+"."+kv.Key] = kv.Value
			}
		}
		err := b.HandleMessage(txHash, msg, kvMap)
		if err != nil {
			panic(err)
		}
	}
}

func (b *BandDB) HandleMessage(txHash []byte, msg sdk.Msg, events map[string]string) error {
	switch msg := msg.(type) {
	// Just proof of concept
	case zoracle.MsgCreateDataSource:
		return b.handleMsgCreateDataSource(txHash, msg, events)
	case zoracle.MsgEditDataSource:
		return b.handleMsgEditDataSource(txHash, msg, events)
	default:
		// TODO: Better logging
		fmt.Println("HandleMessage: There isn't event handler for this type")
		return nil
	}
}
