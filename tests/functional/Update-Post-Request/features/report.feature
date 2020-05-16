# file: $GOPATH/src/have-some-func/features/gitProjects.feature
@gitProjects
Feature: create and update gitProjects test case
  This feature will pass if create and update gitProjects succeeded

  Scenario: Create and update gitProjects test case
    When "report.json" exists
    Then create gitProjects jsons,test cases in gitProjects from user story scenario mapping file and upload test result