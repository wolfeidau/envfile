# envfile [![Build Status](https://drone.io/github.com/wolfeidau/envfile/status.png)](https://drone.io/github.com/wolfeidau/envfile/latest)

This module reads the contents env file and loads it into a map.

The format of the env file is key value pairs seperated by `=` and 
comments prefixed with `#`. White space on either end of the key or
value is trimmed.

```
# Some value
somval=123
# same as above as space on either side of tokens is trimmed
spaceval = 123
```

# Usage

```go

envMap := make(map[string]string)

err := ReadFile("./myapp.env", envMap)

```

# Licence

Copyright (c) 2013 Mark Wolfe and released under the MIT license.
