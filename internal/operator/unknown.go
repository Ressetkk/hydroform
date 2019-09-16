package operator

import (
	"errors"
	"github.com/kyma-incubator/hydroform/types"
)

type Unknown struct {
}

func (u *Unknown) Create(providerType types.ProviderType, configuration map[string]interface{}) error {
	return errors.New("unknown operator")
}

func (u *Unknown) Delete(providerType types.ProviderType, configuration map[string]interface{}) error {
	return errors.New("unknown operator")
}
