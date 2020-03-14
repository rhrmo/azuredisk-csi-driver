// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/api/volume/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

const groupName = "volume"

var version = apiversion.NewVersionOrPanic("v1alpha1")

type Client struct {
	client     v1alpha1.VolumeClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the volume API group version v1alpha1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1alpha1.NewVolumeClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1alpha1.VolumeClient = &Client{}

func (w *Client) DismountVolume(context context.Context, request *v1alpha1.DismountVolumeRequest, opts ...grpc.CallOption) (*v1alpha1.DismountVolumeResponse, error) {
	return w.client.DismountVolume(context, request, opts...)
}

func (w *Client) FormatVolume(context context.Context, request *v1alpha1.FormatVolumeRequest, opts ...grpc.CallOption) (*v1alpha1.FormatVolumeResponse, error) {
	return w.client.FormatVolume(context, request, opts...)
}

func (w *Client) IsVolumeFormatted(context context.Context, request *v1alpha1.IsVolumeFormattedRequest, opts ...grpc.CallOption) (*v1alpha1.IsVolumeFormattedResponse, error) {
	return w.client.IsVolumeFormatted(context, request, opts...)
}

func (w *Client) ListVolumesOnDisk(context context.Context, request *v1alpha1.ListVolumesOnDiskRequest, opts ...grpc.CallOption) (*v1alpha1.ListVolumesOnDiskResponse, error) {
	return w.client.ListVolumesOnDisk(context, request, opts...)
}

func (w *Client) MountVolume(context context.Context, request *v1alpha1.MountVolumeRequest, opts ...grpc.CallOption) (*v1alpha1.MountVolumeResponse, error) {
	return w.client.MountVolume(context, request, opts...)
}

func (w *Client) ResizeVolume(context context.Context, request *v1alpha1.ResizeVolumeRequest, opts ...grpc.CallOption) (*v1alpha1.ResizeVolumeResponse, error) {
	return w.client.ResizeVolume(context, request, opts...)
}