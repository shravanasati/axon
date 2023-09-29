package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func itemInSlice[T comparable](s T, slice []T) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func filter[T any](function func(arg T) bool, slice []T) []T {
	var newSlice []T
	for _, v := range slice {
		if function(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func validPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// FileOrganizer is the base struct for all the functions to organize files.
type FileOrganizer struct {
	path    string
	actions []string
	regex   *regexp.Regexp
}

func NewFileOrganizer(path string, regex *regexp.Regexp) *FileOrganizer {
	return &FileOrganizer{
		path:  path,
		regex: regex,
	}
}
func (fo *FileOrganizer) getFilesInFolder() ([]fs.DirEntry, error) {
	files, err := os.ReadDir(fo.path)
	if err != nil {
		return nil, err
	}
	function := func(arg fs.DirEntry) bool { return !arg.IsDir() && fo.regex.MatchString(arg.Name()) }
	return filter(function, files), nil
}

func (fo *FileOrganizer) prettify(casing string) {
	os.Chdir(fo.path)
	files, err := fo.getFilesInFolder()
	if err != nil {
		fmt.Println("unable to get files in the folder: ", fo.path)
		return
	}

	for _, file := range files {
		name := file.Name()
		if casing == "lower" {
			os.Rename(name, strings.ToLower(name))
		} else if casing == "upper" {
			os.Rename(name, strings.ToUpper(name))
		} else {
			os.Rename(name, strings.Title(name))
		}
	}
	fo.actions = append(fo.actions, "Prettified the directory to "+casing)
}

func (fo *FileOrganizer) createDirs() {
	os.Chdir(fo.path)
	dirs := [...]string{"Images", "Music", "Videos", "Programs", "Compressed Files", "PDFs", "Others"}
	for _, dir := range dirs {
		localDir := fmt.Sprintf("./%s", dir)
		if !validPath(localDir) {
			os.Mkdir(localDir, os.ModePerm)
			fo.actions = append(fo.actions, fmt.Sprintf("Created '%s' directory in '%s'.", dir, fo.path))
		}
	}
}

func (fo *FileOrganizer) organize() {
	os.Chdir(fo.path)
	files, err := fo.getFilesInFolder()
	if err != nil {
		fmt.Println("unable to get files in the folder: ", fo.path)
		return
	}

	imageExt := []string{"jpg", "jpeg", "png", "jfif"}
	musicExt := []string{"mp3", "aac", "ogg", "wav"}
	videoExt := []string{"mp4", "webm", "mov", "mkv", "mpv2", "avi", "gif"}
	programExt := []string{"exe", "msi", "msp", "dll", "out"}
	compressedExt := []string{"rar", "zip", "7z", "tar"}
	pdfExt := []string{"pdf"}
	folders := []string{"compressed files", "programs", "videos", "music", "others", "images"}

	fo.createDirs()

	for _, file := range files {
		filename := file.Name()
		if itemInSlice(strings.ToLower(filename), folders) {
			continue
		}

		split := strings.Split(filename, ".")
		if len(split) == 1 {
			continue
		}
		ext := strings.ToLower(split[len(split)-1])

		if itemInSlice(ext, imageExt) {
			os.Rename(filename, filepath.Join("Images", filename))
		} else if itemInSlice(ext, musicExt) {
			os.Rename(filename, filepath.Join("Music", filename))
		} else if itemInSlice(ext, videoExt) {
			os.Rename(filename, filepath.Join("Videos", filename))
		} else if itemInSlice(ext, compressedExt) {
			os.Rename(filename, filepath.Join("Compressed Files", filename))
		} else if itemInSlice(ext, programExt) {
			os.Rename(filename, filepath.Join("Programs", filename))
		} else if itemInSlice(ext, pdfExt) {
			os.Rename(filename, filepath.Join("PDFs", filename))
		} else {
			os.Rename(filename, filepath.Join("Others", filename))
		}

	}
	fo.actions = append(fo.actions, "Organized the files in "+fo.path)
}

func (fo *FileOrganizer) showActions() string {
	fmt.Println("-----------------------------------------")
	data := fmt.Sprintf("\nAction report for `%s` (%d actions performed): \n", fo.path, len(fo.actions))

	for _, action := range fo.actions {
		data += action + "\n"
	}

	return data
}

func (fo *FileOrganizer) renameDir(newName string) {
	os.Chdir(fo.path)
	files, err := fo.getFilesInFolder()
	if err != nil {
		fmt.Println("unable to get files in ", fo.path)
		return
	}

	for i, file := range files {
		split := strings.Split(file.Name(), ".")
		ext := strings.ToLower(split[len(split)-1])

		os.Rename(file.Name(), fmt.Sprintf("%v %v.%v", newName, i+1, ext))
		// fo.actions = append(fo.actions, fmt.Sprintf("Renamed %s to %s"))
	}
}

func (fo *FileOrganizer) move(targetDir string) {
	files, err := fo.getFilesInFolder()
	if err != nil {
		fmt.Println("unable to get files in ", fo.path)
		return
	}
	if !validPath(targetDir) {
		fmt.Printf("the target directory `%s` doesnt exist\n", targetDir)
		return
	}

	for _, file := range files {
		newpath := filepath.Join(targetDir, file.Name())
		oldpath := filepath.Join(fo.path, file.Name())
		os.Rename(oldpath, newpath)
		fo.actions = append(fo.actions, fmt.Sprintf("Moved %s to %s.", oldpath, newpath))
	}
}
