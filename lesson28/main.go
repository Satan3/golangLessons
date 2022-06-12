package main

import (
	"fmt"
)

func main() {
	/*group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		fmt.Println("first", ctx)
		return nil
	})

	group.Go(func() error {
		return fmt.Errorf("some error")
	})

	if err := group.Wait(); err != nil {
		fmt.Println("error", err)
	}*/

	fmt.Println("Without dependencies")
}
