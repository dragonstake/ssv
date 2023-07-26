package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcessAgents(t *testing.T) {
	a1 := mockAgent{errs: []string{"dummy error", "another error"}}
	a2 := mockAgent{errs: []string{}}
	a3 := mockAgent{errs: []string{"another error 3"}}

	errs := ProcessAgents([]HealthChecker{&a1, &a2, &a3})
	require.Len(t, errs, 3)
}

type mockAgent struct {
	errs []string
}

func (ma *mockAgent) HealthCheck() []string {
	return ma.errs
}
