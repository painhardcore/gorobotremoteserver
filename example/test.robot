*** Settings ***
Library    Remote    http://localhost:${PORT}

*** Variables ***
${HOST}    localhost
${PORT}    8270


*** Test Cases ***
Example Test That Succeed
    Args should be less than 5   first   second   third
Example Test That Fail
    Args should be less than 3   first   second   third

*** Keywords ***
