package main

import (
	"fmt"
	"math"
	"strconv"
)

//给定一个 32 位有符号整数，将整数中的数字进行反转。

//示例 1:
//输入: 123
//输出: 321

// 示例 2:
//输入: -123
//输出: -321

//示例 3:
//输入: 120
//输出: 21

//注意:
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。
//根据这个假设，如果反转后的整数溢出，则返回 0。

func main() {
	var input int = 8464832147
	var output int
	output = reverse(input)
	fmt.Println(output)
}

func reverse(x int) int {
	var temp int
	var instr string
	var outstr string
	var output int64
	if x > 0 {
		temp = x
	} else {
		temp = -x
	}
	instr = strconv.Itoa(temp)
	for i := len(instr) - 1; i >= 0; i-- {
		outstr += string([]byte(instr)[i])
	}
	output32, err := strconv.Atoi(outstr)
	if err != nil {
		fmt.Println(err)
	}
	output = int64(output32)
	if x > 0 {
		output = output
		if output > (int64(math.Pow(2, 31)) - 1) {
			output = 0
		}
	} else {
		output = -output
		if output < int64(-math.Pow(2, 31)) {
			output = 0
		}
	}

	return int(output)
}
