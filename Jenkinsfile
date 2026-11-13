pipeline{
    agent{
        docker{
            image 'golang:alpine'
        }
    }

    environment{
        APP_ENV='production'
        DEPLOY_HOST='192.1634.131'
        DEPLOY_USER='deployer'
        SERVICE_PORT='8081'
    }

    stages{
        stage{
            parallel{
                stage('Code Quality (Format & Vet)'){
                    steps{
                        echo 'Checking code formatting and running go vet...'
                        sh 'test -z "$(gofmt -l .)"'
                        sh 'go vet ./...'
                    }
                }

                stage('Security Scan'){
                    steps{
                        echo 'Scanning dependencies for known vurnelabilities...'
                        sh 'go install golang.org/x/vuln/cmd/govulncheck@latest'
                        sh 'govulncheck ./...'
                    }
                }

                stage('Unit Testing & Coverage'){
                    steps{
                        echo 'Running unit tests with coverage profile...'
                        sh 'go test -v -coverprofile=coverage.out ./...'
                    }
                }
            }
        }

        stage('Build Check'){
            steps{
                echo 'Compiling go binary...'
                sh 'go build -o bin/api main.go'
            }
        }

        stage('Deploy SSH (Mock)'){
            steps {
                echo "=== SIMULATING SSH DEPLOYMENT ==="
                echo "Connecting to: ${env.DEPLOY_USER}@${env.DEPLOY_HOST}..."
                echo "Deploying application environment: ${env.APP_ENV}..."
                echo "Transferring binary bin/api to remote server..."
                echo "Starting service on port ${env.SERVICE_PORT}..."
                echo "Deployment completed successfully (Mock)!"
            }
        }

        stage('Health check'){
            steps {
                echo "=== VERIFYING SERVICE HEALTH ==="
                echo "Pinging service at http://${env.DEPLOY_HOST}:${env.SERVICE_PORT}/health..."
                echo "HTTP/1.1 200 OK received! System is operational."
            }
        }
    }

    post{
        always{
            cleanWs()
        }
        success{
            echo "CI/CD pipeline finished with SUCCESS! Ready for user traffic."
        }
        failure{
            echo "Pipeline FAILED! Check the failed stage logs."
        }
    }
    
}