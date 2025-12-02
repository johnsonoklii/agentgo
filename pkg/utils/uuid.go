package utils

import "github.com/bwmarrin/snowflake"

func NewUUID() string {
	node, _ := snowflake.NewNode(1)
	id := node.Generate()
	return id.String()
}
