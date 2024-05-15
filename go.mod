module tour-service

go 1.22.1

require (
	github.com/google/uuid v1.6.0
	gorm.io/driver/mysql v1.5.5
	gorm.io/gorm v1.25.8
	soa/grpc/proto v0.0.1
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace soa/grpc/proto => /app/proto
