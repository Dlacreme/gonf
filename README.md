# Gonf - Light & Powerful configuration library

Gonf let you easily configure your golang application.

## Install

## Usage

Create a folder `config` and declare your config files from there.
Then set a single `GONFIG_ENV` environment variable matching a file available in `config`.

### API

#### Setup

```golang
// Set the folder hosting your config files (by default "gonfig/")
func SetFolder(pathname string) {}

// Set the base config file to use (by default "base.gonf")
func SetBase(filename string) {}

// Set the override file to use (by default `dev.gonf` or GONF_ENV)
func Use(filename string) {}
```

#### Get configuration

```golang
// Retrieve a single value
func Get(key string) string {}

// Retrieve many values
func GetMany(key []string) []string {}

// Retrieve all values
func GetAll() []string {}
```

#### Preload

You can preload many values to make sure they are cached and available.

```golang
// Load in cache multiple values
func PreloadMany(key []string) error {}

// Load all the config in cache
func PreloadAll() error {}
```

#### Exposed internal functions

##### Load and validate a file as a value
```golang
// Get a file content or empty string ("") if missing
func GetFileContent(key string) string {}

// Get a file content or `def` if missing
func GetFileContentOrDefault(key string, def string) string {}

// Get a file content or panic with appropriate message if missing
func GetFileContentOrPanic(key string) string {}
```

##### Load and validate an env variable

```golang
// Get an env var or empty string ("") if missing
func GetEnv(key string) string {}

// Get an env var or `def` if missing
func GetEnvOrDefault(key string, def string) string {}

// Get an env var or panic with appropriate message if missing
func GetEnvOrPanic(key string) string {}
```