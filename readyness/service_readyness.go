package readyness

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type ServiceReadyness struct {
	components map[string]bool

	logger *logrus.Logger
	mux    sync.RWMutex
}

func NewServiceReadyness(logger *logrus.Logger) *ServiceReadyness {
	return &ServiceReadyness{
		components: make(map[string]bool),
		logger:     logger,
	}
}

func (s *ServiceReadyness) isReady() (bool, string) {
	for component, isReady := range s.components {
		if !isReady {
			return false, component
		}
	}
	return true, ""
}

func (s *ServiceReadyness) IsReady(component ...string) (bool, string) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if len(component) == 0 {
		return s.isReady()
	}

	for _, c := range component {
		if ready, ok := s.components[c]; !ok || !ready {
			return false, c
		}
	}
	return true, ""
}

func (s *ServiceReadyness) SetReady(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.components[component] = true

	s.logger.WithFields(logrus.Fields{
		"component": component,
	}).Warn("Component is ready")
}

func (s *ServiceReadyness) Register(component string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.components[component] = false

	s.logger.WithFields(logrus.Fields{
		"component": component,
	}).Warn("Registered unready component")
}
