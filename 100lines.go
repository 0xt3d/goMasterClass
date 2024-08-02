var (
	a bool
	b float32
	d string
)

 /* User specified types */
 const a int32 = 12         // 32-bit integer
 const b float32 = 20.5      // 32-bit float
 var c complex128 = 1 + 4i  // 128-bit complex number
 var d uint16 = 14          // 16-bit unsigned integer

 /* Default types */
 n := 42              // int
 pi := 3.14           // float64
 x, y := true, false  // bool
 z := "Go is awesome" // string
