package utils

import (
	"errors"
	"github.com/pkg/sftp"
	"github.com/samber/lo"
	"io"
	"os"
	"path/filepath"
)

func GetRecursively(sftpClient *sftp.Client, remotePath string, localPath string, skipDir []string) (fs []string, err error) {
	remoteWalker := sftpClient.Walk(remotePath)
	if remoteWalker == nil {
		err = errors.New("SFTP Walker Error")
		return
	}
	for remoteWalker.Step() {
		err = remoteWalker.Err()
		if err != nil {
			return
		}
		remoteFullFilepath := remoteWalker.Path()
		localFilepath, _ := getRecursivelyPath(localPath, remotePath, remoteFullFilepath)
		if remoteWalker.Stat().IsDir() {
			localStat, localStatErr := os.Stat(localFilepath)
			if lo.IndexOf(skipDir, localStat.Name()) > -1 {
				remoteWalker.SkipDir()
				continue
			}
			// 存在するかつディレクトリではない場合エラー
			if !os.IsNotExist(localStatErr) && !localStat.IsDir() {
				err = errors.New("Cannot create a directry when that file already exists")
				return
			}
			mode := remoteWalker.Stat().Mode()
			if os.IsNotExist(localStatErr) {

				err = os.Mkdir(localFilepath, mode)
				if err != nil {
					return
				}
			}
			continue
		}
		_, err = getTransfer(sftpClient, localFilepath, remoteFullFilepath, nil, nil)
		if err != nil {
			return
		}
		fs = append(fs, localFilepath)
	}

	return
}

func getRecursivelyPath(localPath string, remotePath string, remoteFullFilepath string) (string, error) {
	rel, err := filepath.Rel(filepath.Clean(remotePath), remoteFullFilepath)
	if err != nil {
		return "", err
	}
	localFilepath := filepath.Join(localPath, rel)
	return filepath.ToSlash(localFilepath), nil
}

type IOReaderProgress struct {
	io.Reader
	TransferredBytes *int64 // Total of bytes transferred
}

// getTransfer Download Transfer execute
func getTransfer(client *sftp.Client, localFilepath string, remoteFilepath string, tfBytes *int64, totalBytes *int64) (int64, error) {
	localFile, localFileErr := os.Create(localFilepath)
	if localFileErr != nil {
		return 0, errors.New("localFileErr: " + localFileErr.Error())
	}
	defer localFile.Close()

	remoteFile, remoteFileErr := client.Open(remoteFilepath)
	if remoteFileErr != nil {
		return 0, errors.New("remoteFileErr: " + remoteFileErr.Error())
	}
	defer remoteFile.Close()

	var bytes int64
	var copyErr error
	// withProgress
	if totalBytes != nil {
		f, _ := remoteFile.Stat()
		*totalBytes = f.Size()
	}
	if tfBytes != nil {
		remoteFileWithProgress := &IOReaderProgress{Reader: remoteFile, TransferredBytes: tfBytes}
		bytes, copyErr = io.Copy(localFile, remoteFileWithProgress)
	} else {
		bytes, copyErr = io.Copy(localFile, remoteFile)
	}
	if copyErr != nil {
		return 0, errors.New("copyErr: " + copyErr.Error())
	}

	syncErr := localFile.Sync()
	if syncErr != nil {
		return 0, errors.New("syncErr: " + syncErr.Error())
	}
	return bytes, nil
}
