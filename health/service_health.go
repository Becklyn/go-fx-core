package health

import "sync"

type ServiceHealth struct {
	components       map[string]componentHealth
	componentChanged chan componentChanged

	mux sync.RWMutex
}

func newServiceHealth() *ServiceHealth {
	return &ServiceHealth{
		components:       make(map[string]componentHealth),
		componentChanged: make(chan componentChanged),
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
		if !s.components[c].healthy {
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

	s.componentChanged <- componentChanged{
		name:   component,
		health: componentHealth,
	}
}

func (s *ServiceHealth) SetUnhealthy(component string, reason string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	componentHealth := componentHealth{
		healthy: false,
		reason:  reason,
	}
	s.components[component] = componentHealth

	s.componentChanged <- componentChanged{
		name:   component,
		health: componentHealth,
	}
}
