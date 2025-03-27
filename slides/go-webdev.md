

# Web Development with Go: A Beginner-Friendly Guide

## 🌐 What is HTTP?

HTTP (Hypertext Transfer Protocol) is the backbone of data communication on the web. It follows a **client-server model** where:

- A **client** (browser) sends an HTTP request
- A **server** responds with an HTTP response

### 🔑 Key Characteristics:

- **Stateless**: Each request is independent
- **Resources & URLs**: Access resources via URLs (e.g., `http://example.com/index.html`)
- **Methods**: GET, POST, PUT, DELETE, etc.
- **Messages**: Structured HTTP requests/responses

---

## 🔁 TCP vs UDP

HTTP typically runs on **TCP**, but let’s understand how it differs from **UDP**:

| Feature        | TCP                            | UDP                           |
|----------------|----------------------------------|--------------------------------|
| Connection     | Connection-oriented              | Connectionless                 |
| Reliability    | Guaranteed delivery & ordering   | No guarantees                  |
| Speed          | Slower due to overhead           | Faster, lightweight            |
| Use Cases      | HTTP, file transfer, emails      | Streaming, gaming, DNS         |

**In Web Development:**  
HTTP runs on **TCP** for reliable delivery of web content.

---

## 📟 Common HTTP Status Codes

| Code | Meaning               | Description                                                                 |
|------|------------------------|-----------------------------------------------------------------------------|
| 200  | OK                     | Request succeeded                                                           |
| 201  | Created                | Resource created                                                            |
| 301  | Moved Permanently      | Resource has a new permanent URL                                           |
| 400  | Bad Request            | Client error, invalid request syntax                                       |
| 401  | Unauthorized           | Authentication required                                                    |
| 403  | Forbidden              | Authenticated, but not allowed                                             |
| 404  | Not Found              | Resource not found                                                         |
| 500  | Internal Server Error  | Server error                                                               |
| 503  | Service Unavailable    | Server temporarily unavailable                                             |

---

## ⚙️ HTTP/1.1: Persistent Connections

HTTP/1.1 improved upon HTTP/1.0 by introducing:

- **Persistent connections** (keep-alive): Reuse TCP connection for multiple requests
- **Reduced latency**: Fewer handshakes, faster resource loading
- **Parallel connections**: Browsers use multiple connections for performance
- **Pipelining** (optional, but rarely used)

---

## 🧪 Build a Basic Web Server in Go

### Step 1: Hello World Server

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Run the above code with:

```bash
go run main.go
```

Visit `http://localhost:8080/` to see your server in action.

---

## 🗂️ Serving Static HTML Files

### Directory Structure:

```
project/
├── main.go
└── static/
    └── index.html
```

### Code:

```go
func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
```

This will serve any file from the `./static` directory, including HTML, CSS, and JS.

---

## 🧱 What is a REST API?

**REST (Representational State Transfer)** is an architectural style that uses HTTP methods to operate on resources.

### Key Concepts:

- **Stateless**
- **URI-based resources** (e.g. `/todos/1`)
- **HTTP methods** for CRUD:
    - GET: Read
    - POST: Create
    - PUT: Update
    - DELETE: Delete
- **Standardized status codes**
- **JSON** is the common data format

---

## 🚀 Build a REST API in Go with Gin

### 🔧 What is Gin?

Gin is a fast, lightweight HTTP framework in Go. It helps with:

- Routing
- Middleware support
- JSON handling
- Error recovery

### 📦 Installation

```bash
go get github.com/gin-gonic/gin
```

### ✅ Example: Todo API

#### Step 1: Define Todo Model

```go
type Todo struct {
    ID   string `json:"id"`
    Task string `json:"task"`
    Done bool   `json:"done"`
}

var todos = []Todo{
    {ID: "1", Task: "Learn Go", Done: true},
    {ID: "2", Task: "Build API", Done: false},
}
```

#### Step 2: CRUD Handlers

- **List Todos**

```go
func getTodos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, todos)
}
```

- **Get by ID**

```go
func getTodoByID(c *gin.Context) {
    id := c.Param("id")
    for _, t := range todos {
        if t.ID == id {
            c.IndentedJSON(http.StatusOK, t)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
```

- **Create Todo**

```go
func createTodo(c *gin.Context) {
    var newTodo Todo
    if err := c.BindJSON(&newTodo); err != nil {
        return
    }
    newTodo.ID = strconv.Itoa(len(todos) + 1)
    todos = append(todos, newTodo)
    c.IndentedJSON(http.StatusCreated, newTodo)
}
```

- **Update Todo**

```go
func updateTodo(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo Todo
    if err := c.BindJSON(&updatedTodo); err != nil {
        return
    }
    for i, t := range todos {
        if t.ID == id {
            todos[i].Task = updatedTodo.Task
            todos[i].Done = updatedTodo.Done
            c.IndentedJSON(http.StatusOK, todos[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
```

- **Delete Todo**

```go
func deleteTodo(c *gin.Context) {
    id := c.Param("id")
    for i, t := range todos {
        if t.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
```

#### Step 3: Register Routes and Run Server

```go
func main() {
    r := gin.Default()

    r.GET("/todos", getTodos)
    r.GET("/todos/:id", getTodoByID)
    r.POST("/todos", createTodo)
    r.PUT("/todos/:id", updateTodo)
    r.DELETE("/todos/:id", deleteTodo)

    r.Run(":8080")
}
```

---

## 🧪 Test the API

### List Todos:

```bash
curl http://localhost:8080/todos
```

### Create Todo:

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"task":"Write guide","done":false}'
```

### Update Todo:

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"task":"Updated task","done":true}'
```

### Delete Todo:

```bash
curl -X DELETE http://localhost:8080/todos/1
```

---


