package test

import (
	"bytes"
	"cloud_disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	var err error
	eng, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud_disk")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	if err = eng.Find(&data); err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	b, err := json.Marshal(data)
	dst := new(bytes.Buffer)
	json.Indent(dst, b, "", " ")
	fmt.Println(dst.String())
}
func TestUUID(t *testing.T) {
	v4, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v4.String())
	return
}
