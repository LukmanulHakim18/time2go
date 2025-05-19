package httpcaller

import (
	"github.com/LukmanulHakim18/time2go/repository"
)

type httpCallerConfig struct {
	checkHealth bool
}

func NewHttpCallerConfig(checkHealth bool) repository.RepoConf {
	return &httpCallerConfig{
		checkHealth: checkHealth,
	}
}

func (conf *httpCallerConfig) Init(r *repository.Repository) error {

	httpcallerClient := &HttpCallerClient{
		// TODO: uncomment line below if implementor already defined
		// cli: cli,
	}

	r.HttpCaller = httpcallerClient
	return nil
}

func (conf *httpCallerConfig) GetRepoName() string {
	return "HttpCaller"
}
