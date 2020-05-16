package main

import (
	_ "bytes"
	"github.com/DATA-DOG/godog/gherkin"
	"net"
	"net/http/httptest"
	"walk-the-camino/tests/functional/testutil"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

var host string
var port string

var status bool
var tcpConn net.Conn

// TestCases1 testcases
type TestCases1 struct {
	UserStoryID    string   `json:"userStoryID"`
	TestCaseName   string   `json:"testCaseName"`
	TestStatus     string   `json:"testStatus"`
	TestUpdateFlag string   `json:"testUpdateFlag"`
	TestSteps      []Steps1 `json:"testSteps"`
}

// Steps1 steps
type Steps1 struct {
	StepIndex          string `json:"stepIndex"`
	StepDescription    string `json:"stepDescription"`
	StepExpectedResult string `json:"stepExpectedResult"`
}

// TestCasesUpdate test case updates
type TestCasesUpdate struct {
	UserStoryID  string   `json:"userStoryID"`
	TestCaseName string   `json:"testCaseName"`
	TestStatus   string   `json:"testStatus"`
	Attachments  []string `json:"attachments"`
}

var testCases []TestCases1
var testCasesUpdate []TestCasesUpdate
var currentFeature *gherkin.Feature
var userStoryTCNameMap []testutil.gitProjectsUserStoryTCNameMapping

func theUserstory(userStory string) error {
	return nil
}

func main() {
	/**

	 */
}
