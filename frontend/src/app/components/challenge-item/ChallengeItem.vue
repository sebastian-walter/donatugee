<template>
    <div>
        <v-flex xs12>
            <v-card>
                <v-card-title>
                    <div>
                        <h3 class="headline mb-0">Challenge Headline</h3>
                        <p>Description</p>
                    </div>
                </v-card-title>
                <v-card-title primary-title class="pt-0">
                    <v-flex xs2>
                        <v-avatar
                                :tile="tile"
                                :size="avatarSize"
                                class="grey lighten-4"
                        >
                            <img src="/static/apple-touch-icon-180x180.png" alt="avatar">
                        </v-avatar>
                    </v-flex>
                    <v-flex xs10>
                        <h3 class="mb-0">Trivago</h3>
                        <div>Location</div>
                    </v-flex>
                </v-card-title>
            </v-card>
        </v-flex>
    </div>
</template>

<script>
    export default {
        name: 'ChallengeItem'
    }
</script>