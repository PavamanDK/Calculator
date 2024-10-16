pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                go test -v -cover ./...
                GOOS=linux GOARCH=amd64 go build -o myapp .
                golint ./...
            }
        }
    }
}
