package service

import (
	"github.com/pkg/errors"
)

type ServiceProvider struct {
	opts ServiceOptions

	apiService *APIService

	commonService *CommonService
}

type ServiceOptions struct {
	APIAdvertiseURLs []string
}

func NewServiceProvider(opts ServiceOptions) (*ServiceProvider, error) {
	sp := &ServiceProvider{
		opts: opts,
	}

	var err error

	sp.apiService, err = newAPIService(sp, opts.APIAdvertiseURLs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sp.commonService, err = newCommonService(sp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return sp, nil
}

func (d *ServiceProvider) APIService() *APIService {
	return d.apiService
}

// common

func (d *ServiceProvider) CommonService() *CommonService {
	return d.commonService
}
