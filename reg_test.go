package main

import (
	"fmt"
	"regexp"
	st "strings"
	"testing"
)

const regStr = `查询: 我爱你

拼音: wǒ ài nǐ

Exps:

 1.I love you


网络释义:

   1.我爱你:
    I love you, ich liebe dich, Wuh that I love you

   2.因为我爱你:
    Because I love you, Because you loved me, I love you because, OT Valentine

   3.命中注定我爱你:
    Fated to Love You, Fated to Lov, Destiny Love, DJ White Remix




`
func TestReg(t *testing.T) {
	 reg:=regexp.MustCompile(`Exps:[\s\S]*1.([\s\S]*)网络释义:`)
	 ret := reg.FindAllStringSubmatch(regStr,-1)
	for _, strings := range ret {
		for _, s := range strings {
			fmt.Println("-------",st.TrimSpace(s))
		}
	}
}
func TestReplace(t *testing.T) {
	s := `I_love_you\u001b[38;5;167m`
	fmt.Println(s)
	fmt.Println(st.Replace(s,`\u001b[38;5;167m`,"",-1))
}