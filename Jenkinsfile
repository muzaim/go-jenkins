pipeline {
    agent {
        docker {
            image 'golang:alpine'
        }
    }

    environment {
        // Konfigurasi umum (non-rahasia)
        APP_ENV_DEV       = 'development'
        APP_ENV_PROD      = 'production'
        SERVICE_PORT_DEV  = '8081'
        SERVICE_PORT_PROD = '8082'

        // Data sensitif diambil langsung dari Jenkins Credentials Store
        DEPLOY_HOST_DEV   = credentials('deploy-host-dev')
        DEPLOY_HOST_PROD  = credentials('deploy-host-prod')
        DEPLOY_USER_DEV   = credentials('deploy-user-dev')
        DEPLOY_USER_PROD  = credentials('deploy-user-prod')
    }

    stages {
        stage('Quality & Security Check') {
            parallel {
                stage('Code Quality (Format & Vet)') {
                    steps {
                        echo 'Checking code formatting and running go vet...'
                        sh 'test -z "$(gofmt -l .)"'
                        sh 'go vet ./...'
                    }
                }

                stage('Security Scan') {
                    steps {
                        echo 'Scanning dependencies for known vulnerabilities...'
                        sh 'go install golang.org/x/vuln/cmd/govulncheck@latest'
                        sh 'govulncheck ./...'
                    }
                }

                stage('Unit Testing & Coverage') {
                    steps {
                        echo 'Running unit tests with coverage profile...'
                        sh 'go test -v -coverprofile=coverage.out ./...'
                    }
                }
            }
        }

        stage('Build Check') {
            steps {
                echo 'Compiling go binary...'
                sh 'go build -o bin/api main.go'
            }
        }

        stage('Deploy to DEV Server') {
            when {
                branch 'dev'
            }
            steps {
                echo "=== SIMULATING SSH DEPLOYMENT ==="
                echo "Connecting to: ${env.DEPLOY_USER_DEV}@${env.DEPLOY_HOST_DEV}..."
                echo "Deploying application environment: ${env.APP_ENV_DEV}..."
                echo "Transferring binary bin/api to remote server..."
                echo "Starting service on port ${env.SERVICE_PORT_DEV}..."
                echo "Deployment completed successfully (Mock)!"
            }
        }

        stage('Deploy to PROD Server') {
            when {
                branch 'main'
            }
            steps {
                echo "=== SIMULATING SSH DEPLOYMENT ==="
                echo "Connecting to: ${env.DEPLOY_USER_PROD}@${env.DEPLOY_HOST_PROD}..."
                echo "Deploying application environment: ${env.APP_ENV_PROD}..."
                echo "Transferring binary bin/api to remote server..."
                echo "Starting service on port ${env.SERVICE_PORT_PROD}..."
                echo "Deployment completed successfully (Mock)!"
            }
        }

        stage('Health Check DEV') {
            when { branch 'dev' }
            steps {
                echo "=== VERIFYING DEV SERVICE HEALTH ==="
                echo "Pinging service at http://${env.DEPLOY_HOST_DEV}:${env.SERVICE_PORT_DEV}/health..."
                echo "HTTP/1.1 200 OK received! DEV System is operational."
            }
        }

        stage('Health Check PROD') {
            when { branch 'main' }
            steps {
                echo "=== VERIFYING PROD SERVICE HEALTH ==="
                echo "Pinging service at http://${env.DEPLOY_HOST_PROD}:${env.SERVICE_PORT_PROD}/health..."
                echo "HTTP/1.1 200 OK received! PROD System is operational."
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            echo "CI/CD pipeline finished with SUCCESS! Ready for user traffic."
        }
        failure {
            echo "Pipeline FAILED! Check the failed stage logs."
        }
    }
}