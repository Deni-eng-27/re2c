// Code generated by re2c, DO NOT EDIT.
//line "golang/005_posix_captures.--lang(go).re":1
package main

import "reflect"

var YYMAXNMATCH int = 5


func Lex(str string) (int, []int) {
	var cursor, marker, yynmatch int
	yypmatch := make([]int, YYMAXNMATCH*2)

	var yyt1 int
	var yyt10 int
	var yyt2 int
	var yyt3 int
	var yyt4 int
	var yyt5 int
	var yyt6 int
	var yyt7 int
	var yyt8 int
	var yyt9 int

	
//line "golang/005_posix_captures.--lang(go).go":27
{
	var yych byte
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt1 = cursor
		yyt2 = -1
		yyt3 = -1
		yyt5 = -1
		yyt6 = -1
		yyt7 = -1
		yyt8 = -1
		yyt9 = -1
		yyt10 = -1
		goto yy2
	case 'a':
		yyt1 = cursor
		yyt2 = cursor
		yyt5 = cursor
		yyt6 = cursor
		goto yy6
	case 'b':
		yyt1 = cursor
		yyt2 = cursor
		yyt6 = -1
		yyt7 = -1
		yyt8 = cursor
		goto yy7
	case 'c':
		yyt1 = cursor
		yyt2 = cursor
		yyt6 = -1
		yyt7 = -1
		yyt8 = -1
		yyt9 = -1
		yyt10 = cursor
		goto yy8
	default:
		goto yy4
	}
yy2:
	cursor += 1
	yyt4 = cursor
	yynmatch = 5
	yypmatch[0] = yyt1
	yypmatch[1] = yyt4
	yypmatch[2] = yyt5
	yypmatch[3] = yyt2
	yypmatch[4] = yyt6
	yypmatch[5] = yyt7
	yypmatch[6] = yyt8
	yypmatch[7] = yyt9
	yypmatch[8] = yyt10
	yypmatch[9] = yyt3
//line "golang/005_posix_captures.--lang(go).re":41
	{
		return yynmatch, yypmatch
	}
//line "golang/005_posix_captures.--lang(go).go":86
yy4:
	cursor += 1
yy5:
//line "golang/005_posix_captures.--lang(go).re":29
	{
		return -1, nil
	}
//line "golang/005_posix_captures.--lang(go).go":94
yy6:
	cursor += 1
	marker = cursor
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = -1
		yyt7 = cursor
		yyt8 = -1
		yyt9 = -1
		yyt10 = -1
		goto yy9
	case 'a':
		goto yy11
	case 'b':
		yyt5 = cursor
		yyt6 = cursor
		goto yy14
	case 'c':
		yyt7 = cursor
		yyt8 = -1
		yyt9 = -1
		yyt10 = cursor
		goto yy15
	default:
		goto yy5
	}
yy7:
	cursor += 1
	marker = cursor
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = -1
		yyt9 = cursor
		yyt10 = -1
		goto yy9
	case 'b':
		goto yy16
	default:
		goto yy5
	}
yy8:
	cursor += 1
	marker = cursor
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = cursor
		goto yy9
	case 'c':
		goto yy18
	default:
		goto yy5
	}
yy9:
	cursor += 1
	yyt4 = cursor
	yynmatch = 5
	yypmatch[0] = yyt1
	yypmatch[1] = yyt4
	yypmatch[2] = yyt2
	yypmatch[3] = yyt3
	yypmatch[4] = yyt6
	yypmatch[5] = yyt7
	yypmatch[6] = yyt8
	yypmatch[7] = yyt9
	yypmatch[8] = yyt10
	yypmatch[9] = yyt5
//line "golang/005_posix_captures.--lang(go).re":37
	{
		return yynmatch, yypmatch
	}
//line "golang/005_posix_captures.--lang(go).go":171
yy11:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = -1
		yyt7 = cursor
		yyt8 = -1
		yyt9 = -1
		yyt10 = -1
		goto yy9
	case 'a':
		goto yy11
	case 'b':
		yyt5 = cursor
		yyt6 = cursor
		goto yy20
	default:
		goto yy13
	}
yy13:
	cursor = marker
	goto yy5
yy14:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'c':
		yyt7 = cursor
		yyt8 = cursor
		goto yy21
	default:
		goto yy13
	}
yy15:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt2 = cursor
		yyt3 = cursor
		goto yy2
	case 'a':
		yyt5 = cursor
		yyt6 = cursor
		goto yy22
	default:
		goto yy13
	}
yy16:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = -1
		yyt9 = cursor
		yyt10 = -1
		goto yy9
	case 'b':
		goto yy16
	default:
		goto yy13
	}
yy18:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		yyt5 = cursor
		goto yy9
	case 'c':
		goto yy18
	default:
		goto yy13
	}
yy20:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'c':
		yyt7 = cursor
		yyt8 = cursor
		goto yy23
	default:
		goto yy13
	}
yy21:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'a':
		yyt5 = cursor
		yyt6 = cursor
		goto yy22
	default:
		goto yy24
	}
yy22:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'b':
		yyt7 = cursor
		yyt8 = cursor
		goto yy27
	case 'c':
		yyt7 = cursor
		yyt8 = -1
		yyt9 = -1
		yyt10 = cursor
		goto yy15
	default:
		goto yy13
	}
yy23:
	cursor += 1
	yych = str[cursor]
yy24:
	switch (yych) {
	case 0x00:
		yyt3 = cursor
		goto yy25
	case 'c':
		goto yy23
	default:
		goto yy13
	}
yy25:
	cursor += 1
	yyt9 = cursor
	yynmatch = 4
	yypmatch[0] = yyt1
	yypmatch[1] = yyt9
	yypmatch[2] = yyt2
	yypmatch[3] = yyt5
	yypmatch[4] = yyt6
	yypmatch[5] = yyt7
	yypmatch[6] = yyt8
	yypmatch[7] = yyt3
//line "golang/005_posix_captures.--lang(go).re":33
	{
		return yynmatch, yypmatch
	}
//line "golang/005_posix_captures.--lang(go).go":318
yy27:
	cursor += 1
	yych = str[cursor]
	switch (yych) {
	case 'c':
		yyt9 = cursor
		yyt10 = cursor
		goto yy15
	default:
		goto yy13
	}
}
//line "golang/005_posix_captures.--lang(go).re":44

}

func main() {
	var n int
	var m []int

	n, m = Lex("aabcc\000")
	if !(n == 4 && reflect.DeepEqual(m[0:2*n], []int{0, 6, 0, 2, 2, 3, 3, 5})) {
		panic("failed on aabcc")
	}
	n, m = Lex("aaa\000")
	if !(n == 5 && reflect.DeepEqual(m[0:2*n], []int{0, 4, 0, 3, 0, 3, -1, -1, -1, -1})) {
		panic("failed on aaa")
	}
	n, m = Lex("acabc\000")
	if !(n == 5 && reflect.DeepEqual(m[0:2*n], []int{0, 6, 2, 5, 2, 3, 3, 4, 4, 5})) {
		panic("failed on acabc")
	}
	n, m = Lex("abcac\000")
	if !(n == 5 && reflect.DeepEqual(m[0:2*n], []int{0, 6, 3, 5, 3, 4, -1, -1, 4, 5})) {
		panic("failed on abcac")
	}
	n, m = Lex("ab\000")
	if !(n == -1 && m == nil) {
		panic("failed on ab")
	}
}