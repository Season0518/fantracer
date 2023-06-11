#!/bin/bash

# 创建 bin 目录，如果它不存在
mkdir -p bin

# 编译 server
go build -o bin/server ./server

# 编译 tools
go build -o bin/collector ./tools/collector
# go build -o bin/chatbot ./tools/chatbot