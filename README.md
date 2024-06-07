# go-hello-world-lambda-api-gateway
Simple Go API Gateway Example with AWS SAM Template

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
* [Golang](https://golang.org) - Currently using v.1.22 (can use [Chocolatey](https://community.chocolatey.org/packages/golang) or [Homebrew](https://formulae.brew.sh/formula/go))
* [Bruno](https://www.usebruno.com/) (Optional) - Open source API client. Similar to Postman, but is offline-only and will never require an account. 
* [Make](https://www.gnu.org/software/make/) (Optional) - To run makefile commands 

## Local development

1. **Build the go executable for the [lambda linux environment](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)**
    ```shell
    GOARCH=arm64 GOOS=linux go build -o bootstrap main.go
    ```

2. **Zip project (Windows)**

    `C:\Users\owner\go\bin\build-lambda-zip.exe -o lambda-handler.zip bootstrap` using the provided `build-lambda-zip` package. If needed, you can install with `go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest`. See [To create a .zip deployment package (Windows)](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)

3. **Invoke function locally through local API Gateway**
    ```shell
    sam.cmd local start-api --template-file=template.yaml
    ```

    *Note: Use `sam.cmd` when running AWS SAM on windows*

    To rebuild and apply local changes, use `sam.cmd build --use-container` 

4. **(Optional) Run CRUD Requests Individually**
    
    Can also run CRUD API requests by importing the `/bruno` collection

## Deployment

Ensure you follow the same steps you did to build the executable and zipping the project

1. **Build the go executable for the [lambda linux environment](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)**
    ```shell
    GOARCH=arm64 GOOS=linux go build -o bootstrap main.go
    ```

2. **Zip project (windows)**

    `C:\Users\owner\go\bin\build-lambda-zip.exe -o lambda-handler.zip bootstrap` using the provided `build-lambda-zip` package. If needed, you can install with `go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest`. See [To create a .zip deployment package (Windows)](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)

3. **Deploy CloudFromation stack template**
    ```shell
    sam.cmd deploy --guided --template-file=template.yaml
    ```
