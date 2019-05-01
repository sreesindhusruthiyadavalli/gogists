package main

import("fmt")

func GenDisplaceFn(a,v,d float64) func (float64) float64 {
	fn := func (t float64) float64 {
		s := ((1/2)*a*t*t) + v*t + d
		return s
	}
	return fn
}

func main(){
	//reading an integer
    var a, vel, disp, time float64
    //, 
    fmt.Println("Enter acceleration velocity and displacement")
    //_, err := 
    fmt.Scan(&a)
    fmt.Scan(&vel)
    fmt.Scan(&disp)

    fmt.Println("Your entered accelertion as ", a, "velocity is ", vel,
     "displacement as", disp)

    fmt.Println("Enter time")
    fmt.Scan(&time)

    fmt.Println("Time entered is", time)

    f1 := GenDisplaceFn(a, vel,disp)
    fmt.Println(f1(time))

}

