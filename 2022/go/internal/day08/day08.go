package day08

func PartOne(lines []string) int {
	heightMap := readTreeHeights(lines)

	visibilityMap := heightMap.newVisibleTreesFromTreeHeights()

	return visibilityMap.numberOfVisibleTrees
}

func PartTwo(lines []string) int {
	heightMap := readTreeHeights(lines)

	scenicScoreMap := heightMap.newScenicScoresFromTreeHeights()

	return scenicScoreMap.bestScenicScore
}
