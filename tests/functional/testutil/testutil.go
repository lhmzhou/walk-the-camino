package testutil

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	"github.com/buger/jsonparser"

	_ "github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"log"

	_ "log"
	"mime/multipart"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var statusCode string

// CreateUpdateResponse gitProjects need the same
type CreateUpdateResponse struct {
	TestName     string `json:"testName"`
	UserStory    string `json:"userStory"`
	UpdateResult bool   `json:"updateResult"`
}

// TestCases1 testcase
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

// gitProjectsCreateJSON json
type gitProjectsCreateJSON struct {
	APIKey     string       `json:"apiKey"`
	Workspace  string       `json:"workspace"`
	Project    string       `json:"project"`
	TestFolder string       `json:"testFolder"`
	EmailID    string       `json:"emailId"`
	Iteration  string       `json:"iteration"`
	TestCases  []TestCases1 `json:"testCases"`
}

// gitProjectsUpdateJSON gitProjects update json
type gitProjectsUpdateJSON struct {
	APIKey          string            `json:"apiKey"`
	Workspace       string            `json:"workspace"`
	Project         string            `json:"project"`
	TestFolder      string            `json:"testFolder"`
	EmailID         string            `json:"emailId"`
	Iteration       string            `json:"iteration"`
	TestCasesUpdate []TestCasesUpdate `json:"testCases"`
}

// gitProjectsTestDataFromFile data
type gitProjectsTestDataFromFile struct {
	APIKey      string `json:"apiKey"`
	Workspace   string `json:"workspace"`
	Project     string `json:"project"`
	TestFolder  string `json:"testFolder"`
	EmailID     string `json:"emailId"`
	Iteration   string `json:"iteration"`
	UserStoryID string `json:"userStoryID"`
}

// gitProjectsUserStoryTCNameMapping story
type gitProjectsUserStoryTCNameMapping struct {
	TestCaseName string `json:"testCaseName"`
	UserStoryID  string `json:"userStoryId"`
}

// NEWOutputValidationGATAPIResponse gat
type NEWOutputValidationGATAPIResponse struct {
	Result   string `json:"result"`
	Failures []struct {
		Field  []string `json:"field"`
		Reason string   `json:"reason"`
	} `json:"failures"`
}

// OutputValidationGATAPIResponse output
type OutputValidationGATAPIResponse struct {
	Result   bool     `json:"result"`
	Failures []string `json:"failures"`
}

// gitProjectsAttachmentResponse response
type gitProjectsAttachmentResponse struct {
	Entity   string `json:"entity"`
	Status   int    `json:"status"`
	Metadata struct {
	} `json:"metadata"`
	Annotations                   interface{} `json:"annotations"`
	GenericType                   interface{} `json:"genericType"`
	PostProcessInterceptors       interface{} `json:"postProcessInterceptors"`
	MessageBodyWriterInterceptors interface{} `json:"messageBodyWriterInterceptors"`
	ResourceMethod                interface{} `json:"resourceMethod"`
	ResourceClass                 interface{} `json:"resourceClass"`
}

var testCases []TestCases1
var testCasesUpdate []TestCasesUpdate
var gitProjectsTestDataFromConfig gitProjectsTestDataFromFile
var crtFile, updFile string
var currentFeature *gherkin.Feature
var userStoryTCNameMap []gitProjectsUserStoryTCNameMapping

// FileExists file check
func FileExists(createFile, updateFile string) error {
	absPath, _ := filepath.Abs("gitProjects/" + createFile)
	_, err := os.Stat(absPath)
	if err == nil {
		crtFile = createFile
	} else if os.IsNotExist(err) {
		return fmt.Errorf("File does not exist"+createFile, err)
	} else {
		return fmt.Errorf("File does not exist"+createFile, err)
	}
	absPath, _ = filepath.Abs("gitProjects/" + updateFile)
	_, err = os.Stat(absPath)
	if err == nil {
		updFile = updateFile
	} else if os.IsNotExist(err) {
		return fmt.Errorf("File does not exist"+updateFile, err)
	} else {
		return fmt.Errorf("File does not exist"+updateFile, err)
	}
	return nil
}

// ReportFileExists report check
func ReportFileExists(reportOutput string) error {
	absPath, _ := filepath.Abs("report/in/" + reportOutput)
	_, err := os.Stat(absPath)
	if err == nil {
		crtFile = reportOutput
	} else if os.IsNotExist(err) {
		return fmt.Errorf("File does not exist"+reportOutput, err)
	} else {
		return fmt.Errorf("File does not exist"+reportOutput, err)
	}
	return nil
}

// CreategitProjectsTestCase gitProjects testcase
func CreategitProjectsTestCase(w http.ResponseWriter, gitProjectsCreateJSON gitProjectsCreateJSON) error {
	finalJSONForCreateTestCase, _ := json.Marshal(gitProjectsCreateJSON)
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, &finalJSONForCreateTestCase)
	if err != nil {
		return err
	}
	byteValue := buf.Bytes()
	req, err := http.NewRequest("POST", gitProjectsCreateOrUpdateTestCaseURL,
		bytes.NewBuffer(byteValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	client := &http.Client{Transport: transCfg}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ok(w, resp)
	return nil
}

// UpdategitProjectsTestCase update
func UpdategitProjectsTestCase(w http.ResponseWriter, gitProjectsUpdateJSON gitProjectsUpdateJSON) error {
	finalJSONForUpdateTestCase, _ := json.Marshal(gitProjectsUpdateJSON)
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, &finalJSONForUpdateTestCase)
	if err != nil {
		return err
	}
	byteValue := buf.Bytes()
	req, err := http.NewRequest("POST", gitProjectsUpdateTestResultsURL,
		bytes.NewBuffer(byteValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	client := &http.Client{Transport: transCfg}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ok(w, resp)
	return nil
}

func ok(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

// ValidateCreateResponse response validation
func ValidateCreateResponse(respBytes []byte) error {
	var createResponse []CreateUpdateResponse
	json.Unmarshal([]byte(respBytes), &createResponse)
	//fmt.Println("gitProjects RESPONSE JSON", createResponse)
	ioutil.WriteFile("report/in/gitProjectsCreateResponse.json", respBytes, 0644)
	overallStatus := getOverallgitProjectsCreateUpdateStatus(createResponse)
	if overallStatus != "Fail" {
		return nil
	}
	return nil
}

// ValidateUpdateResponse validate
func ValidateUpdateResponse(respBytes []byte) error {
	var updateResponse []CreateUpdateResponse
	json.Unmarshal([]byte(respBytes), &updateResponse)
	//fmt.Println("gitProjects RESPONSE JSON", updateResponse)
	ioutil.WriteFile("report/in/gitProjectsUpdateResponse.json", respBytes, 0644)
	overallStatus := getOverallgitProjectsCreateUpdateStatus(updateResponse)
	if overallStatus != "Fail" {
		return nil
	}
	return errors.New("error in uploading test results to gitProjects")
}

func getOverallgitProjectsCreateUpdateStatus(dataarr []CreateUpdateResponse) string {
	var overallgitProjectsCreateUpdteStatus string
	overallgitProjectsCreateUpdteStatus = "Pass"
	for _, elem := range dataarr {
		if !elem.UpdateResult {
			overallgitProjectsCreateUpdteStatus = "Fail"
		}
	}
	return overallgitProjectsCreateUpdteStatus
}

func findTestScenarioElements(byStr []byte) bool {
	elementRecs, _, _, elementNoErr := jsonparser.Get(byStr, "elements")
	if elementNoErr != nil {
		return false
	}
	elementLen, err := GetLenForJSONArray(elementRecs)
	if err != nil {
		log.Fatal("no Elements found")
	}
	for i := 0; i < elementLen; i++ {
		var tmpTstCase TestCases1
		var tmpTestCaseUpdate TestCasesUpdate
		var scenarioStatus = "Pass"
		currentOf := "[" + strconv.Itoa(i) + "]"
		testCaseName, _, _, _ := jsonparser.Get(byStr, "elements", currentOf, "name")
		//fmt.Printf("name '%s' \n", string(value))
		userStoryID := GetUserStoryIDForCurrentTestCase(string(testCaseName))
		//fmt.Println("userStoryId" + userStoryId)
		tmpTstCase.UserStoryID = userStoryID        //"US888460"
		tmpTestCaseUpdate.UserStoryID = userStoryID //"US888460"
		tmpTstCase.TestCaseName = string(testCaseName) + strconv.Itoa(i)
		tmpTestCaseUpdate.TestCaseName = string(testCaseName) + strconv.Itoa(i)
		tmpTstCase.TestUpdateFlag = "true"
		stepRecrds, _, _, _ := jsonparser.Get(byStr, "elements", currentOf, "steps")
		stepLen, err := GetLenForJSONArray(stepRecrds)
		if err != nil {
			log.Fatal("no steps found")
		}
		var tmpTestSteps []Steps1
		for j := 0; j < stepLen; j++ {
			var tmpStep Steps1
			tmpStep.StepIndex = strconv.Itoa(j + 1)
			tmpStep.StepExpectedResult = "Pass"
			stepOff := "[" + strconv.Itoa(j) + "]"
			keyword, _, _, _ := jsonparser.Get(byStr, "elements", currentOf, "steps", stepOff, "keyword")
			//fmt.Printf("ss '%s' \n", string(value))
			name, _, _, _ := jsonparser.Get(byStr, "elements", currentOf, "steps", stepOff, "name")
			//fmt.Printf("ss '%s' \n", string(value))
			tmpStep.StepDescription = string(keyword) + string(name)
			stepStatus, _, _, _ := jsonparser.Get(byStr, "elements", currentOf, "steps", stepOff, "result", "status")
			if string(stepStatus) == "failed" {
				scenarioStatus = "Fail"
			}
			//fmt.Printf("ss '%s' \n", string(value))
			tmpTestSteps = append(tmpTestSteps, tmpStep)
		}
		tmpTestCaseUpdate.TestStatus = scenarioStatus
		/*if i == 0 {
			//absPath, _ := filepath.Abs(gitProjectsServerAttachmentPath + "report.html")
			//gitProjectsServerAttachmentPath = getCurrentgitProjectsServerPathToUploadBasedOnOS()
			tmpTestCaseUpdate.Attachments = append(tmpTestCaseUpdate.Attachments, gitProjectsServerAttachmentPath + "report.html")
		} else {
			tmpTestCaseUpdate.Attachments = []string{} //append(tmpTestCaseUpdate.Attachments, tmpAttachment)
		}*/
		tmpTestCaseUpdate.Attachments = append(tmpTestCaseUpdate.Attachments, gitProjectsServerAttachmentPath+"report.html")
		tmpTstCase.TestStatus = scenarioStatus
		tmpTstCase.TestSteps = tmpTestSteps
		testCasesUpdate = append(testCasesUpdate, tmpTestCaseUpdate)
		testCases = append(testCases, tmpTstCase)
	}
	return true
}

func getCurrentgitProjectsServerPathToUploadBasedOnOS() string {
	_, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		gitProjectsServerAttachmentPath = "C:" + gitProjectsServerAttachmentPath
		//+ strings.Split(user.Username, "\\")[1]
	} else {
		gitProjectsServerAttachmentPath = gitProjectsServerAttachmentPath
		// + strings.Split(user.Username, "\\")[1]
	}
	return gitProjectsServerAttachmentPath
}

// GetLenForJSONArray json
func GetLenForJSONArray(jsonObj []byte) (int, error) {
	// Parse the JSON.
	var objs interface{}
	json.Unmarshal(jsonObj, &objs) // Or use json.Decoder.Decode(...)
	// Ensure that it is an array of objects.
	objArr, ok := objs.([]interface{})
	if !ok {
		log.Fatal("expected an array of objects")
		return 0, errors.New("notArray")
	}
	//log.Print(len(objArr))
	return len(objArr), nil
}

func loadgitProjectsDataFromConfig() (gitProjectsTestDataFromFile, error) {
	absPath, err := filepath.Abs("config/options.json")
	var gitProjectsTestDataFromFile gitProjectsTestDataFromFile
	if err != nil {
		return gitProjectsTestDataFromFile, err
	}
	jsonFile, err1 := os.Open(absPath)
	if err1 != nil {
		return gitProjectsTestDataFromFile, err
	}
	var byStr []byte
	byStr, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return gitProjectsTestDataFromFile, err
	}
	err = json.Unmarshal(byStr, &gitProjectsTestDataFromFile)
	if err != nil {
		return gitProjectsTestDataFromFile, fmt.Errorf("unable to load test data for gitProjects from config")
	}
	return gitProjectsTestDataFromFile, nil
}

// MapTestCasegitProjectsUserStoryMap maps
func MapTestCasegitProjectsUserStoryMap(f *gherkin.Feature, tag string) ([]gitProjectsUserStoryTCNameMapping, error) {
	currentFeature = f
	scenarios := make([]interface{}, len(currentFeature.ScenarioDefinitions))
	copy(scenarios, currentFeature.ScenarioDefinitions)
	if currentFeature.Background != nil {
		scenarios = append(scenarios, currentFeature.Background)
	}
	for _, scenario := range scenarios {
		switch t := scenario.(type) {
		case *gherkin.ScenarioOutline:
			if tag != "regression" {
				for _, step := range t.Steps {
					storyTcMap(tag, step.Text, t.Name)
				}
			} else {
				storyTcMap(tag, "", t.Name)
			}
		case *gherkin.Scenario:
			if tag != "regression" {
				for _, step := range t.Steps {
					storyTcMap(tag, step.Text, t.Name)
				}
			} else {
				storyTcMap(tag, "", t.Name)
			}
		case *gherkin.Background:
			if tag != "regression" {
				for _, step := range t.Steps {
					storyTcMap(tag, step.Text, t.Name)
				}
			} else {
				storyTcMap(tag, "", t.Name)
			}
		}
	}
	userStoryTCNameMapJSON, _ := json.Marshal(userStoryTCNameMap)
	fmt.Println("userStoryTCNameMap")
	fmt.Println(userStoryTCNameMap)
	ioutil.WriteFile("report/in/gitProjectsmap.json", userStoryTCNameMapJSON, 0644)
	return userStoryTCNameMap, nil
}

func storyTcMap(tag string, stepText string, testCaseName string) {
	var storyTcMap gitProjectsUserStoryTCNameMapping
	if tag != "regression" {
		if strings.Contains(stepText, "create gitProjects test under userstory") {
			stringsplit := strings.Split(stepText, "\"")[1]
			storyTcMap.UserStoryID = stringsplit
			storyTcMap.TestCaseName = testCaseName
			userStoryTCNameMap = append(userStoryTCNameMap, storyTcMap)
		}
	} else {
		gitProjectsTestDataFromFile, _ := loadgitProjectsDataFromConfig()
		storyTcMap.UserStoryID = gitProjectsTestDataFromFile.UserStoryID
		storyTcMap.TestCaseName = testCaseName
		userStoryTCNameMap = append(userStoryTCNameMap, storyTcMap)
	}

}

// UploadAttachmentsTogitProjectsServer upload
func UploadAttachmentsTogitProjectsServer() error {
	searchDir, err := filepath.Abs("report/out/")
	if err != nil {
		return err
	}
	var fileList []string
	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, file := range fileList {
		fmt.Println(file)
		err = postgitProjectsUploadToServer(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func postgitProjectsUploadToServer(fileNameWithPath string) error {
	//var fileNameWithPath = "C:\\Users\\bvenkatr\\Downloads\\report1.html"
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", fileNameWithPath)
	if err != nil {
		return err
	}
	f, err := os.Open(fileNameWithPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var byStr []byte
	byStr, _ = ioutil.ReadAll(f)
	fileWriter.Write(byStr)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(gitProjectsUploadAttachmentURL, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var gitProjectsAttachmentResponse gitProjectsAttachmentResponse
	json.Unmarshal(respData, &gitProjectsAttachmentResponse)
	fmt.Println(gitProjectsAttachmentResponse.Entity)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s", gitProjectsAttachmentResponse.Entity)
	}
	return nil
}
