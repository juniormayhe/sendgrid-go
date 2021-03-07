# Testing sendgrid with Golang

![](https://media.giphy.com/media/iz30qSwSKKCnC/giphy.gif)
## Create process

create folder and main.go
```
mkdir sendgrid-go
cd sendgrid-go
go mod init github.com/juniormayhe/sendgrid-go
create main.go file
go run main.go
```

## Set environment
```
echo "export SENDGRID_API_KEY='<your key>'" > sendgrid.env
echo "sendgrid.env" >> .gitignore

#Execute commands from a file in the current shell.
source ./sendgrid.env
```

install sendgrid package
```
go get github.com/sendgrid/sendgrid-go
```

create and verify the sender identity
https://app.sendgrid.com/settings/sender_auth/senders


## Send email
```
go run main.go
```

expected result
```
status code: 202

headers: map[
    Access-Control-Allow-Headers:[Authorization, Content-Type, On-behalf-of, x-sg-elas-acl]
    Access-Control-Allow-Methods:[POST]
    Access-Control-Allow-Origin:[https://sendgrid.api-docs.io]
    Access-Control-Max-Age:[600]
    Connection:[keep-alive]
    Content-Length:[0]
    Date:[Thu, 04 Mar 2021 19:49:33 GMT]
    Server:[nginx]
    Strict-Transport-Security:[max-age=600; includeSubDomains]
    X-Message-Id:[gOhVfdaKSXemu_KrEkQBmw]
    X-No-Cors-Reason:[https://sendgrid.com/docs/Classroom/Basics/API/cors.html]
]
```
