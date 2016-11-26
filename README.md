# Go Universal Manager

A command line tool for managing environments for applications.

Helps automate the process of configuring, running, and testing environments for common application configurations.
Agnostic between language (Go, Ruby, Python, ...), and infrastructure (PostgreSQL, MySQL, MongoDB, ...)

This is a work-in-progress and isn't in a working state yet.


## TODO:
- Create a YAML registar for registering various levels of YAML parser (pre-application, in application, in services, ...)
- Build YAML config parser (register settings / sections) and define functions to run on commands with a priority (essentially dictate order) for a) running things b) docker configurations

## Usage

### Example configuration
```yaml
application:
  name: Example Python Web App
  cmd: python3 main.py --port=80
  language:
    python: 3.5
  services:
    postgresql:
      version: 9.4
      address: localhost:5432
  proxies:
    80: localhost:8000
```

### Example usage
Once you have a configuration, such as the example above you can start the application with
```bash
gum start
```

And this will prepare and start all of the required services without relying on any local dependencies. As per the configuration, our application is serving on port `80` and we are binding our `localhost:8000` to the port. If you instead wanted to serve it to others, you can user `80: 0.0.0.0:8000`
