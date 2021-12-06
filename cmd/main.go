package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/quanxiang-cloud/implant/pkg/watcher/overseerrun"
	"github.com/quanxiang-cloud/implant/pkg/watcher/postman"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"github.com/quanxiang-cloud/overseer/pkg/client/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ct "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	leaseLockName      string
	leaseLockNamespace string
	id                 string
	defaultResync      time.Duration
	releaseOnCancel    bool
	leaseDuration      time.Duration
	renewDeadline      time.Duration
	retryPeriod        time.Duration
	concurrency        int
	target             string
	cacheMaxEntries    int
)

func main() {
	flag.StringVar(&id, "id", uuid.New().String(), "the holder identity name")
	flag.StringVar(&leaseLockName, "lease-lock-name", "overseer", "the lease lock resource name")
	flag.StringVar(&leaseLockNamespace, "lease-lock-namespace", "default", "the lease lock resource namespace")
	flag.DurationVar(&defaultResync, "default-resync", time.Duration(30)*time.Second, "")

	flag.BoolVar(&releaseOnCancel, "release-on-cancel", true, "")
	flag.DurationVar(&leaseDuration, "leasese", time.Duration(60)*time.Second, "")
	flag.DurationVar(&renewDeadline, "renew", time.Duration(15)*time.Second, "")
	flag.DurationVar(&retryPeriod, "retry", time.Duration(5)*time.Second, "")

	flag.IntVar(&cacheMaxEntries, "cache-max-entries", 1024, "")
	flag.IntVar(&concurrency, "concurrency", 1, "")
	flag.StringVar(&target, "target", "localhost:8081", "")
	flag.Parse()

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	if target == "" {
		klog.Error("target must be set")
		os.Exit(1)
	}

	config := ctrl.GetConfigOrDie()
	client, err := clientset.NewForConfig(config)
	if err != nil {
		klog.Error(err, "unable to get client set")
		os.Exit(1)
	}
	ctx := context.Background()

	leader := make(chan struct{})
	go HA(ctx, config, leader)
	<-leader
	klog.Info("i am leader,working now")

	MainWithClient(ctx, client)
}

func MainWithClient(ctx context.Context, client *clientset.Clientset) {
	worker, err := postman.New(ctx, target)
	if err != nil {
		klog.Error(err, "unable to get worker client")
		os.Exit(1)
	}

	opts := make([]reconciler.Options, 0)
	errChan := make(chan error)
	for i := 0; i < concurrency; i++ {
		opts = append(opts,
			reconciler.WithConsumer(
				ctx,
				worker.SendFN(errChan),
			),
			reconciler.WithCache(ctx, cacheMaxEntries),
		)
	}
	cc := overseerrun.NewControllerWithConfig(ctx, client, defaultResync, opts...)
	go cc.Run(ctx.Done())

	err = <-errChan
	klog.Error(err)
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
