# go-kit-demo

## go-kit-server

- 启动consul
  ```shell
  docker run -d --name consul -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
  ```

- 访问consul
  <http://localhost:8500/ui/dc1/services>

- 启动服务端
  ```shell
  go run . --name user --port 8080
  ```

- 启动客户端
  ```shell
  go run . --name user
  ```
