package _20

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedAlgorithm = Algorithm("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#")
var expectedImage = &Image{
	Image: map[Coordinates]bool{
		{0, 0}: true,
		{0, 1}: false,
		{0, 2}: false,
		{0, 3}: true,
		{0, 4}: false,

		{1, 0}: true,
		{1, 1}: false,
		{1, 2}: false,
		{1, 3}: false,
		{1, 4}: false,

		{2, 0}: true,
		{2, 1}: true,
		{2, 2}: false,
		{2, 3}: false,
		{2, 4}: true,

		{3, 0}: false,
		{3, 1}: false,
		{3, 2}: true,
		{3, 3}: false,
		{3, 4}: false,

		{4, 0}: false,
		{4, 1}: false,
		{4, 2}: true,
		{4, 3}: true,
		{4, 4}: true,
	},
	LitPixels: 10,
}

func TestParseInput(t *testing.T) {
	a, i, err := ParseInput("testdata/20.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expectedAlgorithm, a)
	assert.Equal(t, expectedImage, i)
}

func TestEnhanceImage(t *testing.T) {
	a, img, err := ParseInput("testdata/20.log")
	if err != nil {
		t.Fail()
	}

	EnhanceImage(2, img, a)

	assert.Equal(t, 35, img.LitPixels)

	EnhanceImage(48, img, a)

	assert.Equal(t, 3351, img.LitPixels)

	fmt.Println(img.String())
}
