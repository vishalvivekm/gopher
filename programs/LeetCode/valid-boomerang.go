func isBoomerang(points [][]int) bool {
    
    var slope1 float64
        
        if isSlicesEqual(points[0],points[1]) || isSlicesEqual(points[1],points[2]) || isSlicesEqual(points[2], points[0]) {
            return false
        }
        if (points[1][0] - points[0][0]) == 0 {
            // line is parallel to y axis
            // check if third point's x coordinate is the same as that of 1st point
            if points[2][0] == points[0][0] {
                return false
            } else {
                return true
            }
        }
        slope1 = (float64(points[1][1]) - float64(points[0][1])) / (float64(points[1][0]) - float64(points[0][0]))
        // y = mx + c => c = y - mx
        fmt.Println(slope1)
        c := float64(points[1][1]) - slope1 * float64(points[1][0])

        c1 := float64(points[2][1]) - slope1 * float64(points[2][0])
        fmt.Println(c, c1)
        if fmt.Sprintf("%.4f", c1) == fmt.Sprintf("%.4f", c) { // test Case: // 85.16666666666667 85.16666666666666 // [[19,82],[73,73],[97,69]]
            return false
        }
        return true
}

func isSlicesEqual(a, b []int) bool {
    // if len(a) != len(b) {
    //     return false
    // }
    
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    
    return true
}
