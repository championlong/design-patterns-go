package main

/*
策略模式
定义⼀族算法类，将每个算法分别封装起来，让它们可以互相替换。解耦的是策略的定义、创建、使⽤这三部分
策略类的定义⽐较简单，包含⼀个策略接⼝和⼀组实现这个接⼝的策略类。
策略的创建根据type创建策略的逻辑抽离出来， 放到⼯⼚类中。(对于无状态的可直接初始化放入map缓存使用，对于有状态的需要根据类型实时创建)
策略的使⽤ 是在程序运⾏期间，根据配置、⽤⼾ 输⼊、计算结果等这些不确定因素，动态决定使⽤哪种策略
可以避免if-else分⽀
 */
func main() {

}
