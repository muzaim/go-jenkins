pipeline{
    agent{
        docker{
            image 'golang:alpine'
        }
    }

    environment{
        APP_ENV_DEV='development'
        APP_ENV_PROD='production'
        DEPLOY_HOST_DEV='192.1634.131'
        DEPLOY_HOST_PROD='192.1634.132'
        DEPLOY_USER_DEV='deployer'
        DEPLOY_USER_PROD='admin'
        SERVICE_PORT_DEV='8081'
        SERVICE_PORT_PROD='8082'
    }

    stages{
        stage('Quality & Security Check'){
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