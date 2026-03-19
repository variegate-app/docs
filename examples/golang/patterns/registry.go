package patterns

import (
	"context"
	"fmt"

	contextpattern "mddocks/examples/golang/patterns/context"
	cancellationpattern "mddocks/examples/golang/patterns/cancellation"
	explicitcancellationpattern "mddocks/examples/golang/patterns/explicit-cancellation"
	deadlinepattern "mddocks/examples/golang/patterns/deadline"
	errgrouppattern "mddocks/examples/golang/patterns/errgroup"
	gracefulshutdownpattern "mddocks/examples/golang/patterns/graceful-shutdown"
	workerpoolpattern "mddocks/examples/golang/patterns/worker-pool"
	pipelinepattern "mddocks/examples/golang/patterns/pipeline"
	boundedparallelismattern "mddocks/examples/golang/patterns/bounded-parallelism"
	ratelimiterpattern "mddocks/examples/golang/patterns/rate-limiter"
	singleflightpattern "mddocks/examples/golang/patterns/singleflight"
	takefirstnpattern "mddocks/examples/golang/patterns/take-first-n"

	forselectpattern "mddocks/examples/golang/patterns/for-select"
	forselectdonepattern "mddocks/examples/golang/patterns/for-select-done"
	orchannelpattern "mddocks/examples/golang/patterns/or-channel"
	teechannelpattern "mddocks/examples/golang/patterns/tee-channel"
	bridgechannelpattern "mddocks/examples/golang/patterns/bridge-channel"
	ringbufferpattern "mddocks/examples/golang/patterns/ring-buffer"
	faninpattern "mddocks/examples/golang/patterns/fan-in"
	fanoutpattern "mddocks/examples/golang/patterns/fan-out"
	pubsubpattern "mddocks/examples/golang/patterns/pub-sub"
	generatorpattern "mddocks/examples/golang/patterns/generator"
	selectstatementwithtimeoutpattern "mddocks/examples/golang/patterns/select-statement-with-timeout"
	waitforresultpattern "mddocks/examples/golang/patterns/wait-for-result"
	waitfortaskpattern "mddocks/examples/golang/patterns/wait-for-task"
	poolingpattern "mddocks/examples/golang/patterns/pooling"
	droppattern "mddocks/examples/golang/patterns/drop"
	mapfilterpattern "mddocks/examples/golang/patterns/map-filter"
	filterpattern "mddocks/examples/golang/patterns/filter"
	queueingpattern "mddocks/examples/golang/patterns/queuing"
	exponentialbackoffpattern "mddocks/examples/golang/patterns/exponential-backoff"
	faulttolerancepattern "mddocks/examples/golang/patterns/fault-tolerance"
	failfastpattern "mddocks/examples/golang/patterns/fail-fast"
	handshakingpattern "mddocks/examples/golang/patterns/handshaking"
	steadystatepattern "mddocks/examples/golang/patterns/steady-state"
	stoppingshortpattern "mddocks/examples/golang/patterns/stopping-short"
	digestingatreepattern "mddocks/examples/golang/patterns/digesting-a-tree"
	paralleldigestionpattern "mddocks/examples/golang/patterns/parallel-digestion"
	conclusionpattern "mddocks/examples/golang/patterns/conclusion"
	roundtripperpattern "mddocks/examples/golang/patterns/round-tripper"
	debouncepattern "mddocks/examples/golang/patterns/debounce"
	tokenbucketleakybucketpattern "mddocks/examples/golang/patterns/token-bucket-leaky-bucket"
	jitterpattern "mddocks/examples/golang/patterns/jitter"
	syncpoolpattern "mddocks/examples/golang/patterns/sync-pool"
)

type Runner func(ctx context.Context) error

// Run executes the runnable example for the given pattern name.
func Run(ctx context.Context, name string) error {
	r, ok := runners[name]
	if !ok {
		return fmt.Errorf("unknown pattern %q", name)
	}
	return r(ctx)
}

var runners = map[string]Runner{
	// Minimal runnable implementations (core set).
	"context":                  contextpattern.Run,
	"cancellation":             cancellationpattern.Run,
	"explicit-cancellation":   explicitcancellationpattern.Run,
	"deadline":                 deadlinepattern.Run,
	"errgroup":                 errgrouppattern.Run,
	"graceful-shutdown":       gracefulshutdownpattern.Run,
	"worker-pool":              workerpoolpattern.Run,
	"pipeline":                 pipelinepattern.Run,
	"bounded-parallelism":     boundedparallelismattern.Run,
	"rate-limiter":            ratelimiterpattern.Run,
	"singleflight":            singleflightpattern.Run,
	"take-first-n":           takefirstnpattern.Run,

	// Remaining patterns: packages exist and are runnable (currently mostly stubs).
	"for-select":                       forselectpattern.Run,
	"for-select-done":                 forselectdonepattern.Run,
	"or-channel":                      orchannelpattern.Run,
	"tee-channel":                     teechannelpattern.Run,
	"bridge-channel":                 bridgechannelpattern.Run,
	"ring-buffer":                    ringbufferpattern.Run,
	"fan-in":                          faninpattern.Run,
	"fan-out":                         fanoutpattern.Run,
	"pub-sub":                         pubsubpattern.Run,
	"generator":                       generatorpattern.Run,
	"select-statement-with-timeout":  selectstatementwithtimeoutpattern.Run,
	"wait-for-result":                waitforresultpattern.Run,
	"wait-for-task":                  waitfortaskpattern.Run,
	"pooling":                         poolingpattern.Run,
	"drop":                            droppattern.Run,
	"map-filter":                     mapfilterpattern.Run,
	"filter":                          filterpattern.Run,
	"queuing":                         queueingpattern.Run,
	"exponential-backoff":            exponentialbackoffpattern.Run,
	"fault-tolerance":               faulttolerancepattern.Run,
	"fail-fast":                      failfastpattern.Run,
	"handshaking":                   handshakingpattern.Run,
	"steady-state":                  steadystatepattern.Run,
	"stopping-short":                stoppingshortpattern.Run,
	"digesting-a-tree":              digestingatreepattern.Run,
	"parallel-digestion":            paralleldigestionpattern.Run,
	"conclusion":                    conclusionpattern.Run,
	"round-tripper":                 roundtripperpattern.Run,
	"debounce":                       debouncepattern.Run,
	"token-bucket-leaky-bucket":     tokenbucketleakybucketpattern.Run,
	"jitter":                         jitterpattern.Run,
	"sync-pool":                      syncpoolpattern.Run,
}

