import ApiConfig from './config';

function postResolve(input: string): Promise<string> {
	return new Promise<string>((resolve, reject) => {
		const url = `http://${ApiConfig.origin}:${ApiConfig.port}/resolve`;
		console.log("url:", url)

		fetch(url, {
			method: 'POST',
			body: input
		})
			.then((response) => {
				response
					.text()
					.then((body) => {
						resolve(body);
					})
					.catch((err) => {
						reject(err);
					});
			})
			.catch((err) => {
				reject(err);
			});
	});
}

export default postResolve;
