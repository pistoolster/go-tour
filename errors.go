package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// `fmt` 包在输出时也会试图匹配 `error`，e变量通过实现Error()的接口函数成为了error类型，
	// 在fmt.Sprint(e)时就会调用e.Error()来输出错误的字符串信息
	if float64(e) < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", fmt.Sprint(float64(e)))
	}
	return fmt.Sprintf("cannot Sqrt imaginaries number: %v", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := x / 2
	for i := 0; ; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(i, z)
		if z*z-x < math.Pow(10, -15) {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
