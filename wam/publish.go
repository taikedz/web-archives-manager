package wam

import (
	"fmt"
)

type Publisher struct {
	label string
	version string
	files []string
}

func (pp *Publisher) Push_with(readme string, readme_altname string) error {
	if err := pp.pushfile(readme, readme_altname); err != nil {
		return err
	}
	if err := pp.Push(); err != nil {
		return err
	}
}

func (pp *Publisher) Push() error {
	tarball_path, err := Tarball_make(pp.files)
	if err != nil {
		return err
	}
	if err = pp.pushfile(tarball_path, fmt.Sprintf("%s-%s", label, version)); err != nil {
		return err
	}
}

func (pp *Publisher) pushfile(prefix string, source_path string, dest_path string) error