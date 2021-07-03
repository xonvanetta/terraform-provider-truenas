package pool

import (
	"context"
	"fmt"
	"net/url"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/http"
)

type DatasetService interface {
	List(ctx context.Context) ([]*Dataset, error)
	Create(ctx context.Context, nfs *Dataset) error
	Update(ctx context.Context, nfs *Dataset) error
	Get(ctx context.Context, id string) (*Dataset, error)
	Delete(ctx context.Context, id string) error
}

func NewDatasetService(client http.Client) DatasetService {
	return &datasetService{
		client: client,
	}
}

type datasetService struct {
	client http.Client
}

type Dataset struct {
	ID                    string `json:"id"`
	Type                  string `json:"type"`
	Name                  string `json:"name"`
	Pool                  string `json:"pool"`
	Encrypted             bool   `json:"encrypted"`
	EncryptionRoot        string `json:"encryption_root"`
	KeyLoaded             bool   `json:"key_loaded"`
	Mountpoint            string `json:"mountpoint"`
	Deduplication         *Value `json:"deduplication"`
	Aclmode               *Value `json:"aclmode"`
	Acltype               *Value `json:"acltype"`
	Xattr                 *Value `json:"xattr"`
	Atime                 *Value `json:"atime"`
	Casesensitivity       *Value `json:"casesensitivity"`
	Exec                  *Value `json:"exec"`
	Sync                  *Value `json:"sync"`
	Compression           *Value `json:"compression"`
	Compressratio         *Value `json:"compressratio"`
	Origin                *Value `json:"origin"`
	Quota                 *Value `json:"quota"`
	Refquota              *Value `json:"refquota"`
	Reservation           *Value `json:"reservation"`
	Refreservation        *Value `json:"refreservation"`
	Copies                *Value `json:"copies"`
	Snapdir               *Value `json:"snapdir"`
	Readonly              *Value `json:"readonly"`
	Recordsize            *Value `json:"recordsize"`
	KeyFormat             *Value `json:"key_format"`
	EncryptionAlgorithm   *Value `json:"encryption_algorithm"`
	Used                  *Value `json:"used"`
	Available             *Value `json:"available"`
	SpecialSmallBlockSize *Value `json:"special_small_block_size"`
	Pbkdf2Iters           *Value `json:"pbkdf2iters"`
	Locked                bool   `json:"locked"`
}

type Value struct {
	Parsed   interface{} `json:"parsed"` // yea how about no
	Rawvalue string      `json:"rawvalue"`
	Value    string      `json:"value"`
	Source   string      `json:"source"`
}

func (v *Value) String() string {
	if v == nil {
		return ""
	}
	return v.Value
}

func (s datasetService) List(ctx context.Context) ([]*Dataset, error) {
	var l []*Dataset
	//TODO: fix without children
	err := s.client.Get(ctx, "/pool/dataset", &l)
	return l, err
}

func (s datasetService) Get(ctx context.Context, id string) (*Dataset, error) {
	dataset := &Dataset{}
	err := s.client.Get(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(id)), dataset)
	return dataset, err
}

func (s datasetService) Update(ctx context.Context, dataset *Dataset) error {
	return http.ErrNotImplemented
	//id := nfs.Id
	//nfs.sanitize()
	//return c.Put(ctx, fmt.Sprint("pool/dataset/id/", dataset), dataset, dataset)
}

func (s datasetService) Delete(ctx context.Context, id string) error {
	return http.ErrNotImplemented
	//return c.Delete(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(id)))
}

func (s datasetService) Create(ctx context.Context, dataset *Dataset) error {
	return http.ErrNotImplemented
	////nfs.sanitize()
	//return c.Post(ctx, "pool/dataset", dataset, dataset)
}
