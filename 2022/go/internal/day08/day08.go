package day08

func PartOne(filename string) int {
	heightMap := readTreeHeights(filename)

	visibilityMap := heightMap.newVisibleTreesFromTreeHeights()

	return visibilityMap.numberOfVisibleTrees
}

func PartTwo(filename string) int {
	heightMap := readTreeHeights(filename)

	scenicScoreMap := heightMap.newScenicScoresFromTreeHeights()

	return scenicScoreMap.bestScenicScore
}
