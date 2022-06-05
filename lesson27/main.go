package main

import (
	"context"
	"fmt"
	"ourPackage"

	"golang.org/x/sync/errgroup"

	"golangLessons/lesson27/ourPackageUnderGit"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	fmt.Println(g, ctx)

	print(ourPackage.Sum(1, 3))

	print(ourPackageUnderGit.DaysInWeekCount)
}
