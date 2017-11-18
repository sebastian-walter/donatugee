<template>
    <div class="mb-4">
        <v-flex xs12>
            <v-card :to="{ path: 'challenge/' + challenge.ChallengeID }">
                <v-card-title>
                    <div>
                        <h3 class="headline mb-2">{{ challenge.Name }}</h3>
                        <p>{{ challenge.Description }}</p>
                        <v-chip color="orange" text-color="white">
                            <v-icon left>laptop</v-icon>
                            Laptop
                        </v-chip>
                        <v-chip color="indigo" text-color="white">
                            <v-icon left>code</v-icon>
                            Udemy
                        </v-chip>
                    </div>
                </v-card-title>
                <v-divider></v-divider>
                <v-card-title primary-title class="pt-3">
                    <v-flex xs2>
                        <v-avatar
                                :tile="tile"
                                :size="avatarSize"
                                class="grey lighten-4"
                        >
                            <img src="https://lorempixel.com/180/180/cats/" alt="avatar">
                        </v-avatar>
                    </v-flex>
                    <v-flex xs9 offset-xs1>
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
        name: 'ChallengeItem',
        props: {
            challenge: { type: Object, required: true }
        }
    }
</script>