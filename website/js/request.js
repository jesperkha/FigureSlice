// Todo add check for incomplete form

async function getMaskedImage() {
	const formData = new FormData(document.getElementById("form"));
	formData.append("Shapes", JSON.stringify(GetAllShapeData()));

	const res = await fetch("/image", {
		method: "POST",
		headers: {
			// Browser adds content type automatically
			// "Content-Type": "multipart/form-data",
		},
		redirect: "follow",
		body: formData,
	});

	if (res.status != 200) {
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

function getImage() {
	const file = document.querySelector("input[type=file]").files[0];
	if (!file) return;

	const reader = new FileReader();
	reader.onloadend = () => {
		const editor = document.querySelector(".editor");
		const img = new Image();
		img.onload = () => {
			ImageWidth = img.width;
			ImageHeight = img.height;
			ImageRatio = img.width / img.height;
			ImageURL = reader.result;

			document.querySelector(".background").src = reader.result;
			editor.style.width = maxHeight * ImageRatio + "vw";
			editor.style.height = maxHeight + "vw";
			editor.style.display = "flex";
		};

		img.src = reader.result;
	};

	reader.readAsDataURL(file);
}
