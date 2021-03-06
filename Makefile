#练习protobuf



 gRPC 默认提供了两种认证方式：

 基于SSL/TLS认证方式

 远程调用认证方式


    grpc服务端提供了interceptor功能，可以在服务端接收到请求时优先对请求中的数据做一些处理后再转交给指定的服务处理并响应，
功能类似middleware，很适合在这里处理验证、日志等流程。

    在自定义Token认证的示例中，认证信息是由每个服务中的方法处理并认证的，如果有大量的接口方法，
这种姿势就太蛋疼了，每个接口实现都要先处理认证信息。这个时候interceptor就站出来解决了这个问题，可以在请求被转到具体接口之前处理认证信息，一处认证，到处无忧


grpc默认提供了客户端和服务端的trace日志，可惜没有提供自定义接口，当前只能查看基本的事件日志和请求日志，对于基本的请求状态查看也是很有帮助的
这个可能和版本有关系，需要添加	grpc.EnableTracing = true在server.go的开头