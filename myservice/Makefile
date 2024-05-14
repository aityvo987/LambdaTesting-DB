.PHONY: build clean deploy

set-environment:
	set GOOS=linux                                             
	set GOARCH=arm64                                           
	set CGO_ENABLED=0        
zipfile:
	cd bin/postEmployee && 7z a ../postEmployee.zip bootstrap
	cd bin/putEmployee && 7z a ../putEmployee.zip bootstrap
	cd bin/deleteEmployee && 7z a ../deleteEmployee.zip bootstrap
	cd bin/getEmployee && 7z a ../getEmployee.zip bootstrap
	cd bin/getEmployees && 7z a ../getEmployees.zip bootstrap
build:
	go build -o bin/postEmployee/bootstrap functions/postEmployee/main.go   
	go build -o bin/putEmployee/bootstrap functions/putEmployee/main.go  
	go build -o bin/deleteEmployee/bootstrap functions/deleteEmployee/main.go
	go build -o bin/getEmployee/bootstrap functions/getEmployee/main.go      
	go build -o bin/getEmployees/bootstrap functions/getEmployees/main.go


