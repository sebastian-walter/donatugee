import Vue from 'vue';
import App from './app/App.vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.css';
import router from './router';
import store from './store';
import Icon from 'vue-awesome';

Vue.component('icon', Icon);

Vue.use(Vuetify);

new Vue({
	el: '#app',
	render: h => h(App),
	router,
	store
});
