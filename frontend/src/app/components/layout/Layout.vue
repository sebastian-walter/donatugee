<template>
    <v-app id="donatugee">
        <v-navigation-drawer
                v-if="isLoggedIn"
                fixed
                v-model="drawer"
                right
                app
        >
            <v-list dense>
                <v-list-tile @click="">
                    <v-list-tile-action>
                        <v-icon>home</v-icon>
                    </v-list-tile-action>
                    <v-list-tile-content>
                        <v-list-tile-title @click="$router.push({ path: '/company/your-challenges' })">
                            All Challenges
                        </v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
                <v-list-tile @click="">
                    <v-list-tile-action>
                        <v-icon>power_settings_new</v-icon>
                    </v-list-tile-action>
                    <v-list-tile-content v-if="isLoggedIn">
                        <v-list-tile-title @click="logout">Logout</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>
        <v-toolbar color="blue accent-4" dark fixed app>
            <div class="logo"><router-link :to="{ path: '/' }"><img src="../../../assets/logo.svg"></router-link></div>
            <v-spacer></v-spacer>
            <v-toolbar-title>HackFugee</v-toolbar-title>
            <v-toolbar-side-icon v-if="isLoggedIn" @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        </v-toolbar>
        <v-content>
            <v-container fluid fill-height>
                <v-layout row>
                    <v-flex xs12 class="app-container">
                        <slot name="content"></slot>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
        <v-card>
            <v-bottom-nav v-if="isLoggedIn && companyId === null" fixed :value="true" color="white">
                <v-btn :class="getClass('/challenges')" :to="{path:'/challenges'}" flat color="primary">
                    <span>Home</span>
                    <v-icon>home</v-icon>
                </v-btn>
                <v-btn :class="getClass('/challenges')" :to="{path:'/your-challenges'}" flat color="primary">
                    <span>Your Challenges</span>
                    <v-icon>star</v-icon>
                </v-btn>
                <v-btn v-if="showRefugeeProfile" :class="getClass('/refugee/profile/')"
                       :to="{path:'/refugee/profile/' + refugeeId}" flat color="primary">
                    <span>Profile</span>
                    <v-icon>person</v-icon>
                </v-btn>
                <v-btn v-if="companyId" :class="getClass('/company/profile/')"
                       :to="{path:'/company/profile/' + refugeeId}" flat color="primary">
                    <span>Profile</span>
                    <v-icon>person</v-icon>
                </v-btn>
            </v-bottom-nav>
        </v-card>
    </v-app>
</template>

<script>
    import { mapActions, mapState } from 'vuex';

	export default {
		name: 'Layout',
		data: () => ({
			drawer: false,
		}),
		props: {
			source: String,
		},
		computed: {
			isLoggedIn() {
				if ((this.refugeeId !== null && typeof this.refugeeId !== 'undefined') ||
                    (this.companyId !== null && this.companyId !== 'undefined')) {
					return true;
				}
				return false;
            },
			refugeeId() {
				return JSON.parse(localStorage.getItem('userId'));
			},
			companyId: () => JSON.parse(window.localStorage.getItem('companyId')),
			showRefugeeProfile() {
				if (this.refugeeId !== null && typeof this.refugeeId !== 'undefined') {
					return true;
				}
				return false;
			},
		},
        methods: {
            ...mapActions([
            	'getRefugeeData',
				'getDonatorData',
            ]),
			getClass(path) {
				let classes = [];
				if (this.$route.path === path) {
					classes.push('active');
				}
				return classes;
			},
			logout() {
				window.localStorage.removeItem('companyId');
				window.localStorage.removeItem('userId');
				window.localStorage.removeItem('wrongAnswers');
				this.$router.push({
					path: '/',
				});
			},
        },
		mounted() {
			if (this.refugeeId !== null) {
				this.getRefugeeData({ id: this.refugeeId });
				return;
			}
			if (this.companyId !== null) {
                this.getDonatorData(this.companyId);
			}
		},
	};
</script>

<style lang="scss" type="text/scss">
    .app-container {
        margin-bottom: 2.5rem;
    }

    .logo {
        height: 100%;
        padding: 10px 0;

        img {
            display: block;
            width: auto;
            height: 100%;
        }
    }
</style>
