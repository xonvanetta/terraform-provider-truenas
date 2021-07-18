package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/http"
)

type ServiceService interface {
	List(ctx context.Context) ([]*Service, error)
	Get(ctx context.Context, name string) (*Service, error)
	Update(ctx context.Context, name string, enable bool) error
	Start(ctx context.Context, name string) error
	Stop(ctx context.Context, name string) error
	Started(ctx context.Context, name string) error
}

type Service struct {
	ID      int    `json:"id"`
	Service string `json:"service"`
	Enable  bool   `json:"enable"`
	State   string `json:"state"`
	PIDS    []int  `json:"pids"`
}

func NewService(client http.Client) ServiceService {
	return serviceService{
		client: client,
	}
}

type serviceService struct {
	client http.Client
}

type service struct {
	Service string `json:"service"`
}

func (s serviceService) List(ctx context.Context) ([]*Service, error) {
	var list []*Service
	err := s.client.Get(ctx, "/service", nil, &list)
	return list, err
}

//Get will use list and search based on name instead of some ID
func (s serviceService) Get(ctx context.Context, name string) (*Service, error) {
	list, err := s.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, service := range list {
		if service.Service == name {
			return service, nil
		}
	}
	return nil, http.ErrNotFound
}

func (s serviceService) Start(ctx context.Context, name string) error {
	b := &service{Service: name}
	err := s.client.Post(ctx, "/service/start", b, nil)
	return err
}

func (s serviceService) Update(ctx context.Context, name string, enable bool) error {
	b := struct {
		Enable bool `json:"enable"`
	}{
		Enable: enable,
	}
	err := s.client.Put(ctx, fmt.Sprint("/service/id/", name), b, nil)
	return err
}

func (s serviceService) Stop(ctx context.Context, name string) error {
	b := &service{Service: name}
	v := bytes.NewBuffer(nil)
	err := s.client.Post(ctx, "/service/stop", b, v)
	fmt.Println(v.String())
	return err
}

func (s serviceService) Started(ctx context.Context, name string) error {
	return http.ErrNotImplemented
	//b := service{ServiceService: name}
	//fmt.Println(string(d))
	//v := bytes.NewBuffer(nil)
	//err := s.client.Get(ctx, "/service/started", b, v)
	//fmt.Println(v.String())
	//return err
}
