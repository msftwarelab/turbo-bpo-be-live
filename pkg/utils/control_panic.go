package utils

import (
	"fmt"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

var nextID = make(chan uint16)

func Counter(num uint16) {
	if num == 0 {
		for i := uint16(1); ; i++ {
			nextID <- i
		}
	}
	for i := uint16(1); ; i++ {
		nextID <- i
		if i == num {
			panic("Many mistakes have happened!")
		}
	}
}

func PanicCtrl(ctx *fasthttp.RequestCtx) {
	if r := recover(); r != nil {
		num := <-nextID
		log.Debug(fmt.Sprintf("[%v] %v ", num, r))
		log.Debug(string(debug.Stack()))
		ctx.SetStatusCode(500)
	}
}
