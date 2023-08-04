package backup

import (
	"context"
	"strings"
	"time"
)

func (s *BackupService) getStoredProducts() ([]string, error) {
	tx := s.txCreator.CreateTransaction(context.TODO())

	idList, err := s.metadataRepo.GetStoredProductList(tx.TransactionExtractor())
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

	for _, productId := range productToBackup {

		_, err := s.tradeDumper.DumpProductBefore(productId, out, now)
		if err != nil {
			return err
		}

	}

	// TODO: write metadata to file
	// - Backup time in unix time stamp and human readable
	// - Backup type
	// - relational directory that actually contains dumped collection file from s.out/{datetime} directory
	// - file list and its hash value

	return nil
}

func (s *BackupService) BackupDataToRemote() error {

	productToBackup, err := s.getStoredProducts()
	if err != nil {
		return err
	}

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")

	for _, productId := range productToBackup {
		f, err := s.tradeDumper.DumpProductBefore(productId, out, now)
		if err != nil {
			return err
		}

		err = s.transmitter.TransmitDataToRemote(f)
		if err != nil {
			return err
		}

	}

	// TODO: write metadata to file
	// - Backup time in unix time stamp and human readable
	// - Backup type
	// - relational directory that actually contains dumped collection file from s.out/{datetime} directory
	// - file list and its hash value

	// TODO: transmit metadata file to remote

	return nil
}

func (s *BackupService) BackupProduct(id string) error {

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")

	_, err := s.tradeDumper.DumpProductBefore(id, out, now)
	if err != nil {
		return err
	}

	// TODO: write metadata to file
	// - Backup time in unix time stamp and human readable
	// - Backup type
	// - relational directory that actually contains dumped collection file from s.out/{datetime} directory
	// - file list and its hash value

	// TODO: transmit metadata file to remote

	return nil
}

func (s *BackupService) BackupProductToRemote(id string) error {

	now := time.Now()
	out := strings.Join([]string{s.backUpDir, now.Format(toolTimeFormatString)}, "/")

	f, err := s.tradeDumper.DumpProductBefore(id, out, now)
	if err != nil {
		return err
	}

	err = s.transmitter.TransmitDataToRemote(f)
	if err != nil {
		return err
	}

	// TODO: write metadata to file
	// - Backup time in unix time stamp and human readable
	// - Backup type
	// - relational directory that actually contains dumped collection file from s.out/{datetime} directory
	// - file list and its hash value

	// TODO: transmit metadata file to remote

	return nil
}
