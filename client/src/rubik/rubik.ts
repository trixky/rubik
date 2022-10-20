import * as THREE from 'three';
import type Moves from '../models/move';
import gsap from 'gsap';
import Move from '../models/move';
import Config from './config';
import type Engine from './engine';
import { tick } from 'svelte';

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

	// push a move to the stack
	pushMove(instruction: string, reverse = false, duration = Config.moves.durations.slow) {
		if (instruction.length === 2) {
			// if handle a "complex" instruction

			// extractt the basic instruction from it
			const basic_instruction = instruction[0];

			if (instruction[1] === '2') {
				// if the instruction extension is "2" (double)
				// split the "complexe" instruction in two basic instruction
				this.pushMove(basic_instruction, reverse, duration);
				this.pushMove(basic_instruction, reverse, duration);
			} else {
				// else is the instruction extension "'" (reverse)
				// call the reverse corresponding basic instruction
				this.pushMove(basic_instruction, !reverse, duration);
			}
		} else {
			// else is a basic instruction
			// push the instruction
			(() => {
				switch (instruction) {
					case 'F':
						return this.moves.push(new Move('x', -1, !reverse, duration));
					case 'R':
						return this.moves.push(new Move('z', 1, reverse, duration));
					case 'U':
						return this.moves.push(new Move('y', 1, reverse, duration));
					case 'B':
						return this.moves.push(new Move('x', 1, reverse, duration));
					case 'L':
						return this.moves.push(new Move('z', -1, !reverse, duration));
					case 'D':
						return this.moves.push(new Move('y', -1, !reverse, duration));
				}
			}).bind(this)();

			// play the stacked moves
			tick().then(() => {
				this.nextMove();
			});
		}
	}

	// play the stacked moves
	async nextMove(force_moving = false) {
		if (!this.moving || force_moving) {
			// if the stack is not already being played
			// block it
			this.moving = true;

			// extract the next move
			const move = this.moves.shift();

			if (move) {
				// if a move as been extracted
				// reset the pivot
				this.pivot.rotation.set(0, 0, 0);

				const attached_cubes: THREE.Mesh[] = [];

				// attach each concerned cubes to the pivot
				this.cubes.forEach((cube) => {
					if (cube.position[move.axe] === move.depth) {
						attached_cubes.push(cube);
						this.pivot.attach(cube);
					}
				});

				// compute the stack speed divider
				// (the more the stack is filled, the faster the movements are)
				const stack_speed_divider = this.moves.length === 0 ? 1 : 1 + Math.log(this.moves.length);

				const instant_speed = this.moves.length > Config.moves.instant;

				const duration = instant_speed ? 0 : move.duration / stack_speed_divider;

				// play a rotation on the pivot
				await gsap.to(this.pivot.rotation, {
					duration,
					[move.axe]: move.reverse ? quarterTurn : -quarterTurn
				});

				// attach back the concerned cubes to the scene
				attached_cubes.forEach((cube) => {
					this.engine.scene.attach(cube);
					move.otherAxis.forEach((axe) => {
						cube.position[axe] = Math.round(cube.position[axe]);
					});
				});

				// play next move with recursivity
				await this.nextMove(true);
			}

			if (!force_moving) this.moving = false;
		}
	}

	// show or hide the rubik
	visible(visible: boolean) {
		this.cubes.forEach((cube) => {
			cube.visible = visible;
		});
	}
}
