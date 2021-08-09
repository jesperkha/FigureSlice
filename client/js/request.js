// Todo: scale w, h, x, y to actual image size

async function getMaskedImage() {
	const file = document.querySelector("input[type=file]").files[0];
	if (!file) return;

	const reader = new FileReader();
	reader.onloadend = async () => {
		const request = {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(GetAllShapeData()),
			// body: reader.result,
		};

		const res = await fetch("/image", request);
		if (res.status != 200) {
			window.location = "/error/" + res.status;
			return;
		}

		const imageBlob = await res.blob();
		reader.onloadend = () => {
			document.getElementById("img").src = reader.result;
		};

		reader.readAsDataURL(imageBlob);
	};

	reader.readAsArrayBuffer(file);
}

function getImage() {
	const file = document.querySelector("input[type=file]").files[0];
	if (!file) return;

	const reader = new FileReader();
	reader.onloadend = () => {
		const editor = document.querySelector(".editor");
		const img = new Image();
		img.onload = () => {
			const ratio = img.width / img.height;
			editor.style.backgroundImage = `url(${reader.result})`;
			editor.style.width = 35 * ratio + "vw";
			editor.style.height = "35vw";
		};
		img.src = reader.result;
	};

	reader.readAsDataURL(file);
}
