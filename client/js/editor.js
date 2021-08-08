const visualizer = document.querySelector(".visualizer");
const ctxMenu = document.querySelector(".context-menu");
const editor = document.getElementById("editor");

const minSize = 50;
const shapes = ["rectangle", "circle"];
const RECTANGLE = 0;
const CIRCLE = 1;

let shapeType = CIRCLE;
let selected = null;
let mouseDown = false;
let startPos = [];
let visualizerWidth = 0;
let visualizerHeight = 0;

// Todo make setting div valules a function
// fix bug where circle gets too big

// dont display ctx menu inside editor
editor.addEventListener("contextmenu", e => e.preventDefault());

// when ctx menu is open it should be closed as normal when clicking
editor.addEventListener("click", e => {
	if (e.target.offsetParent !== ctxMenu) {
		ctxMenu.style.display = "none";
	}
});

// setup for new visual rect
editor.addEventListener("mousedown", e => {
	mouseDown = true;
	startPos = [e.clientX, e.clientY];
	visualizer.style.display = "block";

	visualizer.style.left = e.clientX + "px";
	visualizer.style.top = e.clientY + "px";
	visualizer.style.width = "0px";
	visualizer.style.height = "0px";
	visualizerWidth = 0;
	visualizerHeight = 0;

	if (shapeType == RECTANGLE) {
		visualizer.style.borderRadius = "0%";
	}

	if (shapeType == CIRCLE) {
		visualizer.style.borderRadius = "50%";
	}
});

// show rect and change color based on width/height
editor.addEventListener("mousemove", e => {
	if (mouseDown) {
		const w = e.clientX - startPos[0];
		const h = e.clientY - startPos[1];
		if (visualizerWidth < minSize || visualizerHeight < minSize) visualizer.style.borderColor = "red";
		else visualizer.style.borderColor = "gray";

		if (shapeType == RECTANGLE) {
			visualizer.style.width = w + "px";
			visualizer.style.height = h + "px";
			visualizerWidth = w;
			visualizerHeight = h;
		}

		if (shapeType == CIRCLE) {
			visualizer.style.width = w + "px";
			visualizer.style.height = w + "px";
			visualizerWidth = w;
			visualizerHeight = w;
		}
	}
});

// if mouse is outside editor it should still halt shape drawing
document.addEventListener("mouseup", e => {
	if (mouseDown) {
		mouseDown = false;
		visualizer.style.display = "none";

		// Create new rectangle
		if (visualizerWidth > 50 && visualizerHeight > 50) {
			const div = document.createElement("div");
			div.classList.add(shapes[shapeType]);
			editor.appendChild(div);

			div.style.left = startPos[0] + "px";
			div.style.top = startPos[1] + "px";
			div.style.width = visualizer.style.width;
			div.style.height = visualizer.style.height;

			div.addEventListener("contextmenu", e => {
				e.preventDefault();
				selected = div;
				ctxMenu.style.display = "flex";
				ctxMenu.style.left = `${e.clientX}px`;
				ctxMenu.style.top = `${e.clientY}px`;
			});
		}
	}
});

// ctx menu methods
function removeShape() {
	if (selected) {
		editor.removeChild(selected);
		ctxMenu.style.display = "none";
		selected = null;
	}
}

function changeOpacity(e) {
	selected.style.opacity = e.value / 100;
}
