package sortd

 import "testing"
 //benchmark function looks like a test function, 
 //but with the Benchmark preﬁx and a *testing
 //B parameter that provides most of the 
 // same methods as a *testing.T, 
 //plus a few extra related to performance measurement. 
 //N speciﬁes the number of times toper 
 // form the operation being measured.
 func BenchmarkBBSorting(b *testing.B) {
     for i := 0; i < b.N; i++ {
     	  var items = []int{43,23,76,198,32}	
          BubbleSort(items) 
        }
     }

 func BenchmarkISorting(b *testing.B) {
     for i := 0; i < b.N; i++ {
     	  var items = []int{43,23,76,198,32}	
          InsertionSort(items) 
        }
     }    
			