package main

// 二分法
// 二分查找其实就是折半查找
// 前提条件是数组为有序数组

// 二分查找的思路
// 1、先确定中间位置k
// 2、将要查找的值与array[k]进行比较，若相等，则查找成功
// 如果array[k] < T，则应该向大的方向找 ,k=k+1
// 如果array[k] > T，则向小的方向查找,k=k-1
func main() {

}

func BinarySerach(array []int, key int) int {
	low := 0
	high := len(array) - 1

	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if array[mid] < key {
			low = mid + 1
		} else if array[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
