package sum

import "testing"

type sumTests []sumTest{
         sumTest{1, 1, 2}
         sumTest{500 ,400, 900}
         sumTest{56, 44, 100}
}
func TestSum(t *testing.T){
        for _, dt := range sumTests{
                v := sum(dt.in1, dt.in2)
                if v != dt.out{
                        t.errorf("sum(%d, %d) = %d, want %d.", dt.in1, dt.in2, v, dt.out)
                }
        }
}  
