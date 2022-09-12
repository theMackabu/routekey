import ky from 'ky';

const compactLink = (url: string, expires: string) => {
	return ky.post('/api/v2/links', { json: { target: url, expire_in: expires } }).json();
};

export { compactLink };
