package main

import "fmt"

/*
*
接口是一种抽象类型，是一组方法的集合，里面只声明方法，而没有任何数据成员
在GO语言种实现一个接口也不需要显示的声明，只需要其他类型实现了接口中所有方法
在GO语言中，接口方法名的大小写区别主要体现在可见性和访问权限上：
首字母大写的接口方法‌
方法名首字母大写表示该方法可被其他包访问（导出），属于公开方法。例如Read()或Write()，其他包导入该接口后可直接调用这些方法
首字母小写的接口方法‌
方法名首字母小写表示仅限当前包内访问（私有），外部包无法调用。例如privateMethod()，即使实现了该接口的结构体被导出，外部包也无法通过接口调用此方法
对JSON序列化的影响‌
若接口方法用于结构体字段的JSON序列化，只有首字母大写的字段名会被编码到JSON中，小写字段会被忽略
设计哲学体现‌
GO语言通过大小写简化访问控制，替代了其他语言中的public/private等关键字，强调简洁性
*/

type PayMathod interface {
	Account
	Pay(account int) bool
}

type Account interface {
	GetBalance() int
}

// CreditCard 信用卡
type CreditCard struct {
	balance int
	limit   int
}

// Pay 支付方法
func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Println("信用卡支付成功")
		return true
	}
	fmt.Println("信用卡支付失败，超出额度")
	return false
}

// GetBalance 获取余额
func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 借记卡
type DebitCard struct {
	balance int
}

// Pay 借记卡支付
func (c *DebitCard) Pay(amount int) bool {
	if c.balance >= amount {
		c.balance -= amount
		fmt.Println("借记卡支付成功")
		return true
	}
	fmt.Println("借记卡支付余额不足，支付失败")
	return false
}

// GetBalance 借记卡获取余额的方法
func (c *DebitCard) GetBalance() int {
	return c.balance
}

// purchaseItem 综合购买方法
func purchaseItem(p PayMathod, price int) {
	if p.Pay(price) {
		fmt.Println("购买成功，剩余余额", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

// 参数interface{}实际上是Object类型
func anyPay(param interface{}) {
	fmt.Println("param", param)
}

func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}
	purchaseItem(creditCard, 800)
	purchaseItem(debitCard, 300)
	//类型转化
	var account Account = debitCard
	fmt.Println(account.GetBalance())
	//类型转化
	var account2 Account = creditCard
	fmt.Println(account2.GetBalance())
	//interface{}参数实际上就是Object
	anyPay(account2)
	anyPay(1)
	anyPay("string")
}
