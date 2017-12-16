package executor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
)

func Execute(files, outputPrefix, inPlace, shortPath, overwrite string) error {
	errChan := make(chan error, 0)
	wg := &sync.WaitGroup{}

	for _, file := range strings.Split(files, ",") {
		wg.Add(1)
		go translate(shortPath, file, outputPrefix, inPlace, overwrite, errChan, wg)
	}

	done := make(chan bool, 0)

	go func() {
		wg.Wait()
		done <- true
	}()

	log.Info("Waiting for all translations to complete")

	select {
	case err := <-errChan:
		return err
	case <-done:
		log.Info("Successfully translated all input files")
		return nil
	}
}

func translate(shortPath, file, prefix, inPlace, overwrite string, errChan chan error, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if file == "" {
		return
	}

	log.Infof("Translating file=[%s] prefix=[%s] inPlace=[%s] overwrite=[%s]", file, prefix, inPlace, overwrite)

	cmd := "short"
	if shortPath != "" {
		cmd = shortPath
	}
	args := []string{"-k", "-f", file}

	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		errChan <- err
		log.Errorf("Error translating file %s", output)
		return
	}

	outputFileName := fmt.Sprintf("%s%s", prefix, filepath.Base(file))
	outputFileDir := filepath.Dir(file)

	outputFile := filepath.Join(outputFileDir, outputFileName)
	if inPlace == "true" {
		outputFile = file
	}

	if overwrite != "true" {
		if fileExists(outputFile) {
			errChan <- fmt.Errorf("Output file %s, already exists, and --overwrite is false", outputFile)
			return
		}
	}

	if err := ioutil.WriteFile(outputFile, output, 0644); err != nil {
		errChan <- err
		return
	}
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
