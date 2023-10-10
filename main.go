package main

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/shopspring/decimal"
	"github.com/siddontang/go-log/log"
)

var (
	cursor = 1 // we will store this pos to consume the next position
	name   = "binlog.000002"
)

type Order struct {
	ID         uint64          `gorm:"column:id"`
	ChainID    string          `gorm:"column:chain_id"`
	Nonce      uint64          `gorm:"column:nonce"`
	MakerAsset string          `gorm:"column:maker_asset"`
	TakerAsset string          `gorm:"column:taker_asset"`
	Maker      string          `gorm:"column:maker"`
	Rate       decimal.Decimal `gorm:"column:rate"`
}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	b := BinlogParser{}

	for i := 0; i < len(e.Rows); i++ {
		var order Order
		err := b.GetBinLogData(&order, e, i)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("action: %v, order_id: %d, pos: %d \n\n", e.Action, order.ID, e.Header.LogPos)
		fmt.Printf("%+v \n", order)
	}

	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Password = "123456"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"orders"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	pos, err := c.GetMasterPos()
	if err != nil {
		log.Fatal(err)
	}
	// Start canal
	c.RunFrom(mysql.Position{
		Pos:  uint32(cursor),
		Name: pos.Name,
	})
}
