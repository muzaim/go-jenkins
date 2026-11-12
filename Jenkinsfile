pipeline {
    // Memberitahu Jenkins untuk menjalankan seluruh stage di dalam container Golang
    agent {
        docker {
            image 'golang:1.23-alpine'
        }
    }

    stages {
        stage('Check Environment') {
            steps {
                sh 'go version'
            }
        }

        stage('Download Dependencies') {
            steps {
                sh 'go mod download'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go test -v -cover ./...'
            }
        }

        stage('Build Binary') {
            steps {
                sh 'go build -o bin/api main.go'
            }
        }
    }

    post {
        always {
            cleanWs() // Bersihkan workspace temporary setelah selesai
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}