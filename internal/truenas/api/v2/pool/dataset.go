package pool

import (
	"bytes"
	"context"
	"encoding/json"
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
	Type string `json:"type"`
	Name string `json:"name"`
	//Pool           string      `json:"pool,omitempty"`
	//Encrypted      bool        `json:"encrypted"`
	//EncryptionRoot string      `json:"encryption_root"`
	//KeyLoaded      bool        `json:"key_loaded"`
	//Mountpoint     string      `json:"mountpoint"`
	//Children interface{} `json:"children"`
	Comments *Value `json:"comments,omitempty"`
	//Used for VOLUME
	Volsize      *Value `json:"volsize,omitempty"`
	Volblocksize *Value `json:"volblocksize,omitempty"`
	Sparse       *Value `json:"sparse,omitempty"`
	ForceSize    *Value `json:"force_size"`

	Deduplication   *Value `json:"deduplication"`
	Aclmode         *Value `json:"aclmode"`
	Acltype         *Value `json:"acltype"`
	Xattr           *Value `json:"xattr"`
	Atime           *Value `json:"atime"`
	Casesensitivity *Value `json:"casesensitivity"`
	Exec            *Value `json:"exec"`
	Sync            *Value `json:"sync"`
	Compression     *Value `json:"compression"`
	//Compressratio         *Value      `json:"compressratio"`
	//Origin                *Value      `json:"origin"`
	//Quota                 *Value      `json:"quota"`
	//Refquota              *Value      `json:"refquota"`
	//Reservation           *Value      `json:"reservation"`
	//Refreservation        *Value      `json:"refreservation"`
	//Copies                *Value      `json:"copies"`
	//Snapdir               *Value      `json:"snapdir"`
	Readonly   *Value `json:"readonly"`
	Recordsize *Value `json:"recordsize"`
	//KeyFormat             *Value      `json:"key_format"`
	//EncryptionAlgorithm   *Value      `json:"encryption_algorithm"`
	//Used                  *Value      `json:"used"`
	//Available             *Value      `json:"available"`
	//SpecialSmallBlockSize *Value      `json:"special_small_block_size"`
	//Pbkdf2Iters           *Value      `json:"pbkdf2iters"`
	//Locked bool `json:"locked"`
}

// Value is the weird struct truenas talks with ous
// While we send raw data in the Parsed field to the api we try to use Value as interpretation
// when we get it back as much as we can
type Value struct {
	Parsed   interface{} `json:"parsed"` // yea how about no
	Rawvalue string      `json:"rawvalue"`
	Value    string      `json:"value"`
	Source   string      `json:"source"`
}

func NewValue(raw interface{}) *Value {
	return &Value{
		Parsed: raw,
	}
}

func (v *Value) Raw() string {
	if v == nil {
		return ""
	}
	return v.Rawvalue
}

func (v *Value) String() string {
	if v == nil {
		return ""
	}
	return v.Value
}

func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Parsed)
}

type queryOptions struct {
	Options struct {
		Extra struct {
			Flat             bool `json:"flat"`
			RetrieveChildren bool `json:"retrieve_children"`
		} `json:"extra"`
	} `json:"query-options"`
	Filters []int `json:"query-filters"`
}

func (s datasetService) List(ctx context.Context) ([]*Dataset, error) {
	options := queryOptions{Filters: []int{}}
	options.Options.Extra.Flat = true
	options.Options.Extra.RetrieveChildren = true
	var l []*Dataset
	err := s.client.Get(ctx, "/pool/dataset", options, &l)
	return l, err
}

func (s datasetService) Get(ctx context.Context, name string) (*Dataset, error) {
	dataset := &Dataset{}
	b := bytes.NewBuffer(nil)
	err := s.client.Get(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(name)), nil, b)
	//err := s.client.Get(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(name)), nil, dataset)
	fmt.Println(b.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b.Bytes(), dataset)
	return dataset, err
}

func (s datasetService) Update(ctx context.Context, dataset *Dataset) error {
	return s.client.Put(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(dataset.Name)), dataset, dataset)
}

func (s datasetService) Delete(ctx context.Context, name string) error {
	return s.client.Delete(ctx, fmt.Sprint("pool/dataset/id/", url.QueryEscape(name)))
}

func (s datasetService) Create(ctx context.Context, dataset *Dataset) error {
	return s.client.Post(ctx, "pool/dataset", dataset, nil)
}
