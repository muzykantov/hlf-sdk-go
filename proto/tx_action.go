package proto

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/protoutil"
)

type (
	TransactionAction struct {
		Event                   *peer.ChaincodeEvent
		Endorsers               []*msp.SerializedIdentity
		ReadWriteSets           []*kvrwset.KVRWSet
		ChaincodeInvocationSpec *peer.ChaincodeInvocationSpec
	}

	TransactionsActions []*TransactionAction
)

func ParseTxActions(txActions []*peer.TransactionAction) ([]*TransactionAction, error) {
	var parsedTxActions []*TransactionAction

	for _, action := range txActions {
		txAction, err := ParseTxAction(action)
		if err != nil {
			return nil, fmt.Errorf(`tx action: %w`, err)
		}
		parsedTxActions = append(parsedTxActions, txAction)
	}

	return parsedTxActions, nil
}

func ParseTxAction(txAction *peer.TransactionAction) (*TransactionAction, error) {
	ccAction, err := ParseChaincodeAction(txAction)
	if err != nil {
		return nil, fmt.Errorf("parse transaction chaincode action: %w", err)
	}

	ccEvent, err := ParseTransactionActionEvents(ccAction)
	if err != nil {
		return nil, fmt.Errorf("parse transaction events: %w", err)
	}

	endorsers, err := ParseTransactionActionEndorsers(txAction)
	if err != nil {
		return nil, fmt.Errorf("parse transaction endorsers: %w", err)
	}

	rwSets, err := ParseTransactionActionReadWriteSet(ccAction)
	if err != nil {
		return nil, fmt.Errorf("parse transaction read/write sets: %w", err)
	}

	chaincodeInvocationSpec, err := ParseTransactionActionChaincode(txAction)
	if err != nil {
		return nil, fmt.Errorf("parse transaction chaincode invocation spec: %w", err)
	}

	parsedTxAction := &TransactionAction{
		Event:                   ccEvent,
		Endorsers:               endorsers,
		ReadWriteSets:           rwSets,
		ChaincodeInvocationSpec: chaincodeInvocationSpec,
	}

	return parsedTxAction, nil
}

func ParseChaincodeAction(txAction *peer.TransactionAction) (*peer.ChaincodeAction, error) {
	ccActionPayload, err := protoutil.UnmarshalChaincodeActionPayload(txAction.Payload)
	if err != nil {
		return nil, fmt.Errorf(`chaincode action payload: %w`, err)
	}

	proposalResponsePayload, err := protoutil.UnmarshalProposalResponsePayload(
		ccActionPayload.Action.ProposalResponsePayload)
	if err != nil {
		return nil, fmt.Errorf(`proposal response payload:  %w`, err)
	}

	chaincodeAction, err := protoutil.UnmarshalChaincodeAction(proposalResponsePayload.Extension)
	if err != nil {
		return nil, fmt.Errorf(`chaincode action from proposal response: %w`, err)
	}

	return chaincodeAction, nil
}

func ParseTransactionActionEvents(chaincodeAction *peer.ChaincodeAction) (*peer.ChaincodeEvent, error) {
	ccEvent, err := protoutil.UnmarshalChaincodeEvents(chaincodeAction.Events)
	if err != nil {
		return nil, fmt.Errorf(`event from chaincode action: %w`, err)
	}
	return ccEvent, nil
}

func ParseTransactionActionEndorsers(txAction *peer.TransactionAction) ([]*msp.SerializedIdentity, error) {
	ccActionPayload, err := protoutil.UnmarshalChaincodeActionPayload(txAction.Payload)
	if err != nil {
		return nil, fmt.Errorf(`chaincode action payload: %w`, err)
	}

	endorsers := make([]*msp.SerializedIdentity, 0)
	for _, en := range ccActionPayload.Action.Endorsements {
		endorser := &msp.SerializedIdentity{}

		if err := proto.Unmarshal(en.Endorser, endorser); err != nil {
			return nil, fmt.Errorf("failed to get endorser: %w", err)
		}

		endorsers = append(endorsers, endorser)
	}

	return endorsers, nil
}

func ParseTransactionActionReadWriteSet(chaincodeAction *peer.ChaincodeAction) ([]*kvrwset.KVRWSet, error) {
	txReadWriteSet := &rwset.TxReadWriteSet{}
	if err := proto.Unmarshal(chaincodeAction.Results, txReadWriteSet); err != nil {
		return nil, fmt.Errorf("failed to get txReadWriteSet: %w", err)
	}

	kvReadWriteSets := make([]*kvrwset.KVRWSet, 0)
	for _, rw := range txReadWriteSet.NsRwset {
		kvReadWriteSet := &kvrwset.KVRWSet{}
		if err := proto.Unmarshal(rw.Rwset, kvReadWriteSet); err != nil {
			return nil, fmt.Errorf("failed to get kvReadWriteSet: %w", err)
		}
		kvReadWriteSets = append(kvReadWriteSets, kvReadWriteSet)
	}

	return kvReadWriteSets, nil
}

func ParseTransactionActionChaincode(txAction *peer.TransactionAction) (*peer.ChaincodeInvocationSpec, error) {
	actionPayload, err := protoutil.UnmarshalChaincodeActionPayload(txAction.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get chaincode action from action payload: %w", err)
	}

	chaincodeProposalPayload, err := protoutil.UnmarshalChaincodeProposalPayload(actionPayload.ChaincodeProposalPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to get chaincode proposal from action payload: %w", err)
	}

	chaincodeInvocationSpec := &peer.ChaincodeInvocationSpec{}
	if err = proto.Unmarshal(chaincodeProposalPayload.Input, chaincodeInvocationSpec); err != nil {
		return nil, fmt.Errorf("failed to get chaincode invocation spec from action payload: %w", err)
	}
	return chaincodeInvocationSpec, nil
}
