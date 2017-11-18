import Router from 'vue-router';
import Vue from 'vue';
import Home from '../app/pages/refugees/Home.vue';

Vue.use(Router);

const routes = [
	{
		path: '/',
		component: Home,
	},
];

const router = new Router({
	mode: 'hash',
	routes,
});

export default router;
