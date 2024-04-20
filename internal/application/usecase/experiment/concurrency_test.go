package experiment

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRunConcurrencyTest(t *testing.T) {
    // Run the concurrency test
    RunConcurrencyTest()

    assert.Equal(t, expectedX, x)
    assert.Equal(t, expectedY, y)
    assert.Equal(t, expectedSum, x+y)
}