# Inifinit Test

## Commands

Requirements : [GoLang](https://go.dev/)

Run for development
```sh
go run .
```

Run tests
```sh
go test .
```

Build and run for production
```sh
go build . && ./infinit-task
```

## Given Requirements

The goal is for you to write a program that connects via the GitHub API to the GitHub - lodash/lodash repository and gathers a couple of statistics for us.

We would like the program to output statistics on how often each letter is present in the JavaScript/TypeScript files of the repository, in decreasing order.

## Requirment breakdown

Write a program that:
- connects via the GitHub API to github repo (lodash)
- gathers/outputs how often each letter is present in the JavaScript/TypeScript files

### Connecting via the GitHub API to github repo 

Questions
What might "connect" mean ? ... anyways basically we need to get the file content to be able to look into them.

Github api research :

```
Possibilities on github api docs:
1 - fine grained aproach: fetch file list -> and query content file by file on the js/ts files to be processed
    Pros: uses github api a lot / Cons: Overkill, there might be thousands of files to fetch, will we need to manage rate limiting ?  
2 - one shot: fetch a zip of the repo on a certain version and do the processing in bulk next 
    Pros: Simple, uses github api / Cons: will download all repo files, depends on zip

also it can be done using git cli but the requirement clearly states that we use github api

```

Proceeding with 2

### Charachter occurency stat

Needs repository js/ts file
- what version (branch/commit/tag?) -> configurable
- what about mjs or tsx -> configurable

Needs to count charater occurences for a file
- Load file into memory (how many at the same time?)
- process file content (get the stats) (can we parallelize ?)
- aggregate results for the whole repository


