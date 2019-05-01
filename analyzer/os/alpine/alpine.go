package alpine

import (
	"bufio"
	"bytes"
	"errors"

	"github.com/knqyf263/fanal/analyzer/os"

	"github.com/knqyf263/fanal/analyzer"
	"github.com/knqyf263/fanal/extractor"
)

func init() {
	analyzer.RegisterOSAnalyzer(&alpineOSAnalyzer{})
}

type alpineOSAnalyzer struct{}

func (a alpineOSAnalyzer) Analyze(fileMap extractor.FileMap) (analyzer.OS, error) {
	for _, filename := range a.RequiredFiles() {
		file, ok := fileMap[filename]
		if !ok {
			continue
		}
		scanner := bufio.NewScanner(bytes.NewBuffer(file))
		for scanner.Scan() {
			line := scanner.Text()
			return analyzer.OS{Family: os.Alpine, Name: line}, nil
		}
	}
	return analyzer.OS{}, errors.New("alpine: Not match")
}

func (a alpineOSAnalyzer) RequiredFiles() []string {
	return []string{"etc/alpine-release"}
}
