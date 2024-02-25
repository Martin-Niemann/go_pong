package main

import "math"

const paddle_height = 120
const paddle_width = 10
const paddle_edge_padding = 60
const circle_radius = 6
const ball_speed = 500
const max_bounce_angle = 5 * math.Pi / 16
const paddle_speed = 800

var delta float32 = 1

var player_score = 0
var cpu_score = 0
