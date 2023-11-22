package driver

import (
	"os"
	"reflect"
	"testing"
)

func TestInitCfg(t *testing.T) {
	// 创建一个临时的配置文件
	content := []byte(`
Bot: 
  http: "http://localhost:25566"
  websocket: "ws://localhost:10086"
  access_token: "A1b2C3d4e5F6g7H&*()!@#$%^&*(;/.,<>?"
MySQL:
  account: "root"
  password: "ex!a$m*pLepAsSw0rd"
  port: 3306

Greeting:
  - target: 1234567890
    text: "✨🙏😩*&^%$$#你好Helloこんにちは안녕하세요🙏✨💗❤️"
    media_url:
      - "https://bkimg.cdn.bcebos.com/pic/4b90f603738da9773912aed66105ef198618377a458b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5/format,f_auto"
      - "https://bkimg.cdn.bcebos.com/pic/728da9773912b31bb051e7d4574c217adab44bed4b8b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UxNTA=,g_7,xp_5,yp_5/format,f_auto"
      - "https://bkimg.cdn.bcebos.com/pic/f703738da9773912b31bae9c294d9118367adbb4448b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyMjA=,g_7,xp_5,yp_5/format,f_auto"
  - target: 9876543210
    text: "PlainText，一条朴素的中文English测试"
#  - text: ""
#    media_url:
#      - ""
#      - ""

Mail:
  post_interval: 48
  smtp_address: "smtp.example.com"
  smtp_port: 465
  sender_account: "notify@exmple.com"
  sender_key: "secretKey4authorization"
  mail_to: "receiver@example1.com"
`)
	expected := &Config{
		Bot: Bot{
			Http:        "http://localhost:25566",
			WebSocket:   "ws://localhost:10086",
			AccessToken: "A1b2C3d4e5F6g7H&*()!@#$%^&*(;/.,<>?",
		},
		MySQL: MySQL{
			Account:  "root",
			Password: "ex!a$m*pLepAsSw0rd",
			Port:     3306,
		},
		Greeting: []Greeting{
			{
				Target: 1234567890,
				Text:   "✨🙏😩*&^%$$#你好Helloこんにちは안녕하세요🙏✨💗❤️",
				MediaURL: []string{
					"https://bkimg.cdn.bcebos.com/pic/4b90f603738da9773912aed66105ef198618377a458b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5/format,f_auto",
					"https://bkimg.cdn.bcebos.com/pic/728da9773912b31bb051e7d4574c217adab44bed4b8b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UxNTA=,g_7,xp_5,yp_5/format,f_auto",
					"https://bkimg.cdn.bcebos.com/pic/f703738da9773912b31bae9c294d9118367adbb4448b?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyMjA=,g_7,xp_5,yp_5/format,f_auto",
				},
			},
			{
				Target: 9876543210,
				Text:   "PlainText，一条朴素的中文English测试",
			},
		},
		Mail: Mail{
			Interval: 48,
			Addr:     "smtp.example.com",
			Port:     465,
			Account:  "notify@exmple.com",
			Key:      "secretKey4authorization",
			To:       "receiver@example1.com",
		},
	}

	tempFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Failed to release temp file: %v", err)
		}
	}(tempFile.Name())

	// 写入测试数据到临时文件
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("写入临时文件失败: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("关闭临时文件失败: %v", err)
	}

	err = InitCfg(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(Base, expected) {
		t.Errorf("LoadConfig() = %v, want %v", Base, expected)
	}
}
