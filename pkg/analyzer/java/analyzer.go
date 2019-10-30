package java

import (
	"os"

	"github.com/stackrox/scanner/pkg/analyzer"
	"github.com/stackrox/scanner/pkg/component"
	"github.com/stackrox/scanner/pkg/tarutil"
)

type analyzerImpl struct{}

func (a analyzerImpl) Match(fullPath string, fileInfo os.FileInfo) bool {
	return match(fullPath)
}

func match(fullPath string) bool {
	return javaRegexp.MatchString(fullPath)
}

func (a analyzerImpl) Analyze(fileMap tarutil.FilesMap) ([]*component.Component, error) {
	var allComponents []*component.Component
	for filePath, contents := range fileMap {
		if !match(filePath) {
			continue
		}
		packages, err := parseContents(filePath, contents)
		if err != nil {
			return nil, err
		}
		for _, p := range packages {
			allComponents = append(allComponents, &component.Component{
				JavaPkgMetadata: p,
			})
		}
	}
	return allComponents, nil
}

func Analyzer() analyzer.Analyzer {
	return analyzerImpl{}
}
