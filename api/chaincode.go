package api

import (
	"context"

	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/msp"
)

type ChaincodeTx string

type TransArgs map[string][]byte

// Chaincode describes common operations with chaincode
type Chaincode interface {
	// GetPeers returns chaincode peers
	GetPeers() []Peer
	// Invoke returns invoke builder for presented chaincode function
	Invoke(fn string) ChaincodeInvokeBuilder
	// Query returns query builder for presented function and arguments
	Query(fn string, args ...string) ChaincodeQueryBuilder
	// Deprecated: Install fetches chaincode from repository and installs it on local peer
	Install(version string)
	// Subscribe returns subscription on chaincode events
	Subscribe(ctx context.Context) (EventCCSubscription, error)
}

type ChaincodePackage interface {
	// Latest allows to get latest version of chaincode
	Latest(ctx context.Context) (*peer.ChaincodeDeploymentSpec, error)
	// Install chaincode using defined chaincode fetcher
	Install(ctx context.Context, path, version string) error
	// Instantiate chaincode on channel with presented params
	Instantiate(ctx context.Context, channelName, path, version, policy string, args [][]byte, transArgs TransArgs) error
}

type ChaincodeInvokeResponse struct {
	TxID    ChaincodeTx
	Payload []byte
	Err     error
}

// TxWaiter is interface for build your custom function for wait of result of tx after endorsement
type TxWaiter interface {
	Wait(ctx context.Context, channel string, txId ChaincodeTx) error
}

type DoOptions struct {
	Identity msp.SigningIdentity
	Pool     PeerPool

	TxWaiter TxWaiter
	// necessary only for 'tx waiter all'
	EndorsingMspIDs []string
}

type DoOption func(opt *DoOptions) error

func WithEndorsingMpsIDs(mspIDs []string) DoOption {
	return func(opt *DoOptions) error {
		opt.EndorsingMspIDs = mspIDs

		return nil
	}
}

func WithIdentity(identity msp.SigningIdentity) DoOption {
	return func(opt *DoOptions) error {
		opt.Identity = identity

		return nil
	}
}

// ChaincodeInvokeBuilder describes possibilities how to get invoke results
type ChaincodeInvokeBuilder interface {
	// WithIdentity allows invoking chaincode from custom identity
	WithIdentity(identity msp.SigningIdentity) ChaincodeInvokeBuilder
	// Transient allows passing arguments to transient map
	Transient(args TransArgs) ChaincodeInvokeBuilder
	// ArgBytes set slice of bytes as argument
	ArgBytes([][]byte) ChaincodeInvokeBuilder
	// ArgJSON set slice of JSON-marshalled data
	ArgJSON(in ...interface{}) ChaincodeInvokeBuilder
	// ArgString set slice of strings as arguments
	ArgString(args ...string) ChaincodeInvokeBuilder
	// Do makes invoke with built arguments
	Do(ctx context.Context, opts ...DoOption) (*peer.Response, ChaincodeTx, error)
}

// ChaincodeQueryBuilder describe possibilities how to get query results
type ChaincodeQueryBuilder interface {
	// WithIdentity allows invoking chaincode from custom identity
	WithIdentity(identity msp.SigningIdentity) ChaincodeQueryBuilder
	// WithArguments allows querying chaincode with arguments
	WithArguments(argBytes [][]byte) ChaincodeQueryBuilder
	// Transient allows passing arguments to transient map
	Transient(args TransArgs) ChaincodeQueryBuilder
	// AsBytes allows getting result of querying chaincode as byte slice
	AsBytes(ctx context.Context) ([]byte, error)
	// AsJSON allows getting result of querying chaincode to presented structures using JSON-unmarshalling
	AsJSON(ctx context.Context, out interface{}) error
	// AsProposalResponse allows getting raw peer response
	AsProposalResponse(ctx context.Context) (*peer.ProposalResponse, error)
	// Do makes query with built arguments
	Do(ctx context.Context) (*peer.Response, error)
}

type CCFetcher interface {
	Fetch(ctx context.Context, id *peer.ChaincodeID) (*peer.ChaincodeDeploymentSpec, error)
}
