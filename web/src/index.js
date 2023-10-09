const gridDimensions = {
  height: 5,
  width: 5,
};

const Face = {
  NORTH: "NORTH",
  EAST: "EAST",
  SOUTH: "SOUTH",
  WEST: "WEST",
};

const Commands = {
  MOVE: "MOVE",
  LEFT: "LEFT",
  RIGHT: "RIGHT",
};

class Robot {
  constructor({ x, y }, initialFacing, { gridWidthBounds, gridHeightBounds }) {
    this.x = parseInt(x);
    this.y = parseInt(y);
    this.facing = initialFacing;
    this.gridHeightBounds = parseInt(gridHeightBounds);
    this.gridWidthBounds = parseInt(gridWidthBounds);
    this.icon = this.currentIcon();
  }

  currentIcon() {
    switch (this.facing) {
      case Face.NORTH:
        this.icon = "⬆️";
        break;
      case Face.EAST:
        this.icon = "➡️";
        break;
      case Face.SOUTH:
        this.icon = "⬇️";
        break;
      case Face.WEST:
        this.icon = "⬅️";
        break;
      default:
        this.icon = "X";
    }

    return this.icon;
  }

  move() {
    switch (this.facing) {
      case Face.NORTH:
        if (this.y == 0) {
          console.error("invalid move");
        } else {
          this.y--;
        }
        break;

      case Face.EAST:
        if (this.x == this.gridWidthBounds - 1) {
          console.error("invalid move");
        } else {
          this.x++;
        }
        break;
      case Face.WEST:
        if (this.x == 0) {
          console.error("invalid move");
        } else {
          this.x--;
        }
        break;
      case Face.SOUTH:
        if (this.y == this.gridHeightBounds - 1) {
          console.error("invalid move");
        } else {
          this.y++;
        }
    }
  }

  turn(command) {
    let newFace;

    switch (command) {
      case Commands.LEFT:
        switch (this.facing) {
          case Face.NORTH:
            newFace = Face.WEST;
            break;
          case Face.EAST:
            newFace = Face.NORTH;
            break;
          case Face.SOUTH:
            newFace = Face.EAST;
            break;
          case Face.WEST:
            newFace = Face.SOUTH;
            break;
        }
        break;

      case Commands.RIGHT:
        switch (this.facing) {
          case Face.NORTH:
            newFace = Face.EAST;
            break;
          case Face.EAST:
            newFace = Face.SOUTH;
            break;
          case Face.SOUTH:
            newFace = Face.WEST;
            break;
          case Face.WEST:
            newFace = Face.NORTH;
            break;
        }
        break;
    }

    this.facing = newFace;
  }
}

let robot;

function place() {
  const [x, y, facing] = document
    .getElementById("coordinates")
    .value.trim()
    .split(";");

  robot = new Robot({ x, y }, facing.toUpperCase(), {
    gridHeightBounds: gridDimensions.height,
    gridWidthBounds: gridDimensions.width,
  });

  drawGrid();
}

function move() {
  robot.move();
  drawGrid();
}

function turn(command) {
  robot.turn(command);
  drawGrid();
}

function drawGrid() {
  const gridContainer = document.querySelector(".grid-container");
  gridContainer.innerHTML = "";

  for (let row = 0; row < gridDimensions.height; row++) {
    for (let col = 0; col < gridDimensions.width; col++) {
      const cell = document.createElement("div");
      cell.classList.add("grid-cell");

      if (robot && row == robot.y && col == robot.x) {
        cell.textContent = robot.currentIcon();
      } else {
        cell.textContent = `${col}, ${row}`;
      }

      gridContainer.appendChild(cell);
    }
  }
}
