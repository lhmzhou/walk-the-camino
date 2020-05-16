package main

import (
	"errors"
	"fmt"
	_ "log"
	"net/http/httptest"
	"walk-the-camino/tests/functional/testutil"

	"github.com/DATA-DOG/godog"
)

var reportLocation string

func fileExists(reportOutput string) error {
	err := testutil.ReportFileExists(reportOutput)
	reportLocation = "report/in/" + reportOutput
	if err != nil {
		return errors.New("unable to find test output for the current run to upload in to gitProjects. missing output.json")
	}
	err = testutil.ReportFileExists("gitProjectsmap.json")
	if err != nil {
		return errors.New("unable to find gitProjects userstory scenario mapping. missing gitProjectsmap.json")
	}
	return nil
}

func (a *apiFeature) createATestIngitProjectsUnderTheUserstoryAndUploadTestResult() (err error) {
	//fmt.Println("reportLocation", reportLocation)
	err = testutil.UploadAttachmentsTogitProjectsServer()
	if err != nil {
		fmt.Println("unable to upload attachments to gitProjects server ", err)
		//return err
	}
	a.resp = httptest.NewRecorder()
	gitProjectsCreateJSON, gitProjectsUpdateJSON, err := testutil.CreategitProjectsJsons(reportLocation)
	if err != nil {
		return err
	}
	err = testutil.CreategitProjectsTestCase(a.resp, gitProjectsCreateJSON)
	if err != nil {
		fmt.Println("failed to CreategitProjectsTestCase ", err)
		return err
	}
	//z, _ := ioutil.ReadAll(a.resp.Body)
	if a.resp.Code != 200 {
		if a.resp.Code != 201 {
			return errors.New("error in creating test case in gitProjects" + a.resp.Body.String())
		}
		//return errors.New("error in creating test case in gitProjects" + a.resp.Body.String())
	}
	err = testutil.ValidateCreateResponse(a.resp.Body.Bytes())
	if err != nil {
		return err
	}

	a.resp = httptest.NewRecorder()
	err = testutil.UpdategitProjectsTestCase(a.resp, gitProjectsUpdateJSON)
	if err != nil {
		fmt.Println("failed to UpdategitProjectsTestCase ", err)
		return err
	}
	if a.resp.Code != 200 {
		if a.resp.Code != 201 {
			return errors.New("error in updating test results in gitProjects" + a.resp.Body.String())
		}
		//return errors.New("error in updating test results in gitProjects" + a.resp.Body.String())
	}
	//z, _ := ioutil.ReadAll(a.resp.Body)
	err = testutil.ValidateUpdateResponse(a.resp.Body.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

func gitProjectsFeatureContext(s *godog.Suite) {
	api := &apiFeature{}
	s.BeforeScenario(api.resetResponse)
	s.Step(`^"([^"]*)" exists$`, fileExists)
	s.Step(`^create gitProjects jsons,test cases in gitProjects from user story scenario mapping file and upload test result$`, api.createATestIngitProjectsUnderTheUserstoryAndUploadTestResult)
}
