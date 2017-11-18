export const hasFurtherDetails = () => {
	debugger;
	const introduction = window.localStorage.getItem('introduction');
	const city = window.localStorage.getItem('city');
	if (introduction === null || city === null || typeof introduction === 'undefined' || typeof city === 'undefined' || introduction === '' || city === '') {
		return false;
	}
	return true;
};