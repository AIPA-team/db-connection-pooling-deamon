# Contributing to db-connection-pooling-daemon

### All contributions are more than welcome!  

### How to contribute:

1. To contribute fork original repository and start your work in feature branch
2. Create meaningful commit history while commiting. We use labels described below to make it more readable
3. After Your work is done, rebase your branch to be in sync with upstream
4. Create a pull request
5. If all tests pass and at least one of maintainers accepts PR (we use LGTM), it will be rebased to master branch and PR will be closed (we use rebase workflow instead of merge for clear history)

#### Labels
- main
- config
- tests
- doc  
... (it's just a general rule to use them)

so commit message should look like this:  
`doc: fix typo in README.md`

more info about git commit messages: [`git messages`](http://chris.beams.io/posts/git-commit/)

### Coding styleguide:

1. try to write idiomatic GO code (refer [`effective GO`](https://golang.org/doc/effective_go.html))
2. use gofmt -s
3. use [`golint`](https://github.com/golang/lint)
4. comment all functions and methods
5. write tests
