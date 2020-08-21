package config

import "go-elastic-job-lite/reg"

type ConfigurationService struct {
	reg.Registry
}

func (svc *ConfigurationService) Load() LiteJobConfiguration {
	return nil
}

func (svc *ConfigurationService) Persist(config LiteJobConfiguration) error {
	return nil
}
