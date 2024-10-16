pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                go test -v -cover ./..
            }
        }
    }
}
