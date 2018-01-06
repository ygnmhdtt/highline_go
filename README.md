# highline_go - string masking

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [License](#license)

## Installation

This library doesn't depends on any other packages.

```
$ go get github.com/ygnmhdtt/highline_go
```

## Usage

```
package main

import (
  "os"
  "github.com/ygnmhdtt/highline
)

func main() {
  // You can specify where to output
	m := NewMaskWriter(os.Stdout)

  // Set any phrase(s) you want to filter
	m.SetFilter([]string{"password"})
	log.SetOutput(m)

	str := `{
"id": 1,
"password": "password",
"display_name": "test",
"email": "test@example.com"
 }`
	log.Print(str)
```

Output will be

```
{
"id": 1,
"password": "[FILTERED]",
"display_name": "test",
"email": "test@example.com"
}
```

## License
MIT
