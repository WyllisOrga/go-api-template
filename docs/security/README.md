# Security

## How to set up Sonarqube CI

- "Set up sonarcloud" section
- "Set up Github project" section
- "Set up project files configuration" section

### Set up sonarcloud

- Transfer your current repository to an organization (sonarcloud analyses repos from organization).
- Go in sonarcloud `https://sonarcloud.io/`, and create an account.
- Click in `+` -> `analyze new project`, continue to configure project as you want.
- Go in `https://sonarcloud.io/projects`, you should see your project, go inside.

### Set up Github project

- Create two environements variables
  Go in you github project -> `settings` center tab -> `secrets` left tab -> `New repository secret` button.
  Add two variables :

  - `SONAR_TOKEN`: see `Get SONAR_TOKEN` section
  - `SONAR_HOST_URL` which is `https://sonarcloud.io/`

### Set up project files configuration

- add a `sonar-project.properties` file in you project at `/` and add

```
sonar.projectKey=
sonar.projectName=
sonar.organization=

sonar.sources=.
sonar.exclusions=**/*_test.go

sonar.tests=.
sonar.test.inclusions=**/*_test.go
```

You can find `projectKey`, `projectName`, `organization` in sonarcloud

- Create a pipeline like this

```yml
name: Go

on:
  push:
    branches:
      - "*"
  pull_request:
    branches:
      - "*"

jobs:
  security:
    name: Publish to SonarCloud
    runs-on: ubuntu-18.04
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0

      - uses: docker://sonarsource/sonar-scanner-cli:latest
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}
          SONAR_HOST_URL: ${{ secrets.SONAR_HOST_URL }}
```

### Get SONAR_TOKEN

Go in `https://sonarcloud.io/`

- Log in
- Click on your picture profile at the top in the header -> "My Account" -> "Security" tab
- Under "Generate Tokens" enter your token and generate it.
