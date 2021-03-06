package helperimage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_windowsInfo_create(t *testing.T) {
	revision := "4011f186"
	tests := []struct {
		operatingSystem string
		expectedInfo    Info
		expectedErr     error
	}{
		{
			operatingSystem: "Windows Server 2019 Datacenter Evaluation Version 1809 (OS Build 17763.316)",
			expectedInfo: Info{
				Architecture:            windowsSupportedArchitecture,
				Name:                    name,
				Tag:                     fmt.Sprintf("%s-%s-%s", windowsSupportedArchitecture, revision, baseImage1809),
				IsSupportingLocalImport: false,
				Cmd:                     powerShellCmd,
			},
			expectedErr: nil,
		},
		{
			operatingSystem: "Windows Server Datacenter Version 1809 (OS Build 1803.590)",
			expectedInfo: Info{
				Architecture:            windowsSupportedArchitecture,
				Name:                    name,
				Tag:                     fmt.Sprintf("%s-%s-%s", windowsSupportedArchitecture, revision, baseImage1809),
				IsSupportingLocalImport: false,
				Cmd:                     powerShellCmd,
			},
			expectedErr: nil,
		},
		{
			operatingSystem: "Windows Server Datacenter Version 1803 (OS Build 17134.590)",
			expectedInfo: Info{
				Architecture:            windowsSupportedArchitecture,
				Name:                    name,
				Tag:                     fmt.Sprintf("%s-%s-%s", windowsSupportedArchitecture, revision, baseImage1803),
				IsSupportingLocalImport: false,
				Cmd:                     powerShellCmd,
			},
			expectedErr: nil,
		},
		{
			operatingSystem: "some random string",
			expectedErr:     newUnsupportedWindowsVersionError("some random string"),
		},
	}

	for _, test := range tests {
		t.Run(test.operatingSystem, func(t *testing.T) {
			w := new(windowsInfo)

			image, err := w.Create(revision, Config{OperatingSystem: test.operatingSystem})

			assert.Equal(t, test.expectedInfo, image)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestNewUnsupportedWindowsVersionError(t *testing.T) {
	for _, expectedVersion := range []string{"random1", "random2"} {
		err := newUnsupportedWindowsVersionError(expectedVersion)
		require.Error(t, err)
		assert.Equal(t, expectedVersion, err.version)
	}
}
