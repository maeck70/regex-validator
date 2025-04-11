````markdown
# regex-validator

A simple Go library for validating strings against regular expressions.

[![Go Reference](https://pkg.go.dev/badge/github.com/maeck70/regex-validator.svg)](https://pkg.go.dev/github.com/maeck70/regex-validator)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## Introduction

This Go library provides a straightforward function to check if a given string matches a provided regular expression. It leverages Go's built-in `regexp` package for efficient and reliable pattern matching.

## Features

* **Simple API:** Exposes a single, easy-to-use function for validation.
* **Uses Go's `regexp`:** Built on the standard `regexp` package, ensuring performance and compatibility.
* **Clear Return Value:** Returns a boolean indicating whether the string matches the regex.

## Installation

To use this library in your Go project, simply use `go get`:

```bash
go get [github.com/maeck70/regex-validator](https://github.com/maeck70/regex-validator)
````

## Usage

Import the `regex-validator` package into your Go code:

```go
import "[github.com/maeck70/regex-validator](https://github.com/maeck70/regex-validator)"
```

Then, use the `Validate` function:

```go
package main

import (
	"fmt"
	"[github.com/maeck70/regex-validator](https://github.com/maeck70/regex-validator)"
)

func main() {
	regex := "^[a-z]+$"
	testString1 := "hello"
	testString2 := "hello99"

	isValid1 := regexvalidator.Validate(regex, testString1)
	fmt.Printf("'%s' matches '%s': %t\n", testString1, regex, isValid1) // Output: 'hello' matches '^[a-z]+$': true

	isValid2 := regexvalidator.Validate(regex, testString2)
	fmt.Printf("'%s' matches '%s': %t\n", testString2, regex, isValid2) // Output: 'hello99' matches '^[a-z]+$': false

	emailRegex := "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}"
	email1 := "test@example.com"
	email2 := "invalid-email"

	isValidEmail1 := regexvalidator.Validate(emailRegex, email1)
	fmt.Printf("'%s' matches '%s': %t\n", email1, emailRegex, isValidEmail1) // Output: 'test@example.com' matches '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}': true

	isValidEmail2 := regexvalidator.Validate(emailRegex, email2)
	fmt.Printf("'%s' matches '%s': %t\n", email2, emailRegex, isValidEmail2) // Output: 'invalid-email' matches '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}': false
}
```

## Contributing

Contributions are welcome\! If you find a bug or have an idea for improvement, please:

1.  Fork the repository.
2.  Create a new branch for your changes.
3.  Make your modifications and ensure they are well-tested.
4.  Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://www.google.com/search?q=LICENSE) file for more details.

-----

**Note:** This README.md is generated based on the provided Go file path. If there are other files or functionalities in the repository, you might want to update this README to reflect the complete project structure and features. For example, if there's a command-line interface, you should add a section explaining its usage.

```
```
