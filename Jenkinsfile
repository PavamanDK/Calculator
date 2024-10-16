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
                echo "Welcome to my first CICD pipeline"
            }
        }
        stage('Compile') {
            steps{
                sh 'ls -ltrha'
                /*sh 'go mod init github.com/pavaman'*/
                sh 'go build'
                sh 'ls -ltrha'
            }
        }
        stage('Test') {
            steps{
                sh 'go test'
            }
        }
    }
}
