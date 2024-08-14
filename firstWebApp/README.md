# Go `fmt` Package Overview

The `fmt` package in Go is a core package used for formatted I/O (input/output). It's one of the most commonly used packages in Go, providing a variety of functions to format and print text. Here's a breakdown of the most commonly used functions within the `fmt` package, along with when and why to use them:

## 1. `fmt.Println`

- **Usage**: `fmt.Println(a ...interface{}) (n int, err error)`
- **Details**: Prints the provided arguments followed by a newline (`\n`). Automatically adds a space between arguments if there are multiple.
- **Best Suited For**: Printing multiple values in a human-readable format, with each output on a new line. Ideal for quick debugging or simple console output.
- **Example**:
  ```go
  fmt.Println("Hello", "World") // Output: Hello World
  fmt.Println(123, 456)         // Output: 123 456
  ```

## 2. `fmt.Print`

- **Usage**: `fmt.Print(a ...interface{}) (n int, err error)`
- **Details**: Similar to `Println`, but it doesn't add a newline at the end of the output. Doesn't automatically insert spaces between arguments.
- **Best Suited For**: Printing multiple values without starting a new line or when precise formatting of output is important.
- **Example**:
  ```go
  fmt.Print("Hello") // Output: Hello
  fmt.Print("World") // Output: HelloWorld
  ```

## 3. `fmt.Printf`

- **Usage**: `fmt.Printf(format string, a ...interface{}) (n int, err error)`
- **Details**: Allows formatted output using a format specifier string. Use verbs like `%s`, `%d`, `%v`, etc., to control the output format.
- **Best Suited For**: More control over output formatting, such as with complex data structures or specific formatting requirements.
- **Example**:
  ```go
  name := "John"
  age := 25
  fmt.Printf("Name: %s, Age: %d\n", name, age) // Output: Name: John, Age: 25
  ```

## 4. `fmt.Sprint`

- **Usage**: `fmt.Sprint(a ...interface{}) string`
- **Details**: Returns the formatted string instead of printing it directly. Useful for constructing strings without immediately outputting them.
- **Best Suited For**: Concatenating or building strings for later use rather than direct printing.
- **Example**:
  ```go
  greeting := fmt.Sprint("Hello", " ", "World")
  fmt.Println(greeting) // Output: Hello World
  ```

## 5. `fmt.Sprintf`

- **Usage**: `fmt.Sprintf(format string, a ...interface{}) string`
- **Details**: Works like `Printf` but returns the formatted string instead of printing it.
- **Best Suited For**: Creating a formatted string with specific content for later use, such as logging, error messages, or building complex strings.
- **Example**:
  ```go
  name := "Jane"
  message := fmt.Sprintf("Hello, %s!", name)
  fmt.Println(message) // Output: Hello, Jane!
  ```

## 6. `fmt.Errorf`

- **Usage**: `fmt.Errorf(format string, a ...interface{}) error`
- **Details**: Formats according to a format specifier and returns an error type.
- **Best Suited For**: Creating custom error messages that include formatted data, useful in error handling.
- **Example**:
  ```go
  err := fmt.Errorf("an error occurred: %s", "file not found")
  fmt.Println(err) // Output: an error occurred: file not found
  ```

## 7. `fmt.Fprintln`

- **Usage**: `func Fprintln(w io.Writer, args ...interface{}) (n int, err error)`
- **Details**: Formats and writes arguments to the `io.Writer`, followed by a newline (`\n`). Always includes a newline at the end.
- **Best Suited For**: Writing data to an `io.Writer` with a guaranteed newline.
- **Example**:

  ```go
  file, err := os.Create("output.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  fmt.Fprintln(file, "Hello World")
  // Writes "Hello World" followed by a newline to "output.txt"
  ```

## 8. `fmt.Fprintf`

- **Usage**: `func Fprintf(w io.Writer, format string, args ...interface{}) (n int, err error)`
- **Details**: Formats and writes data to the `io.Writer` using format specifiers.
- **Best Suited For**: Formatting data before writing it to an output destination.
- **Example**:

  ```go
  file, err := os.Create("output.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  fmt.Fprintf(file, "Hello %s!\n", "World")
  // Writes "Hello World!" to "output.txt"
  ```

## 9. `fmt.Fprint`

- **Usage**: `func Fprint(w io.Writer, args ...interface{}) (n int, err error)`
- **Details**: Writes arguments to the `io.Writer` without formatting.
- **Best Suited For**: Writing raw data to an output destination.
- **Example**:

  ```go
  file, err := os.Create("output.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  fmt.Fprint(file, "Hello World")
  // Writes "Hello World" to "output.txt" without newlines or extra spaces
  ```

## 10. `fmt.Fscan`, `fmt.Fscanf`, `fmt.Fscanln`

- **Usage**: Read formatted input from an `io.Reader`:
  - `Fscan` reads with whitespace as a separator.
  - `Fscanf` reads formatted input using format specifiers.
  - `Fscanln` reads with newlines as record separators.
- **Details**: Useful for reading structured input data from files or other streams.
- **Example**:

  ```go
  file, err := os.Open("input.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  var name string
  var age int
  fmt.Fscanf(file, "%s %d", &name, &age)
  // Reads formatted data from "input.txt"
  ```

## Summary

- **`Println`**: Simple output with automatic newlines.
- **`Print`**: Simple output without newlines.
- **`Printf`**: Formatted output with control over layout.
- **`Sprint`**: Constructs strings without printing.
- **`Sprintf`**: Formats strings for later use.
- **`Errorf`**: Creates formatted error messages.
- **`Fprintln`**: Writes arguments to an `io.Writer` with a newline at the end.
- **`Fprintf`**: Formats and writes arguments to an `io.Writer`.
- **`Fprint`**: Writes arguments to an `io.Writer` without formatting.
- **`Fscan`, `Fscanf`, `Fscanln`**: Read formatted input from an `io.Reader`.

These functions are used for a range of tasks, from simple debugging to complex string manipulation and error handling.

---

# Go `net/http` Package Overview

The `net/http` package in Go provides HTTP client and server functionality. It’s commonly used for creating web servers and making HTTP requests. Here’s a summary of the most commonly used functions and types in the `net/http` package, including details on when and why to use them:

## 1. `http.Get`

- **Usage**: `func Get(url string) (resp *http.Response, err error)`
- **Details**: Sends a GET request to the specified URL and returns the response.
- **Best Suited For**: Simple HTTP GET requests to fetch resources from a URL.
- **Example**:
  ```go
  resp, err := http.Get("https://example.com")
  if err != nil {
      log.Fatal(err)
  }
  defer resp.Body.Close()
  ```

## 2. `http.Post`

- **Usage**: `func Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)`
- **Details**: Sends a POST request to the specified URL with the given content type and body.
- **Best Suited For**: Sending data to a server, such as form submissions or API requests.
- **Example**:
  ```go
  resp, err := http.Post("https://example.com/post", "application/json", bytes.NewBuffer([]byte(`{"key":"value"}`)))
  if err != nil {
      log.Fatal(err)
  }
  defer resp.Body.Close()
  ```

## 3. `http.NewRequest`

- **Usage**: `func NewRequest(method, url string, body io.Reader) (*http.Request, error)`
- **Details**: Creates a new HTTP request with the specified method, URL, and body.
- **Best Suited For**: Customizing HTTP requests with methods other than GET or POST, or adding headers or other configurations.
- **Example**:
  ```go
  req, err := http.NewRequest("GET", "https://example.com", nil)
  if err != nil {
      log.Fatal(err)
  }
  req.Header.Set("Accept", "application/json")
  ```

## 4. `http.Do`

- **Usage**: `func (c *Client) Do(req *http.Request) (*http.Response, error)`
- **Details**: Executes a given HTTP request and returns the response. Used with an `http.Client`.
- **Best Suited For**: Executing a custom request created with `http.NewRequest` and handling responses.
- **Example**:
  ```go
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      log.Fatal(err)
  }
  defer resp.Body.Close()
  ```

## 5. `http.HandleFunc`

- **Usage**: `func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))`
- **Details**: Registers a handler function for a given route pattern. This function is called when an HTTP request matches the pattern.
- **Best Suited For**: Setting up route handling in a web server.
- **Example**:
  ```go
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, World!")
  })
  ```

## 6. `http.ListenAndServe`

- **Usage**: `func ListenAndServe(addr string, handler http.Handler) error`
- **Details**: Starts an HTTP server with the given address and handler.
- **Best Suited For**: Starting a web server that listens for incoming HTTP requests.
- **Example**:
  ```go
  http.ListenAndServe(":8080", nil)
  ```

## 7. `http.ResponseWriter`

- **Usage**: Interface provided to handler functions to construct HTTP responses.
- **Details**: Used within handler functions to write the response to the client. Methods include `Write`, `Header`, and `WriteHeader`.
- **Best Suited For**: Building and sending HTTP responses in a web server.
- **Example**:
  ```go
  func handler(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(http.StatusOK)
      fmt.Fprintf(w, "Hello, World!")
  }
  ```

## 8. `http.Request`

- **Usage**: Struct representing an HTTP request.
- **Details**: Contains all information about the incoming request, such as method, URL, headers, and body.
- **Best Suited For**: Accessing details of the request in your handler functions.
- **Example**:
  ```go
  func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Println(r.Method) // e.g., GET
      fmt.Println(r.URL)    // e.g., /path
  }
  ```

## Summary

- **`Get`**: Simple HTTP GET requests.
- **`Post`**: HTTP POST requests with a body.
- **`NewRequest`**: Customizing HTTP requests.
- **`Do`**: Executing custom HTTP requests.
- **`HandleFunc`**: Routing HTTP requests to handler functions.
- **`ListenAndServe`**: Starting an HTTP server.
- **`ResponseWriter`**: Writing responses in handlers.
- **`Request`**: Accessing request details.

These functions and types are crucial for building and interacting with web services in Go, whether you're making HTTP requests or creating a web server.
