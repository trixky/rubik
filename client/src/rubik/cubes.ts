import * as THREE from 'three';
import Faces from './faces';

/*
      g g g
      g g g
      g g g
r r r w w w o o o
r r r w w w o o o
r r r w w w o o o
      b b b
      b b b
      b b b
      y y y
      y y y
      y y y

      a b c
      d e f
      g h i
a d g g h i i f c
j k l l m n n o p
q r s s t u u v w
      s t u
      r x v
      q y w
      q y w
      j z p
      a b c
*/

function generateFacesFromPosition(position: THREE.Vector3) {
	return [
		position.x === 1 ? Faces.orange : Faces.black, // >
		position.x === -1 ? Faces.red : Faces.black, // <
		position.y === 1 ? Faces.white : Faces.black, // ^
		position.y === -1 ? Faces.yellow : Faces.black, // v
		position.z === 1 ? Faces.blue : Faces.black, // *v
		position.z === -1 ? Faces.green : Faces.black // *^
	];
}

function generateCubes(): THREE.Mesh[] {
	const geometry = new THREE.BoxGeometry(1, 1, 1);

	const cubes: THREE.Mesh[] = [];

	for (let i = -1; i < 2; i++)
		for (let j = -1; j < 2; j++)
			for (let k = -1; k < 2; k++) {
				const position = new THREE.Vector3(i, j, k);
				const mesh = new THREE.Mesh(geometry, generateFacesFromPosition(position));
				mesh.position.copy(position);
				cubes.push(mesh);
			}

	return cubes;
}

const cubes: THREE.Mesh[] = generateCubes();

export default cubes;
