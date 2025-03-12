package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func organizeFiles(dir string) error {
	newDir, err := genRandDirName()
	if err != nil {
		return err
	}

	newDir = filepath.Join(dir, fmt.Sprintf("%s_organized", newDir))

	for {
		if _, err := os.Stat(newDir); os.IsNotExist(err) {

			err = os.Mkdir(newDir, 0755)
			if err != nil {
				return err
			}
			break
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var uniqueExtentions []string

	for _, file := range entries {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		extension := getExtension(path)

		if extension != "" && !slices.Contains(uniqueExtentions, extension) {
			uniqueExtentions = append(uniqueExtentions, extension)
		}
	}

	for _, ext := range uniqueExtentions {
		path := filepath.Join(newDir, ext)
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}

	newOrganizedDir, err := os.ReadDir(newDir)
	if err != nil {
		return err
	}

	for _, organizedDir := range newOrganizedDir {
		for _, file := range entries {

			if file.IsDir() {
				continue
			}

			ext := getExtension(filepath.Join(dir, file.Name()))

			newDirName := strings.TrimSuffix(organizedDir.Name(), "/")

			if ext == newDirName {
				_, err := os.Create(filepath.Join(newDir, ext, filepath.Base(file.Name())))
				if err != nil {
					continue
				}

			}

		}
	}

	return nil
}

func getExtension(path string) string {
	extension := filepath.Ext(path)

	extension = strings.TrimPrefix(extension, ".")
	extension = strings.ToLower(extension)

	return extension
}

func genRandDirName() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(buf), err
}
