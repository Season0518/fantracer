# FanTracer

给长春Blossom（地下偶像团体）设计的粉丝群使用的调查/宣传工具，模块包括宣群机器人;水群情况查询等

## 开发状态

- [x] 宣群机器人 [/tools/chatbot]
- [x] 信息收集器 [/tools/collector]
- [ ] 网站后端(backend) [/server]
- [ ] 网站前端(frontend)

## 关于Config

config请放在和可执行文件的同级目录下

``config.ini``
``` ini

[CQHttp]
baseurl=http://[你的正向Http服务地址]
websocket=[你的正向WebSocket服务地址]
accessToken=[你的AccessToken]

[MySQL]
account=[MySQL账户]
password=[MySQL密码]
port=[MySQL服务端口]
```

``MsgConfig.json``
``` json
{
    "text" : "[宣群文案]",
    "media_url" : [
        "[宣群图床地址1]",
        "[宣群图床地址2]" 
    ]
}
```
