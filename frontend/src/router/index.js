import Router from 'vue-router';
import Vue from 'vue';
import { maxNumberOfIncorrectAnswers } from '../library/questions';
import Home from '../app/pages/shared/Home.vue';
import ChallengeList from '../app/pages/shared/ChallengeList.vue';
import YourChallengeList from '../app/pages/refugees/YourChallengeList.vue';
import ChallengeDetail from '../app/pages/shared/ChallengeDetail.vue';
import CompanyChallengeDetail from '../app/pages/companies/ChallengeDetail.vue';
import CompanyYourChallengeList from '../app/pages/companies/YourChallengeList.vue';
import RefugeeProfile from '../app/pages/refugees/Profile.vue';
import CompanyProfile from '../app/pages/companies/Profile.vue';
import Authentication from '../app/pages/refugees/Authentication.vue';
import CreateChallenge from '../app/pages/companies/CreateChallenge.vue';
import RefugeeRegister from '../app/pages/refugees/Register.vue';
import RefugeeLogin from '../app/pages/refugees/Login.vue';
import CompanyRegister from '../app/pages/companies/Register.vue';
import CompanyLogin from '../app/pages/companies/Login.vue';
import Interests from '../app/pages/refugees/Interests.vue';
import TechQuestions from '../app/pages/refugees/TechQuestions.vue';
import Rejected from '../app/pages/refugees/Rejected.vue';
import FurtherDetails from '../app/pages/refugees/FurtherDetails.vue';

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
		path: '/your-challenges',
		component: YourChallengeList,
	},
	{
		path: '/company/your-challenges',
		component: CompanyYourChallengeList,
	},
	{
		path: '/challenge/:id',
		component: ChallengeDetail,
	},
	{
		path: '/refugee/profile',
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
		path: '/refugee/login',
		component: RefugeeLogin,
	},
	{
		path: '/company/register',
		component: CompanyRegister,
	},
	{
		path: '/companies/login',
		component: CompanyLogin,
	},
	{
		path: '/company/challenge/create',
		component: CreateChallenge,
	},
	{
		path: '/company/challenge/:id',
		component: CompanyChallengeDetail,
	},
	{
		path: '/authentication/:question',
		component: Authentication,
	},
	{
		path: '/interests/add/:idUser',
		component: Interests,
	},
	{
		path: '/tech-questions/:step',
		component: TechQuestions,
	},
	{
		path: '/rejected',
		component: Rejected,
	},
	{
		path: '/refugee/further-details',
		component: FurtherDetails,
	},
];

const router = new Router({
	mode: 'hash',
	routes,
});

router.beforeEach((to, from, next) => {
	if (to.path === '/rejected') {
		return next();
	}

	let wrongAnswers = parseInt(window.localStorage.getItem('wrongAnswers'));
	if (wrongAnswers > maxNumberOfIncorrectAnswers) {
		next('/rejected');
	}
	next();
});

export default router;
