package main

import (
	"version1_copy/model"
	"version1_copy/router"
	"version1_copy/utils"
)

func main() {
		utils.Init()
		model.InitDB()
		r := router.SetRouter()
		r.Run(utils.HttpPort)
}
//const n = 2
//
//func main() {
//
//	var dp = make([]int, n*5+1)
//	var res = make([]float64, 1)
//	num := math.Pow(5, n)
//
//	for j := 1; j <= 5; j++ {
//		dp[j] = 1
//	}
//
//	for i := 2; i <= n; i++ {
//		for s := i * 5; s >= i; s-- {
//			sum := 0
//			for t := 1; t <= 5; t++ {
//				if s-t >= i-1 {
//					sum += dp[s-t]
//				} else {
//					break
//				}
//				dp[s] = sum
//			}
//		}
//		for k := n; k <= n*5; k++ {
//			res = append(res, float64(float64(dp[k])/num))
//		}
//		res = res[1:]
//	}
//	fmt.Println(res)
	////
	//type sttr struct {
	//	stri string
	//	len  int
	//}
	//var sttrs []sttr
	//s := "Welcome to join the HundSun family"
	//splits := strings.Fields(s)
	//var strs = make(map[string]int, 1)
	//for _, split := range splits {
	//	strs[split] = len(split)
	//}
	//for k, v := range strs {
	//	sttrs = append(sttrs, sttr{k, v})
	//}
	//sort.Slice(sttrs, func(i, j int) bool {
	//	return sttrs[i].len < sttrs[j].len
	//})
	//fmt.Println(sttrs)

	//fmt.Println(strs)
//}
