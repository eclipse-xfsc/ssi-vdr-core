package types

import (
	"errors"
	"fmt"
	"io"
)

type VerifiableDataRegistry interface {
	Get(*DataIdentifier) (*VDROutput, error)
	Put(*DataIdentifier, io.Reader) (*DataIdentifier, error)
	Update(*DataIdentifier, io.Reader) (*DataIdentifier, error)
	Delete(*DataIdentifier) error
	IsAlive() bool
	Configure(interface{}) error
	List() ([]*DataIdentifier, error)
}

type VerifiableDataRegistryModule interface {
	GetVerifiableDataRegistry() (VerifiableDataRegistry, error)
}

type VDROutput struct {
	Data []byte
}

type DataIdentifier struct {
	Format string
	Value  string
}

func (id *DataIdentifier) String() string {
	return fmt.Sprintf("DataIdentifier Format: %s Value: %s", id.Format, id.Value)
}

var DataIdentifierNotFound = errors.New("DataIdentifier not found")
