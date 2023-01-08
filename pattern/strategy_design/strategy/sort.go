package strategy

import "fmt"

type ISortAlg interface {
	sort(filePath string)
}

// QuickSort 快速排序
type QuickSort struct {
}

func (self QuickSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_QUICK_SORT)
}
 // ExternalSort 外部排序
type ExternalSort struct {
}

func (self ExternalSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_EXTERNAl_SORT)
}

// ConcurrentExternalSort 多线程外部排序
type ConcurrentExternalSort struct {
}

func (self ConcurrentExternalSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_CONCURRENT_SORT)
}

// MapReduceSort 利⽤MapReduce多机排序
type MapReduceSort struct {
}

func (self MapReduceSort) sort(filePath string) {
	fmt.Println(ALGS_KEY_MAP_REDUCE_SORT)
}
