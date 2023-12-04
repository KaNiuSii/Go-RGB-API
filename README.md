# API Test Project

This project is a simple Go server that generates a color image based on RGB values received from a POST request.

## Setup and Running

To run this project, you need Go installed on your machine.

1. Clone the repository:
git clone https://github.com/KaNiuSii/Go-RGB-API

2. Navigate to the project directory:
cd ApiTest

3. Run the server:
go run main.go

4. Access the server at `http://localhost:8081`.

## Usage

Send a POST request to `/getColor` with JSON payload containing `r`, `g`, and `b` fields (each ranging from 0 to 255) to receive a PNG image of the corresponding color. Use index.html for simple API testing.

## Dependencies

- Go Programming Language