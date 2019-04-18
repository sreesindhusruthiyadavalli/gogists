package main

import

"uilive"

writer := uilive.New()
// start listening for updates and render
writer.Start()

for i := 0; i <= 100; i++ {
  fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
  time.Sleep(time.Millisecond * 5)
}

fmt.Fprintln(writer, "Finished: Downloaded 100GB")
writer.Stop() // flush and stop rendering