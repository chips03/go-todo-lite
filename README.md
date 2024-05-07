<div align='center'>

<h1>Go Todo Lite</h1>
<p>Simple Restful API for todo list using Golang and SQLite. Using fiber as HTTP Server and GORM to interact with the database.</p>

</div>

### :running: Run Locally

Clone the project

```bash
https://github.com/chips03/go-todo-lite
```
Start the server
```bash
go run main.go
```

### API Endpoint
#### url : http://127.0.0.1:3000

#### /todos
* `GET` : Get all todo's
* `POST` : Create a new todo
  
#### /todos/:id
* `GET` : Get a single todo
* `PUT` : Update a single todo
  
