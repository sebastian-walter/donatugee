<template>
    <v-app id="donatugee">
        <v-navigation-drawer
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
                        <v-list-tile-title>All Challenges</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
                <v-list-tile @click="">
                    <v-list-tile-action>
                        <v-icon>power_settings_new</v-icon>
                    </v-list-tile-action>
                    <v-list-tile-content>
                        <v-list-tile-title @click="logout">Logout</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>
        <v-toolbar color="cyan" dark fixed app>
            <v-spacer></v-spacer>
            <v-toolbar-title>Application</v-toolbar-title>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        </v-toolbar>
        <v-content>
            <v-container fluid fill-height>
                <v-layout row>
                    <v-flex xs12>
                        <slot name="content"></slot>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
        <v-card>
            <v-bottom-nav fixed :value="true" color="white">
              <v-btn :class="getClass('/challenges')" :to="{path:'/challenges'}" flat color="primary">
                <span>Home</span>
                <v-icon>home</v-icon>
              </v-btn>
              <v-btn :class="getClass('/challenges')" :to="{path:'/challenges'}" flat color="primary">
                <span>Your Challenges</span>
                <v-icon>star</v-icon>
              </v-btn>
              <v-btn v-if="refugeeId" :class="getClass('/refugee/profile/')" :to="{path:'/refugee/profile/' + refugeeId}" flat color="primary">
                <span>Profile</span>
                <v-icon>person</v-icon>
              </v-btn>
              <v-btn v-if="companyId" :class="getClass('/company/profile/')" :to="{path:'/company/profile/' + refugeeId}" flat color="primary">
                <span>Profile</span>
                <v-icon>person</v-icon>
              </v-btn>
            </v-bottom-nav>
        </v-card>
    </v-app>
</template>

<script>
	export default {
		name: 'Layout',
		data: () => ({
			drawer: false,
            refugeeId: window.localStorage.getItem('userId'),
            companyId: window.localStorage.getItem('userIdCompany'),
		}),
        methods: {
            getClass(path) {
                let classes = [];
                if (this.$route.path === path) {
                    classes.push('active');
                }
                return classes;
            },
        },
		props: {
			source: String,
		},
        methods: {
			logout() {
				window.localStorage.removeItem('idChallenge');
				window.localStorage.removeItem('userId');
				window.localStorage.removeItem('email');
				window.localStorage.removeItem('name');
				window.localStorage.removeItem('skills');
				window.localStorage.removeItem('wrongAnswers');
				window.localStorage.removeItem('authenticated');
				this.$router.push({
                    path: '/',
                });
            }
        }
	};
</script>
