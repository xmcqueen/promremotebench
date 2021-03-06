// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package generators

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHostsSimulator(t *testing.T) {
	start := time.Now()
	s := NewHostsSimulator(1, start, HostsSimulatorOptions{})

	{
		// Without offset
		series, err := s.Generate(0, time.Second, 0)
		require.NoError(t, err)
		assert.True(t, len(series) > 0)
	}

	{
		// With offset
		series, err := s.Generate(time.Second, time.Second, 0)
		require.NoError(t, err)
		assert.True(t, len(series) > 0)
	}
}

func TestHostsSimulatorTenSeconds(t *testing.T) {
	start := time.Now()
	s := NewHostsSimulator(13, start, HostsSimulatorOptions{})
	for i := 0; i < 100; i++ {
		series, err := s.Generate(time.Second, 10*time.Second, 0.01)
		require.NoError(t, err)
		assert.True(t, len(series) > 0)
	}
}
