@functional
Feature: Add new employee and then update employee
  In order to be successful
  I must be able to add new employee
  I must be able to update new employee
  And validate whether employee added in database
  Scenario Outline: POST new employee and validate
    Given the app ip is "<host>"
    And the app port is "<port>"
    And the testcase is "<testcase>"
    And Add new employee
    And update employee and validate response
    #And create gitProjects test under userstory "US1230644" and upload test result
    Examples:
      | testcase    | host   | port  |
      | POST Request| walk-the-camino  | 4444  |