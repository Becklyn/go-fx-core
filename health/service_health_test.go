package health

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func componentChangedConsumer(service *ServiceHealth) {
	for component := range service.componentChanged {
		fmt.Printf("%s: %v\n", component.name, component.health)
	}
}

func Test_IsHealthy_ReturnsHealthyByDefault(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	healthy, _ := serviceHealth.IsHealthy()
	is.True(healthy) // must be healthy by default
}

func Test_IsHealthy_ReturnsNoReasonByDefault(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	_, reason := serviceHealth.IsHealthy()
	is.Equal("", reason) // must have no reason by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsHealthy(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	healthy, _ := serviceHealth.IsHealthy("foo")
	is.True(healthy) // must be healthy by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsNoReason(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	_, reason := serviceHealth.IsHealthy("foo")
	is.Equal("", reason) // must have no reason by default
}

func Test_IsHealthy_ReturnsUnhealthy_IfAnyComponentIsUnhealthy(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")
	serviceHealth.SetHealthy("bar")

	healthy, _ := serviceHealth.IsHealthy()
	is.True(!healthy) // must be unhealthy
}

func Test_IsHealthy_ForAHealthyComponent_ReturnsHealthy(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	serviceHealth.SetHealthy("foo")

	healthy, _ := serviceHealth.IsHealthy("foo")
	is.True(healthy) // must be healthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsUnhealthy(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")

	healthy, _ := serviceHealth.IsHealthy("foo")
	is.True(!healthy) // must be unhealthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsReason(t *testing.T) {
	is := is.New(t)

	serviceHealth := NewServiceHealth()
	go componentChangedConsumer(serviceHealth)

	serviceHealth.SetUnhealthy("foo", "can't foo without bar")

	_, reason := serviceHealth.IsHealthy("foo")
	is.Equal("can't foo without bar", reason) // must be the reason of foo
}
