package backup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

//HACK:
// There are many duplications. braking function for each use case may cause these redundancy
// Consider adapting abstractions based on the evolving progress of this project.
// However, avoid abstracting too hastily, as incorrect abstractions can make code maintenance more challenging.

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

func (s *BackupService) BackupTradeFull() error {
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

// BackupTradeDiff backs up the differential trade data of last full backup to the local storage.
func (s *BackupService) BackupTradeDiff() error {
	panic("not implemented") // TODO: Implement
}

// BackupTradeFullToRemote backs up all trade data to a remote storage.
func (s *BackupService) BackupTradeFullToRemote() error {
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

// BackupTradeDiffToRemote backs up the differential trade data to a remote storage.
func (s *BackupService) BackupTradeDiffToRemote() error {
	panic("not implemented") // TODO: Implement
}

// BackupProductFull backs up all trade data
// related to a specific product (identified by 'id') to the local storage.
func (s *BackupService) BackupProductFull(id string) error {
	ctx := context.TODO()
	tx := s.txCreator.CreateTransaction(ctx)

	isStored, err := s.metadataRepo.IsProductStored(ctx, tx.TransactionExtractor(), id)

	if err != nil {
		return err
	}
	if !isStored {
		return fmt.Errorf("find product: %s is not stored", id)
	}

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
		} // BackupTradeFull backs up all trade data to the local storage.

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

// BackupProductDiff backs up the differential trade data of last full backup
// related to a specific product (identified by 'id') to the local storage.
func (s *BackupService) BackupProductDiff(id string) error {
	panic("not implemented") // TODO: Implement
}

// BackupProductFullToRemote backs up all trade data
// related to a specific product (identified by 'id') to a remote storage.
func (s *BackupService) BackupProductFullToRemote(id string) error {
	ctx := context.TODO()
	tx := s.txCreator.CreateTransaction(ctx)

	isStored, err := s.metadataRepo.IsProductStored(ctx, tx.TransactionExtractor(), id)

	if err != nil {
		return err
	}
	if !isStored {
		return fmt.Errorf("find product: %s is not stored", id)
	}

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

// BackupProductDiffToRemote backs up the differential trade data
// related to a specific product (identified by 'id') to a remote storage.
func (s *BackupService) BackupProductDiffToRemote(id string) error {
	panic("not implemented") // TODO: Implement
}
