import Vue from 'vue';
import App from './app/App.vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.css';
import router from './router';

Vue.use(Vuetify);

new Vue({
	el: '#app',
	render: h => h(App),
	router
});
