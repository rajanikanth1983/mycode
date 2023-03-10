/* Alta3 Research | RZFeeser
   Variables */

// In Go, variables are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls

package main

import "fmt"

func main() {

    // variables declared inside of main() are at the "function" level
    // var name type = expression
    var instructor string = "rzfeeser"
    fmt.Println(instructor)
    
    // `var` declares 1 or more variables
    // here we let the compiler infer the type
    var a = "Goodbye"
    fmt.Println(a)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "apple"` in this case.
    f := "apple"
    fmt.Println(f)
}

