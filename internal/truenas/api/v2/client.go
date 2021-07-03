package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/pool"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2/sharing"
)

type Client interface {
	Get(context context.Context, path string, v interface{}) error

	SharingNFS() sharing.NFSService
	PoolDataset() pool.DatasetService

	//ListNFS(ctx context.Context) ([]*NFS, error)
	//CreateNFS(ctx context.Context, nfs *NFS) error
	//UpdateNFS(ctx context.Context, nfs *NFS) error
	//GetNFS(ctx context.Context, id int) (*NFS, error)
	//DeleteNFS(ctx context.Context, id int) error
	//
	//ListDataset(ctx context.Context) ([]*Dataset, error)
	//GetDataset(ctx context.Context, id string) (*Dataset, error)
	//UpdateDataset(ctx context.Context, dataset *Dataset) error
	//DeleteDataset(ctx context.Context, id string) error
	//CreateDataset(ctx context.Context, dataset *Dataset) error
}

var (
	ErrNotFound = errors.New("not found")
)

type client struct {
	host   string
	apiKey string
	http   *http.Client

	sharingNFSService  sharing.NFSService
	poolDatasetService pool.DatasetService
}

func NewClient(host, apiKey string) Client {
	client := &client{
		host:   host,
		apiKey: apiKey,
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	client.sharingNFSService = sharing.NewNFSService(client)
	client.poolDatasetService = pool.NewDatasetService(client)

	return client
}

func (c client) SharingNFS() sharing.NFSService {
	return c.sharingNFSService
}

func (c client) PoolDataset() pool.DatasetService {
	return c.poolDatasetService
}

func (c client) Post(ctx context.Context, path string, body interface{}, v interface{}) error {
	return c.request(ctx, http.MethodPost, path, body, v)
}

func (c client) Get(ctx context.Context, path string, v interface{}) error {
	return c.request(ctx, http.MethodGet, path, nil, v)
}

func (c client) Put(ctx context.Context, path string, body interface{}, v interface{}) error {
	return c.request(ctx, http.MethodPut, path, body, v)
}

func (c client) Delete(ctx context.Context, path string) error {
	return c.request(ctx, http.MethodDelete, path, nil, nil)
}

func (c client) request(ctx context.Context, method, path string, body interface{}, v interface{}) error {
	url := fmt.Sprintf("%s/api/v2.0/%s", c.host, path)

	b := bytes.NewBuffer(nil)
	if body != nil {
		err := json.NewEncoder(b).Encode(body)
		log.Println("[DEBUG] ", b.String())
		if err != nil {
			return fmt.Errorf("failed to encode body: %w", err)
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, url, b)
	if err != nil {
		return fmt.Errorf("failed to create request for url: %s, err: %w", url, err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "terraform-provider-truenas")

	response, err := c.http.Do(request)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}
	if response.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(response.Body)
		//TODO: list codes in allowed
		return fmt.Errorf("wrong status code, got: %d, body: %s", response.StatusCode, string(b))
		//return responseError(response, fmt.Errorf("wrong status code"))
	}

	if v == nil {
		return nil
	}

	buffer, ok := v.(*bytes.Buffer)
	if ok {
		_, err := io.Copy(buffer, response.Body)
		return err
	}

	// TODO: handle json errors
	return json.NewDecoder(response.Body).Decode(v)

}
