function submitForm() {
	const formData = new FormData(document.getElementById("file"));
	formData.append("Shapes", JSON.stringify(GetAllShapeData()));
	formData.append("Trim", document.getElementById("trim").value);

	if (document.getElementById("img").files.length > 0) {
		getMaskedImage(formData);
	}
}

async function getMaskedImage(formData) {
	const res = await fetch("/image", {
		method: "POST",
		headers: {
			// Browser adds content type automatically
			// "Content-Type": "multipart/form-data",
		},
		redirect: "follow",
		body: formData,
	});

	if (res.status !== 200) {
		window.location = "/error/" + res.status;
		return;
	}

	// Get response image
	const blob = await res.blob();
	const reader = new FileReader();
	reader.onloadend = () => {
		document.getElementById("preview").setAttribute("src", reader.result);
	};

	reader.readAsDataURL(blob);
}

function loadImage() {
	const file = document.querySelector("input[type=file]").files[0];
	if (!file) return;

	// Set filename
	const filename = document.getElementById("img").value.split("\\").pop();
	document.getElementById("filename").textContent = `(${filename})`;

	const reader = new FileReader();
	reader.onloadend = () => {
		const editor = document.querySelector(".editor");
		const img = new Image();
		img.onload = () => {
			ImageWidth = img.width;
			ImageHeight = img.height;
			ImageRatio = img.width / img.height;
			ImageURL = reader.result;
			document.querySelector(".editor-image").src = reader.result;

			// Correct for wider images
			let correction = 1;
			if (ImageRatio > 2) {
				let newRatio = ImageRatio;
				while (newRatio > 2) {
					correction *= 0.9;
					newRatio = (img.width / img.height) * correction;
				}
			}

			CssWidth = maxHeight * ImageRatio * correction;
			CssHeight = maxHeight * correction;
			editor.style.width = CssWidth + "vh";
			editor.style.height = CssHeight + "vh";
			editor.style.display = "flex";
		};

		img.src = reader.result;
	};

	reader.readAsDataURL(file);
}
