/*
포인터를 인자로 받는 함수는 값(value)을 전달받을 수 없고, 값(value)을 인자로 받는 함수는 포인터를 전달받을 수 없음

var v Vertex
ScaleFunc(v)  // 컴파일 에러
ScaleFunc(&v) // 정상

반면에 포인터 리시버는 값(value)과 포인터 모두 접근 가능

var v Vertex
v.Scale(5)  // 값으로도 접근할 수 있음
p := &v
p.Scale(10) // 포인터로도 접근할 수 있음

포인터 리시버를 쓰는 이유는 다음과 같다.
1. 메소드 내부에서 리시버가 가리키는 값을 바꾸고 싶은 경우
2. 메소드가 호출될 때, 값이 복사(copy)되는 것을 피하고 싶은 경우

*주의* value리시버나 포인터 리시버 모두 자유롭게 쓸 수 있지만, 보통 한 메소드 안에서 둘을 한꺼번에 쓰지는 않는다.
*/

package main 

import "fmt"

type Vertex struct {
    X, Y float64
}

// ① Vertex 포인터 리시버가 있는 메소드 임. Vertex 혹은 Vertex 포인터로 접근할 수 있음
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// ② Vertex 포인터 인자가 있는 함수임
// Vertex 포인터만 인자로 들어올 수 있음
func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2)
    ScaleFunc(&v, 10)
    fmt.Println("① Vertex{3, 4}로 접근했을 때 :", v)
    
    p := &Vertex{3, 4}
    p.Scale(2)
    ScaleFunc(p, 10)
    fmt.Println("② &Vertex{3, 4}로 접근했을 때 :", p)
}
