const visualizer = document.querySelector(".visualizer");
const ctxMenu = document.querySelector(".context-menu");
const editor = document.getElementById("editor");
const editorOffsetX = editor.offsetLeft;
const editorOffsetY = editor.offsetTop;

const RECTANGLE = 0;
const CIRCLE = 1;

const minSize = 20;
const shapes = ["rectangle", "circle"];
const shapeConfig = {
	rectangle: div => {
		div.style.borderRadius = "0%";
	},
	circle: div => {
		div.style.borderRadius = "50%";
	},
};

let listShapes = [];
let vSize = [0, 0];
let shapeType = CIRCLE;
let selected = null;
let mouseDown = false;
let startPos = [];

// Shorthand for setting style values
function setDiv(div, w, h, x = null, y = null) {
	div.style.width = w + "px";
	div.style.height = h + "px";

	if (x && y) {
		div.style.left = x + "px";
		div.style.top = y + "px";
	}
}

// Toggle mouse events and visibility of outline
function toggleDraw(state) {
	mouseDown = state;
	visualizer.style.display = state ? "block" : "none";
}

// Returns new size for outline on mousemove
function vgetNewSize(x, y, shape) {
	let w = x - startPos[0];
	let h = y - startPos[1];
	if (shape == CIRCLE) {
		const s = Math.max(w, h);
		(w = s), (h = s);
	}

	return [w, h];
}

// Context menu in editor is reserved for ctxMenu
editor.addEventListener("contextmenu", e => e.preventDefault());

// Close ctx menu
editor.addEventListener("click", e => {
	if (e.target.offsetParent !== ctxMenu) {
		ctxMenu.style.display = "none";
	}
});

// Resets drawing and the border rect/circle
editor.addEventListener("mousedown", e => {
	toggleDraw(true);
	startPos = [e.clientX, e.clientY];
	setDiv(visualizer, 0, 0, e.clientX, e.clientY);
	shapeConfig[shapes[shapeType]](visualizer);
	vSize = [0, 0];
});

// Resize shape outline
editor.addEventListener("mousemove", e => {
	if (mouseDown) {
		if (vSize[0] < minSize || vSize[1] < minSize) visualizer.style.borderColor = "red";
		else visualizer.style.borderColor = "gray";

		const [w, h] = vgetNewSize(e.clientX, e.clientY, shapeType);
		setDiv(visualizer, w, h);
		vSize = [w, h];
	}
});

// Reset / create new shape at outline
document.addEventListener("mouseup", e => {
	if (mouseDown) {
		toggleDraw(false);

		// Create new shape
		if (vSize[0] > minSize && vSize[1] > minSize) {
			const div = document.createElement("div");
			div.classList.add(shapes[shapeType]);
			div.dataset.type = shapeType;
			div.style.opacity = 1;
			setDiv(div, vSize[0], vSize[1], startPos[0], startPos[1]);
			editor.appendChild(div);
			listShapes.push(div);

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

// From ctxMenu
function removeShape() {
	if (selected) {
		editor.removeChild(selected);
		ctxMenu.style.display = "none";
		selected = null;
	}
}

// From ctxMenu
function changeOpacity(e) {
	selected.style.opacity = e.value / 100;
}

// Gets all the shapes as array of objects with same key types as Shape struct
// Called when shape data should be exported
function GetAllShapeData() {
	const data = [];
	const nmap = (n, a, b, x, y) => ((n - a) / (b - a)) * (y - x) + x;

	for (const shape of listShapes) {
		const t = shape.dataset.type;
		const X = shape.offsetLeft - editorOffsetX;
		const Y = shape.offsetTop - editorOffsetY;
		const w = shape.offsetWidth;
		const h = shape.offsetHeight;
		const o = shape.style.opacity;

		data.push({
			Type: Number(t),
			Pos: {
				X,
				Y,
			},
			Size: {
				X: w,
				Y: h,
			},
			Opacity: nmap(Number(o), 0, 1, 0, 255),
		});
	}

	return data;
}
