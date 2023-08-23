package backup

import (
	"context"
	"strings"
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

func (s *BackupService) getStoredProducts() ([]string, error) {
	ctx := context.TODO()
	tx := s.txCreator.CreateTransaction(ctx)

	idList, err := s.metadataRepo.GetStoredProductList(ctx, tx.TransactionExtractor())
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return idList, err
}

func (s *BackupService) BackupData() error {

	productToBackup, err := s.getStoredProducts()
	if err != nil {
		return err
	}

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")
	meta := entity.BackupMeta{
		BackupType:   entity.DiffBackup,
		BackupDbList: productToBackup,
		Timestamp:    now.Unix(),
		Date:         now.Format(toolTimeFormatString),
		HashVer:      hashVer,
		FileList:     []entity.FileNameWithHash{},
	}

	for _, productId := range productToBackup {
		f, err := s.tradeDumper.DumpProductBefore(productId, out, now)
		if err != nil {
			return err
		}

		for i := range f {
			h, err := s.fileOperator.CalculateFileHash(f[i])
			if err != nil {
				return err
			}

			meta.FileList = append(meta.FileList, entity.FileNameWithHash{
				Name: f[i].Name,
				Hash: h,
			})
		}
	}

	metaFile := entity.File{
		Name:    "metadata.json",
		DirPath: out,
	}

	err = s.backupMetaPort.StoreBackupMeta(meta, metaFile)
	if err != nil {
		return err
	}

	return nil
}

func (s *BackupService) BackupDataToRemote() error {

	productToBackup, err := s.getStoredProducts()
	if err != nil {
		return err
	}

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")
	remoteDir := "/" + now.Format(toolTimeFormatString)
	meta := entity.BackupMeta{
		BackupType:   entity.DiffBackup,
		BackupDbList: productToBackup,
		Timestamp:    now.Unix(),
		Date:         now.Format(toolTimeFormatString),
		HashVer:      hashVer,
		FileList:     []entity.FileNameWithHash{},
	}

	s.transmitter.CreateRemoteDir(remoteDir)

	for _, productId := range productToBackup {
		f, err := s.tradeDumper.DumpProductBefore(productId, out, now)
		if err != nil {
			return err
		}

		for i := range f {

			h, err := s.fileOperator.CalculateFileHash(f[i])
			if err != nil {
				return err
			}

			err = s.transmitter.TransmitDataToRemote(f[i], remoteDir)
			if err != nil {
				return err
			}

			meta.FileList = append(meta.FileList, entity.FileNameWithHash{
				Name: f[i].Name,
				Hash: h,
			})
		}
	}

	metaFile := entity.File{
		Name:    "metadata.json",
		DirPath: out,
	}

	s.backupMetaPort.StoreBackupMeta(meta, metaFile)
	if err != nil {
		return err
	}

	s.transmitter.TransmitDataToRemote(metaFile, remoteDir)
	if err != nil {
		return err
	}

	return nil
}

func (s *BackupService) BackupProduct(id string) error {

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")
	meta := entity.BackupMeta{
		BackupType:   entity.FullBackup,
		Timestamp:    now.Unix(),
		Date:         now.Format(toolTimeFormatString),
		HashVer:      hashVer,
		BackupDbList: []string{id},
	}

	f, err := s.tradeDumper.DumpProductBefore(id, out, now)
	if err != nil {
		return err
	}

	for i := range f {
		h, err := s.fileOperator.CalculateFileHash(f[i])
		if err != nil {
			return err
		}

		meta.FileList = append(meta.FileList, entity.FileNameWithHash{
			Name: f[i].Name,
			Hash: h,
		})
	}

	MetadataFile := entity.File{
		Name:    "metadata.json",
		DirPath: out,
	}

	s.backupMetaPort.StoreBackupMeta(meta, MetadataFile)
	if err != nil {
		return err
	}

	return nil
}

func (s *BackupService) BackupProductToRemote(id string) error {

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")
	remoteDir := "/" + now.Format(toolTimeFormatString)
	meta := entity.BackupMeta{
		BackupType:   entity.FullBackup,
		Timestamp:    now.Unix(),
		Date:         now.Format(toolTimeFormatString),
		HashVer:      hashVer,
		BackupDbList: []string{id},
	}

	s.transmitter.CreateRemoteDir(remoteDir)

	f, err := s.tradeDumper.DumpProductBefore(id, out, now)
	if err != nil {
		return err
	}

	for i := range f {
		err = s.transmitter.TransmitDataToRemote(f[i], remoteDir)
		if err != nil {
			return err
		}

		h, err := s.fileOperator.CalculateFileHash(f[i])

		if err != nil {
			return err
		}

		meta.FileList = append(meta.FileList, entity.FileNameWithHash{
			Name: f[i].Name,
			Hash: h,
		})
	}

	MetadataFile := entity.File{
		Name:    "metadata.json",
		DirPath: out,
	}

	s.backupMetaPort.StoreBackupMeta(meta, MetadataFile)
	if err != nil {
		return err
	}

	s.transmitter.TransmitDataToRemote(MetadataFile, remoteDir)
	if err != nil {
		return err
	}

	return nil
}
