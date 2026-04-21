# GitHub Actions CI/CD Setup

This directory contains GitHub Actions workflows for the Indigo rules engine project.

## Workflows

### 🔧 CI (`ci.yml`)
**Triggers:** Push to main/master/develop, Pull requests
**Purpose:** Continuous Integration pipeline

- **Multi-version testing:** Tests against Go 1.22, 1.23, and 1.24
- **Comprehensive testing:** Unit tests, race condition tests
- **Code quality:** Go formatting, `go vet`, staticcheck
- **Build verification:** Ensures the project builds successfully
- **Coverage reporting:** Uploads coverage reports to Codecov

### 📊 Benchmark (`benchmark.yml`)
**Triggers:** Push to main/master, Pull requests, Daily at 2 AM UTC
**Purpose:** Performance monitoring and regression detection

- **Automated benchmarking:** Runs comprehensive benchmarks using the project's Makefile
- **Performance comparison:** Compares current benchmarks with previous runs
- **PR comments:** Automatically comments benchmark results on pull requests
- **Artifact storage:** Stores benchmark results for historical analysis
- **Scheduled runs:** Daily benchmarks to track performance over time

### 🚀 Release (`release.yml`)
**Triggers:** Git tags matching `v*`
**Purpose:** Automated release process

- **Quality assurance:** Runs full test suite before release
- **Changelog generation:** Automatically generates changelog from git history
- **GitHub releases:** Creates GitHub releases with proper versioning
- **Release benchmarks:** Captures performance baseline for each release
- **Prerelease detection:** Automatically marks alpha/beta/rc versions as prereleases

### 🔒 Security (`security.yml`)
**Triggers:** Push to main/master, Pull requests, Weekly on Sundays
**Purpose:** Security scanning and vulnerability detection

- **Static analysis:** Uses gosec for Go-specific security issues
- **Dependency scanning:** Trivy scanner for known vulnerabilities
- **SARIF integration:** Results integrated with GitHub Security tab
- **Dependency review:** Automated review of dependency changes in PRs

### 🔍 CodeQL (`codeql.yml`)
**Triggers:** Push to main/master, Pull requests, Weekly on Sundays
**Purpose:** Advanced code analysis

- **Semantic analysis:** GitHub's CodeQL engine for deep code analysis
- **Security finding:** Identifies potential security vulnerabilities
- **Code quality:** Detects code quality issues and patterns

### 📚 Documentation (`docs.yml`)
**Triggers:** Changes to Go files, Markdown files, or module files
**Purpose:** Documentation quality assurance

- **Documentation validation:** Checks for proper Go documentation
- **README validation:** Ensures README examples are valid
- **Code quality:** Identifies TODO/FIXME comments for review

## Configuration Files

### 🤖 Dependabot (`dependabot.yml`)
**Purpose:** Automated dependency updates

- **Go modules:** Weekly updates for Go dependencies
- **GitHub Actions:** Weekly updates for action versions
- **Automated PRs:** Creates pull requests for dependency updates
- **Reviewer assignment:** Automatically assigns reviewers

### 📝 Templates
**Purpose:** Standardized contribution process

- **Pull Request Template:** Ensures consistent PR descriptions
- **Issue Templates:** 
  - Bug reports with detailed information
  - Feature requests with use case details
  - Performance issues with benchmarking context

## Setup Requirements

### Repository Settings
1. **Branch Protection:** Configure branch protection rules for main/master
2. **Secrets:** Add any required secrets (e.g., for external services)
3. **Permissions:** Ensure actions have necessary permissions

### Optional Integrations
1. **Codecov:** For coverage reporting (requires CODECOV_TOKEN secret)
2. **Slack/Discord:** For notification integration
3. **Deployment:** Add deployment workflows as needed

## Badge Integration

Add these badges to your README.md:

```markdown
[![Go Reference](https://pkg.go.dev/badge/github.com/ezachrisen/indigo.svg)](https://pkg.go.dev/github.com/ezachrisen/indigo)
[![CI](https://github.com/ezachrisen/indigo/workflows/CI/badge.svg)](https://github.com/ezachrisen/indigo/actions/workflows/ci.yml)
[![Security](https://github.com/ezachrisen/indigo/workflows/Security/badge.svg)](https://github.com/ezachrisen/indigo/actions/workflows/security.yml)
[![CodeQL](https://github.com/ezachrisen/indigo/workflows/CodeQL/badge.svg)](https://github.com/ezachrisen/indigo/actions/workflows/codeql.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ezachrisen/indigo)](https://goreportcard.com/report/github.com/ezachrisen/indigo)
```

## Customization

### Adjusting Go Versions
Edit the matrix in `ci.yml` to test against different Go versions:
```yaml
strategy:
  matrix:
    go-version: ['1.22', '1.23', '1.24']
```

### Benchmark Frequency
Modify the cron schedule in `benchmark.yml`:
```yaml
schedule:
  - cron: '0 2 * * *'  # Daily at 2 AM UTC
```

### Security Scanning
Add additional security tools or customize existing ones in `security.yml`.

## Troubleshooting

### Common Issues
1. **Permission Errors:** Ensure GITHUB_TOKEN has sufficient permissions
2. **Cache Issues:** Clear GitHub Actions cache if builds fail unexpectedly
3. **Dependency Issues:** Check dependabot configuration for conflicts

### Debug Mode
Enable debug logging by setting repository secrets:
- `ACTIONS_STEP_DEBUG=true`
- `ACTIONS_RUNNER_DEBUG=true`

## Best Practices

1. **Keep workflows focused:** Each workflow has a specific purpose
2. **Use caching:** All workflows use Go module caching for speed
3. **Fail fast:** Workflows are designed to fail quickly on errors
4. **Security first:** Multiple security scanning tools integrated
5. **Performance monitoring:** Continuous benchmarking prevents regressions
