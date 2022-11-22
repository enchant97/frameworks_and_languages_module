Technical Report
================

(intro)


Server Framework Features
-------------------------

### Data Binding & Validation

Data Binding in Gin is where raw data is serialised into a specific layout, for example when serializing a JSON string into a struct. It can also be used to ensure that the data matches the given schema. In the Gin-Gonic; framework struct field tags are used to add rules for the validation library (Validator) to apply.

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

This is useful for the project as it is required to add CORS support to the routes. Using middleware cuts down on the amount of repeated code that would be required to add CORS to multiple routes. By having the same code run over all CORS routes we can ensure that if the code works for one route; it will work for all.

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

Having reduced "boilerplate" code allows for more time to be used for implementing new features. For example adding a new "GET" route only requires one line of code. Having less code will also aid future developers when they are maintaining/adding to the existing code, as there will be less duplicated code.

- [Gin-Gonic Quickstart](https://gin-gonic.com/docs/quickstart/)

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

Having an error as a type means that Go's type assertion can be used to check specific errors, this is useful as it allows different errors to be handled differently. Since all created errors use the error type, handling all/unknown errors is possible; this will ensure that there are less unexpected results.

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
<h1>Welcome {username()}</h1>
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

Components are useful as they allow us to maintain DRY code, this reduces the amount of code written. If the project was worked on by multiple developers it also allows for different teams to work on different sections of the project simultaneously; which makes the workflow more efficient.

- [SolidJS Component](https://www.solidjs.com/tutorial/introduction_components)

### Signals
To make an app reactive in SolidJS signals can be created. Signals can track modifications and update listening components when they change, this enables the app to update data and have it display updated on the users screen.

```jsx
import { createSignal } from "solid-js";

function Counter() {
  const [count, setCount] = createSignal(0);
  setInterval(() => setCount(count() + 1), 1000);
  return <div>Count: {count}</div>;
}
```

Signals ensure that all components that "listen" to one get updated when the value is modified elsewhere. A difference with using signals compared to other frameworks (like React) is that they cannot be used like a normal variable, instead they are accessed using a get/set method, this ensures that the developer is always using the reactivity.

- [SolidJS Signals](https://www.solidjs.com/tutorial/introduction_signals)

Client Language Features
------------------------

### Strongly Typed
TypeScript adds type checking into JavaScript to help catch errors before running any code. For the example shown below a "User" type is created, which can be added to the argument for the `log_user()` method. TypeScript will then ensure that any value passed to this method has that type, it will also ensure that usage of this value is supported by that type.

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
TypeScript has support for modules which allows a TypeScript file to import specific features from another TypeScript file. In TypeScript modules are also locally scoped; unless functionality is specifically "exported", allowing external usage.

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

Using modules is an important feature for a larger project with lots of code, as it allows for code to be separated into an organised structure. This aids developers in understanding where specific features are being used, allowing for better readability. Modules also allows for features to be used in multiple places.

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

The client prototype creates and modifies HTML elements manually. This is a problem it is not clear what each function is doing. If another developer took over maintaining the implemented functions it would be harder to determine what these functions are doing, especially since the code is not commented. This could have been fixed by using a client framework which allows the use of JSX or HTML templates, removing the need to create elements using JavaScript.

- [Original Code Source](https://github.com/enchant97/frameworks_and_languages_module/blob/13eed800212051ec10804221e2aab8317f60e587/example_client/index.html)

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
