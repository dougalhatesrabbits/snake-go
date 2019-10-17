package trisnake

import (
	tl "github.com/JoelOtter/termloop"
)

// NewSnake will create a new snake and is called when the game is initialized.
func NewSnake() *Snake {
	snake := new(Snake)
	// Create a new entity for a 1x1 pixel.
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	// Sets a standard direction to right, do not change this to up or left as the snake.
	// will crash into a wall right after the game starts.
	snake.Direction = right
	// Creates a snake containing of 3 entities with the given coordinates.
	snake.Bodylength = []Coordinates{
		{1, 6}, // Tail (The tail of the snake will stay the same unless the snake is not colliding with food)
		{2, 6}, // Body (The body will grow taller when a new head is created, the last peice of the body will become the tail if there is no collision with food)
		{3, 6}, // Head (Will become a piece of the body when a new head is created)
	}

	return snake
}

// Head is the snake head wich is used to move the snake around.
// The head is also the hitbox for food, border and the snake itself.
func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

// BorderCollision checks if the arena border contains the snakes head, if so it will return true.
func (snake *Snake) BorderCollision() bool {
	return arena.Contains(*snake.Head())
}

// FoodCollision checks if the food contains the snakes head, if so it will return true.
func (snake *Snake) FoodCollision() bool {
	return food.Contains(*snake.Head())
}

// SnakeCollision checks if the snakes body contains its head, if so it will return true.
func (snake *Snake) SnakeCollision() bool {
	return snake.Contains()
}

// Draw will check every tick and draw the snake on the screen, it also checks if the snake has any collisions
// using the funtions above.
func (snake *Snake) Draw(screen *tl.Screen) {
	// This will create a new head give the direction the snake is heading.
	nHead := *snake.Head()
	// Checks the current direction of the snake.
	switch snake.Direction {
	// If the snakedirection is up, the snake will move up every tick.
	case up:
		// The Y coorddinates will be lowered, making the snake move up.
		nHead.Y--
	// If the snakedirection is down, the snake will move down every tick.
	case down:
		// The Y coorddinates will be incresed, making the snake move down.
		nHead.Y++
	// If the snakedirection is left, the snake will move left every tick.
	case left:
		// The X coorddinates will be lowered, making the snake move left.
		nHead.X--
	// If the snakedirection is right, the snake will move right every tick.
	case right:
		// The X coorddinates will be incresed, making the snake move right.
		nHead.X++
	}

	// Checks for a food collision using the collision function.
	if snake.FoodCollision() {
		// This switch case checks if the food emoji is a special kind of food
		switch food.Emoji {
		// If the food is a mouse, you will score extra points!
		case '🐁':
			// This is the Unicode version of the mouse.
			// This calls the UpdateScore funtion and will give a score of 5.
			// You can change the scores at any point if you want!
			UpdateScore(5)
			// The old head will be moved to a body and a new head will become the head.
			snake.Bodylength = append(snake.Bodylength, nHead)
		case 'R':
			// This is the Non UTF-8 version of the rat.
			// This calls the UpdateScore funtion and will give a score of 5.
			// You can change the scores at any point if you want!
			UpdateScore(5)
			// The old head will be moved to a body and a new head will become the head.
			snake.Bodylength = append(snake.Bodylength, nHead)
		default:
			// If the food emoji is any other emoji it will give a score of one.
			UpdateScore(1)
			// The old head will be moved to a body and a new head will become the head.
			snake.Bodylength = append(snake.Bodylength, nHead)
		}
		// If there is a food collision the food it will call the MoveFood funtion to move the food
		food.MoveFood()
	} else {
		// If there is no collision with food the snake will add the new head but exclude the tail from the body
		// keeping the snake the same size as before.
		snake.Bodylength = append(snake.Bodylength[1:], nHead)
	}

	// The position of the snake will be moved after the new heads coordinates.
	snake.SetPosition(nHead.X, nHead.Y)

	// Check if the snake is colliding with the border or itself using the collision functions.
	if snake.BorderCollision() || snake.SnakeCollision() {
		// Calls the GameOver function to take the player to the game over screen.
		Gameover()
	}

	// This for loop will range over the snakebody and print out the snake given the body coordinates.
	// This will update every tick so the snake keeps moving.
	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: ParseUserSettingsColor(snakecolor),
			Ch: snakerune,
		})
	}
}

// Contains checks if the snake contains the head of the snake, if so it will return true.
func (snake *Snake) Contains() bool {
	// This for loop will check if the head is in any part of the body.
	for i := 0; i < len(snake.Bodylength)-1; i++ {
		// If the head is in any part of the body, it will return true.
		if *snake.Head() == snake.Bodylength[i] {
			return true
		}
	}
	// It will return false if the snake is not colliding with itself.
	return false
}
