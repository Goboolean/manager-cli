package backup

import (
	"context"
	"time"
)

func (s BackupService) BackupData() error {
	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	_, err := s.tradeRepo.DumpTradeRepo()

	if err == nil {
		s.tx.Commit()
		return nil
	} else {
		s.tx.Rollback()
		return err
	}
}

func (s BackupService) BackupDataBefore(time time.Time) error {
	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	_, err := s.tradeRepo.DumpTradeRepoBefore(time)

	if err == nil {
		s.tx.Commit()
		return err
	} else {
		s.tx.Rollback()
		return nil
	}
}

func (s BackupService) BackupDataToRemote() error {
	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	f, err := s.tradeRepo.DumpTradeRepo()
	defer s.fileRemover.RemoveFile(f)

	if err != nil {
		s.tx.Rollback()
		return err
	}

	err = s.transmit.TransmitDataToRemote(f)

	if err != nil {
		s.tx.Rollback()
		return err
	}

	s.tx.Commit()
	return nil

}

func (s BackupService) BackupDataToRemoteBefore(time time.Time) error {
	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	f, err := s.tradeRepo.DumpTradeRepoBefore(time)
	defer s.fileRemover.RemoveFile(f)

	if err != nil {
		s.tx.Rollback()
		return err
	}

	err = s.transmit.TransmitDataToRemote(f)

	if err != nil {
		s.tx.Rollback()
		return err
	}

	s.tx.Commit()
	return nil
}

func (s BackupService) BackupProduct(id string) error {
	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	_, err := s.tradeRepo.DumpProduct(id)

	if err != nil {
		s.tx.Rollback()
		return err
	} else {
		s.tx.Commit()
		return nil
	}

}

func (s BackupService) BackupProductToRemote(id string) error {

	ctx, _ := context.WithCancel(context.Background())
	s.tx.Begin(ctx)

	f, err := s.tradeRepo.DumpProduct(id)
	defer s.fileRemover.RemoveFile(f)

	if err != nil {
		s.tx.Rollback()
		return err
	}

	err = s.transmit.TransmitDataToRemote(f)

	if err != nil {
		s.tx.Rollback()
		return err
	} else {
		s.tx.Commit()
		return nil
	}

}
