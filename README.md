# go-blog-api
API service for a blog.

To setup this project in local please follow the below-mentioned steps,

1. Checkout this repo using `git clone repo_url`

2. Make sure you installed GoLang in your system.

3. After checkout this repo, open the root folder and execute command `go install` to install/update required dependecies.

4. Execute command `go mod tidy` to ensures that the go.mod file matches the source code.

5. Set the below-mentioned varaibles in OS environment.
	`os.Setenv("DB_USERNAME", "root")`
	`os.Setenv("DB_PASSWORD", "")`
	`os.Setenv("DB_HOST", "127.0.0.1")`
	`os.Setenv("DB_PORT", "3306")`
	`os.Setenv("DB_NAME", "db_name")`
	`os.Setenv("HOST", "localhost")`
	`os.Setenv("PORT", "8078")`

6. Execute command `go build .`

7. Execute command `go run .`  //it will start the server based on set host and port like "http://localhost:8078"

8. We have total 6 APIs, you can find postman collection from this link. https://drive.google.com/file/d/1cTOFVHToV6IHeKsMUdByDNMKf3Qe1dUR/view?usp=sharing

9. We have all test cases for unit test in folder `/test`
	and then execute command `go test -v`. it will display the test cases are failed or passed.