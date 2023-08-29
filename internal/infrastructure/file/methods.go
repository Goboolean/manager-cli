package fileInf

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"os"

	"github.com/cespare/xxhash/v2"
)

// This method removes the file or directory at the specified path.
func (inf *FileInfra) RemoveFileOrDir(path string) error {
	return os.Remove(path)
}

func (inf *FileInfra) IsDirExist(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// This method retrieves a list of file names in the specified directory.
func (inf *FileInfra) GetFileNameList(dir string) ([]string, error) {
	dirInfo, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	res := make([]string, len(dirInfo))

	for _, info := range dirInfo {
		if !(info.IsDir()) {
			res = append(res, info.Name())
		}
	}

	return res, nil
}

func (inf *FileInfra) CreateDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// This method calculates the xxHash checksum for the file at the given path.
func (inf *FileInfra) CalculateXxhashChecksum(path string) (string, error) {

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	h := xxhash.New()

	for {

		bytesRead, err := f.Read(buf)

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		h.Write((buf[:bytesRead]))
	}

	return hex.EncodeToString(h.Sum(nil)), nil

}

// This method stores the given JSON data into a file at the specified path.
func (inf *FileInfra) StoreJsonToFile(data []byte, path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// This method decodes JSON data from the file at the specified path into the target interface.
func (inf *FileInfra) DecodeJsonFromFile(path string, target anyPointer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(target)
	if err != nil {
		return err
	}

	return nil
}
