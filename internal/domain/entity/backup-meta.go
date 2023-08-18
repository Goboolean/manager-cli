package entity

const (
	FullBackup = "full"
	DiffBackup = "diff"
)

type BackupMeta struct {
	BackupType   string           `json:"backupType"` // Ex:)full,diff
	BackupDbList []string         `json:"backupDb"`   // Database list
	Timestamp    int64            `json:"timestamp"`  // Unix timestamp of backup time
	Date         string           `json:"date"`       //Human readable datetime of backup time
	HashVer      string           `json:"hashVer"`    // Version of hash algorithm
	FileInfoList []BackupFileInfo `json:"fileInfo"`
}

type BackupFileInfo struct {
	Name string `json:"fileName"`
	Hash string `json:"fileHash"`
}
