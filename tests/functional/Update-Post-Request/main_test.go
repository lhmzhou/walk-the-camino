package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/DATA-DOG/godog/gherkin"
	"io"
	"os"
	"path/filepath"
	"walk-the-camino/tests/functional/testutil"
	"testing"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}
var out io.Writer

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()
	if opt.Tags != "gitProjects" {
		var f *os.File
		fileCreationPath, _ := filepath.Abs("report/in")
		err := os.MkdirAll(fileCreationPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		fileCreationPath, err = filepath.Abs("report/in/report.json")
		fmt.Println(fileCreationPath)
		f, err = os.Create(fileCreationPath)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		out = bufio.NewWriter(f)
		opt.Output = out
	}
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	if status != 0 {
		fmt.Println("Status is not 0. exit1. Original status: ", status)
		os.Exit(1)
	}
}

func createUserStoryTestCaseMapping(userStoryID string) error {
	return nil
}
func storegitProjectsUserStoryList(f *gherkin.Feature) {
	var err error
	if opt.Tags != "gitProjects" && opt.Tags != "regression" {
		userStoryTCNameMap, err = testutil.MapTestCasegitProjectsUserStoryMap(f, "functional")
		if err != nil {
			fmt.Println("unable to create gitProjects userstory map. gitProjects update could be incorrect")
		}
	} else if opt.Tags != "gitProjects" && opt.Tags == "regression" {
		userStoryTCNameMap, err = testutil.MapTestCasegitProjectsUserStoryMap(f, "regression")
		if err != nil {
			fmt.Println("unable to create gitProjects userstory map. gitProjects update could be incorrect")
		}
	}
}

func afterFeature(ft *gherkin.Feature) {

}

func FeatureContext(s *godog.Suite) {
	//s.BeforeFeature(storegitProjectsUserStoryList)
	PostFeatureContext(s)
	//gitProjectsFeatureContext(s)
	//s.Step(`^create gitProjects test under userstory "([^"]*)" and upload test result$`, createUserStoryTestCaseMapping)
	s.AfterFeature(afterFeature)
}
