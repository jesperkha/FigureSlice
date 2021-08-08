let selected = null;

const ctxMenu = document.querySelector(".context-menu");
const editor = document.getElementById("editor");

editor.addEventListener("contextmenu", e => e.preventDefault());
editor.addEventListener("click", e => {
	if (e.target.offsetParent !== ctxMenu) {
		ctxMenu.style.display = "none";
	}
});

function removeShape() {
	if (selected) {
		editor.removeChild(selected.div);
		ctxMenu.style.display = "none";
		selected = null;
	}
}

function resizeShapeX(e) {
	selected.w = e.value * 10 + 50;
	selected.div.style.width = `${selected.w}px`;
}

function resizeShapeY(e) {
	selected.h = e.value * 10 + 50;
	selected.div.style.height = `${selected.h}px`;
}

class Shape {
	constructor() {
		this.selected = false;
		this.grabOffset = [0, 0];

		this.div = document.createElement("div");
		this.div.classList.add("rectangle");
		editor.appendChild(this.div);

		this.div.addEventListener("mousedown", e => {
			if (e.button == 0) {
				this.selected = true;
				selected = this;
				this.grabOffset = [e.clientX - this.x, e.clientY - this.y];
			}
		});

		this.div.addEventListener("mouseup", e => (this.selected = false));
		this.div.addEventListener("mouseleave", e => (this.selected = false));

		this.div.addEventListener("contextmenu", e => {
			e.preventDefault();
			selected = this;
			ctxMenu.style.display = "flex";
			ctxMenu.style.left = `${e.clientX}px`;
			ctxMenu.style.top = `${e.clientY}px`;
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
		this.w = 550;
		this.h = 550;
	}
}

new Shape();
new Shape();
