package conf

import (
	"github.com/pkg/errors"
)

type Block struct {
	ProjectName       string
	HostAlias         string
	HostName          string
	PortNumber        string
	AdditionalOptions []string
}

func (b *Block) SetBlock(block []interface{}, additional []string) error {
	var ok bool
	b.ProjectName, ok = block[0].(string)
	if !ok {
		return errors.New("failed to convert project name into string")
	}

	b.HostAlias, ok = block[3].(string)
	if !ok {
		return errors.New("failed to convert host alias into string")
	}

	b.HostName, ok = block[5].(string)
	if !ok {
		return errors.New("failed to convert host name into string")
	}

	b.PortNumber, ok = block[6].(string)
	if !ok {
		return errors.New("failed to convert port nunber into string")
	}

	b.AdditionalOptions = additional

	return nil
}
