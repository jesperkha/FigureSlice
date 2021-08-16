async function submitForm() {
	const formData = new FormData(getE("file"));
	formData.append("Shapes", JSON.stringify(getAllShapeData()));
	formData.append("Trim", getE("trim").value);

	if (getE("img").files.length > 0) {
		const status = await getNewImage(formData);
		if (status !== 200) {
			window.location = "/error/" + status;
		}
	}
}

async function getNewImage(formData) {
	const res = await fetch("/image", {
		method: "POST",
		headers: {
			// Browser adds content type automatically
			// "Content-Type": "multipart/form-data",
		},
		redirect: "follow",
		body: formData,
	});

	const blob = await res.blob();
	const reader = new FileReader();
	reader.onloadend = () => {
		getE("preview").setAttribute("src", reader.result);
		getE("download").setAttribute("href", reader.result);
	};

	reader.readAsDataURL(blob);
	return res.status;
}

function loadImage() {
	const file = getQ("input[type=file]").files[0];
	if (!file) return;

	// Set filename
	const filename = getE("img").value.split("\\").pop();
	getE("filename").textContent = `(${filename})`;

	const reader = new FileReader();
	reader.onloadend = () => {
		const editor = getQ(".editor");
		const img = new Image();
		img.onload = () => {
			ImageWidth = img.width;
			ImageHeight = img.height;
			ImageRatio = img.width / img.height;
			ImageURL = reader.result;
			getQ(".editor-image").src = reader.result;

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
