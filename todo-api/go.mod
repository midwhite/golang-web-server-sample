module github.com/midwhite/golang-web-server-sample/todo-api

go 1.19

replace github.com/midwhite/golang-web-server-sample/todo-api/router => ./router

replace github.com/midwhite/golang-web-server-sample/todo-api/controllers => ./controllers

require github.com/midwhite/golang-web-server-sample/todo-api/router v0.0.0-00010101000000-000000000000

require github.com/midwhite/golang-web-server-sample/todo-api/controllers v0.0.0-00010101000000-000000000000 // indirect
