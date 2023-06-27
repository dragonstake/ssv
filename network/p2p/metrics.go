package p2pv1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bloxapp/ssv/logging/fields"
	"github.com/bloxapp/ssv/network/records"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"

	"github.com/bloxapp/ssv/utils/format"
)

var (
	// MetricsAllConnectedPeers counts all connected peers
	MetricsAllConnectedPeers = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "ssv_p2p_all_connected_peers",
		Help: "Count connected peers",
	})
	// MetricsConnectedPeers counts connected peers for a topic
	MetricsConnectedPeers = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ssv_p2p_connected_peers",
		Help: "Count connected peers for a validator",
	}, []string{"pubKey"})
	// MetricsPeersIdentity tracks peers identity
	MetricsPeersIdentity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ssv:network:peers_identity",
		Help: "Peers identity",
	}, []string{"pubKey", "operatorID", "v", "pid", "type"})
	metricsRouterIncoming = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ssv:network:router:in",
		Help: "Counts incoming messages",
	}, []string{"identifier", "mt"})
)

func init() {
	if err := prometheus.Register(MetricsAllConnectedPeers); err != nil {
		log.Println("could not register prometheus collector")
	}
	if err := prometheus.Register(MetricsPeersIdentity); err != nil {
		log.Println("could not register prometheus collector")
	}
	if err := prometheus.Register(MetricsConnectedPeers); err != nil {
		log.Println("could not register prometheus collector")
	}
	if err := prometheus.Register(metricsRouterIncoming); err != nil {
		log.Println("could not register prometheus collector")
	}
}

var unknown = "unknown"

func (n *p2pNetwork) reportAllPeers(logger *zap.Logger) func() {
	return func() {
		pids := n.host.Network().Peers()
		logger.Debug("connected peers status", fields.Count(len(pids)))
		MetricsAllConnectedPeers.Set(float64(len(pids)))
	}
}

func (n *p2pNetwork) reportPeerIdentities(logger *zap.Logger) func() {
	return func() {
		pids := n.host.Network().Peers()
		for _, pid := range pids {
			n.reportPeerIdentity(logger, pid)
		}
	}
}

func (n *p2pNetwork) reportTopics(logger *zap.Logger) func() {
	return func() {
		topics := n.topicsCtrl.Topics()
		nTopics := len(topics)
		logger.Debug("connected topics", fields.Count(nTopics))
		for _, name := range topics {
			n.reportTopicPeers(logger, name)
		}
	}
}

func namedPeers(ids []peer.ID) []string {
	var known = map[string]string{
		"16Uiu2HAkzXdRuiGgkMV7A4GdftAdPb5jESQLfrPbCJtG8CK3zxDf": "node-4",
		"16Uiu2HAm8Gs8C13HmmWq7QH4hWYo6R4vvBNK4pr3iaBrvNjzzfup": "node-1",
		"16Uiu2HAmLog6qkUjnmSnJMGhPQwpHU7ohgiPYW5ffqyRJGHbW9kM": "node-3",
		"16Uiu2HAkyZhiBL9RWuZwSSrzV8sSEF7J7eZcWangvEvtKF4MeoBN": "node-2",
		"16Uiu2HAkwLUMbQ7q8jpQMGcALrLFM5R1F3kWSPLDxSzEfPcmcM7D": "full-node-1",
	}

	names := make([]string, len(ids))
	for i, id := range ids {
		name, ok := known[id.String()]
		if !ok {
			names[i] = id.String()
		} else {
			names[i] = name
		}
	}
	return names
}

func (n *p2pNetwork) reportTopicPeers(logger *zap.Logger, name string) {
	peers, err := n.topicsCtrl.Peers(name)
	if err != nil {
		logger.Warn("could not get topic peers", fields.Topic(name), zap.Error(err))
		return
	}
	allPeers := n.host.Network().Peers()

	var errs []string
	subnet, err := strconv.Atoi(name)
	if err != nil {
		errs = append(errs, fmt.Sprintf("could not convert topic name to subnet: %s", err.Error()))
	}
	getSubnetPeers := n.idx.GetSubnetPeers(subnet)

	var metadataPeers []peer.ID
	for _, p := range allPeers {
		n, err := n.idx.GetNodeInfo(p)
		if err != nil {
			errs = append(errs, fmt.Sprintf("could not get node info (peer id: %s): %s", p, err.Error()))
			continue
		}
		subnets, err := records.Subnets{}.FromString(n.Metadata.Subnets)
		if err != nil {
			errs = append(errs, fmt.Sprintf("could not parse subnets (peer id: %s): %s", p, err.Error()))
			continue
		}
		if subnets[subnet] == 1 {
			metadataPeers = append(metadataPeers, p)
		}
	}

	logger.Debug("topic peers status",
		fields.Topic(name),
		fields.Count(len(peers)),
		zap.Any("peers", namedPeers(peers)),
		zap.Int("subnet", subnet),

		zap.Int("get_subnet_peers_count", len(getSubnetPeers)),
		zap.Any("get_subnet_peers", namedPeers(getSubnetPeers)),

		zap.Int("metadata_subnet_peers_count", len(metadataPeers)),
		zap.Any("metadata_subnet_peers", namedPeers(metadataPeers)),

		zap.Int("all_peers_count", len(allPeers)),
		zap.Any("all_peers", namedPeers(allPeers)),

		zap.Strings("errors", errs),
	)
	MetricsConnectedPeers.WithLabelValues(name).Set(float64(len(peers)))
}

func (n *p2pNetwork) reportPeerIdentity(logger *zap.Logger, pid peer.ID) {
	opPKHash, opIndex, forkv, nodeVersion, nodeType := unknown, unknown, unknown, unknown, unknown
	ni, err := n.idx.GetNodeInfo(pid)
	if err == nil && ni != nil {
		opPKHash = unknown
		nodeVersion = unknown
		forkv = ni.ForkVersion.String()
		if ni.Metadata != nil {
			opPKHash = ni.Metadata.OperatorID
			nodeVersion = ni.Metadata.NodeVersion
		}
		nodeType = "operator"
		if len(opPKHash) == 0 && nodeVersion != unknown {
			nodeType = "exporter"
		}
	}

	if pubKey, ok := n.operatorPKCache.Load(opPKHash); ok {
		operatorData, found, opDataErr := n.nodeStorage.GetOperatorDataByPubKey(logger, pubKey.([]byte))
		if opDataErr == nil && found {
			opIndex = strconv.FormatUint(operatorData.ID, 10)
		}
	} else {
		operators, err := n.nodeStorage.ListOperators(logger, 0, 0)
		if err != nil {
			logger.Warn("failed to get all operators for reporting", zap.Error(err))
		}

		for _, operator := range operators {
			pubKeyHash := format.OperatorID(operator.PublicKey)
			n.operatorPKCache.Store(pubKeyHash, operator.PublicKey)
			if pubKeyHash == opPKHash {
				opIndex = strconv.FormatUint(operator.ID, 10)
			}
		}
	}

	nodeState := n.idx.State(pid)
	logger.Debug("peer identity",
		fields.PeerID(pid),
		zap.String("forkv", forkv),
		zap.String("nodeVersion", nodeVersion),
		zap.String("opPKHash", opPKHash),
		zap.String("opIndex", opIndex),
		zap.String("nodeType", nodeType),
		zap.String("nodeState", nodeState.String()),
	)
	MetricsPeersIdentity.WithLabelValues(opPKHash, opIndex, nodeVersion, pid.String(), nodeType).Set(1)
}

//
// func reportLastMsg(pid string) {
//	MetricsPeerLastMsg.WithLabelValues(pid).Set(float64(timestamp()))
//}
//
// func timestamp() int64 {
//	return time.Now().UnixNano() / int64(time.Millisecond)
//}
