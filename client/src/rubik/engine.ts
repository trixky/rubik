import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';

export default class engine {
	sizes: {
		width: number;
		height: number;
	};
	scene: THREE.Scene;
	camera: THREE.Camera;
	renderer: THREE.WebGLRenderer;
	controls: OrbitControls;

	tick() {
		this.renderer.render(this.scene, this.camera);

		window.requestAnimationFrame(this.tick.bind(this));
		this.controls.update();
	}

	constructor(width: number, height: number, el: HTMLCanvasElement | undefined) {
		this.sizes = { height, width };
		this.scene = new THREE.Scene();
		this.camera = new THREE.PerspectiveCamera(75, this.sizes.width / this.sizes.height);
		this.camera.position.z = 5;
		this.scene.add(this.camera);

		this.renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true, canvas: el });
		this.renderer.setSize(this.sizes.width, this.sizes.height);
		this.renderer.render(this.scene, this.camera);
		this.controls = new OrbitControls(this.camera, el);
		this.controls.enableZoom = false;
		this.controls.enablePan = false;
		this.controls.enableDamping = true;

		this.tick();
	}
}
