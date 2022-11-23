package main

import "fmt"

func main(){
	var a []int
	fmt.Println(a,len(a),cap(a))
	a = []int{1,2,3}
	fmt.Println(a,len(a),cap(a))
	fmt.Println(a)
	

	b := make([]int,0)
	fmt.Println(b)
	b = []int{1,2,3}
	fmt.Println(b)
	
	c:= make([]int,10)
	fmt.Println(c)
	c = []int{1,2,3}
	fmt.Println(c)
	
/*
	var da = [3]int{1,2,3}
	var ds = da[:]
	var ds1 = da[:]
	fmt.Println(ds,da,ds1,len(ds),cap(ds))
	ds[2]=11
	fmt.Println(ds,da,ds1)
	fmt.Println(ds,da,ds1,len(ds),cap(ds))
	ds1[2]=122
	fmt.Println(ds,da,ds1)
	fmt.Println(ds,da,ds1,len(ds),cap(ds))
*/
	var da = [5]int{1,2,3,4,5}
	var ds = da[:3]
	var ds1 = da[:2]
	fmt.Println(ds,da,ds1,len(ds),cap(ds))
	ds[2]=11
	fmt.Println(ds,da,ds1)
	fmt.Println(ds,da,ds1,len(ds),cap(ds))


	var es = []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(es,len(es),cap(es))
	var es1 = es[2:5]
	fmt.Println(es1,len(es1),cap(es1))
	var es2 = es1[:7]
	fmt.Println(es2,len(es2),cap(es2))

	fmt.Println("COPY")
	var fs = []int{1,2,3,4,5}
	var fs2 = make([]int,len(fs))
	fmt.Println(fs,fs2)
	copy(fs2,fs)
	fmt.Println(fs2,fs)
	fs2[1]=-1
	fmt.Println(fs2,fs)

	fmt.Println("APPEND")
	var gs = []int{1,2,3}
	fmt.Println(gs)
//	append(gs,99)
//	fmt.Println(gs)

	var gs1 = []int{1,2,3}
	var gs2 = []int{4,5,6}
	fmt.Println(gs1,gs2)
	gs1 = append(gs1,gs2...)
	fmt.Println(gs1,gs2)
	gs2[1]=11
	fmt.Println(gs1,gs2)
	arr := [7]string{"1", "2", "3", "4", "5", "6", "7"}

slcVar := arr[1:6]
fmt.Println(slcVar)
}
