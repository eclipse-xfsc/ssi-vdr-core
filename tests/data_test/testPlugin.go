package main

import (
	"io"

	"github.com/eclipse-xfsc/ssi-vdr-core/types"
)

var Plugin = TestVerifiableDataRegistryModule{}

type TestVerifiableDataRegistryModule struct{}

func (t *TestVerifiableDataRegistryModule) GetVerifiableDataRegistry() (types.VerifiableDataRegistry, error) {
	return &TestVerifiableDataRegistry{}, nil
}

type TestVerifiableDataRegistry struct{}

func (r *TestVerifiableDataRegistry) Get(identifier *types.DataIdentifier) (*types.VDROutput, error) {
	return nil, nil
}
func (r *TestVerifiableDataRegistry) Put(identifier *types.DataIdentifier, data io.Reader) (*types.DataIdentifier, error) {
	return nil, nil
}
func (r *TestVerifiableDataRegistry) Update(identifier *types.DataIdentifier, data io.Reader) (*types.DataIdentifier, error) {
	return nil, nil
}
func (r *TestVerifiableDataRegistry) Delete(identifier *types.DataIdentifier) error {
	return nil
}
func (r *TestVerifiableDataRegistry) IsAlive() bool {
	return true
}
func (r *TestVerifiableDataRegistry) Configure(interface{}) error {
	return nil
}

func (r *TestVerifiableDataRegistry) List() ([]*types.DataIdentifier, error) {
	return nil, nil
}
