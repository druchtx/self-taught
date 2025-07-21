package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

// 使用bitmap统计活跃用户和付费用户

const userCount = 100_000

func main() {
	// 每个用户一个bit，100000个用户需要 100000/64 = 1563个uint64
	n := (userCount + 63) / 64
	active := make([]uint64, n)
	paid := make([]uint64, n)

	// 随机生成活跃用户和付费用户
	rand.NewSource(time.Now().UnixNano())
	for i := range userCount {
		if rand.Intn(2) == 1 {
			active[i/64] |= 1 << (i % 64)
		}
		if rand.Intn(3) == 1 {
			paid[i/64] |= 1 << (i % 64)
		}
	}

	// 交集：既活跃又付费
	both := make([]uint64, n)
	for i := range n {
		both[i] = active[i] & paid[i]
	}

	// 并集：活跃或付费
	either := make([]uint64, n)
	for i := range n {
		either[i] = active[i] | paid[i]
	}

	// 统计交集和并集人数
	// fmt.Println("活跃用户数：", countBits(active))
	// fmt.Println("付费用户数：", countBits(paid))
	// fmt.Println("既活跃又付费用户数：", countBits(both))
	// fmt.Println("活跃或付费用户数：", countBits(either))

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', tabwriter.DiscardEmptyColumns)
	const format = "%s\t%v\t%v\t%v\t%v\t\n"
	_, _ = fmt.Fprintf(tw, format, "中文", "users", "active", "paid", "both")
	_, _ = fmt.Fprintf(tw, format, "这个是中文", countBits(active), countBits(paid), countBits(both), countBits(either))
	_ = tw.Flush()
}

// 统计bitmap中1的个数
func countBits(bits []uint64) int {
	count := 0
	for _, word := range bits {
		for word != 0 {
			count++
			// Brian Kernighan 算法
			word &= word - 1 // 清除最低位的1
		}
	}
	return count
}
