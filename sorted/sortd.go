package sortd

func BubbleSort( items []int)  {

   L:=len(items)

   for  i:=0;i<L;i++{

      for j:=0;j< (L-1-i);j++{
         if items[j] > items[j+1]{
            items[j], items[j+1] = items[j+1], items[j]
         }
      }
   }

}


func InsertionSort(items []int) {

   L := len(items)

   for i := 1; i < L; i++ {

      j := i
      for j > 0 && items[j] < items[j-1] {
         items[j], items[j-1] = items[j-1], items[j]
         j -= 1
      }

   }
}   
