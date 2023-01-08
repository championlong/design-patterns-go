package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetStateName())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetStateName())
	m.Approval()
}
