package app

import (
	"time"

	tmcfg "github.com/tendermint/tendermint/config"
)

func SetTendermintConfigs(config *tmcfg.Config) {
	// Peer Configs
	config.P2P.MaxNumInboundPeers = 100
	config.P2P.MaxNumOutboundPeers = 100
	config.P2P.SendRate = 20480000
	config.P2P.RecvRate = 20480000
	config.P2P.MaxPacketMsgPayloadSize = 10240
	config.P2P.FlushThrottleTimeout = 10 * time.Millisecond
	// Consensus Configs
	config.Consensus.TimeoutPrevote = 250 * time.Millisecond
	config.Consensus.TimeoutPrecommit = 250 * time.Millisecond
	config.Consensus.TimeoutCommit = 250 * time.Millisecond
	config.Consensus.SkipTimeoutCommit = true
}
