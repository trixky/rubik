import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
import Cubes from './cubes';

const sizes = {
	width: 266,
	height: 266
};

const scene = new THREE.Scene();

const camera = new THREE.PerspectiveCamera(75, sizes.width / sizes.height);
camera.position.z = 5;

scene.add(camera);

let renderer: THREE.WebGLRenderer;

let controls: OrbitControls;

function tick() {
	renderer.render(scene, camera);
	window.requestAnimationFrame(tick);
	controls.update();
}

export const createScene = async (el: HTMLCanvasElement | undefined) => {
	renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true, canvas: el });
	renderer.setSize(sizes.width, sizes.height);
	renderer.render(scene, camera);

	controls = new OrbitControls(camera, el);
	controls.enableDamping = true;
	tick();

	Cubes.forEach((mesh) => {
		const group = new THREE.Group();
		group.add(mesh);
		scene.add(group);
	});
};
