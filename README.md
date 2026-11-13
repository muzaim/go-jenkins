# Go Jenkins CI/CD Pipeline

A lightweight Go REST API integrated with an enterprise-grade, containerized Jenkins CI/CD multibranch pipeline. This project demonstrates modern automation practices, including ephemeral Docker-based build agents, parallel code quality and vulnerability checks, and automated branch-specific deployment workflows.

---

## Architecture Overview

```
[ Developer ] --( git push )--> [ GitHub Repository ]
                                        |
                             ( Webhook / Real-Time )
                                        v
                            [ Jenkins Controller ]
                             (Docker Socket Mount)
                                        |
                          ( Spins up Ephemeral Agent )
                                        v
                            [ golang:alpine Container ]
                                        |
       +--------------------------------+--------------------------------+
       |                                |                                |
       v                                v                                v
[ Code Quality ]                [ Security Scan ]               [ Unit Testing ]
(gofmt & go vet)                 (govulncheck)                 (coverage report)
       |                                |                                |
       +--------------------------------+--------------------------------+
                                        | (All Passed)
                                        v
                               [ Binary Compilation ]
                                        |
                 +----------------------+----------------------+
                 | (Branch: dev)                               | (Branch: main)
                 v                                             v
       [ Deploy to DEV Server ]                      [ Deploy to PROD Server ]
       [ DEV Health Verification ]                   [ PROD Health Verification ]
```

---

## Key Features

- **Containerized Infrastructure**: Jenkins runs inside Docker with host Docker socket access (`/var/run/docker.sock`), eliminating the need to install runtime dependencies directly on the host.
- **Ephemeral Build Agents**: Pipeline stages run inside temporary `golang:alpine` containers, ensuring clean, reproducible, and isolated builds.
- **Parallel Quality & Security Gates**:
  - **Code Quality**: Code formatting checks via `gofmt` and static analysis via `go vet`.
  - **Vulnerability Scanning**: Automated dependency vulnerability audits using Go's official `govulncheck`.
  - **Unit Testing**: Automated unit tests with coverage reporting (`go test -cover`).
- **Multibranch Pipeline**: Automatically discovers branches and applies environment-specific deployment logic (`dev` vs `main`) within a single declarative `Jenkinsfile`.
- **Real-Time Webhook Triggering**: Instant build execution triggered by GitHub push events.

---

## Project Structure

```text
.
├── Dockerfile.jenkins      # Custom Jenkins image with Docker CLI client
├── Jenkinsfile             # Declarative multi-stage multibranch pipeline
├── docker-compose.yml      # Local Jenkins server orchestration
├── go.mod                  # Go module definition
├── main.go                 # Simple HTTP health-check API
├── main_test.go            # Unit tests for the HTTP handler
└── README.md
```

---

## Pipeline Workflow

1. **Quality & Security Checks (Parallel)**
   - `Code Quality`: Validates Go formatting rules and checks for code anomalies.
   - `Security Scan`: Scans external dependencies for known CVEs.
   - `Unit Testing & Coverage`: Executes unit test suite and outputs coverage data.
2. **Build Check**: Compiles the Go binary to `bin/api`.
3. **Environment-Specific Deployment**:
   - `Deploy to DEV Server`: Triggered only on the `dev` branch.
   - `Deploy to PROD Server`: Triggered only on the `main` branch.
4. **Health Check**: Validates post-deployment service availability.
5. **Workspace Cleanup**: Automatically wipes temporary files in the `post` stage.

---

## Getting Started Locally

### Prerequisites
- Docker & Docker Compose
- Go 1.23+ (for local application execution)
- Git

### 1. Running the API Locally
```bash
go run main.go
# API is accessible at http://localhost:8081/health
```

### 2. Running Local Unit Tests
```bash
go test -v -cover ./...
```

### 3. Starting the Jenkins Server
```bash
docker compose up -d --build
```
Access the Jenkins dashboard at `http://localhost:8080`. Retrieve the initial unlock key using:
```bash
docker logs jenkins-server
```
