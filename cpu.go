package main

type CPU struct {
	paddle Paddle
}

func (c *CPU) draw() {
	c.paddle.draw()
}

func (c *CPU) update(ball *Ball) {
	if c.paddle.y+c.paddle.height/2 > ball.y {
		c.paddle.y -= c.paddle.speed * delta
	}
	if c.paddle.y+c.paddle.height/2 < ball.y {
		c.paddle.y += c.paddle.speed * delta
	}

	c.paddle.limit_movement()
}
