package main

import (
    "os"
	"context"
	"github.com/go-kit/kit/log"
    "net/url"

	"github.com/improbable-eng/thanos/pkg/reloader"
    "flag"
)

func main() {
    baseurl := flag.String("url", "http://localhost:9090", "The base url to Prometheus")
    input := flag.String("input", "", "The input template file")
    output := flag.String("output", "", "The output file")
    flag.Parse()
    rules := flag.Args()

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)

	u, err := url.Parse(*baseurl)
	if err != nil {
		logger.Log("err", err)
        os.Exit(1)
	}

	rl := reloader.New(
		logger,
		reloader.ReloadURLFromBase(u),
		*input,
		*output,
		rules,
	)

	ctx, cancel := context.WithCancel(context.Background())
	if err := rl.Watch(ctx); err != nil {
		cancel()
		logger.Log("err", err)
        os.Exit(1)
	}
}
