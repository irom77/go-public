package remote
import "fmt"

// Args is a data structure for the incoming arguments
// This needs to be exported for the RPC to be valid/work
type Args struct {
	A, B int
}

// Arith is our functions return type
// This also needs to be exported
type Arith int

// Multiply does simply multiplication on provided arguments
// This also needs to be exported
func (t *Arith) Multiply(args *Args, reply *int) error {
	fmt.Printf("Args received: %+v\n", args)
	*reply = args.A * args.B
	return nil
}
