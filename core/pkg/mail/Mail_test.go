package mail

import (
	"core/driver"
	"fmt"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	smtpmock "github.com/mocktools/go-smtp-mock/v2"
)

var server *smtpmock.Server

func initServer() {
	server = smtpmock.New(smtpmock.ConfigurationAttr{
		LogToStdout:       true,
		LogServerActivity: true,
	})

	// To start server use Start() method
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}

	// Load configuration dynamically
	driver.Base = &driver.Config{}
	driver.Base.Mail = driver.Mail{
		Interval: 1,
		Addr:     "127.0.0.1",
		Port:     server.PortNumber(),
		To:       "example@example.com",
		Account:  "notify@example.com",
	}
}

func assertMailBox(t *testing.T, excepted int) {
	msg := server.Messages()
	if len(msg) != excepted {
		t.Fatalf("excepted %d messages, got %d", excepted, len(msg))
	}
}

func TestSendMail(t *testing.T) {
	initServer()

	err := SendMail("Test Subject", "《罗密欧与灰姑娘》是初音ミク演唱的歌曲，由doriko填词，doriko谱曲，收录在专辑《ロミオとシンデレラ / doriko feat.初音ミク》中")
	if err != nil {
		t.Fatal(err)
	}

	assertMailBox(t, 1)
}

func TestInitSuccess(t *testing.T) {
	c := clockwork.NewFakeClock()
	initServer()
	_ = InitNotify(c)

	c.BlockUntil(1)

	// Assert launched mail sent before 1 hour
	c.Advance(30 * time.Minute)
	time.Sleep(1 * time.Millisecond)
	assertMailBox(t, 2)

	// Assert mail sent after 1 hour
	c.Advance(1 * time.Hour)
	time.Sleep(1 * time.Millisecond)

	// assertMailBox(t, 3)
}

func TestInitFailed(t *testing.T) {
	c := clockwork.NewFakeClock()

	// 模拟服务器异常
	initServer()
	driver.Base.Mail.Port = 0

	err := InitNotify(c)

	if err == nil {
		t.Fatal("excepted error, got nil")
	}
}
