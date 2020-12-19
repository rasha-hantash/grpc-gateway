# grpc-gateway
## Build Binaries
```bash
make all 
```

## Run Binaries
Start the server:
```bash
./bin/server
```
Then run the client a few times
```bash
./bin/client
```
Now visit ```http://localhost:8080/swagger-ui``` to make REST calls interactively. <br />
Alternatively, use curl: <br />
```bash
curl -X GET "http://localhost:8080/api/users" -H "accept: application/json" ; echo
```
```bash
curl -X GET "http://localhost:8080/api/v1/users" -H "accept: application/json" ; echo
```
## Build Container
```bash
make image
```
## Run Container
```bash
docker run -p 5566:5566 grpc-docker:latest
```

Now visit ```http://localhost:8080/swagger-ui``` to make REST calls interactively. <br />
Alternatively, use curl: <br />
```bash
curl -X POST "http://localhost:8080/api/users" -H "accept: application/json" ; echo
```
```bash
curl -X GET "http://localhost:8080/api/v1/users" -H "accept: application/json" ; echo
```
## Regenerating code after proto file changes
The gRPC and REST bindings plus swagger file are generated automatically from the proto file. The generated files are committed to the repo so you don't need to run these commands to try the code. <br />

However, if you make changes to the proto file you'll need to regenerate like this: <br />
```bash
make generate
```
