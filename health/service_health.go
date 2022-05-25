package health

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type ServiceHealth struct {
	components map[string]componentHealth

	logger *logrus.Logger
	mux    sync.RWMutex
}

func newServiceHealth(logger *logrus.Logger) *ServiceHealth {
	return &ServiceHealth{
		components: make(map[string]componentHealth),
		logger:     logger,
	}
}

func (s *ServiceHealth) isHealthy() (bool, string) {
	for _, component := range s.components {
		if !component.healthy {
			return false, component.reason
		}
	}
	return true, ""
}

func (s *ServiceHealth) IsHealthy(component ...string) (bool, string) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if len(component) == 0 {
		return s.isHealthy()
	}

	for _, c := range component {
		if ch, ok := s.components[c]; ok && !ch.healthy {
			return false, s.components[c].reason
		}
	}
	return true, ""
}

func (s *ServiceHealth) SetHealthy(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	componentHealth := componentHealth{
		healthy: true,
	}
	s.components[component] = componentHealth

	s.logger.WithFields(logrus.Fields{
		"component": component,
	}).Warn("Component is healthy (again)")
}

func (s *ServiceHealth) SetUnhealthy(component string, reason string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	componentHealth := componentHealth{
		healthy: false,
		reason:  reason,
	}
	s.components[component] = componentHealth

	s.logger.WithFields(logrus.Fields{
		"component": component,
		"reason":    reason,
	}).Warn("Component became unhealthy")
}
