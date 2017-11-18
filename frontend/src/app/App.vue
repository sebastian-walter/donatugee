<template>
    <layout>
        <template slot="content">
            <transition name="fade">
                <router-view :key="$route.params.step"></router-view>
            </transition>
        </template>
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

<style lang="scss" type="text/scss">
    @import '../assets/index.scss';

    .fade-enter-active, .fade-leave-active {
        transition-property: opacity;
        transition-duration: .25s;
    }

    .fade-enter-active {
        transition-delay: .25s;
    }

    .fade-enter, .fade-leave-active {
        opacity: 0
    }
</style>
