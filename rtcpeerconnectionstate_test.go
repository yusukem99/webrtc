package webrtc

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewRTCPeerConnectionState(t *testing.T) {
	testCases := []struct {
		stateString   string
		expectedState RTCPeerConnectionState
	}{
		{"unknown", RTCPeerConnectionState(Unknown)},
		{"new", RTCPeerConnectionStateNew},
		{"connecting", RTCPeerConnectionStateConnecting},
		{"connected", RTCPeerConnectionStateConnected},
		{"disconnected", RTCPeerConnectionStateDisconnected},
		{"failed", RTCPeerConnectionStateFailed},
		{"closed", RTCPeerConnectionStateClosed},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			NewRTCPeerConnectionState(testCase.stateString),
			testCase.expectedState,
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestRTCPeerConnectionState_String(t *testing.T) {
	testCases := []struct {
		state          RTCPeerConnectionState
		expectedString string
	}{
		{RTCPeerConnectionState(Unknown), "unknown"},
		{RTCPeerConnectionStateNew, "new"},
		{RTCPeerConnectionStateConnecting, "connecting"},
		{RTCPeerConnectionStateConnected, "connected"},
		{RTCPeerConnectionStateDisconnected, "disconnected"},
		{RTCPeerConnectionStateFailed, "failed"},
		{RTCPeerConnectionStateClosed, "closed"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.state.String(),
			testCase.expectedString,
			"testCase: %d %v", i, testCase,
		)
	}
}