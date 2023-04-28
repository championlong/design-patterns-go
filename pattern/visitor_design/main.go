package main

/*
访问者模式：
	允许⼀个或者多个操作应⽤到⼀组对象上，解耦操作和对象本⾝。
访问者模式针对的是⼀组类型不同的对象（PdfFile、PPTFile、WordFile）。不过，尽管这组对 象的类型是不同的，但是，它们继承相同的⽗类（ResourceFile）或者实现相同的接⼝。
在不同的应⽤场景 下，我们需要对这组对象进⾏⼀系列不相关的业务操作（抽取⽂本、压缩等），但为了避免不断添加功能导 致类（PdfFile、PPTFile、WordFile）不断膨胀，
职责越来越不单⼀，以及避免频繁地添加功能导致的频繁 代码修改，我们使⽤访问者模式，将对象与操作解耦，将这些业务操作抽离出来，定义在独⽴细分的访问者 类（Extractor、Compressor）中。

难理解、难实现，应⽤它会导致代码的可读性、可维护性变差，在没有特别必要的情况下，建议你不要使 ⽤访问者模式。
*/
func main() {

}
