pipeline {
    agent any
    tools {
        go 'go-1.23'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Hello') {
            steps {
                sh 'echo "Hello World"'
            }
        }
    }
}
