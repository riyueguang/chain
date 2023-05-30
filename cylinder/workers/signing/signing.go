package signing

import (
	"fmt"

	"github.com/bandprotocol/chain/v2/cylinder"
	"github.com/bandprotocol/chain/v2/cylinder/client"
	"github.com/bandprotocol/chain/v2/pkg/event"
	"github.com/bandprotocol/chain/v2/pkg/logger"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Signing is a worker responsible for signing process of TSS module
type Signing struct {
	context *cylinder.Context

	logger *logger.Logger
	client *client.Client

	eventCh <-chan ctypes.ResultEvent
}

var _ cylinder.Worker = &Signing{}

// New creates a new instance of the Signing workes.
// It initializes the necessary components and returns the created Signing instance or an error if initialization fails.
func New(ctx *cylinder.Context) (*Signing, error) {
	cli, err := client.New(ctx.Config, ctx.Keyring)
	if err != nil {
		return nil, err
	}

	return &Signing{
		context: ctx,
		logger:  ctx.Logger.With("worker", "signing"),
		client:  cli,
	}, nil
}

// subscribe subscribes to the round2 events and initializes the event channel for receiving events.
// It returns an error if the subscription fails.
// TODO-CYLINDER: subscribe the real event
func (s *Signing) subscribe() error {
	var err error
	s.eventCh, err = s.client.Subscribe(
		"signing",
		fmt.Sprintf(
			"tm.event = 'Tx' AND %s.%s = '%s'",
			types.EventTypeCreateGroup,
			types.AttributeKeyMember,
			s.context.Config.Granter,
		),
		1000,
	)
	return err
}

// handleTxResult handles the result of a transaction.
// It extracts the relevant message logs from the transaction result and processes the events.
func (s *Signing) handleTxResult(txResult abci.TxResult) {
	msgLogs, err := event.GetMessageLogs(txResult)
	if err != nil {
		s.logger.Error("Failed to get message logs: %s", err.Error())
		return
	}

	for _, log := range msgLogs {
		event, err := ParseEvent(log)
		if err != nil {
			s.logger.Error(":cold_sweat: Failed to parse event with error: %s", err.Error())
			return
		}

		go s.handleSigning(
			event.GroupID,
			event.SigningID,
			event.MemberIDs,
			event.Data,
			event.Bytes,
			event.GroupPubNonce,
			event.PubDE,
		)
	}
}

// handleGroup processes an incoming group.
// TODO-CYLINDER: use the real value
func (s *Signing) handleSigning(
	gid tss.GroupID,
	sid tss.GroupID,
	mids []tss.MemberID,
	data []byte,
	bytes []byte,
	groupPubNonce tss.PublicKey,
	pubDE types.DE,
) {
	logger := s.logger.With("gid", gid).With("sid", sid)

	// Log
	logger.Info(":delivery_truck: Processing incoming group")

	// Set group data
	group, err := s.context.Store.GetGroup(gid)
	if err != nil {
		logger.Error(":cold_sweat: Failed to find group in store: %s", err.Error())
		return
	}

	// Get private keys of DE
	privDE, err := s.context.Store.GetDE(pubDE)
	if err != nil {
		logger.Error(":cold_sweat: Failed to private DE from store: %s", err.Error())
		return
	}

	// Compute Lo value
	lo := tss.ComputeOwnLo(group.MemberID, data, bytes)

	// Compute own private nonce
	privNonce, err := tss.ComputeOwnPrivateNonce(privDE.PrivD, privDE.PrivE, lo)
	if err != nil {
		logger.Error(":cold_sweat: Failed to compute private nonce: %s", err.Error())
		return
	}

	// Compute lagrange
	lagrange := tss.ComputeLagrangeCoefficient(group.MemberID, mids)

	// Sign the singing
	sig, err := tss.SignSigning(groupPubNonce, group.PubKey, data, lagrange, privNonce, group.PrivKey)
	if err != nil {
		logger.Error(":cold_sweat: Failed to sign signing: %s", err.Error())
		return
	}

	// Send MsgSigning
	// TODO-CYLINDER: use the correct message
	s.context.MsgCh <- &types.MsgConfirm{
		MemberID:     group.MemberID,
		OwnPubKeySig: sig,
		Member:       s.context.Config.Granter,
	}
}

// Start starts the Signing worker.
// It subscribes to signing request events and starts processing incoming events.
func (s *Signing) Start() {
	s.logger.Info("start")

	err := s.subscribe()
	if err != nil {
		s.context.ErrCh <- err
		return
	}

	for ev := range s.eventCh {
		go s.handleTxResult(ev.Data.(tmtypes.EventDataTx).TxResult)
	}
}

// Stop stops the Signing workes.
func (s *Signing) Stop() {
	s.logger.Info("stop")
	s.client.Stop()
}
