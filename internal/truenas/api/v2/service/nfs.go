package service

import (
	"context"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/http"
)

type NFSService interface {
	Get(ctx context.Context) (*NFS, error)
	Update(ctx context.Context, nfs *NFS) error
}

func NewNFSService(client http.Client) NFSService {
	return &nfsService{
		client: client,
	}
}

type nfsService struct {
	client http.Client
}

type NFS struct {
	ID              int         `json:"id"`
	Servers         int         `json:"servers"`
	UDP             bool        `json:"udp"`
	AllowNonroot    bool        `json:"allow_nonroot"`
	V4              bool        `json:"v4"`
	V4V3Owner       bool        `json:"v4_v3owner"`
	V4Krb           bool        `json:"v4_krb"`
	Bindip          []string    `json:"bindip"`
	MountdPort      interface{} `json:"mountd_port"`
	RpcstatdPort    interface{} `json:"rpcstatd_port"`
	RpclockdPort    interface{} `json:"rpclockd_port"`
	MountdLog       bool        `json:"mountd_log"`
	StatdLockdLog   bool        `json:"statd_lockd_log"`
	V4Domain        string      `json:"v4_domain"`
	V4KrbEnabled    bool        `json:"v4_krb_enabled"`
	UserdManageGids bool        `json:"userd_manage_gids"`
}

func (s nfsService) Get(ctx context.Context) (*NFS, error) {
	nfs := &NFS{}
	err := s.client.Get(ctx, "nfs", nil, nfs)
	return nfs, err
}

func (s nfsService) Update(ctx context.Context, nfs *NFS) error {

	return nil
}
