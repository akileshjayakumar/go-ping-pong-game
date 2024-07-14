# Ping Pong Game

A Ping Pong game built with Go using the `pixel` library.

## Overview

This project implements a basic Ping Pong game where two players can control paddles to bounce a ball back and forth. The game window is created using the `pixel` library, and the game mechanics include paddle movement, ball collision, and randomized ball direction.

## Installation

1. Make sure you have Go installed. If not, you can download and install it from [here](https://golang.org/dl/).
2. Clone this repository or download the source code.
3. Navigate to the project directory.

```sh
cd ping_pong_game
```

4. Install the required packages.

```sh
go get github.com/faiface/pixel
go get github.com/faiface/pixel/pixelgl
go get golang.org/x/image/colornames
```

## Usage

1. Run the game using the following command:

```sh
go run main.go
```

2. The game window will open, and you can start playing.

## Controls

- **Left Paddle**: Use the `W` key to move up and the `S` key to move down.
- **Right Paddle**: Use the `Up Arrow` key to move up and the `Down Arrow` key to move down.

## Code Explanation

### Paddle

The `Paddle` struct represents a paddle in the game. It includes properties like `Rect` for the paddle's rectangle, `Color` for the paddle's color, `Speed` for the movement speed, and keys for controlling the paddle.

### Ball

The `Ball` struct represents the ball in the game. It includes properties like `Circle` for the ball's shape, `Color` for the ball's color, and `Velocity` for the ball's speed and direction. The ball's movement and collision with the paddles and walls are handled in the `Update` method.

### Main Function

The `run` function initializes the game window, paddles, and ball. It contains the main game loop where the game entities are updated and drawn on each frame.

## Contributing

Your contributions are welcome! If you have ideas for improvements or new features:

1. **Fork the Repository**
   - Click the "Fork" button at the top right of the repository page to create a copy in your GitHub account.

2. **Create a New Branch**
   - Open your terminal and switch to the repository directory.
   - Create and switch to a new branch:
     ```bash
     git checkout -b feature-branch
     ```

3. **Commit Your Changes**
   - Make your desired changes in the code.
   - Stage and commit your changes with a descriptive message:
     ```bash
     git commit -am 'Add new feature: description'
     ```

4. **Push to the Branch**
   - Push your changes to the newly created branch in your forked repository:
     ```bash
     git push origin feature-branch
     ```

5. **Submit a Pull Request**
   - Go to the original repository and click on the "Pull Requests" tab.
   - Click the "New Pull Request" button.
   - Select your branch and provide a clear description of your changes.

## Contact Information

For more information or if you have any questions, please reach out:

- **Email**: [jayakuma006@mymail.sim.edu.sg](mailto:jayakuma006@mymail.sim.edu.sg)
- **LinkedIn**: [Akilesh Jayakumar on LinkedIn](https://www.linkedin.com/in/akileshjayakumar/)
- **GitHub**: [Akilesh Jayakumar on GitHub](https://github.com/akileshjayakumar)
