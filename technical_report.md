Technical Report
================

(intro)


Server Framework Features
-------------------------

### Data Binding & Validation

Data Binding in Gin is where raw data is serialised into a specific layout, for example when serializing a JSON string into a struct. It can also be used to ensure that the data is in a valid format. In the Gin-Gonic; framework struct field tags are used to add rules for the validation library (Validator) to apply.

```go
type Person struct {
    // Field with a data-type and bindings for validation
    Name string `json:"name" binding:"required"`

func createPerson(c *gin.Context) {
    var person Person
    if err := c.ShouldBindJSON(&person); err != nil {
        // unsuccessful code here...
        return
    }
    // successful code here...
}
```

This is useful for this project as we have a REST API that accepts data from the client in JSON format. When serialising this data into a struct we can ensure that it is valid and the given fields also match the requirements. For the code example the field "Name" is marked as required, so it will produce an error when it is missing. Data binding has reduced the amount of code needed for validating the data, as we can simply specify the requirements inside a struct tag.

- [Gin-Gonic Data-Binding](https://gin-gonic.com/docs/examples/bind-query-or-post/)
- [Validator Lib](https://pkg.go.dev/gopkg.in/validator.v2)

### Middleware Support

Middleware in relation to the project is where we can add pre-processing and post-processing to requests and responses to add functionality and perform checks on data. It also allows us to add features without interacting directly with the middleware; as it simply wraps around the existing code.

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // add cors middleware, with default config
    r.Use(cores.Default())

    r.Run()
}
```

This is useful for the project as it is required to add CORS support to the routes. Using middleware cuts down on the amount of repeated code that would be required to add CORS to multiple routes. By having the same code run over all CORS routes we can ensure that if the code works for one route; it will work for all, this reduces bugs.

- [Gin-Gonic Middleware](https://gin-gonic.com/docs/examples/custom-middleware/)

### Minimal Code Required

Selecting Gin as the framework for the project allows for reduced code for main functionality. For example implementing routes to accept different methods (GET, POST, etc) is done by adding a method to the Gin engine for example:

```go
import "github.com/gin-gonic/gin"

func getIndex(c *gin.Context) {
    // single line to send a response back
    c.String(200, "Hello World")
}

func main() {
    // create a app using defaults (no extra config required)
    r := gin.Default()
    // single line to add a new GET route
    r.GET("/", getIndex)
    // Run the server with one line
    r.Run()
}
```

Having reduced "boilerplate" code allows for quicker implementation of new features and less bugs, due to repeated tasks like adding a new "GET" route with one line of code. It also helps future developers when they are maintaining/adding to the existing code, as it will look more consistent.

- [Gin-Gonic Quickstart](https://gin-gonic.com/docs/quickstart/)

Server Language Features
-----------------------

### (name of Feature 1)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


### (name of Feature 2)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


Client Framework Features
-------------------------

### (name of Feature 1)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


### (name of Feature 2)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


### (name of Feature 3)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


Client Language Features
------------------------

### (name of Feature 1)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)

### (name of Feature 2)

(Technical description of the feature - 40ish words - 1 mark)
(A code block snippet example demonstrating the feature - 1 mark)
(Explain the problem-this-is-solving/why/benefits/problems - 40ish words - 1 mark)
(Provide reference urls to your sources of information about the feature - required)


Critique of Server/Client prototype
---------------------

### (name of Issue 1)

(A code snippet example demonstrating the feature - 1 mark)
(Explain why this pattern is problematic - 40ish words 1 mark)

### (name of Issue 2)

(A code snippet example demonstrating the feature - 1 mark)
(Explain why this pattern is problematic - 40ish words 1 mark)


Future Technology Suggestions
-----------------------------

### (name of technology/feature 1)

(Description of a feature or tool - 40ish words - 1 mark)
(Why/benefits/problems with using this - 40ish words - 1 mark)
(Provide reference urls to your source of information about this technology - required)


### (name of technology/feature 2)

(Description of a feature or tool - 40ish words - 1 mark)
(Why/benefits/problems with using this - 40ish words - 1 mark)
(Provide reference urls to your source of information about this technology - required)


### (name of technology/feature 3)

(Description of a feature or tool - 40ish words - 1 mark)
(Why/benefits/problems with using this - 40ish words - 1 mark)
(Provide reference urls to your source of information about this technology - required)
