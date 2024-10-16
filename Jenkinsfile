pipeline {
    agent any
    tools {
        go 'go-1.11'
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
