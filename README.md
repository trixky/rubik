# rubik

A 3D rubik resolver online, using the [Morwen Thistlethwaite](https://en.wikipedia.org/wiki/Morwen_Thistlethwaite) and [IDA*](https://en.wikipedia.org/wiki/Iterative_deepening_A*) algorithms, rendered with the [Three.js](https://threejs.org/) browser [WebGL](https://en.wikipedia.org/wiki/WebGL) engine.

[client demo.webm](https://user-images.githubusercontent.com/50099834/197333015-2a1fd24c-623f-4ab2-8ecf-6fe04b4cbc5a.webm)

## Notation

The notation used is the one used globally:

- F: front
- R: right
- U: up
- B: back
- L: left
- D: down

You can add "__F2__" *(double)* instead of "__F F__", and "__F'__" *(reverse)* instead of "__F F F__".

## Usage

### Full-stack (prod)

```bash
source env.sh
docker-compose -f docker-compose.prod.yaml up
```

### Full-stack (dev)

```bash
source env.sh
docker-compose up
```

### Api only

```bash
cd api
go build
./rubik
```

### Command line

```bash
cd api
go build
./rubik "F2 D' F D R' U L B2"
```

## Online

This project is online, so you can visit it by clicking [here](https://rubik.trixky.com/)!