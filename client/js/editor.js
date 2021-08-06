let shapes = [];
let selected = null;

const editor = document.getElementById("editor");
const ex = editor.offsetLeft;
const ey = editor.offsetTop;
const ew = editor.offsetWidth;
const eh = editor.offsetHeight;

class Shape {
	constructor() {
		this.selected = false;
		this.grabOffset = [0, 0];

		this.div = document.createElement("div");
		this.div.classList.add("rectangle");
		editor.appendChild(this.div);

		this.div.addEventListener("mousedown", e => {
			this.selected = true;
			selected = this;
			this.grabOffset = [e.clientX - this.x, e.clientY - this.y];
		});

		this.div.addEventListener("mouseup", e => {
			this.selected = false;
			selected = null;
		});

		this.div.addEventListener("mouseleave", e => {
			this.selected = false;
			selected = null;
		});

		this.div.addEventListener("mousemove", e => {
			if (this.selected) {
				this.x = e.clientX - this.grabOffset[0];
				this.y = e.clientY - this.grabOffset[1];

				this.div.style.left = `${this.x}px`;
				this.div.style.top = `${this.y}px`;
			}
		});

		this.x = 0;
		this.y = 0;
		this.w = 300;
		this.h = 150;
	}
}

new Shape();
