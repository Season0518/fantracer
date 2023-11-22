# FanTracer

给长春Blossom（地下偶像团体）设计的粉丝群使用的调查/宣传工具，模块包括宣群机器人;水群情况查询等

## 开发状态

- [x] 宣群机器人 [/tools/chatbot]
- [x] 信息收集器 [/tools/collector]
- [ ] 网站后端(backend) [/server]
- [ ] 网站前端(frontend)

## 关于Config

config请放在和可执行文件的同级目录下,以下是一个`config.yaml`的示例
```
# config.yaml

# 请填写基于Onebot V11协议的机器人地址和token
Bot: 
  http: "http://localhost:25566"
  websocket: "ws://localhost:10086"
  access_token: "A1b2C3d4e5F6g7H&*()!@#$%^&*(;/.,<>?"
 
MySQL:
  account: "root"
  password: "ex!a$m*pLepAsSw0rd"
  port: 3306

# 填写欢迎文案，当group_increase事件发生时被机器人推送。
# media_url为图片链接，可选，当不填写时不会发送图片。
# 当media_url为多个链接时，会随机选择一个发送。

Greeting:
  - target: 1234567890
    text: "我是一条带海报的文案️"
    media_url:
      - "https://bkimg.cdn.bcebos.com/pic/4b90f603738da9773912aed66105ef198618377a458b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5/format,f_auto"
      - "https://bkimg.cdn.bcebos.com/pic/728da9773912b31bb051e7d4574c217adab44bed4b8b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UxNTA=,g_7,xp_5,yp_5/format,f_auto"
      - "https://bkimg.cdn.bcebos.com/pic/f703738da9773912b31bae9c294d9118367adbb4448b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyMjA=,g_7,xp_5,yp_5/format,f_auto"
  - target: 9876543210
    text: "PlainText，一条朴素的文案"

# 邮件推送服务
# 其中服务状态会间隔post_interval的小时向指定邮箱发送。
# 当post_interval为0时，邮件服务将被停用。
Mail:
  post_interval: 48
  smtp_address: "smtp.example.com"
  smtp_port: 465
  sender_account: "notify@exmple.com"
  sender_key: "secretKey4authorization"
  mail_to: "receiver@example1.com"
```
