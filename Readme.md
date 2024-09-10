# init project
go mod init gotest

## create main.go

## create repositories

## create services

## create handlers

# test all in services (module of go.mod/package of services)
go test gotest/services -v

# test some service
go test gotest/services -v -run=TestHello

# code tdd 
# create for function test before function in services

# config vscode
# hilight expected cover test case
cmd + shift + p 
open setting user (JSON)

# assign this code in to file
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}# Intro_Unit_Test_Go
