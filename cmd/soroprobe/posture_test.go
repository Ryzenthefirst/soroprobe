package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/stellar/stellartest"
)

func TestCheckAdminKeyReportsPosture(t *testing.T) {
	// The recorded Wasm contract keeps its admin in instance storage.
	out, err := runWithFake(t, stellartest.NewFake(t), "check", "--admin-key", "sym:admin", stellartest.WasmContract)
	require.NoError(t, err, out)
	assert.Contains(t, out, "admin")
	assert.Contains(t, out, "instance storage sym:admin")
}

func TestCheckPostureFlagsConflicts(t *testing.T) {
	fake := stellartest.NewFake(t)
	_, err := runWithFake(t, fake, "check", "--admin-fn", "admin", "--admin-key", "sym:Admin", stellartest.WasmContract)
	require.Error(t, err)
}
