const getE = id => document.getElementById(id);
const getQ = sel => document.querySelector(sel);

const visualizer = getQ(".visualizer");
const ctxMenu = getQ(".context-menu");
const editor = getE("editor");

const RECTANGLE = 0;
const CIRCLE = 1;
const POLYGON = 2;

const minSize = 20;
const maxHeight = 70; // vh

let CssWidth;
let CssHeight;
let ImageWidth;
let ImageHeight;
let ImageRatio;
let ImageURL;
let ImageBlob;

let listShapes = [];
let vSize = [0, 0];
let startPos = [0, 0];
let selected = null;
let mouseDown = false;
let shapeType = RECTANGLE;

function getMousePos(e) {
	const x = e.clientX;
	const y = e.clientY + window.scrollY;
	return [x, y];
}

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

// Context menu in editor is reserved for ctxMenu
editor.addEventListener("contextmenu", e => e.preventDefault());

// Close ctx menu
document.addEventListener("click", e => {
	if (e.target.offsetParent !== ctxMenu) {
		ctxMenu.style.display = "none";
	}
});

// Resets drawing and the border rect/circle
editor.addEventListener("mousedown", e => {
	toggleDraw(true);
	startPos = getMousePos(e);
	const [x, y] = getMousePos(e);
	setDiv(visualizer, 0, 0, x, y);
	visualizer.style.borderRadius = shapeType == 1 ? "50%" : "0%";
	vSize = [10, 10];
});

// Resize shape outline
editor.addEventListener("mousemove", e => {
	if (mouseDown) {
		if (vSize[0] < minSize || vSize[1] < minSize) visualizer.style.borderColor = "red";
		else visualizer.style.borderColor = "gray";

		const [x, y] = getMousePos(e);
		let w = x - startPos[0];
		let h = y - startPos[1];
		if (shapeType == CIRCLE) {
			const s = Math.max(w, h);
			(w = s), (h = s);
		}

		// For circle to not go past border
		if (startPos[1] + h > editor.offsetTop + editor.offsetHeight)
			h = w = editor.offsetTop + editor.offsetHeight - startPos[1];
		else if (startPos[0] + w > editor.offsetLeft + editor.offsetWidth)
			h = w = editor.offsetLeft + editor.offsetWidth - startPos[0];

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
			div.classList.add(["rectangle", "circle", "polygon"][shapeType]);
			div.dataset.type = shapeType;
			div.style.opacity = 1;

			const ox = -(startPos[0] - editor.offsetLeft);
			const oy = -(startPos[1] - editor.offsetTop);
			div.style.backgroundPosition = `${ox}px ${oy}px`;
			div.style.backgroundImage = `url(${ImageURL})`;
			div.style.backgroundSize = `${CssWidth}vh ${CssHeight}vh`;

			setDiv(div, vSize[0], vSize[1], startPos[0], startPos[1]);
			editor.appendChild(div);
			listShapes.push(div);

			div.addEventListener("contextmenu", e => {
				e.preventDefault();
				selected = div;
				getE("opacity-slider").value = div.style.opacity * 100;
				ctxMenu.style.display = "flex";
				const [x, y] = getMousePos(e);
				ctxMenu.style.left = `${x}px`;
				ctxMenu.style.top = `${y}px`;
			});
		}
	}

	mouseDown = false;
});

const nmap = (n, a, b, x, y) => ((n - a) / (b - a)) * (y - x) + x;

// Gets all the shapes as array of objects with same key types as Shape struct
function getAllShapeData() {
	const data = [];
	const scale = editor.offsetWidth / ImageWidth;

	for (const shape of listShapes) {
		const t = Number(shape.dataset.type);
		const w = shape.offsetWidth;
		const h = shape.offsetHeight;
		const x = shape.offsetLeft - editor.offsetLeft + (t == CIRCLE ? w / 2 : 0);
		const y = shape.offsetTop - editor.offsetTop + (t == CIRCLE ? h / 2 : 0);
		const o = shape.style.opacity;

		// All numbers are expected to be integers
		data.push({
			Type: t,
			Pos: {
				X: Math.floor(x / scale),
				Y: Math.floor(y / scale),
			},
			Size: {
				// Width is radius for circles
				X: Math.floor(w / scale),
				Y: Math.floor(h / scale),
			},
			Opacity: Math.floor(nmap(Number(o), 0, 1, 0, 255)),
		});
	}

	return data;
}
