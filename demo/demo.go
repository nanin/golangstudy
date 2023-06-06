package demo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	model "test/demo/model"
)

func Cal() {
	// fmt.Printf("%c%c%c%c%c%c", 38472, 29642, 10084, 39532, 26480, 26143)
	fmt.Printf("%c%c%c%c\n", 10084, 39532, 26480, 26143)
	fmt.Printf("%v\n", string([]rune{10084, 39532, 26480, 26143}))
	// fmt.Printf("num=%v", strings.Count("seeeee", "ee"))
	// fmt.Println(strings.Trim("abaa1aba", "ab"))
}

func IntToString(n1 int) {
	var str string = fmt.Sprintf("%d", n1)
	fmt.Printf("n1 type is %T,value is %q\n", n1, n1)
	fmt.Printf("fmt.Sprintf(n1)后 type is %T,value is %q\n", str, str)
	var str2, _ = strconv.Atoi(str)
	fmt.Printf("strconv.Atoi(str)后 type is %T,value is %q\n", str2, str2)
}

func FordemoWithRune() {
	var str string = "hello北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) //会乱码，因为是按字符输出，但是中文占3个字符
	}
	fmt.Println()
	var str2 = []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i]) //
	}
	fmt.Println()
	for index, val := range str {
		fmt.Printf("第%d个字符是%c\n", index, val)
	}
}

func SwitchCaseDemo() {

	var b byte
	var c byte = 'c'
	fmt.Println("请输入字符：")
	// fmt.Scanln(&b)
	fmt.Scanf("%c", &b)
	fmt.Printf("type of b is %T ,value is%v \n", b, b)
	switch b {
	case 'a':
		fmt.Println("case 1")
	case 'b', c:
		fmt.Println("case 2")
	case 'c':
		fmt.Println("case 3")
	default:
		fmt.Println("case default")

	}

}

func Recursion(n int) {
	if n > 2 {
		n--
		Recursion(n)
	}
	fmt.Println("n=", n)
}

func SumALl(n1 int, args ...int) (sum int) {
	sum += n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return
}

func Swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func AddUpper(n1 int) func(int) int {
	var n int = n1
	return func(x int) int {
		n += x
		return n
	}
}

func Sub(n1 int, n2 int) int {
	return n1 - n2
}

func Makesuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

func PrintPyramid(level int, isNoHeart bool) {
	for i := 1; i <= level; i++ {
		//在打印* 前先打印空格
		for j := 1; j <= level-i; j++ {
			fmt.Print(" ")
		}

		//打印每一层的*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == level {
				fmt.Print("*")
			} else {
				if !isNoHeart {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}

			}
		}
		fmt.Println()
	}
}

func BubbleSort(arr *[]int) {
	fmt.Println("排序前：", *arr)

	arrLength := len(*arr)
	tmp := 0

	for i := 0; i < arrLength-1; i++ {

		for j := 0; j < arrLength-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
	fmt.Println("排序后：", *arr)
}

func BinarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int) {

	// fmt.Println(*arr)
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2

	if (*arr)[middleIndex] > findVal {
		newArr := (*arr)[:middleIndex]
		// fmt.Println("left", newArr)
		BinarySearch(&newArr, leftIndex, middleIndex-1, findVal)
	} else if (*arr)[middleIndex] < findVal {
		newArr := (*arr)[middleIndex+1:]
		// fmt.Println("right", newArr)
		BinarySearch(&newArr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到%v了，下标为%v \n", findVal, middleIndex)
	}

}

func SortSlice() {

	var students model.StudentSlice
	for i := 0; i < 10; i++ {
		s := model.Student{
			Name: fmt.Sprintf("student~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		students = append(students, s)
	}

	fmt.Println("------排序前-------")
	for _, v := range students {
		fmt.Println(v)
	}
	//此处因为StudentSlice实现了 sort.Interface的三个接口，
	//因此可以当做参数传给sort.Sort(data sort.Interface)
	sort.Sort(students)
	fmt.Println("------排序后-------")
	for _, v := range students {
		fmt.Println(v)
	}
}

//tag 可让序列化的json重命名
func JsonParse() {
	stu := model.Student{
		Name: "晓明",
		Age:  18,
	}

	data, _ := json.Marshal(stu)
	fmt.Printf("student序列化后=%v\n", string(data))

	str := "{\"nickname\":\"晓明\",\"age\":18}"
	var student model.Student
	//注意，这里需要引用传递，否则无法改变student的值
	erro := json.Unmarshal([]byte(str), &student)
	if erro != nil {
		fmt.Printf("unmashal error=%v", erro)
	} else {
		fmt.Printf("反序列化成功,student=%v\n", student)
	}
}
