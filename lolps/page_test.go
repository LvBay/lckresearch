package lolps

import (
	"log"
	"testing"
)

func Test_getPageInfo(t *testing.T) {
	p := getPageInfo(100, 105, 100)
	log.Println(ToString(p))
}
