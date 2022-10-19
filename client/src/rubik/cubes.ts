import * as THREE from 'three';

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

async function generateFacesFromPosition(position: THREE.Vector3) {
	const faces = await import('./faces');

	return [
		position.x === 1 ? faces.default.orange : faces.default.black, // >
		position.x === -1 ? faces.default.red : faces.default.black, // <
		position.y === 1 ? faces.default.white : faces.default.black, // ^
		position.y === -1 ? faces.default.yellow : faces.default.black, // v
		position.z === 1 ? faces.default.blue : faces.default.black, // *v
		position.z === -1 ? faces.default.green : faces.default.black // *^
	];
}

// generate cubes for a rubik
export default async function generateCubes(): Promise<THREE.Mesh[]> {
	const geometry = new THREE.BoxGeometry(1, 1, 1);

	const cubes: THREE.Mesh[] = [];

	for (let i = -1; i < 2; i++)
		for (let j = -1; j < 2; j++)
			for (let k = -1; k < 2; k++) {
				const position = new THREE.Vector3(i, j, k);
				const mesh = new THREE.Mesh(geometry, await generateFacesFromPosition(position));

				mesh.position.copy(position);
				cubes.push(mesh);
			}

	return cubes;
}
