func numberOfSteps(num int) int {
    i := 0

    for {
            if num == 0 {
                break
    }
        if num % 2 == 0 {
       
        num = num / 2
         i++

    } else {
      num = num -1
      i++
    }
        
    }

return i
}


