async function uploadImage() {
	const file = document.querySelector("input[type=file]").files[0];
	if (!file) return;

	const reader = new FileReader();
	reader.onloadend = async () => {
		const request = {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: reader.result,
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
