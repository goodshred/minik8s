package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kubelet/app/options"
	"k8s.io/kubernetes/pkg/kubelet"
	"k8s.io/kubernetes/pkg/kubelet/config"
	"os"
)

// NewKubeletCommand creates a *cobra.Command object with default parameters
func NewKubeletCommand() *cobra.Command {
	kubeletFlags := options.NewKubeletFlags()

	kubeletConfig, err := options.NewKubeletConfiguration()
	// programmer error
	if err != nil {
		klog.ErrorS(err, "Failed to create a new kubelet configuration")
		os.Exit(1)
	}
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			// construct a KubeletServer from kubeletFlags and kubeletConfig
			kubeletServer := &options.KubeletServer{
				KubeletFlags:         *kubeletFlags,
				KubeletConfiguration: *kubeletConfig,
			}
			kubeletDeps, err := UnsecuredDependencies(kubeletServer, utilfeature.DefaultFeatureGate)

			return Run(ctx, kubeletServer, kubeletDeps, utilfeature.DefaultFeatureGate)

		},
	}
	return nil
}

func Run(ctx context.Context, s *options.KubeletServer, kubeDeps *kubelet.Dependencies, featureGate featuregate.FeatureGate) error {
	if err := run(ctx, s, kubeDeps, featureGate); err != nil {
		return fmt.Errorf("failed to run Kubelet: %w", err)
	}
	return nil
}

func run(ctx context.Context, s *options.KubeletServer, kubeDeps *kubelet.Dependencies, featureGate featuregate.FeatureGate) (err error) {
	if err := kubelet.PreInitRuntimeService(&s.KubeletConfiguration, kubeDeps); err != nil {
		return err
	}
	if err := RunKubelet(ctx, s, kubeDeps, s.RunOnce); err != nil {
		return err
	}
}

// UnsecuredDependencies returns a Dependencies suitable for being run, or an error if the server setup
// is not valid.  It will not start any background processes, and does not include authentication/authorization
func UnsecuredDependencies(s *options.KubeletServer, featureGate featuregate.FeatureGate) (*kubelet.Dependencies, error) {
	// Initialize the TLS Options
	tlsOptions, err := InitializeTLS(&s.KubeletFlags, &s.KubeletConfiguration)
	if err != nil {
		return nil, err
	}

	mounter := mount.New(s.ExperimentalMounterPath)
	subpather := subpath.New(mounter)
	hu := hostutil.NewHostUtil()
	pluginRunner := exec.New()

	plugins, err := ProbeVolumePlugins(featureGate)
	if err != nil {
		return nil, err
	}
	var tp oteltrace.TracerProvider
	tp = noopoteltrace.NewTracerProvider()
	if utilfeature.DefaultFeatureGate.Enabled(features.KubeletTracing) {
		tp, err = newTracerProvider(s)
		if err != nil {
			return nil, err
		}
	}
	return &kubelet.Dependencies{
		Auth:                nil, // default does not enforce auth[nz]
		CAdvisorInterface:   nil, // cadvisor.New launches background processes (bg http.ListenAndServe, and some bg cleaners), not set here
		Cloud:               nil, // cloud provider might start background processes
		ContainerManager:    nil,
		KubeClient:          nil,
		HeartbeatClient:     nil,
		EventClient:         nil,
		TracerProvider:      tp,
		HostUtil:            hu,
		Mounter:             mounter,
		Subpather:           subpather,
		OOMAdjuster:         oom.NewOOMAdjuster(),
		OSInterface:         kubecontainer.RealOS{},
		VolumePlugins:       plugins,
		DynamicPluginProber: GetDynamicPluginProber(s.VolumePluginDir, pluginRunner),
		TLSOptions:          tlsOptions}, nil
}

// RunKubelet is responsible for setting up and running a kubelet.  It is used in three different applications:
//
//	1 Integration tests
//	2 Kubelet binary
//	3 Standalone 'kubernetes' binary
//
// Eventually, #2 will be replaced with instances of #3
func RunKubelet(ctx context.Context, kubeServer *options.KubeletServer, kubeDeps *kubelet.Dependencies, runOnce bool) error {
	// process pods and exit.
	if runOnce {
		//if _, err := k.RunOnce(podCfg.Updates()); err != nil {
		//	return fmt.Errorf("runonce failed: %w", err)
		//}
		//klog.InfoS("Started kubelet as runonce")
	} else {
		startKubelet(k, podCfg, &kubeServer.KubeletConfiguration, kubeDeps, kubeServer.EnableServer)
		klog.InfoS("Started kubelet")
	}
}

func startKubelet(k kubelet.Bootstrap, podCfg *config.PodConfig, kubeCfg *kubeletconfiginternal.KubeletConfiguration, kubeDeps *kubelet.Dependencies, enableServer bool) {
	// start the kubelet
	go k.Run(podCfg.Updates())
}
