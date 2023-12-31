# Project 2 Milestone

I found this to be an easier assignment than [Project 1](../project1/README.md). I struggled a bit with static file handling. My solution differs a bit from the provided solution but seems to work. The output of [index.js](./webapp/static/js/index.js) is rendered in the DevTools console and [styles.css](./webapp/static/css/styles.css) impacts the output of the default page when viewed in a browser.

## Working with the Code

The code used in this project is the solution provided by the LiveLesson.

### Building the Application

```bash
cd webapp && go build -o webapp server.go
```

## Accessing the Application

Once the application is built, it is run using the following command:

```bash
./webapp
```

Once the application is running, you can run the following commands.

```bash
# Get default page
curl http://localhost:8080/
```

```bash
# Get API endpoint
curl http://localhost:8080/api
```

```bash
# Get Healthcheck endpoint
curl http://localhost:8080/healthcheck
```
