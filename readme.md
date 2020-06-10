## protobuf

### 分析

* 缺点  
1、 序列化之后可读性很差   
2、 需要额外的开发成本  
3、 适用性很小，不如json、xml  

* 优点   
1、支持多语言、跨平台，多个平台只需要维护一套协议  
2、序列号之后是二进制流，体积比json&xml小，适用数据量比较大的场景  
3、序列化与反序列化很快   
场景就是数据量大，响应速度要求很高。

### step

install tools
```
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go

// peformance is high
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/protoc-gen-gofast

```

生成数据模版

```
// 官方goprotobuf
protoc --go_out=. *.proto
// gogoprotobuf
protoc --gofast_out=. *.proto
```

### question 

* 性能比较与json、xml如何进行测试

### 注意点

1、数据类型使用需要考虑

### doc
* https://developers.google.com/protocol-buffers/docs/proto3