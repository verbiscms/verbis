// Load 3D Scene
import gsap from "gsap";
import * as THREE from "three";
import {GLTFLoader} from "three/examples/jsm/loaders/GLTFLoader";

let scene = new THREE.Scene();

// Load Camera Perspective
const camera = new THREE.PerspectiveCamera( 10, window.innerWidth / window.innerHeight, 1, 20000 );
camera.position.set( 1, 1, 20);
camera.position.set( 9.023588180541993, 4.472665786743164,  9.152990341186524);
camera.rotation.set(-0.4367282469153242, 0.7325030998379315, 0.3025729450328216)
camera.zoom = 1.5;

// Load a Renderer
const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true });
renderer.setClearColor( 0xC5C5C3 );
renderer.setPixelRatio( window.devicePixelRatio );
renderer.setSize(window.innerWidth, window.innerHeight);
document.querySelector('.home-hero-canvas').appendChild(renderer.domElement);

const directionalLight = new THREE.DirectionalLight(0xffffff, 0.5);
directionalLight.position.set( 2, 1, 3).normalize();
scene.add(directionalLight);

const pointLight = new THREE.PointLight( 0xF6E4BF,  0.20, 0);
pointLight.position.set( 1.080, 0.440, 0.040 );
pointLight.decay = 1.94;
scene.add(pointLight);

renderer.outputEncoding = THREE.sRGBEncoding;
renderer.physicallyCorrectLights = true;

// Load a glTF resource
var loader = new GLTFLoader();
loader.load(
    // Resource URL
    '/assets/3d/how-to-sleep-arm.glb',
    // Called when the resource is loaded
    function (gltf) {
        scene.add(gltf.scene)
        camera.updateProjectionMatrix();
        animate()
        console.log(scene.children[2].children[5].children[2]);

    },
    // Called while loading is progressing
    function (xhr) {
        console.log( (xhr.loaded / xhr.total * 100 ) + '% loaded');
    },
    // Called when loading has errors
    function ( error ) {
        console.log(error);
        console.log( 'An error happened' );
    }
);

function animate() {
    render();
    let moon = scene.children[2].children[2].children[1];




    //moon.rotation.y -= 0.014
    requestAnimationFrame(animate);
}

function render() {
    renderer.render(scene, camera);
}

setInterval(rotateClock, 1000)


function rotateClock() {
    const currentDate = new Date(),
        secondsRatio = currentDate.getSeconds() / 60,
        minutesRatio = (secondsRatio + currentDate.getMinutes()) / 60,
        hoursRatio = (minutesRatio + currentDate.getHours()) / 12;

    let secondArm = scene.children[2].children[7],
        minutesArm = scene.children[2].children[5],
        hourArm = scene.children[2].children[6]

    //console.log(secondsRatio * 360);

    secondArm.rotation.x = (Date.now() / 1000 ) % 60 * 6
    //minutesArm.rotation.x = (Date.now()/60000)%60 * 6

    // console.log(currentDate.getSeconds())
    // console.log(secondsRatio * Math.PI)
    //
    //secondArm.rotation.x = secondsRatio * 360 / 3;
    // minutesArm.rotation.x = minutesRatio * 360;
   // hourArm.rotation.x = hoursRatio * 360;
}




// function clock() {
//     let arm = scene.children[2].children[5].children[2];
//     gsap.to(arm.rotation, {x: 2, duration: 0.2})
// }



