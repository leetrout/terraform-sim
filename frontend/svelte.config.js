import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

const isLocal = process.env.SK_LOCAL == 'development';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: adapter(),
		paths: {
			base: isLocal ? '' : '/static'
		},
		// Uses _app by default but Go will not serve
		// directories with leading underscores so
		// change the appDir when we publish
		appDir: 'internal'
	}
};

export default config;
