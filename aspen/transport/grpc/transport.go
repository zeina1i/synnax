package grpc

import (
	"context"
	"github.com/synnaxlabs/aspen/internal/cluster/gossip"
	"github.com/synnaxlabs/aspen/internal/cluster/pledge"
	"github.com/synnaxlabs/aspen/internal/kv"
	"github.com/synnaxlabs/aspen/internal/node"
	"github.com/synnaxlabs/aspen/transport"
	aspenv1 "github.com/synnaxlabs/aspen/transport/grpc/gen/proto/go/v1"
	"github.com/synnaxlabs/freighter/fgrpc"
	"github.com/synnaxlabs/x/address"
	"github.com/synnaxlabs/x/signal"
	"go/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type (
	pledgeClient = fgrpc.UnaryClient[
		node.ID,
		*aspenv1.ClusterPledge,
		node.ID,
		*aspenv1.ClusterPledge,
	]
	pledgeServer = fgrpc.UnaryServer[
		node.ID,
		*aspenv1.ClusterPledge,
		node.ID,
		*aspenv1.ClusterPledge,
	]
	clusterGossipClient = fgrpc.UnaryClient[
		gossip.Message,
		*aspenv1.ClusterGossip,
		gossip.Message,
		*aspenv1.ClusterGossip,
	]
	clusterGossipServer = fgrpc.UnaryServer[
		gossip.Message,
		*aspenv1.ClusterGossip,
		gossip.Message,
		*aspenv1.ClusterGossip,
	]
	batchClient = fgrpc.UnaryClient[
		kv.BatchRequest,
		*aspenv1.BatchRequest,
		kv.BatchRequest,
		*aspenv1.BatchRequest,
	]
	batchServer = fgrpc.UnaryServer[
		kv.BatchRequest,
		*aspenv1.BatchRequest,
		kv.BatchRequest,
		*aspenv1.BatchRequest,
	]
	leaseClient = fgrpc.UnaryClient[
		kv.BatchRequest,
		*aspenv1.BatchRequest,
		types.Nil,
		*emptypb.Empty,
	]
	leaseServer = fgrpc.UnaryServer[
		kv.BatchRequest,
		*aspenv1.BatchRequest,
		types.Nil,
		*emptypb.Empty,
	]
	feedbackClient = fgrpc.UnaryClient[
		kv.FeedbackMessage,
		*aspenv1.FeedbackMessage,
		types.Nil,
		*emptypb.Empty,
	]
	feedbackServer = fgrpc.UnaryServer[
		kv.FeedbackMessage,
		*aspenv1.FeedbackMessage,
		types.Nil,
		*emptypb.Empty,
	]
)

var (
	_ pledge.TransportServer             = (*pledgeServer)(nil)
	_ pledge.TransportClient             = (*pledgeClient)(nil)
	_ aspenv1.PledgeServiceServer        = (*pledgeServer)(nil)
	_ gossip.TransportClient             = (*clusterGossipClient)(nil)
	_ gossip.TransportServer             = (*clusterGossipServer)(nil)
	_ aspenv1.ClusterGossipServiceServer = (*clusterGossipServer)(nil)
	_ kv.BatchTransportClient            = (*batchClient)(nil)
	_ kv.BatchTransportServer            = (*batchServer)(nil)
	_ aspenv1.BatchServiceServer         = (*batchServer)(nil)
	_ kv.LeaseTransportClient            = (*leaseClient)(nil)
	_ kv.LeaseTransportServer            = (*leaseServer)(nil)
	_ aspenv1.LeaseServiceServer         = (*leaseServer)(nil)
	_ kv.FeedbackTransportClient         = (*feedbackClient)(nil)
	_ kv.FeedbackTransportServer         = (*feedbackServer)(nil)
	_ aspenv1.FeedbackServiceServer      = (*feedbackServer)(nil)
)

func New(pool *fgrpc.Pool) transport.Transport {
	return &transportImpl{
		pledgeClient: &pledgeClient{
			Pool:               pool,
			RequestTranslator:  pledgeTranslator{},
			ResponseTranslator: pledgeTranslator{},
			Client: func(
				ctx context.Context,
				conn grpc.ClientConnInterface,
				req *aspenv1.ClusterPledge,
			) (*aspenv1.ClusterPledge, error) {
				return aspenv1.NewPledgeServiceClient(conn).Exec(ctx, req)
			},
		},
		pledgeServer: &pledgeServer{
			RequestTranslator:  pledgeTranslator{},
			ResponseTranslator: pledgeTranslator{},
			ServiceDesc:        &aspenv1.PledgeService_ServiceDesc,
		},
		gossipClient: &clusterGossipClient{
			Pool:               pool,
			RequestTranslator:  clusterGossipTranslator{},
			ResponseTranslator: clusterGossipTranslator{},
			Client: func(
				ctx context.Context,
				conn grpc.ClientConnInterface,
				req *aspenv1.ClusterGossip,
			) (*aspenv1.ClusterGossip, error) {
				return aspenv1.NewClusterGossipServiceClient(conn).Exec(ctx, req)
			},
		},
		gossipServer: &clusterGossipServer{
			RequestTranslator:  clusterGossipTranslator{},
			ResponseTranslator: clusterGossipTranslator{},
			ServiceDesc:        &aspenv1.ClusterGossipService_ServiceDesc,
		},
		batchClient: &batchClient{
			Pool:               pool,
			RequestTranslator:  batchTranslator{},
			ResponseTranslator: batchTranslator{},
			Client: func(
				ctx context.Context,
				conn grpc.ClientConnInterface,
				req *aspenv1.BatchRequest,
			) (*aspenv1.BatchRequest, error) {
				return aspenv1.NewBatchServiceClient(conn).Exec(ctx, req)
			},
		},
		batchServer: &batchServer{
			RequestTranslator:  batchTranslator{},
			ResponseTranslator: batchTranslator{},
			ServiceDesc:        &aspenv1.BatchService_ServiceDesc,
		},
		leaseClient: &leaseClient{
			Pool:               pool,
			RequestTranslator:  batchTranslator{},
			ResponseTranslator: fgrpc.EmptyTranslator{},
			Client: func(
				ctx context.Context,
				conn grpc.ClientConnInterface,
				req *aspenv1.BatchRequest,
			) (*emptypb.Empty, error) {
				return aspenv1.NewLeaseServiceClient(conn).Exec(ctx, req)
			},
		},
		leaseServer: &leaseServer{
			RequestTranslator:  batchTranslator{},
			ResponseTranslator: fgrpc.EmptyTranslator{},
			ServiceDesc:        &aspenv1.LeaseService_ServiceDesc,
		},
		feedbackClient: &feedbackClient{
			Pool:               pool,
			RequestTranslator:  feedbackTranslator{},
			ResponseTranslator: fgrpc.EmptyTranslator{},
			Client: func(
				ctx context.Context,
				conn grpc.ClientConnInterface,
				req *aspenv1.FeedbackMessage,
			) (*emptypb.Empty, error) {
				return aspenv1.NewFeedbackServiceClient(conn).Exec(ctx, req)
			},
		},
		feedbackServer: &feedbackServer{
			RequestTranslator:  feedbackTranslator{},
			ResponseTranslator: fgrpc.EmptyTranslator{},
			ServiceDesc:        &aspenv1.FeedbackService_ServiceDesc,
		},
	}
}

// transportImpl implements the aspen.transportImpl interface.
type transportImpl struct {
	pledgeServer   *pledgeServer
	pledgeClient   *pledgeClient
	gossipServer   *clusterGossipServer
	gossipClient   *clusterGossipClient
	batchServer    *batchServer
	batchClient    *batchClient
	leaseServer    *leaseServer
	leaseClient    *leaseClient
	feedbackServer *feedbackServer
	feedbackClient *feedbackClient
}

func (t *transportImpl) PledgeServer() pledge.TransportServer {
	return t.pledgeServer
}

func (t *transportImpl) PledgeClient() pledge.TransportClient {
	return t.pledgeClient
}

func (t *transportImpl) GossipServer() gossip.TransportServer {
	return t.gossipServer
}

func (t *transportImpl) GossipClient() gossip.TransportClient {
	return t.gossipClient
}

func (t *transportImpl) BatchServer() kv.BatchTransportServer {
	return t.batchServer
}

func (t *transportImpl) BatchClient() kv.BatchTransportClient {
	return t.batchClient
}

func (t *transportImpl) LeaseServer() kv.LeaseTransportServer {
	return t.leaseServer
}

func (t *transportImpl) LeaseClient() kv.LeaseTransportClient {
	return t.leaseClient
}

func (t *transportImpl) FeedbackServer() kv.FeedbackTransportServer {
	return t.feedbackServer
}

func (t *transportImpl) FeedbackClient() kv.FeedbackTransportClient {
	return t.feedbackClient
}

func (t *transportImpl) BindTo(reg grpc.ServiceRegistrar) {
	t.pledgeServer.BindTo(reg)
	t.gossipServer.BindTo(reg)
	t.batchServer.BindTo(reg)
	t.leaseServer.BindTo(reg)
	t.feedbackServer.BindTo(reg)
}

func (t *transportImpl) Configure(ctx signal.Context, addr address.Address, external bool) error {
	if external {
		return nil
	}
	server := grpc.NewServer()
	t.BindTo(server)
	lis, err := net.Listen("tcp", addr.String())
	if err != nil {
		return err
	}
	ctx.Go(func(ctx context.Context) (err error) {
		go func() { err = server.Serve(lis) }()
		if err != nil {
			return err
		}
		defer server.Stop()
		<-ctx.Done()
		return ctx.Err()
	})
	return nil
}
