import Router from 'vue-router';
import Vue from 'vue';
import Home from '../app/pages/shared/Home.vue';
import ChallengeList from '../app/pages/shared/ChallengeList.vue';
import ChallengeDetail from '../app/pages/shared/ChallengeDetail.vue';
import RefugeeProfile from '../app/pages/refugees/Profile.vue';
import CompanyProfile from '../app/pages/companies/Profile.vue';
import Authentication from '../app/pages/refugees/Authentication.vue';
import CreateChallenge from '../app/pages/companies/CreateChallenge.vue';
import RefugeeRegister from '../app/pages/refugees/Register.vue';
import CompanyRegister from '../app/pages/companies/Register.vue';
import Interests from '../app/pages/refugees/Interests.vue';

Vue.use(Router);

const routes = [
	{
		path: '/',
		component: Home,
	},
	{
		path: '/challenges',
		component: ChallengeList,
	},
	{
		path: '/challenge/:id',
		component: ChallengeDetail,
	},
	{
		path: '/refugee/profile/:id',
		component: RefugeeProfile,
	},
	{
		path: '/company/profile/:id',
		component: CompanyProfile,
	},
	{
		path: '/refugee/register',
		component: RefugeeRegister,
	},
	{
		path: '/company/register',
		component: CompanyRegister,
	},
	{
		path: '/challenge/create',
		component: CreateChallenge,
	},
	{
		path: '/authentication/:question',
		component: Authentication,
	},
	{
		path: '/interests/add/:idUser',
		component: Interests,
	},
];

const router = new Router({
	mode: 'hash',
	routes,
});

export default router;
