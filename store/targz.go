package store

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type tarGz struct {
	fs afero.Fs
}

func (t *tarGz) extractTarGz(dest string, r io.Reader) error {
	if t.fs == nil {
		return errors.New("fs is nil")
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// no more files
		case err == io.EOF:
			return nil

		// unknown error
		case err != nil:
			return err

		// no header
		case header == nil:
			continue

		default:
			if err := t.extractItem(tr, dest, header); err != nil {
				return errors.Wrap(err, "extract item from tar")
			}
		}
	}
}

func (t *tarGz) extractItem(tr *tar.Reader, dest string, header *tar.Header) error {
	target := filepath.Join(dest, header.Name)

	switch header.Typeflag {

	// ensure dir exists
	case tar.TypeDir:
		if _, err := t.fs.Stat(target); err != nil {
			if err := t.fs.MkdirAll(target, dirMode); err != nil {
				return err
			}
		}

	case tar.TypeReg:
		f, err := t.fs.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, tr); err != nil {
			return err
		}
	}

	return nil
}
