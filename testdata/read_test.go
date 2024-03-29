// This is a Golang source file used only for testing the imports package.

package testdata

import "github.com/user/pkg"
import p "github.com/user2/pkg2"
import (
    "github.com/user3/pkg3"
    q "github.com/user4/pkg4"

    // Comment
    "fmt"
    i "io"
    "runtime"
)

// Test comment ignore
// import "fakeimport"
// import (
//   "anotherfake"
//   x "andanother"
// )

// Another comment
import "os"

// Test C style comment ignore
/* import "nope" */
/*
import "nopeagain"
*/

func aFunction() bool {

    return false
}

