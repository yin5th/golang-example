package main

import "fmt"

func testScanf() {
	var (
		a int
		b string
		c float32

		d int
		e string
		f float32
	)

	//在一行内 用空格作为分隔符输入
	fmt.Scanf("%d %s %f", &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)

	//多行输入
	fmt.Scanf("%d\n", &d)
	fmt.Scanf("%s\n", &e)
	fmt.Scanf("%f\n", &f)
	fmt.Printf("d=%d e=%s f=%f\n", d, e, f)

	//结果
	var result = `
		123 你好 334.2
		a=123 b=你好 c=334.200012
		34
		猪是你
		83
		d=34 e=猪是你 f=83.000000
`
	fmt.Printf("结果：%s", result)
}

func testScan() {
	var (
		a int
		b string
		c float32
	)

	//多行输入 以换行为分割
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("多行结果：a=%d b=%s c=%f\n", a, b, c)
}

func testScanln() {
	var a int

	fmt.Scanln(&a)
	fmt.Printf("a=%d\n", a)
}

func main() {
	//testScanf()
	//testScan()
	testScanln()
}
