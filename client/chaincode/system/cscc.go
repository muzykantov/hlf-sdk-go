package system

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"

	"github.com/s7techlab/hlf-sdk-go/api"
	"github.com/s7techlab/hlf-sdk-go/client/chaincode"
	hlfproto "github.com/s7techlab/hlf-sdk-go/proto"
)

//go:embed qscc.swagger.json
var CSCCServiceSwagger []byte

// These are function names from Invoke first parameter
const (
	JoinChain        string = "JoinChain"
	GetConfigBlock   string = "GetConfigBlock"
	GetChannels      string = "GetChannels"
	GetConfigTree    string = `GetConfigTree`    // HLF Peer V1.x
	GetChannelConfig string = "GetChannelConfig" // HLF Peer V2 +
)

type CSCCService struct {
	UnimplementedCSCCServiceServer

	Querier         *chaincode.ProtoQuerier
	ChannelsFetcher *ChannelsFetcher
	FabricVersion   hlfproto.FabricVersion
}

type ChannelsFetcher struct {
	Querier *chaincode.ProtoQuerier
}

func NewCSCCFromClient(client api.Core) *CSCCService {
	return NewCSCC(client, hlfproto.FabricVersionIsV2(client.FabricV2()))
}

func NewCSCC(querier api.Querier, version hlfproto.FabricVersion) *CSCCService {
	return &CSCCService{
		Querier:         chaincode.NewProtoQuerier(querier, ``, CSCCName),
		ChannelsFetcher: NewChannelsFetcher(querier),
		FabricVersion:   version,
	}
}

func NewChannelsFetcher(querier api.Querier) *ChannelsFetcher {
	return &ChannelsFetcher{
		Querier: chaincode.NewProtoQuerier(querier, ``, CSCCName),
	}
}

func (f *ChannelsFetcher) GetChannels(ctx context.Context) (*peer.ChannelQueryResponse, error) {
	res, err := f.Querier.QueryStringsProto(ctx, []string{GetChannels}, &peer.ChannelQueryResponse{})
	if err != nil {
		return nil, err
	}
	return res.(*peer.ChannelQueryResponse), nil
}

func (c *CSCCService) ServiceDef() ServiceDef {
	return NewServiceDef(
		_CSCCService_serviceDesc.ServiceName,
		CSCCServiceSwagger,
		&_CSCCService_serviceDesc,
		c,
		RegisterCSCCServiceHandlerFromEndpoint,
	)
}

func (c *CSCCService) GetChannels(ctx context.Context, _ *empty.Empty) (*peer.ChannelQueryResponse, error) {
	return c.ChannelsFetcher.GetChannels(ctx)
}

func (c *CSCCService) JoinChain(ctx context.Context, request *JoinChainRequest) (*empty.Empty, error) {
	blockBytes, err := proto.Marshal(request.GenesisBlock)
	if err != nil {
		return nil, fmt.Errorf("marshal genesis block: %w", err)
	}

	if _, err = c.Querier.Query(ctx, JoinChain, blockBytes); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *CSCCService) GetConfigBlock(ctx context.Context, request *GetConfigBlockRequest) (*common.Block, error) {
	res, err := c.Querier.QueryProto(ctx, []interface{}{GetConfigBlock, request.Channel}, &common.Block{})
	if err != nil {
		return nil, err
	}
	return res.(*common.Block), nil
}

func (c *CSCCService) GetChannelConfig(ctx context.Context, request *GetChannelConfigRequest) (*common.Config, error) {
	switch c.FabricVersion {

	case hlfproto.FabricV1:
		res, err := c.Querier.QueryStringsProto(ctx, []string{GetConfigTree, request.Channel}, &peer.ConfigTree{})
		if err != nil {
			return nil, err
		}
		return res.(*peer.ConfigTree).ChannelConfig, nil

	case hlfproto.FabricV2:

		res, err := c.Querier.QueryStringsProto(ctx, []string{GetChannelConfig, request.Channel}, &common.Config{})
		if err != nil {
			return nil, err
		}
		return res.(*common.Config), nil

	default:
		return nil, fmt.Errorf(`fabric version=%s: %w`, c.FabricVersion, hlfproto.ErrUnknownFabricVersion)
	}
}
