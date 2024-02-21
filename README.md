# Template

Create your project with one command

![](assets/ghostemane.gif)

## Parameters

To see the available commands

```bash
    go run github.com/perebaj/jjtemplate@latest --help
```

To run the project

```bash
    go run github.com/perebaj/jjtemplate@latest --name jjisawesome --registry perebaj
```

* **name(required)**: Set up here the name of your new project.
* **registry**: Docker all the way. **Default**: `fakeregistry`
* **output**: Set up the folder where you need to save your things. **Default**: `./`
* **compose**: Set up if you need to generate a docker-compose.yml file. **Default**: `false`


# Usage example

[![asciicast](https://asciinema.org/a/tPNMTcoPLqsAyf3VWjUjsbyty.svg)](https://asciinema.org/a/tPNMTcoPLqsAyf3VWjUjsbyty)

# Run tests

`make test`

then you can access the output folder and run the make commands to verify if all it's working well
