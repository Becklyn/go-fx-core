package health

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_IsHealthy_ReturnsHealthyByDefault(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	healthy, _ := serviceHealth.IsHealthy()
	assert.True(t, healthy) // must be healthy by default
}

func Test_IsHealthy_ReturnsNoReasonByDefault(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	_, reason := serviceHealth.IsHealthy()
	assert.Equal(t, "", reason) // must have no reason by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsHealthy(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	healthy, _ := serviceHealth.IsHealthy("foo")
	assert.True(t, healthy) // must be healthy by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsNoReason(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	_, reason := serviceHealth.IsHealthy("foo")
	assert.Equal(t, "", reason) // must have no reason by default
}

func Test_IsHealthy_ReturnsUnhealthy_IfAnyComponentIsUnhealthy(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")
	serviceHealth.SetHealthy("bar")

	healthy, _ := serviceHealth.IsHealthy()
	assert.True(t, !healthy) // must be unhealthy
}

func Test_IsHealthy_ForAHealthyComponent_ReturnsHealthy(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	serviceHealth.SetHealthy("foo")

	healthy, _ := serviceHealth.IsHealthy("foo")
	assert.True(t, healthy) // must be healthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsUnhealthy(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")

	healthy, _ := serviceHealth.IsHealthy("foo")
	assert.True(t, !healthy) // must be unhealthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsReason(t *testing.T) {
	logger := logrus.New()
	serviceHealth := newServiceHealth(logger)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")

	_, reason := serviceHealth.IsHealthy("foo")
	assert.Equal(t, "can't foo without bar", reason) // must be the reason of foo
}
