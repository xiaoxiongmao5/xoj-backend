package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20230925_180113 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20230925_180113{}
	m.Created = "20230925_180113"

	migration.Register("User_20230925_180113", m)
}

// Run the migrations
func (m *User_20230925_180113) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) DEFAULT NULL,`user_account` varchar(128) NOT NULL,` user_password` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *User_20230925_180113) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
