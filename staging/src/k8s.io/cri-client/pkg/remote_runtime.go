package cri

import (
	internalapi "cri-api/pkg/apis"
	runtimeapi "cri-api/pkg/apis/runtime/v1"
	"errors"
	"fmt"
	"google.golang.org/grpc"
)

type remoteRuntimeService struct {
	runtimeClient runtimeapi.RuntimeServiceClient
}

func NewRemoteRuntimeService(endpoint string) (internalapi.RuntimeService, error) {
	conn, err := grpc.DialContext(ctx, addr, dialOpts...)
	if err != nil {
		internal.LogErr(logger, err, "Connect remote runtime failed", "address", addr)
		return nil, err
	}
	return &remoteRuntimeService{
		runtimeClient: runtimeapi.NewRuntimeServiceClient(conn),
	}, nil
}

// CreateContainer creates a new container in the specified PodSandbox.
func (r *remoteRuntimeService) CreateContainer(ctx context.Context, podSandBoxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {

	r.log(10, "[RemoteRuntimeService] CreateContainer", "podSandboxID", podSandBoxID, "timeout", r.timeout)
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	return r.createContainerV1(ctx, podSandBoxID, config, sandboxConfig)
}

// StartContainer starts the container.
func (r *remoteRuntimeService) StartContainer(ctx context.Context, containerID string) (err error) {
	r.log(10, "[RemoteRuntimeService] StartContainer", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	if _, err := r.runtimeClient.StartContainer(ctx, &runtimeapi.StartContainerRequest{
		ContainerId: containerID,
	}); err != nil {
		r.logErr(err, "StartContainer from runtime service failed", "containerID", containerID)
		return err
	}
	r.log(10, "[RemoteRuntimeService] StartContainer Response", "containerID", containerID)

	return nil
}

func (r *remoteRuntimeService) createContainerV1(ctx context.Context, podSandBoxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	resp, err := r.runtimeClient.CreateContainer(ctx, &runtimeapi.CreateContainerRequest{
		PodSandboxId:  podSandBoxID,
		Config:        config,
		SandboxConfig: sandboxConfig,
	})
	if err != nil {
		r.logErr(err, "CreateContainer in sandbox from runtime service failed", "podSandboxID", podSandBoxID)
		return "", err
	}

	r.log(10, "[RemoteRuntimeService] CreateContainer", "podSandboxID", podSandBoxID, "containerID", resp.ContainerId)
	if resp.ContainerId == "" {
		errorMessage := fmt.Sprintf("ContainerId is not set for container %q", config.Metadata)
		err := errors.New(errorMessage)
		r.logErr(err, "CreateContainer failed")
		return "", err
	}

	return resp.ContainerId, nil
}
