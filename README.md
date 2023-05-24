# In-Memory Datastore

This is a simple implementation of an in-memory datastore in Go. It provides basic data storage and retrieval functionalities without the need for an external database.

## Features

- Stores key-value pairs in memory
- Supports basic CRUD operations: create, read, update, and delete
- Thread-safe access and concurrent operations
- Simple and easy to use API

## Installation

To use the in-memory datastore in your Go project, you need to have Go installed and set up. Then, follow these steps:

1. Open a terminal and navigate to your project directory.
2. Use the following command to install the package:

```shell
go get github.com/rAJ-1312/in-memory-datastore
```

3. Import the package in your Go code:

```go
import "github.com/rAJ-1312/in-memory-datastore"
```

## Usage

Here's a basic example of how to use the in-memory datastore:

```go
package main

import (
	"fmt"

	"github.com/rAJ-1312/in-memory-datastore"
)

func main() {
	// Create a new datastore
	datastore := inmemorydatastore.New()

	// Set a value
	datastore.Set("key1", "value1")

	// Get a value
	value, exists := datastore.Get("key1")
	if exists {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}

	// Update a value
	datastore.Set("key1", "new-value1")

	// Delete a value
	datastore.Delete("key1")
}
```

For more detailed usage instructions and available methods, please refer to the [API documentation](https://github.com/rAJ-1312/in-memory-datastore/).

## Contributing

Contributions are welcome! If you find a bug or want to suggest an enhancement, please open an issue or submit a pull request on the [GitHub repository](https://github.com/rAJ-1312/in-memory-datastore).

## License

This project is licensed under the [MIT License](LICENSE).

---
