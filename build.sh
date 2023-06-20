#!/bin/bash

# 创建 bin 目录，如果它不存在
mkdir -p bin

# 设定目标操作系统和体系结构
platforms=("windows/amd64" "linux/amd64" "darwin/arm64")

# 设定需要编译的模块
modules=("./server" "./tools/collector" "./tools/chatbot")

# 循环遍历所有平台
for platform in "${platforms[@]}"
do
    # 使用/分隔符分割平台变量
    platform_split=(${platform//\// })

    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    # 输出当前编译平台信息
    echo "Building for $GOOS/$GOARCH"

    # 遍历所有模块
    for module in "${modules[@]}"
    do
        # 获取模块名称
        module_name=$(basename $module)

        # 编译模块
        GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/$module_name-$GOOS-$GOARCH" $module
    done
done

