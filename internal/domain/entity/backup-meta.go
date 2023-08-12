package entity

type BackupMeta struct {
	BackupType   string   `json:"backupType"` // Ex:)full,diff
	BackupDbList []string `json:"backupDb"`   // Database list
	Timestamp    int      `json:"timestamp"`  // Unix timestamp of backup time
	Date         string   `json:"date"`       //Human readable datetime of backup time
	HashVer      string   `json:"hashVer"`    // Version of hash algorithm
	FileInfoList []struct {
		Name string `json:"fileName"`
		Hash string `json:"fileHash"`
	} `json:"fileInfo"`
}
