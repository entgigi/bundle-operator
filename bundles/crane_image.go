package bundles

import (
	"archive/tar"
	"fmt"
	"io"
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func ExtractImageTo(repoSrc string, pathDir string) error {
	pathTarFile := generateFileTarPath(pathDir)
	f, err := os.Create(pathTarFile)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", pathTarFile, err)
	}
	defer f.Close()

	// pull image
	var img v1.Image
	var options *[]crane.Option = &[]crane.Option{}
	img, err = crane.Pull(repoSrc, *options...)
	if err != nil {
		return fmt.Errorf("pulling %s: %w", repoSrc, err)
	}

	// mutate extract to pathDir.tar
	err = crane.Export(img, f)
	if err != nil {
		return fmt.Errorf("failed to export %s: %w", repoSrc, err)
	}

	// untar pathDir.tar
	return extractTar(pathDir, f)

}

func generateFileTarPath(pathDir string) string {

	return pathDir + ".tar"
}

func extractTar(pathDir string, tarFile *os.File) error {
	var fileReader io.ReadCloser = tarFile
	tarBallReader := tar.NewReader(fileReader)
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				return err
			}

		case tar.TypeReg:
			// handle normal file
			fmt.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				return err
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				return err
			}

			writer.Close()
		default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}
	return nil
}
