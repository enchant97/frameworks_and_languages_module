Technical Report
================

This technical report contains descriptions for language and framework features which will guide the engineering team to decide on the most effient development path to take and why using frameworks can benefit over writing an application from scratch. It will also indicate future technology that could be used to aid development of the FreeCycle software for production use.

Server Framework Features
-------------------------

### Data Binding & Validation

Data Binding in Gin is where raw data is serialised from JSON or XML using a predefined schema that is defined using a struct. Using data binding ensures that the raw data can be validated against the required schema. In the Gin-Gonic framework; struct field tags are used to add rules for the validation library (Validator) to check.

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

This is useful for this project as we have a REST API that accepts data from the client in JSON format. When serialising this data into a struct we can ensure that it is valid and the given fields also match the schema. For the code example the field "Name" is marked as required, so it will produce a specific error when it is missing. Because validation rules are applied using struct tags; the amount of code needed to be written is reduced, as the library implements the rules.

- [Gin-Gonic Data-Binding](https://gin-gonic.com/docs/examples/bind-query-or-post/)
- [Validator Lib](https://pkg.go.dev/gopkg.in/validator.v2)

### Middleware Support

Middleware allows adding pre-processing and post-processing to requests and responses to add functionality and perform checks on data. It also allows us to add features without interacting directly with the middleware; as it simply wraps around the existing code.

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

This is useful for the project as it allows for adding CORS support to the routes. Using middleware cuts down on the amount of repeated code that would be required to add CORS to multiple routes. By having the same code run over all CORS routes we can ensure that if the code works for one route; it will work for all.

- [Gin-Gonic Middleware](https://gin-gonic.com/docs/examples/custom-middleware/)

### Routing
Gin allows a url path to be mapped to a method of a specific route. It can also be configured to accept different methods (GET, POST, etc). A single route in Gin can be configured in a single line. Unlike Django, routing does not have to be setup in a specific file, it also keeps the method with the route mapping.

```go
// main.go
package main

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

```py
# urls.py
from django.urls import path

from . import views

urlpatterns = [
    # unclear whether this is GET, POST, etc
    path("/", views.index),
]
```

Keeping the route mappings with the accepted http method; reduces the need to look in multiple files to see what method will be called for a given route definition. Having a single line required to add a route; reduces "boilerplate" code, allowing for less lines of code to be required for a basic app to be implemented.

- [Gin-Gonic Quickstart](https://gin-gonic.com/docs/quickstart/)
- [Django Urls](https://docs.djangoproject.com/en/4.0/topics/http/urls/)

Server Language Features
-----------------------

### Error Handling
In Go errors are returned as a error type rather than being "raised and caught". Having errors use a specific type ensures that any errors that are returned can be handled by the calling code directly.

```go
func MyFunc() error {
    // method code here ...
}

func main() {
    if err := MyFunc(); err != nil {
        // handle errors here
    }
}
```

Having an error as a type means that Go's type assertion can be used to check and handle specific errors, this is useful as it allows specific errors to be handled differently. Since all created errors use the error type, handling all/unknown errors is possible; this will ensure that there are less unexpected results.

- [Go Error Handling](https://go.dev/blog/error-handling-and-go)

### Concurrency
Go is a concurrent language, meaning any functions created will always have the ability to be called asynchronously. To call a function asynchronously it can be wrapped in a goroutine, for example in the code below `go longRunningTask()` will run concurrently with any code after the method call.

```go
func longRunningTask() {
    // some long task ...
}

func main() {
    // run func in "background"
    go longRunningTask()
    // other code ...
}
```

Concurrency is an important factor for increasing performance. In an IO heavy environment, such as a web server; client requests can be paused and resumed while waiting for IO tasks to finish. This reduces the amount of cpu blocking while waiting and can reduce the amount of hardware needed for handling the same number of simultaneous clients.

Client Framework Features
-------------------------

### JSX
JSX is a syntax for building dynamic components using a syntax similar to HTML. SolidJS has modifications to ensure it looks closer to the HTML standards. If TypeScript is being used it will also be type checked.

```jsx
<div class="container">
  <h1>Welcome {username()}</h1>
  <button onclick={onClickMeClicked}>Click Me!</button>
</div>
```

JSX makes it more readable for developers when writing dynamic HTML for a component as it looks very similar to HTML. JSX also allows for Solid's reactivity to be added directly into the JSX for a component. Unlike HTML it does not allow unclosed tags, this will ensure that there are less unexpected results.

- [SolidJS JSX](https://www.solidjs.com/tutorial/introduction_jsx)

### Components
Components in SolidJS allows for code to be refactored, ensuring better separation and reuse without code duplication. As components can be detached from other code, they can even be imported from other locations meaning a library can be imported to use new components in a project.

```js
function User(props) {
  return <li>{props.username}</li>
}

function Users() {
  return (
    <ul>
      <User username="Leo" />
      <User username="Steve" />
    </ul>
  )
}
```

Components are useful as they allow us to maintain DRY code, this reduces the amount of code written. If the project was worked on by multiple developers it also allows for different teams to work on different sections of the project simultaneously; making the workflow more efficient.

- [SolidJS Component](https://www.solidjs.com/tutorial/introduction_components)

### Signals
To make an app reactive in SolidJS signals can be created. Signals can track and react to modifications, updating listening components when they change.

```jsx
import { createSignal } from "solid-js";

function Counter() {
  const [count, setCount] = createSignal(0);
  setInterval(() => setCount(count() + 1), 1000);
  return <div>Count: {count}</div>;
}
```

Signals ensure that all components that are "listening" get updated when the value is modified elsewhere. A difference with using signals compared to other frameworks (like React) is that they cannot be used like a normal variable, instead they are accessed using a get/set method, this ensures that the developer is always using the reactivity.

- [SolidJS Signals](https://www.solidjs.com/tutorial/introduction_signals)

Client Language Features
------------------------

### Strongly Typed
TypeScript adds type checking into JavaScript to help catch syntax and type errors before runtime. For the example shown below a "User" type is created, which can be added to the argument for the `log_user()` method. TypeScript will then ensure that any value passed to this method has that type, it will also ensure that usage of this value is supported by that type.

```ts
type User = {
  name: string
}

function log_user(user: User) {
  // Compiler knows this exists, so no warning
  console.log(user.name)
  // This will give a warning, as 'age' does not exist in the 'User' type
  console.log(user.age)
}
```

This is useful for developers (and the project) as it catches possible usage errors without running any of the code. It can also enable auto complete and suggestions for a developers editor.

- [TypeScript Static Type Checking](https://www.typescriptlang.org/docs/handbook/2/basic-types.html)

### Modules
TypeScript has support for modules which allows a TypeScript file to import specific features from another TypeScript file. These files could be in the same project or from an external library. These modules can also be imported multiple times in other modules.

```ts
// file: a.ts
export function hello() {
  console.log("hello world");
}

// file: main.ts
import { hello } from "./a";

function main() {
  hello()
}
```

Using modules with good naming conventions and an organised structure; allows for clear navigation to specific features, preventing the need to navigate through a single large file with many lines. This aids new and existing developers to understand where specific functionally is defined.

- [TypeScript Modules](https://www.typescriptlang.org/docs/handbook/modules.html)

Critique of Server/Client prototype
---------------------

### Server Handling of Split Packets

```py
# Snippet illustrating the issue, some code has been removed that is not related to the issue
def serve_app(func_app, port, host=''):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        with conn:
            data = conn.recv(65535)  # <-- Start of issue

            try:
                request = parse_request(data)
            except InvalidHTTPRequest as ex:
                log.exception("InvalidHTTPRequest")
                continue

            while int(request.get('content-length', 0)) > len(request['body']):  # <-- Fix that does not always work
                request['body'] += conn.recv(65535).decode('utf8')
```

The prototype server implementation has an issue, where if concurrent requests are sent and there are split packets; the server starts to operate non-deterministically and causes certain packets to be ignored. This is a problem with not only multiple users, but even having one active user making concurrent requests will cause these issues, this may result in data loss. This issue could of been solved by following the WSGI specification which most well known web frameworks follow.

- [Original Code Source](https://github.com/enchant97/frameworks_and_languages_module/blob/13eed800212051ec10804221e2aab8317f60e587/example_server/app/http_server.py#L112)

### Client Manual Creation/Modification Of Elements

```js
function renderItems(data) {
  const $item_list = document.querySelector(`[data-page="items"] ul`);
  const new_item_element = () => document.querySelector(`[data-page="items"] li`).cloneNode(true);

  for (let item_data of data) {
    const $new_item_element = new_item_element();
    $item_list.appendChild($new_item_element);
    renderDataToTemplate(item_data, $new_item_element, renderItemListFieldLookup);
    attachDeleteAction($new_item_element);
  }
}
```

The client prototype creates and modifies HTML elements manually. This is a problem as it is not clear what each function is doing. If another developer took over maintaining the implemented functions it would be harder to determine what these functions are doing, especially since the code is not commented. This could have been fixed by using a client framework which allows the use of JSX or HTML templates, removing the need to create elements using JavaScript.

- [Original Code Source](https://github.com/enchant97/frameworks_and_languages_module/blob/13eed800212051ec10804221e2aab8317f60e587/example_client/index.html)

Future Technology Suggestions
-----------------------------

### Serverless (AWS Lambda)
Using Serverless computing such as AWS Lambda allows for rapid deployment of application code without needing direct management of server infrastructure as the cloud provider handles this for you. Serverless does not have a idle running cost instead opting for paying per execution model, this can reduce costs when there is no active users. Using Serverless can also remove the need to think about scaling to keep up with users as it handles this automatically, with the advantage of directly running requests closest to users on demand; thus a service can be provided globally.

- [IBM Serverless](https://www.ibm.com/cloud/learn/serverless)

### NoSQL (MongoDB)
Using a NoSQL document database such as MongoDB allows for a more performant alternative to a traditional SQL database, that is also flexible to schema changes allowing for future application updates. MongoDB could be used as it supports both sharding and replication, which allows for distribution of data across multiple servers. This has the benefit of providing failover redundancy and for increasing performance with horizontal scaling. MongoDB also features a cloud service of their database called Atlas, which handles the maintenance of the database and provides automatic scaling. It can also integrate with Serverless computing.

- [MongoDB Sharding](https://www.mongodb.com/docs/manual/sharding/)
- [MongoDB Atlas](https://www.mongodb.com/atlas)

### Login Provider (auth0)
An external login provider/system will allow the handling of user accounts and user login. Using a service such as auth0 can increase security as it has been certified for compliance by third-party security professionals. Using a login provider also allows for integrations with other login services such as a Google login, this can make it easier for user to sign-up without entering different credentials. It also supports MFA which can further secure user accounts without needing any extra code to be written.

- [auth0](https://auth0.com/)
