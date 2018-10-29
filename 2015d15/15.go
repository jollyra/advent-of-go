package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

type score struct {
	score, calories int
}

func inputLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func fromLine(line []string) (ingredient, error) {
	capacity, err := strconv.Atoi(line[1])
	if err != nil {
		return ingredient{}, err
	}
	durability, err := strconv.Atoi(line[2])
	if err != nil {
		return ingredient{}, err
	}
	flavour, err := strconv.Atoi(line[3])
	if err != nil {
		return ingredient{}, err
	}
	texture, err := strconv.Atoi(line[4])
	if err != nil {
		return ingredient{}, err
	}
	calories, err := strconv.Atoi(line[5])
	if err != nil {
		return ingredient{}, err
	}

	return ingredient{
		name:       line[0],
		capacity:   capacity,
		durability: durability,
		flavour:    flavour,
		texture:    texture,
		calories:   calories,
	}, nil
}

func addIngredient(recipe *[5]int, ing ingredient, tsps int) {
	recipe[0] += tsps * ing.capacity
	recipe[1] += tsps * ing.durability
	recipe[2] += tsps * ing.flavour
	recipe[3] += tsps * ing.texture
	recipe[4] += tsps * ing.calories
}

func scoreRecipe(recipe [5]int) (int, int) {
	if recipe[0] <= 0 || recipe[1] <= 0 || recipe[2] <= 0 || recipe[3] <= 0 {
		return 0, recipe[4]
	}

	score := recipe[0] * recipe[1] * recipe[2] * recipe[3]
	calories := recipe[4]
	return score, calories
}

func max(scores []score) score {
	max := scores[0]
	for _, score := range scores {
		if score.score >= max.score {
			max = score
		}
	}
	return max
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("A single filepath arg is required")
	}

	lines := inputLines(os.Args[1])
	ingredients := make([]ingredient, 0)
	for _, line := range lines {
		words := strings.Split(line, ",")
		ing, err := fromLine(words)
		if err != nil {
			log.Fatalf("Unable to read line:\n%v\n%v", line, err)
		}
		ingredients = append(ingredients, ing)
	}

	maxTsps := 100
	calorieRestriction := 500
	scores := make([]score, maxTsps^2)
	for i := 0; i <= maxTsps; i++ {
		for j := 0; j <= maxTsps-i; j++ {
			for k := 0; k <= maxTsps-i-j; k++ {
				recipe := [5]int{}
				addIngredient(&recipe, ingredients[0], i)
				addIngredient(&recipe, ingredients[1], j)
				addIngredient(&recipe, ingredients[2], k)
				addIngredient(&recipe, ingredients[3], maxTsps-i-j-k)
				if recipe[4] == calorieRestriction {
					s, cals := scoreRecipe(recipe)
					scores = append(scores, score{s, cals})
				}
			}
		}
	}

	max := max(scores)
	fmt.Println(max)
}
