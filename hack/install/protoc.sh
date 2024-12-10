#!/usr/bin/env bash

# max安装命令protoc命令
# export HOMEBREW_NO_AUTO_UPDATE=1 禁用brew自动更新或者命令后面加上参数--no-upgrade
# export HOMEBREW_BOTTLE_DOMAIN='mirrors.tuna.tsinghua.edu.cn' 更换镜像源
# 参考https://cloud.tencent.com/developer/article/1961001
brew install grpc
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc


# [配置go使用protobuf](https://www.cnblogs.com/jxzCoding/articles/16092494.html)
# 安装使用 protoc-gen-gogo
# go get github.com/gogo/protobuf/protoc-gen-gogo
# 进入到库的根目录【比如：cd $GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protoc-gen-gogo】
# 编译生成protoc-gen-gogo 【命令：go build】，并把生成的protoc-gen-gogo文件移到 $GOPATH/bin 目录下【命令mv protoc-gen-gogo $GOPATH】(mv protoc-gen-gogo $GOPATH/bin我的会移动到bin/bin这样还会报protoc-gen-gogo: program not found)方便后续使用
# 使用：protoc --gogo_out=. *.proto
# 注意：此处如果不走第2，3步，使用时可能会报错：protoc-gen-gogo: program not found or is not executable

