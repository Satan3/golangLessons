package ourPackageUnderGit

const (
	DaysInWeekCount = 7
)

type Name string

func Sum(x, y int64) int64 {
	return x + y
}

func getWeeksCount(daysCount int) int {
	return daysCount / DaysInWeekCount
}
