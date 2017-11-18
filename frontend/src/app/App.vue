<template>
    <div>
        <router-view></router-view>
    </div>
</template>

<script>
	export default {
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

<style scoped lang="scss" type="text/scss">
  .test {
    .test {
      color: pink;
    }
  }
</style>
