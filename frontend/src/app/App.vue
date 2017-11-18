<template>
    <layout>
        <router-view slot="content"></router-view>
    </layout>
</template>

<script>
    import Layout from './components/layout/Layout.vue';

	export default {
		name: 'App',
		components: {
			Layout,
        },
		data () {
			return {
				clipped: false,
				drawer: true,
				fixed: false,
				items: [
					{ icon: 'bubble_chart', title: 'Inspire' }
				],
				miniVariant: false,
				right: true,
				rightDrawer: false,
				title: 'Vuetify.js'
			}
		}
	}
</script>
