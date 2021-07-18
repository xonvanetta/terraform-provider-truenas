package sharing

import (
	"context"
	"fmt"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/http"
)

type NFSService interface {
	List(ctx context.Context) ([]*NFS, error)
	Create(ctx context.Context, nfs *NFS) error
	Update(ctx context.Context, nfs *NFS) error
	Get(ctx context.Context, id int) (*NFS, error)
	Delete(ctx context.Context, id int) error
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
	Id           int      `json:"id,omitempty"`
	Paths        []string `json:"paths"`
	Aliases      []string `json:"aliases"`
	Comment      string   `json:"comment"`
	Networks     []string `json:"networks"`
	Hosts        []string `json:"hosts"`
	Alldirs      bool     `json:"alldirs"`
	Ro           bool     `json:"ro"`
	Quiet        bool     `json:"quiet"`
	MaprootUser  string   `json:"maproot_user"`
	MaprootGroup string   `json:"maproot_group"`
	MapallUser   string   `json:"mapall_user"`
	MapallGroup  string   `json:"mapall_group"`
	Security     []string `json:"security"`
	Enabled      bool     `json:"enabled"`
}

//sanitize will make nil slices to empty slices, truenas doesn't like null in json
//TODO: maybe create custom marshaller
func (n *NFS) sanitize() {
	n.Id = 0
	if n.Paths == nil {
		n.Paths = []string{}
	}
	if n.Aliases == nil {
		n.Aliases = []string{}
	}
	if n.Networks == nil {
		n.Networks = []string{}
	}
	if n.Hosts == nil {
		n.Hosts = []string{}
	}
	if n.Security == nil {
		n.Security = []string{}
	}
}

func (s nfsService) List(ctx context.Context) ([]*NFS, error) {
	var list []*NFS
	err := s.client.Get(ctx, "sharing/nfs", nil, &list)
	return list, err
}

func (s nfsService) Get(ctx context.Context, id int) (*NFS, error) {
	nfs := &NFS{}
	err := s.client.Get(ctx, fmt.Sprint("sharing/nfs/id/", id), nil, nfs)
	return nfs, err
}

func (s nfsService) Update(ctx context.Context, nfs *NFS) error {
	id := nfs.Id
	nfs.sanitize()
	return s.client.Put(ctx, fmt.Sprint("sharing/nfs/id/", id), nfs, nfs)
}

func (s nfsService) Delete(ctx context.Context, id int) error {
	return s.client.Delete(ctx, fmt.Sprint("sharing/nfs/id/", id))
}

func (s nfsService) Create(ctx context.Context, nfs *NFS) error {
	nfs.sanitize()
	return s.client.Post(ctx, "sharing/nfs", nfs, nfs)
}
