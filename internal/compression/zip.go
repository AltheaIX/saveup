package compression

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Zip(sourceDir, destination string) error {
	zipFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(
		sourceDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip the root directory itself.
			if path == sourceDir {
				return nil
			}

			// Relative path inside the zip.
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}

			// ZIP format always uses forward slashes.
			relPath = filepath.ToSlash(relPath)

			// Create header.
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = relPath

			// Preserve directory entries.
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		},
	)
}
