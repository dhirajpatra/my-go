# Getting Started with Go

Welcome to the world of Go (also known as Golang)! Whether you're a beginner or an experienced developer, this guide will help you get started with Go.

## Table of Contents

1. [Installing Go](#installing-go)
2. [Creating Your First Go Program](#creating-your-first-go-program)
3. [Working with Go Modules](#working-with-go-modules)
4. [Running Go Files](#running-go-files)
5. [Building Executables](#building-executables)

## Installing Go

1. **Download and Install Go**:
   - Visit the [official Go website](https://golang.org/dl/) and download the latest stable version for your operating system.
   - Follow the installation instructions for your platform.

2. **Verify Installation**:
   - Open a terminal or command prompt and run:
     ```bash
     go version
     ```
   - You should see the installed Go version.

## Creating Your First Go Program

1. **Create a New Go Workspace**:
   - Create a directory for your Go workspace (e.g., `hello`):
     ```bash
     mkdir hello
     cd hello
     ```

2. **Initialize a Go Module**:
   - Run the following command to create a `go.mod` file:
     ```bash
     go mod init hello
     ```

3. **Write Your First Go Program**:
   - Create a file named `helloworld.go` and add the following code:
     ```go
     package main

     import "fmt"

     func main() {
         fmt.Println("Hello, World!")
     }
     ```

4. **Run Your Program**:
   - Execute the following command to run your program:
     ```bash
     go run helloworld.go
     ```
   - You should see the output: "Hello, World!"

## Working with Go Modules

1. **Dependencies and Modules**:
   - Go uses modules to manage dependencies.
   - Create a `go.mod` file in your project directory:
     ```bash
     go mod init <module-name>
     ```

2. **Download Dependencies**:
   - Use `go get` to download dependencies listed in your `go.mod` file:
     ```bash
     go get
     ```

3. **Vendor Dependencies (Optional)**:
   - To save a copy of your dependencies locally, use:
     ```bash
     go mod vendor
     ```

## Running Go Files

1. **Run a Go File**:
   - Use `go run` to execute a Go file:
     ```bash
     go run filename.go
     ```

## Building Executables

1. **Build an Executable**:
   - Use `go build` to create an executable binary:
     ```bash
     go build hello
     ```
   - The executable will be named `hello`.

2. **Customize Executable Name**:
   - To specify a custom name for the executable:
     ```bash
     go build -o myapp
     ```
   - The executable will be named `myapp`.


## Developing Go Microservices Application With Docker

### Create MySql Container

1. **Pull the MySQL Docker Image**:
    - Open a command prompt or terminal.
    - Pull the appropriate MySQL image using the following command:
        ```
        docker pull mysql/mysql-server:latest
        ```
        Replace `latest` with the specific version if you want a particular MySQL version.
    - Create a docker volume
      ```
      docker volume create mysql-data
      ```

2. **Deploy the MySQL Container**:
    - Deploy a new MySQL container using the following command:
        ```
        docker run -d --name my-mysql -e MYSQL_ROOT_PASSWORD=pass -p 3306:3306 -v mysql-data:/var/lib/mysql mysql:8.0
        ```
        Replace `my-mysql-container` with your preferred container name. If you don't provide a name, Docker generates a random one. The `-d` option runs the container as a service in the background.

3. **Connect to the MySQL Container**:
    - To access the MySQL container, use the following command:
        ```
        docker exec -it my-mysql-container mysql -uroot -p
        ```
        You'll be prompted to enter the MySQL root password (which you set during container deployment).
    - To know the IP address of the my-mysql-container, use the following command:
        ```
        docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-mysql
        ```



### Some more links to learn Go

Source: Conversation with Bing, 14/3/2024
(1) How to run a Golang Program? - TutorialKart. https://www.tutorialkart.com/golang-tutorial/run-go-program/.
(2) How to Run a Golang File. https://www.golang101.com/how-tos/how-to-run-golang-file/.
(3) Running a one-off File - Go Cookbook. https://golangcookbook.com/chapters/running/one-off/.
(4) Tutorial: Get started with Go - The Go Programming Language. https://go.dev/doc/tutorial/getting-started.
(5) How To Write Your First Program in Go | DigitalOcean. https://www.digitalocean.com/community/tutorials/how-to-write-your-first-program-in-go.
(6) How To Build Go Executables for Multiple Platforms on Ubuntu 20.04. https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04.
(7) How to build executable with name other than Golang package. https://stackoverflow.com/questions/42706246/how-to-build-executable-with-name-other-than-golang-package.
(8) How To Build and Install Go Programs | DigitalOcean. https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs.
(9) How to compile Go program consisting of multiple files?. https://stackoverflow.com/questions/10838469/how-to-compile-go-program-consisting-of-multiple-files.
(10) Get Started - The Go Programming Language. https://go.dev/learn/.
(11) The Golang Handbook – A Beginner's Guide to Learning Go. https://www.freecodecamp.org/news/learn-golang-handbook/.
(12) How to Learn Golang – A Beginner's Guide to the Basics - freeCodeCamp.org. https://www.freecodecamp.org/news/golang-for-beginners/.
(13) Installing all dependencies from a go.mod file - Stack Overflow. https://stackoverflow.com/questions/69474609/installing-all-dependencies-from-a-go-mod-file.
(14) go modules installing go tools - Stack Overflow. https://stackoverflow.com/questions/53368187/go-modules-installing-go-tools.
(15) go.mod file reference - The Go Programming Language. https://go.dev/doc/modules/gomod-ref.
(16) Installing go modules from Github repository - Medium. https://medium.com/@yussufshaikh/installing-go-modules-from-github-repository-5e381cbd5683.
(17) How to Use Go Modules | DigitalOcean. https://www.digitalocean.com/community/tutorials/how-to-use-go-modules.
(18) undefined. https://go.dev/ref/mod.
