export const cn = (...classes: (string | false | undefined)[]): string => {
	return classes.filter(Boolean).join(' ');
};
