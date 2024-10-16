pipeline {
    agent any
    tools {
        go 'go-1.4'
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
