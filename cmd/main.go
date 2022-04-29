package main

import (
	"context"
	"flag"
	"os"
	"time"

	fnClientset "github.com/openfunction/pkg/client/clientset/versioned"
	id2 "github.com/quanxiang-cloud/cabin/id"
	"github.com/quanxiang-cloud/cabin/tailormade/client"
	"github.com/quanxiang-cloud/implant/pkg/watcher"
	"github.com/quanxiang-cloud/implant/pkg/watcher/broadcaster/bus"
	"github.com/quanxiang-cloud/implant/pkg/watcher/informers"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	tkClientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ct "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	klog "k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	leaseLockName      string
	leaseLockNamespace string
	namespace          string
	id                 string
	defaultResync      time.Duration
	releaseOnCancel    bool
	leaseDuration      time.Duration
	renewDeadline      time.Duration
	retryPeriod        time.Duration
	concurrency        int
	fnUpdate           string
	docUpdate          string
	cacheMaxEntries    int
	pubsubName         string

	timeout      time.Duration
	maxIdleConns int
)

func main() {
	flag.StringVar(&id, "id", id2.BaseUUID(), "the holder identity name")
	flag.StringVar(&leaseLockName, "lease-lock-name", "faas", "the lease lock resource name")
	flag.StringVar(&leaseLockNamespace, "lease-lock-namespace", "default", "the lease lock resource namespace")
	flag.StringVar(&namespace, "namespace", "default", "")
	flag.DurationVar(&defaultResync, "default-resync", time.Duration(30)*time.Second, "")

	flag.BoolVar(&releaseOnCancel, "release-on-cancel", true, "")
	flag.DurationVar(&leaseDuration, "leasese", time.Duration(60)*time.Second, "")
	flag.DurationVar(&renewDeadline, "renew", time.Duration(15)*time.Second, "")
	flag.DurationVar(&retryPeriod, "retry", time.Duration(5)*time.Second, "")

	flag.IntVar(&cacheMaxEntries, "cache-max-entries", 1024, "")
	flag.IntVar(&concurrency, "concurrency", 1, "")
	flag.StringVar(&fnUpdate, "fn-update", "localhost:8080", "")
	flag.StringVar(&docUpdate, "doc-update", "localhost:8080", "")
	flag.DurationVar(&timeout, "timeout", time.Duration(20)*time.Second, "")
	flag.IntVar(&maxIdleConns, "maxIdleConns", 10, "")
	flag.StringVar(&pubsubName, "pubsub", "", "")
	flag.Parse()

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	if fnUpdate == "" || docUpdate == "" {
		klog.Error("target must be set")
		os.Exit(1)
	}
	config := ctrl.GetConfigOrDie()
	ctx := context.Background()

	leader := make(chan struct{})
	go HA(ctx, config, leader)
	<-leader
	klog.Info("i am leader")

	c := &client.Config{
		Timeout:      timeout,
		MaxIdleConns: maxIdleConns,
	}

	errChan := make(chan error)
	// TODO: remove
	bus, err := bus.NewDaprClient(ctx, errChan, bus.WithPubsubName(pubsubName))
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	watch(ctx, c, config, bus)

	err = <-errChan
	klog.Error(err)
}

func watch(ctx context.Context, cc *client.Config, rc *rest.Config, bus *bus.EventBus) {
	oper := informers.Oper{
		Namespace:     namespace,
		DefaultResync: defaultResync,
	}
	watcher.NewWatcherWithOper(ctx, oper).
		Cache(cacheMaxEntries).
		Bus(bus, concurrency).
		Opts(reconciler.WithFunction(ctx)).
		RunOrDie(fnClientset.NewForConfigOrDie(rc))

	watcher.NewWatcherWithOper(ctx, oper).
		Cache(cacheMaxEntries).
		Bus(bus, concurrency).
		Opts(reconciler.WithPipelineRun(ctx)).
		RunOrDie(tkClientset.NewForConfigOrDie(rc))
}

func HA(ctx context.Context, config *rest.Config, going chan<- struct{}) {
	client := ct.NewForConfigOrDie(config)

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaseLockName,
			Namespace: leaseLockNamespace,
		},
		Client: client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: id,
		},
	}

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: releaseOnCancel,
		LeaseDuration:   leaseDuration,
		RenewDeadline:   renewDeadline,
		RetryPeriod:     retryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				going <- struct{}{}
			},
			OnStoppedLeading: func() {
				klog.Infof("leader lost: %s", id)
				os.Exit(0)
			},
			OnNewLeader: func(identity string) {
				// we're notified when new leader elected
				klog.Infof("new leader elected: %s", identity)
				if identity == id {
					return
				}
			},
		},
	})
}
