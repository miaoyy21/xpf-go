package md

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

type TargetDirectory string

var (
	TargetDirectoryFile      TargetDirectory = "file"
	TargetDirectoryShare     TargetDirectory = "share"
	TargetDirectoryThumbnail TargetDirectory = "thumbnail"
)

// Create Directory
func (s TargetDirectory) createDirectory(userId int64) error {
	dir := filepath.Join("store", string(s), strconv.FormatInt(userId, 10))
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

// Get User's File Name
func (s TargetDirectory) getPath(userId int64, fileId int64) string {
	if s == TargetDirectoryShare {
		return filepath.Join("store", string(s), strconv.FormatInt(fileId, 10))
	}

	return filepath.Join("store", string(s), strconv.FormatInt(userId, 10), strconv.FormatInt(fileId, 10))
}

// Write File
func (s TargetDirectory) writeFile(userId int64, fileId int64, bytes []byte) error {
	err := os.WriteFile(s.getPath(userId, fileId), bytes, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Delete File
func (s TargetDirectory) deleteFile(userId int64, fileId int64) error {
	err := os.Remove(s.getPath(userId, fileId))
	if err != nil {
		return err
	}

	return nil
}

// Copy File
func (s TargetDirectory) copyFile(d TargetDirectory, userId int64, fileId int64) error {

	// Src
	sRead, err := os.OpenFile(s.getPath(userId, fileId), os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return err
	}

	// Dst
	dWrite, err := os.Create(d.getPath(userId, fileId))
	if err != nil {
		return err
	}

	// Copy
	if _, err := io.Copy(dWrite, sRead); err != nil {
		return err
	}

	return nil
}
