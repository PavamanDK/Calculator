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
                sh 'go build'
            }
        }
    }
}
