package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	bin := getBinary()
	os.Remove(bin)

	runCmd(t, &cmd{
		Cmd:  "go",
		Args: []string{"build", "-o", bin},
		Dir:  filepath.Join("..", "cmd", "snclient"),
		Env: map[string]string{
			"CGO_ENABLED": "0",
		},
		ErrLike: []string{`.*`},
		Timeout: 5 * time.Minute,
	})

	require.FileExistsf(t, bin, "snclient binary must exist")
}
