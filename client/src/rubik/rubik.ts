import * as THREE from 'three';
import type Moves from '../models/move';
import gsap from 'gsap';
import Move from '../models/move';
import Config from './config';
import type Engine from './engine';

const quarterTurn = Math.PI / 2;

export class Rubik {
	cubes: THREE.Mesh[];
	engine: Engine;
	pivot: THREE.Mesh;
	moves: Moves[];
	moving: boolean;

	constructor(cubes: THREE.Mesh[], engine: Engine) {
		this.engine = engine;
		this.cubes = cubes;
		this.cubes.forEach((cube) => {
			this.engine.scene.add(cube);
		});
		this.pivot = new THREE.Mesh();
		this.engine.scene.add(this.pivot);
		this.moves = [];
		this.moving = false;
	}

	pushMove(instruction: string, reverse = false) {
		if (instruction.length === 2) {
			const basic_instruction = instruction[0];

			if (instruction[1] === '2') {
				this.pushMove(basic_instruction, reverse);
				this.pushMove(basic_instruction, reverse);
			} else {
				this.pushMove(basic_instruction, !reverse);
			}
		} else {
			(() => {
				switch (instruction) {
					case 'F':
						return this.moves.push(new Move('x', -1, !reverse));
					case 'R':
						return this.moves.push(new Move('z', 1, reverse));
					case 'U':
						return this.moves.push(new Move('y', 1, reverse));
					case 'B':
						return this.moves.push(new Move('x', 1, reverse));
					case 'L':
						return this.moves.push(new Move('z', -1, !reverse));
					case 'D':
						return this.moves.push(new Move('y', -1, !reverse));
				}
			}).bind(this)();

			this.nextMove();
		}
	}

	async nextMove(force_moving = false) {
		if (!this.moving || force_moving) {
			this.moving = true;

			const move = this.moves.shift();

			if (move) {
				this.pivot.rotation.set(0, 0, 0);

				const attached_cubes: THREE.Mesh[] = [];

				this.cubes.forEach((cube) => {
					if (cube.position[move.axe] === move.depth) {
						attached_cubes.push(cube);
						this.pivot.attach(cube);
					}
				});

				await gsap.to(this.pivot.rotation, {
					duration: Config.moves.duration,
					[move.axe]: move.reverse ? quarterTurn : -quarterTurn
				});

				attached_cubes.forEach((cube) => {
					this.engine.scene.attach(cube);
					move.otherAxis.forEach((axe) => {
						cube.position[axe] = Math.round(cube.position[axe]);
					});
				});

				await this.nextMove(true);
			}

			if (!force_moving) this.moving = false;
		}
	}

	visible(visible: boolean) {
		this.cubes.forEach((cube) => {
			cube.visible = visible;
		});
	}
}
