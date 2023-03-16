package paint_test

import (
	"paint"
	"testing"
)

func TestLog(t *testing.T) {
	t.Log(paint.G("[debug]"), "rainbows are  so beautiful !")
	t.Log(paint.Concat(paint.R("rainbows"), paint.G("are"), paint.Y("so"), paint.B("beautiful"), paint.M("!")))
	t.Log(paint.Y("[warn]"), "quick brown fox jumps over the lazy dog")
	t.Log(paint.Concat(paint.R("quick"), paint.G("brown"), paint.Y("fox"), paint.B("jumps"), paint.M("over"), paint.C("the"), paint.W("lazy"), paint.B("dog")))
	t.Log(paint.R("[error]"), "to be or not to be, that is the question")
	t.Log(paint.Concat(paint.R("to"), paint.G("be"), paint.Y("or"), paint.B("not"), paint.M("to"), paint.C("be,"), paint.W("that"), paint.B("is"), paint.R("the"), paint.G("question")))
	t.Log(paint.G("[info]"), "you can either choose to use font standard, bold, light, italic, or underline")
	t.Log(paint.Concat(paint.R("you"), paint.G("can"), paint.Y("either"), paint.B("choose"), paint.M("to"), paint.C("use"), paint.B("font"), paint.HY("bold"), paint.LR("light"), paint.IG("italic"), paint.M("or"), paint.UC("underline")))
}
