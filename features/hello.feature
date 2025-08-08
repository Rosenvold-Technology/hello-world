Feature: Hello endpoint
  Background:
    Given the API base url is "http://localhost:8080"

  Scenario: Root responds
    When I GET "/"
    Then the response code is 200
    And the body contains "Hello, Kubernetes!"
