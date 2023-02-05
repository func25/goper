package development

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type GenVersion interface {
	GenVersion(data any) error
}

type Version string

func (v *Version) GenVersion(data any) error {
	version, err := genHash(data)
	if err != nil {
		return err
	}

	*v = Version(version)
	return nil
}

func genHash(data any) (string, error) {
	raw, err := json.Marshal(data)
	return fmt.Sprintf("%x", md5.Sum(raw)), err
}
