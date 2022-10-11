import * as THREE from 'three';

const textureLoader = new THREE.TextureLoader();

export default {
	white: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/white.png')
	}),
	blue: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/blue.png')
	}),
	orange: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/orange.png')
	}),
	red: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/red.png')
	}),
	green: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/green.png')
	}),
	yellow: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/yellow.png')
	}),
	black: new THREE.MeshBasicMaterial({
		map: textureLoader.load('/faces/black.png')
	})
};
