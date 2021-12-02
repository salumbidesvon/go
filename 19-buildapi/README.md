## 🏁 Getting Started <a name = "getting_started"></a>

Run 
```
go get -u github.com/gorilla/mux

go mod tidy // bring all the packages for this project and remove that are not being used 

go mod verify

go list -m all

go list -m -versions github.com/gorilla/mux

go mod vendor

go build -mod=readonly -o ./backend
```

## Run docker compose to build the project in your local machine
```
make
```

## Prerequisite
Go mux
https://github.com/gorilla/mux