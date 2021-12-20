package main

import (
	"antlr4-demo/3-nums/parser"
	"log"
)

type DataListerner struct {
	*parser.BaseDataListener
}

func NewDataListerner() * DataListerner{
	return &DataListerner{
		BaseDataListener: &parser.BaseDataListener{},
	} 
}

// func (d *DataListerner) EnterSequence(ctx *parser.SequenceContext) {
// 	log.Println(ctx.GetI(), ctx.GetN())
// }

func (d *DataListerner) ExitGroup(ctx *parser.GroupContext) {
	//log.Println("i: ", ctx.Sequence().GetI())
	log.Println(ctx.INT().GetText(), " -> ", ctx.Sequence().GetText())
}