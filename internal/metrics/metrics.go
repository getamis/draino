package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	NodesCordoned       = promauto.NewCounterVec(prometheus.CounterOpts{Name: "draino_cordoned_nodes_total", Help: "Number of nodes cordoned."}, []string{"result"})
	NodesUncordoned     = promauto.NewCounterVec(prometheus.CounterOpts{Name: "draino_uncordoned_nodes_total", Help: "Number of nodes uncordoned."}, []string{"result"})
	NodesDraining       = promauto.NewCounterVec(prometheus.CounterOpts{Name: "draino_draining_nodes", Help: "Number of nodes draining."}, []string{"result"})
	NodesDrained        = promauto.NewCounterVec(prometheus.CounterOpts{Name: "draino_drained_nodes_total", Help: "Number of nodes drained."}, []string{"result"})
	NodesDrainScheduled = promauto.NewCounterVec(prometheus.CounterOpts{Name: "draino_drain_scheduled_nodes_total", Help: "Number of nodes drain scheduled."}, []string{"result"})
)
