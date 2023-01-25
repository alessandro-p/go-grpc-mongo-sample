## Readme

This project shows an example of gRPC combined with a simple MongoDB data layer. 

The server exposes a gRPC endpoint for managing a basic Blog: 
- Create a Post (Author, Title, Content)
- Get a post
- Update a post 
- Delete a post 
- List all posts (stream api)

In order to call those APIs, you can both use the given client, or you can use reflection through [evans](https://github.com/ktr0731/evans) (A CLI to run gRPC commands from your terminal)

Note: this project has been done for educational purpose only and can be considered as a starting point and not a as production ready application 

## Compile
In order to compile the project run from the root folder the following command: 

```bash
make blog
```

This command generates the required protos for gRPC and compile the binaries for the client and the server application.

## Run 
In order to run, you first need an up and running container running MongoDB. You can use the following command: 

```bash 
docker run -d -p 27017:27017 --name example-mongo mongo:latest 
```

Then open 2 different terminals. First run the server: 

```bash
 ./bin/blog/server
```

Then run the client

```bash
 ./bin/blog/client
```

## Example output
Example output from the terminal is: 
```
XXXX/XX/XX XX:XX:XX Creating posts...
XXXX/XX/XX XX:XX:XX Retrieving post list:
XXXX/XX/XX XX:XX:XX Post 0: id:"63d14e8f0e7fcc89deded37e"  author:"John Doe #0"  title:"Test Post #0"  content:"Test Content #0"
XXXX/XX/XX XX:XX:XX Post 1: id:"63d14e8f0e7fcc89deded37f"  author:"John Doe #1"  title:"Test Post #1"  content:"Test Content #1"
XXXX/XX/XX XX:XX:XX Post 2: id:"63d14e8f0e7fcc89deded380"  author:"John Doe #2"  title:"Test Post #2"  content:"Test Content #2"
XXXX/XX/XX XX:XX:XX Post 3: id:"63d14e8f0e7fcc89deded381"  author:"John Doe #3"  title:"Test Post #3"  content:"Test Content #3"
XXXX/XX/XX XX:XX:XX Post 4: id:"63d14e8f0e7fcc89deded382"  author:"John Doe #4"  title:"Test Post #4"  content:"Test Content #4"
XXXX/XX/XX XX:XX:XX Updating first post...
XXXX/XX/XX XX:XX:XX Post updated successfully
XXXX/XX/XX XX:XX:XX Getting first post...
XXXX/XX/XX XX:XX:XX Post now is: id:"63d14e8f0e7fcc89deded37e"  author:"Updated Author"  title:"Updated Title"  content:"Updated Content"
XXXX/XX/XX XX:XX:XX Clean up
XXXX/XX/XX XX:XX:XX Clean up successful
```

## Final notes & Credits 

The project is SSL ready. In order to introduce SSL you need to: 
- set `tlsEnabled` flag to true in both server and client `main.go` file
- add a `ssl` folder containing the certificates

Inspiration for this project has been taken from this [Udemy course](https://www.udemy.com/course/grpc-golang/learn/lecture/12351200?start=225#overview) 