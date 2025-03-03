package main

import "fmt"

var count int = 0;

func main()  {

	//Explicit Declaration การประกาศแบบ type ชัดเจน ไม่สามารถเปลี่ยนtype ทีหลังได้
	fmt.Println("begin");
	var tmp1 int = 0;
	tmp1 = 10;
	var tmp2 string = "สวัสดี";
	var tmp3 bool = true;

	// const tmp4 int = 0;
	// tmp4 = 5;
 
	// implicit Declaration ประกาศตัวแปรแบบ ไม่ระบุType
	tmp5 := 0;
	tmp6 := "test"

	fmt.Println(tmp1);
	fmt.Println(tmp2);
	fmt.Println(tmp3);
	fmt.Println(tmp5);
	fmt.Println(tmp6);
	// run(1);
	tmp5 = run(1);
	fmt.Println("last tmp 5",tmp5);
	

}

func run(count2 int) (result int)   {
	count++;
	fmt.Println(count);
	count2 += count2;
	count2++;
	return count2;
}