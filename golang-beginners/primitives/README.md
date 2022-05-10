# Primitives

### Agenda
* Boolean type
* Numeric types
   * Integers
   * Floating point
   * Complex numbers
* Text types

## Boolean
It represents two states, you either have `true` or you have `false`.

```go
func main() {
   var n bool = true
   fmt.Println("%v, %T\n", n, n) // true, bool
}
```

There's a couple of uses for Boolean variables in our applications, perhaps one of the most common is to use them as state flags. The more common case for using boolean in Go is as a result of logical tests. Now, the `==` operatore is called *equals operator* and that's basically checking to see if the item on the left is equivalent to the item on the right.
```go
func main() {
   n := 1 == 1
   m := 1 == 2
   fmt.Println("%v, %T\n", n, n) // true, bool
   fmt.Println("%v, %T\n", m, m) // false, bool
}
```

The other thing that's important to know about the primitives is that every time you initialize a variable in Go, it actually has a zero value. The zero value for Boolean is false.
```go
func main() {
   var n bool
   fmt.Println("%v, %T\n", n, n) // false, bool
}
```

## Numeric types
The zero value for all numeric types is going to be zero (0) or the equivalent of zero for that numeric type.

### Integer
* Signed integers
   * `int`
      - integer of unspecified size
      - one thing that you're guaranteed is regardless of your environment, it will be at least 32 bits, but it could be 64 or even 128 bits depending on the system that you're running on
      - this will be the default integer type
      ```go
      func main() {
         n := 42
         fmt.Println("%v, %T\n", n, n) // 42, int
      }
      ```

      **Other types**:
      <table>
         <tbody>
            <tr>
               <td>int8</td>
               <td> -128 — 127 </td>
            </tr>
            <tr>
               <td>int16</td>
               <td> -32, 768 — 32, 767 </td>
            </tr>
            <tr>
               <td>int32</td>
               <td> -2, 147, 648 — 2, 147, 647 </td>
            </tr>
            <tr>
               <td>int64</td>
               <td> -9, 223, 372, 036, 854, 775, 808 — 9, 223, 372, 036, 854, 775, 807 </td>
            </tr>
         </tbody>
      </table>

* Unsigned integers
   * `uint`
      ```go
      func main() {
         var n uint16 = 42
         fmt.Println("%v, %T\n", n, n) // 42, uin16
      }
      ```

      **Other types**:
      <table>
         <tbody>
            <tr>
               <td>uint8</td>
               <td> 0 — 255 </td>
            </tr>
            <tr>
               <td>uint16</td>
               <td> 0 — 65, 535 </td>
            </tr>
            <tr>
               <td>uint32</td>
               <td> 0 — 4, 294, 295 </td>
            </tr>
         </tbody>
      </table>

* Byte
   * `byte`
      - an alias for an eight bit unsigned integer and the reason we have that is because the unsigned eight bit integer is very common, because that's what a lot of data streams are used to encode their data

* Arithmetic Operations
```go
func main() {
   a := 10
   b := 3
   fmt.Println(a + b) // Output: 13
   fmt.Println(a - b) // Output: 7
   fmt.Println(a * b) // Output: 30
   fmt.Println(a / b) // Output: 3; This is called integer division
   fmt.Println(a % b) // Output: 1; This returns the remainder

   var x int = 10
   var y int8  = 3
   fmt.Println(x + y) // You will get an error of `mismatched type int and int8`. To make this work, you need to do a type conversion `int(b)`
}
```

* Bit Operations
   * AND means 1 if the corresponding bits of two operands is 1. If eaither bit of an operand is 0, the result of the corresponding bit is evaluated to 0.
   * OR means 1 if at least one corresponding bit of two operands is 1
   * Exclusive OR means 1 if the corresponding bits of two operands are opposite
   * AND NOT means 1 if neither one of the numbers have the bit set
      ```go
      func main(){
         a := 10 // 1010
         b := 3  // 0011

         fmt.Println(a & b) // Output: 2 -> 0010
         fmt.Println(a | b) // Output: 11 -> 1011
         fmt.Println(a ^ b) // Output: 9 -> 1001
         fmt.Println(a &^ b) // Output: 8 -> 0100
      }
      ```
      > Detailed bitwise operation
      >```
      >   a = 10 = 1010
      >   b =  3 = 0011
      >   
      >   Bit & (AND) operation
      >      1010
      >   &  0011
      >   ________
      >      0010 = 2 in decimal
      >   
      >   Bit | (OR) operation
      >      1010
      >   |  0011
      >   ________
      >      1011 = 11 in decimal
      >   
      >   Bit ^ (Exclusive OR) Operation
      >      1010
      >   ^  0011
      >   ________
      >      1001
      >
      >   Bit &^ (AND NOT) Operation
      >      1010
      >   &^ 0011
      >   ________
      >      0100
      >```

* Bit shifting
The first statement means, it is going to bit shift a left three places. The second statement means, going a bit shift a right three places. When we do bit shifting, we're basically adding to that exponent, as long as we're dealing with the power of two.

```go
func main() {
   a := 8
   fmt.Println(a << 3) // Output: 63; (2^3) * (2^3) = 2^6 = 64
   fmt.Println(a >> 3) // Output: 1; (2^3) / (2^3) = 2^0 = 1
}
```
> Detailed bit shift operations
> ```
>    a = 8 = 1000
>
>    // Bit shift to the left
>    a << 3 = 0100 0000
>
>    // Bit shift to the right
>    a >> 3 = 0000 0001
> ```


### Float
In order to store decimal numbers in Go, we're going to use the floating point numbers. The floating point numbers in Go follow IEEE 754 Stanndard. Two types under IEEE 754 standard:
* float32
   * You can store ± 1.18E-38 — ± 3.4E38
* float64
   * You can store ± 2.23E-308 — ± 1.8E308

#### Defining floatinfg point literal
The `13.7e72` uses an exponential notation. So we can use 13.7 * 10^72 and that's going to use the short form IEEE 72nd to stand for that 10 to the 72nd. If you're going to use the initializer syntax on decimal, it's always going to be initialized to a `float64`. Keep in mind that you can't do arithmetic operations between float64 and float32.

```go
func main() {
   n := 3.14
   n = 13.7e72
   n = 2.1E14
   fmt.Printf("%v, %T", n, n) // 2.1e+14, float64
}
```