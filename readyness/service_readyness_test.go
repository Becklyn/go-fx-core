package readyness

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_IsReady_ReturnsHealthyByDefault(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	ready, _ := serviceReadyness.IsReady()
	assert.True(t, ready) // must be ready by default
}

func Test_IsReady_ReturnsNoComponentByDefault(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	_, component := serviceReadyness.IsReady()
	assert.Equal(t, "", component) // must have no component by default
}

func Test_IsReady_ForANewComponent_ReturnsNotReady(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	ready, _ := serviceReadyness.IsReady("foo")
	assert.False(t, ready) // must be not ready by default
}

func Test_IsReady_ForANewComponent_ReturnsComponent(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	_, component := serviceReadyness.IsReady("foo")
	assert.Equal(t, "foo", component) // must have component by default
}

func Test_IsReady_ReturnsNotReady_IfAnyComponentIsNotReady(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	serviceReadyness.Register("foo")
	serviceReadyness.SetReady("bar")

	ready, _ := serviceReadyness.IsReady()
	assert.True(t, !ready) // must be not ready
}

func Test_IsReady_ForAReadyComponent_ReturnsReady(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	serviceReadyness.SetReady("foo")

	ready, _ := serviceReadyness.IsReady("foo")
	assert.True(t, ready) // must be ready
}

func Test_IsReady_ForANotReadyComopnent_ReturnsNotReady(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	serviceReadyness.Register("foo")

	ready, _ := serviceReadyness.IsReady("foo")
	assert.True(t, !ready) // must be not ready
}

func Test_IsReady_ForANotReadyComopnent_ReturnsComponent(t *testing.T) {
	logger := logrus.New()
	serviceReadyness := NewServiceReadyness(logger)

	serviceReadyness.SetReady("foo")

	_, component := serviceReadyness.IsReady("foo", "bar")
	assert.Equal(t, "bar", component) // must be the bar component
}
