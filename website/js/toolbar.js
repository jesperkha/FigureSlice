// From ctxMenu
function removeShape() {
	if (selected) {
		listShapes.splice(listShapes.indexOf(selected), 1);
		editor.removeChild(selected);
		ctxMenu.style.display = "none";
		selected = null;
	}
}

// From ctxMenu
function changeOpacity(e) {
	selected.style.opacity = e.value / 100;
}

function toggleTool(n) {
	shapeType = n;
}

function clearAll() {
	for (const s of listShapes) {
		editor.removeChild(s);
	}

	listShapes = [];
}
